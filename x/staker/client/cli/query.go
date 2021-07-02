package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/dfinance/dstation/pkg"
	"github.com/dfinance/dstation/x/staker/types"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module.
func 	GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the Staker module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	queryCmd.AddCommand(
		GetCmdQueryCall(),
		GetCmdQueryCallByUniqueId(),
		GetCmdQueryAccCalls(),
		GetCmdQueryParams(),
	)

	return queryCmd
}

// GetCmdQueryCall returns query command that implement keeper querier.
func GetCmdQueryCall() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "call [callID]",
		Short:   "Get Call by ID",
		Example: "call 100",
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			// Parse inputs
			id, err := pkg.ParseSdkUintParam("callID", args[0], pkg.ParamTypeCliArg)
			if err != nil {
				return err
			}

			// Build and send request
			res, err := queryClient.CallById(cmd.Context(), &types.QueryCallByIdRequest{Id: id})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryCallByUniqueId returns query command that implement keeper querier.
func GetCmdQueryCallByUniqueId() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "call-unique [uniqueID]",
		Short:   "Get Call by unique operation ID",
		Example: "call-unique 0x0BAF",
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			// Parse inputs
			uniqueId := args[0]

			// Build and send request
			res, err := queryClient.CallByUniqueId(cmd.Context(), &types.QueryCallByUniqueIdRequest{UniqueId: uniqueId})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryAccCalls returns query command that implement keeper querier.
func GetCmdQueryAccCalls() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "account-calls [address]",
		Short:   "Get Calls for account",
		Example: "account-calls wallet1jk4ld0uu6wdrj9t8u3gghm9jt583hxx7xp7he8",
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			// Parse inputs
			accAddr, err := pkg.ParseSdkAddressParam("address", args[0], pkg.ParamTypeCliArg)
			if err != nil {
				return err
			}

			// Build and send request
			res, err := queryClient.CallsByAccount(cmd.Context(), &types.QueryCallsByAccountRequest{Address: accAddr.String()})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryParams returns query command that implement keeper querier.
func GetCmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "params",
		Short:   "Get Staker module parameters",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			// Build and send request
			res, err := queryClient.Params(cmd.Context(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}