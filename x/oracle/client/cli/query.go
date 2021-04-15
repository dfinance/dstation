package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/dfinance/dstation/pkg"
	"github.com/dfinance/dstation/x/oracle/types"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module.
func GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the Oracle module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	queryCmd.AddCommand(
		GetCmdQueryOracles(),
		GetCmdQueryAssets(),
		GetCmdQueryCurrentPrice(),
		GetCmdQueryCurrentPrices(),
	)

	return queryCmd
}

// GetCmdQueryOracles returns query command that implement keeper querier.
func GetCmdQueryOracles() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "oracles",
		Short: "Get all registered Oracle sources",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			// Build and send request
			res, err := queryClient.Oracles(cmd.Context(), &types.QueryOraclesRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryAssets returns query command that implement keeper querier.
func GetCmdQueryAssets() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "assets",
		Short: "Get all registered Assets",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			// Build and send request
			res, err := queryClient.Assets(cmd.Context(), &types.QueryAssetsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryCurrentPrice returns query command that implement keeper querier.
func GetCmdQueryCurrentPrice() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "price [assetCode]",
		Short: "Get the latest aggregated from all registered source price",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			// Parse inputs
			assetCode, err := pkg.ParseAssetCodeParam("assetCode", args[0], pkg.ParamTypeCliArg)
			if err != nil {
				return err
			}
			leftDenom, rightDenom := assetCode.Split()

			// Build and send request
			res, err := queryClient.CurrentPrice(cmd.Context(), &types.QueryCurrentPriceRequest{
				LeftDenom:  leftDenom,
				RightDenom: rightDenom,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	pkg.BuildCmdHelp(cmd, []string{
		"asset code symbol (eth_usdt)",
	})

	return cmd
}

// GetCmdQueryCurrentPrices returns query command that implement keeper querier.
func GetCmdQueryCurrentPrices() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "prices",
		Short: "Get the latest aggregated from all registered source prices",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			// Build and send request
			res, err := queryClient.CurrentPrices(cmd.Context(), &types.QueryCurrentPricesRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
