package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/types"

	"github.com/dfinance/dstation/app"
	dnConfig "github.com/dfinance/dstation/cmd/dstation/config"
	vmConfig "github.com/dfinance/dstation/x/vm/config"
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

			consParams, err := dnConfig.SetConsensusDefaults(genDoc.ConsensusParams)
			if err != nil {
				return fmt.Errorf("failed to set default consesnsus params:: %w", err)
			}
			genDoc.ConsensusParams = consParams

			var genState app.GenesisState
			if err := json.Unmarshal(genDoc.AppState, &genState); err != nil {
				return fmt.Errorf("genDoc.AppState json unmarshal: %w", err)
			}

			appState, err := dnConfig.SetGenesisDefaults(clientCtx.JSONMarshaler.(codec.Marshaler), genState)
			if err != nil {
				return fmt.Errorf("failed to set default genesis params:: %w", err)
			}

			appStateBz, err := json.MarshalIndent(appState, "", " ")
			if err != nil {
				return fmt.Errorf("appState json marshal: %w", err)
			}
			genDoc.AppState = appStateBz

			if err = genutil.ExportGenesisFile(genDoc, genFile); err != nil {
				return fmt.Errorf("failed to export gensis file: %w", err)
			}

			vmConfig.ReadVMConfig(serverCtx.Config.RootDir)

			return clientCtx.PrintString(string(appStateBz))
		},
	}

	cmd.Flags().String(cli.HomeFlag, defaultNodeHome, "node's home directory")

	return cmd
}
