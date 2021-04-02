package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	authClient "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/dfinance/dstation/x/vm/types"

	"github.com/dfinance/dstation/pkg"
)

// GetQueryCmd returns the cli query commands for this module.
func GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the VM module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	queryCmd.AddCommand(
		GetCmdQueryData(),
		GetCmdQueryTxVmStatus(),
		GetCmdQueryCompile(),
	)

	return queryCmd
}

// GetCmdQueryData returns query command that implement keeper querier.
func GetCmdQueryData() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "data [address] [path]",
		Short:   "Get write set data from the VMStorage by address and path",
		Example: "data wallet1jk4ld0uu6wdrj9t8u3gghm9jt583hxx7xp7he8 0019b01c2cf3c2160a43e4dcad70e3e5d18151cc38de7a1d1067c6031bfa0ae4d9",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			// Parse inputs
			address, err := pkg.ParseSdkAddressParam("address", args[0], pkg.ParamTypeCliArg)
			if err != nil {
				return err
			}

			_, path, err := pkg.ParseHexStringParam("path", args[1], pkg.ParamTypeCliArg)
			if err != nil {
				return err
			}

			// Build and send request
			res, err := queryClient.Data(cmd.Context(), &types.QueryDataRequest{
				Address: types.Bech32ToLibra(address),
				Path:    path,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	pkg.BuildCmdHelp(cmd, []string{
		"VM address [Bech32 / HEX string]",
		"VM path [HEX string]",
	})

	return cmd
}

// GetCmdQueryTxVmStatus returns query command that implement keeper querier.
func GetCmdQueryTxVmStatus() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "tx [hash]",
		Short:   "Get VM status for the Tx by hash",
		Example: "tx 6D5A4D889BCDB4C71C6AE5836CD8BC1FD8E0703F1580B9812990431D1796CE34",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			// Get Tx
			tx, err := authClient.QueryTx(clientCtx, args[0])
			if err != nil {
				return err
			}
			if tx == nil || tx.Empty() {
				return fmt.Errorf("transaction not found")
			}

			// Build and send request
			res, err := queryClient.TxVmStatus(cmd.Context(), &types.QueryTxVmStatusRequest{
				TxMeta: *tx,
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	pkg.BuildCmdHelp(cmd, []string{
		"transaction hash code [HEX string]",
	})

	return cmd
}

// GetCmdQueryCompile returns query command that implement keeper querier.
func GetCmdQueryCompile() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "compile [moveFile] [account]",
		Short:   "Compile script / module using source code from Move file",
		Example: "compile script.move wallet196udj7s83uaw2u4safcrvgyqc0sc3flxuherp6 --to-file script.json",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			// Parse inputs
			moveContent, err := pkg.ParseFilePath("moveFile", args[0], pkg.ParamTypeCliArg)
			if err != nil {
				return err
			}

			address, err := pkg.ParseSdkAddressParam("account", args[1], pkg.ParamTypeCliArg)
			if err != nil {
				return err
			}

			// Build and send request
			res, err := queryClient.Compile(cmd.Context(), &types.QueryCompileRequest{
				Address: address,
				Code:    string(moveContent),
			})
			if err != nil {
				return err
			}

			// Output
			outputFilePath := viper.GetString(FlagOutput)
			outputBz, err := json.MarshalIndent(res, "", "    ")
			if err != nil {
				return fmt.Errorf("output: json marshal: %w", err)
			}

			if outputFilePath == "" {
				return clientCtx.PrintBytes(outputBz)
			}

			if err := ioutil.WriteFile(outputFilePath, outputBz, 0644); err != nil {
				return fmt.Errorf("output: file save: %w", err)
			}

			return clientCtx.PrintString(fmt.Sprintf("Result saved to file: %s", outputFilePath))
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	AddOutputFlagToCmd(cmd)

	pkg.BuildCmdHelp(cmd, []string{
		"path to .move file",
		"account address [Bech32 / HEX string]",
	})

	return cmd
}
