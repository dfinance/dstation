package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/dfinance/dstation/pkg"
	"github.com/dfinance/dstation/x/staker/types"
	"github.com/spf13/cobra"
)

const (
	flagSrcEthAddress = "src-eth-address"
	flagSrcChainId    = "src-chain-id"
)

// GetTxCmd returns a root CLI command handler for all module transaction commands.
func GetTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Staker transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		GetCmdTxDeposit(),
		GetCmdTxWithdraw(),
		GetCmdTxGenAddNominee(),
	)

	return txCmd
}

// GetCmdTxDeposit returns tx command that implement keeper handler.
func GetCmdTxDeposit() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "deposit [uniqueId] [targetAccAddress] [amount]",
		Short:   "Deposit tokens to the target account",
		Example: "deposit wallet1jk4ld0uu6wdrj9t8u3gghm9jt583hxx7xp7he8 1xfi,1btc --from nominee_account --fees 10000xfi --gas 500000",
		Args:    cobra.MinimumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags()).WithTxConfig(clientCtx.TxConfig).WithAccountRetriever(clientCtx.AccountRetriever)

			// Parse inputs
			uniqueId := args[0]

			fromAddr, err := pkg.ParseFromFlag(clientCtx)
			if err != nil {
				return err
			}

			accAddr, err := pkg.ParseSdkAddressParam("targetAccAddress", args[1], pkg.ParamTypeCliArg)
			if err != nil {
				return err
			}

			amount, err := pkg.ParseCoinsParam("amount", args[2], pkg.ParamTypeCliArg)
			if err != nil {
				return err
			}

			srcMeta := parseCallSource(cmd)

			// Build msg
			msg := types.NewMsgDepositCall(uniqueId, fromAddr, accAddr, srcMeta.EthAddress, srcMeta.ChainId, amount)

			return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)
	addSourceFlags(cmd)

	pkg.BuildCmdHelp(cmd, []string{
		"unique operation ID (Ethereum Tx hash for example)",
		"target account address",
		"operation amount (coins)",
	})

	return cmd
}

// GetCmdTxWithdraw returns tx command that implement keeper handler.
func GetCmdTxWithdraw() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "withdraw [uniqueId] [targetAccAddress] [amount]",
		Short:   "Withdraw tokens from the target account",
		Example: "withdraw wallet1jk4ld0uu6wdrj9t8u3gghm9jt583hxx7xp7he8 1xfi,1btc --from nominee_account --fees 10000xfi --gas 500000",
		Args:    cobra.MinimumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags()).WithTxConfig(clientCtx.TxConfig).WithAccountRetriever(clientCtx.AccountRetriever)

			// Parse inputs
			uniqueId := args[0]

			fromAddr, err := pkg.ParseFromFlag(clientCtx)
			if err != nil {
				return err
			}

			accAddr, err := pkg.ParseSdkAddressParam("targetAccAddress", args[1], pkg.ParamTypeCliArg)
			if err != nil {
				return err
			}

			amount, err := pkg.ParseCoinsParam("amount", args[2], pkg.ParamTypeCliArg)
			if err != nil {
				return err
			}

			srcMeta := parseCallSource(cmd)

			// Build msg
			msg := types.NewMsgWithdrawCall(uniqueId, fromAddr, accAddr, srcMeta.EthAddress, srcMeta.ChainId, amount)

			return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)
	addSourceFlags(cmd)

	pkg.BuildCmdHelp(cmd, []string{
		"unique operation ID (Ethereum Tx hash for example)",
		"target account address",
		"operation amount (coins)",
	})

	return cmd
}

func addSourceFlags(cmd *cobra.Command) {
	cmd.Flags().String(flagSrcEthAddress, "", "operation source: Ethereum address (HEX string, optional)")
	cmd.Flags().String(flagSrcChainId, "", "operation source: chain ID (optional)")
}

func parseCallSource(cmd *cobra.Command) types.CallSourceMeta {
	srcMeta := types.CallSourceMeta{}

	if value, err := cmd.Flags().GetString(flagSrcEthAddress); err == nil {
		srcMeta.EthAddress = value
	}
	if value, err := cmd.Flags().GetString(flagSrcChainId); err == nil {
		srcMeta.ChainId = value
	}

	return srcMeta
}
