package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dfinance/dstation/pkg"
	"github.com/dfinance/dstation/x/oracle/types"
	"github.com/spf13/cobra"
)

// GetTxCmd returns a root CLI command handler for all module transaction commands.
func GetTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Oracle transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		GetCmdTxSetOracle(),
		GetCmdTxSetAsset(),
		GetCmdTxPostPrice(),
		GetCmdTxGenAddNominee(),
		GetCmdTxGenAddOracle(),
		GetCmdTxGenAddAsset(),
	)

	return txCmd
}

// GetCmdTxSetOracle returns tx command that implement keeper handler.
func GetCmdTxSetOracle() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "set-oracle [oracleAccAddress] [priceMaxBytes] [priceDecimals] [sourceDescription]",
		Short:   "Add/update Oracle source",
		Example: "set-oracle wallet1jk4ld0uu6wdrj9t8u3gghm9jt583hxx7xp7he8 8 8 \"Binance\" --from nominee_account --fees 10000xfi --gas 500000",
		Args:    cobra.MinimumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags()).WithTxConfig(clientCtx.TxConfig).WithAccountRetriever(clientCtx.AccountRetriever)

			// Parse inputs
			fromAddr, err := pkg.ParseFromFlag(clientCtx)
			if err != nil {
				return err
			}

			oracleAddr, err := pkg.ParseSdkAddressParam("oracleAccAddress", args[0], pkg.ParamTypeCliArg)
			if err != nil {
				return err
			}

			priceMaxBytes, err := pkg.ParseUint8Param("priceMaxBytes", args[1], pkg.ParamTypeCliArg)
			if err != nil {
				return err
			}

			priceDecimals, err := pkg.ParseUint8Param("priceDecimals", args[2], pkg.ParamTypeCliArg)
			if err != nil {
				return err
			}

			oracleDesc := ""
			if len(args) > 3 {
				oracleDesc = args[3]
			}

			// Build msg
			msg := types.NewMsgSetOracle(fromAddr, oracleAddr, oracleDesc, uint32(priceMaxBytes), uint32(priceDecimals))

			return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	pkg.BuildCmdHelp(cmd, []string{
		"Oracle account address",
		"max number of bytes for ask/bid price values (limit)",
		"number of decimals for ask/bid price values (values send from Oracle are big.Int)",
		"Oracle description (optional)",
	})

	return cmd
}

// GetCmdTxSetAsset returns tx command that implement keeper handler.
func GetCmdTxSetAsset() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "set-asset [assetCode] [rateDecimals] [oracleAddress1...]",
		Short:   "Add/update Asset",
		Example: "set-asset eth_usdt 8 wallet1jk4ld0uu6wdrj9t8u3gghm9jt583hxx7xp7he8 --from nominee_account --fees 10000xfi --gas 500000",
		Args:    cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags()).WithTxConfig(clientCtx.TxConfig).WithAccountRetriever(clientCtx.AccountRetriever)

			// Parse inputs
			fromAddr, err := pkg.ParseFromFlag(clientCtx)
			if err != nil {
				return err
			}

			assetCode, err := pkg.ParseAssetCodeParam("assetCode", args[0], pkg.ParamTypeCliArg)
			if err != nil {
				return err
			}

			decimals, err := pkg.ParseUint8Param("rateDecimals", args[1], pkg.ParamTypeCliArg)
			if err != nil {
				return err
			}

			var oracles []sdk.AccAddress
			if len(args) > 2 {
				addrs, err := pkg.ParseSpaceSepSdkAddressesParams("oracleAddresses", args[2:], pkg.ParamTypeCliArg)
				if err != nil {
					return err
				}
				oracles = addrs
			}

			// Build msg
			msg := types.NewMsgSetAsset(fromAddr, assetCode, uint32(decimals), oracles...)

			return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	pkg.BuildCmdHelp(cmd, []string{
		"asset code symbol (eth_usdt)",
		"number of decimals for ask/bid price values",
		"source (Oracle) account addresses (optional, if none - asset is disabled)",
	})

	return cmd
}

// GetCmdTxPostPrice returns tx command that implement keeper handler.
func GetCmdTxPostPrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "post-price [assetCode] [askPrice] [bidPrice] [receivedAt]",
		Short:   "Post RawPrice from Oracle (used only for debug purposes)",
		Example: "post-price eth_usdt wallet1jk4ld0uu6wdrj9t8u3gghm9jt583hxx7xp7he8 10000000000 11000000000 1594732456 --from oracleAddress --fees 10000xfi --gas 500000",
		Args:    cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags()).WithTxConfig(clientCtx.TxConfig).WithAccountRetriever(clientCtx.AccountRetriever)

			// Parse inputs
			oracleAddr, err := pkg.ParseFromFlag(clientCtx)
			if err != nil {
				return err
			}

			assetCode, err := pkg.ParseAssetCodeParam("assetCode", args[0], pkg.ParamTypeCliArg)
			if err != nil {
				return err
			}

			askPrice, err := pkg.ParseSdkIntParam("askPrice", args[1], pkg.ParamTypeCliArg)
			if err != nil {
				return err
			}

			bidPrice, err := pkg.ParseSdkIntParam("bidPrice", args[2], pkg.ParamTypeCliArg)
			if err != nil {
				return err
			}

			receivedAt, err := pkg.ParseUnixTimestamp("receivedAt", args[3], pkg.ParamTypeCliArg)
			if err != nil {
				return err
			}

			// Build msg
			msg := types.NewMsgPostPrice(assetCode, oracleAddr, askPrice, bidPrice, receivedAt)

			return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, &msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(flags.FlagFrom)

	pkg.BuildCmdHelp(cmd, []string{
		"asset code symbol (eth_usdt)",
		"Ask price [Int]",
		"Bid price [Int]",
		"price received at UNIX timestamp in seconds [int]",
	})

	return cmd
}
