package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/dfinance/dstation/x/namespace/types"
	"github.com/spf13/cobra"
)

// GetQueryCmd returns the cli query commands for this module.
func 	GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the  Namespace module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	queryCmd.AddCommand(
		DomainsAccount(),
	)

	return queryCmd
}

// GetCmdQueryCall returns query command that implement keeper querier.
func DomainsAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "domains [accountAddress]",
		Short:   "Get domains by account",
		Example: "domains wallet1jk4ld0uu6wdrj9t8u3gghm9jt583hxx7xp7he8",
		Args:    cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			address := args[0]

			// Build and send request
			res, err := queryClient.DomainsAccount(cmd.Context(), &types.DomainsAccountRequest{Address: address})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}