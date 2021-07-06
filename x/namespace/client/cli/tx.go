package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/dfinance/dstation/pkg"
	"github.com/dfinance/dstation/x/namespace/types"
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
		BuyCmd(),
		DeleteCmd(),
	)

	return txCmd
}

// BuyCmd returns tx command that implement keeper handler.
func BuyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "buy [value] [address] [amount]",
		Short:   "Buy whois",
		Example: "buy xxx.com wallet1jk4ld0uu6wdrj9t8u3gghm9jt583hxx7xp7he8 1xfi,1btc",
		Args:    cobra.MinimumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags()).WithTxConfig(clientCtx.TxConfig).WithAccountRetriever(clientCtx.AccountRetriever)

			// Parse inputs
			value := args[0]

			address, err := pkg.ParseSdkAddressParam("address", args[1], pkg.ParamTypeCliArg)
			if err != nil {
				return err
			}

			amount, err := pkg.ParseCoinsParam("amount", args[2], pkg.ParamTypeCliArg)
			if err != nil {
				return err
			}

			// Build msg
			msg := types.NewMsgBuyCall(address, value, amount)

			return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, &msg)
		},
	}

	pkg.BuildCmdHelp(cmd, []string{
		"Domain name",
		"target account address",
		"operation amount (coins)",
	})

	return cmd
}

// DeleteCmd returns tx command that implement keeper handler.
func DeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete [value] [targetAccAddress]",
		Short:   "Delete tokens from the target account",
		Example: "delete xxx.com wallet1jk4ld0uu6wdrj9t8u3gghm9jt583hxx7xp7he8",
		Args:    cobra.MinimumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags()).WithTxConfig(clientCtx.TxConfig).WithAccountRetriever(clientCtx.AccountRetriever)

			// Parse inputs
			value := args[0]

			address, err := pkg.ParseSdkAddressParam("address", args[1], pkg.ParamTypeCliArg)
			if err != nil {
				return err
			}

			// Build msg
			msg := types.NewMsgDeleteCall(address, value)

			return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, &msg)
		},
	}

	pkg.BuildCmdHelp(cmd, []string{
		"domain name",
		"target account address",
	})

	return cmd
}


