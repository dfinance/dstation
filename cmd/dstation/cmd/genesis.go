package cmd

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/types"

	dnConfig "github.com/dfinance/dstation/cmd/dstation/config"
)

// SetGenesisDefaultsCmd returns set-genesis-defaults cobra Command.
func SetGenesisDefaultsCmd(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-genesis-defaults",
		Short: "Update an existing genesis file with Dfinance default params",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			serverCtx := server.GetServerContextFromCmd(cmd)

			genFile := serverCtx.Config.GenesisFile()
			genDoc, err := types.GenesisDocFromFile(genFile)
			if err != nil {
				return fmt.Errorf("failed to read genesis doc from file: %w", err)
			}

			appState, err := dnConfig.SetGenesisDefaults(clientCtx.JSONMarshaler.(codec.Marshaler), genDoc.AppState)
			if err != nil {
				return fmt.Errorf("failed to set default genesis params:: %w", err)
			}
			genDoc.AppState = appState

			if err = genutil.ExportGenesisFile(genDoc, genFile); err != nil {
				return fmt.Errorf("failed to export gensis file: %w", err)
			}

			return clientCtx.PrintString(string(appState))
		},
	}

	cmd.Flags().String(cli.HomeFlag, defaultNodeHome, "node's home directory")

	return cmd
}
