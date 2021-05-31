package client

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dfinance/dstation/pkg"
	"github.com/dfinance/dstation/x/migration/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	tmJson "github.com/tendermint/tendermint/libs/json"
	tmTypes "github.com/tendermint/tendermint/types"
)

const (
	flagGenesisTime = "genesis-time"
	flagChainID     = "chain-id"
	flagOutputPath  = "output"
)

// MigrateGenesisCmd returns a command to execute genesis state migration.
func MigrateGenesisCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "migrate [targetVersion] [genesisFile]",
		Short:   "Migrate genesis state to a specified target version",
		Example: "migrate v1.0.0 ./genesis_old.json --chain-id=mainnet --genesis-time=2021-05-31T15:00:00Z",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			// Parse inputs
			targetVersion := args[0]
			migrationHandler := types.MigrationMap[targetVersion]
			if migrationHandler == nil {
				return pkg.BuildError("targetVersion", targetVersion, pkg.ParamTypeCliArg, "migration handler not found")
			}

			genFile := args[1]
			if err := pkg.CheckFileExists("genesisFile", genFile, pkg.ParamTypeCliArg); err != nil {
				return err
			}

			updChainID := cmd.Flag(flagChainID).Value.String()

			var updGenesisTime time.Time
			if value := cmd.Flag(flagGenesisTime).Value.String(); value != "" {
				if err := updGenesisTime.UnmarshalText([]byte(value)); err != nil {
					return pkg.BuildError(flagGenesisTime, value, pkg.ParamTypeCliFlag, fmt.Sprintf("failed to unmarshal time: %v", err))
				}
			}

			// Retrieve the app state
			genDoc, err := tmTypes.GenesisDocFromFile(genFile)
			if err != nil {
				return fmt.Errorf("reading initial appState (%s): %w", genFile, err)
			}

			// Migrate
			genDocMigrated, err := migrationHandler(genDoc, updChainID, updGenesisTime, clientCtx)
			if err != nil {
				return fmt.Errorf("migration to %s failed: %w", targetVersion, err)
			}

			// Marshal the resulting genDoc and do the output
			bz, err := tmJson.Marshal(genDocMigrated)
			if err != nil {
				return fmt.Errorf("genesisDoc JSON marshal: %w", err)
			}

			sortedBz, err := sdk.SortJSON(bz)
			if err != nil {
				return fmt.Errorf("genesisDoc JSON sorting: %w", err)
			}

			if err := processOutput(sortedBz); err != nil {
				return fmt.Errorf("output processing:: %w", err)
			}

			return nil
		},
	}
	pkg.BuildCmdHelp(cmd, []string{
		"target migration version (Golang semver format without build version)",
		"path to an exported genesis state file for the current state version",
	})
	cmd.Flags().String(flagGenesisTime, "", "override genesis_time")
	cmd.Flags().String(flagChainID, "", "override chain_id")
	cmd.Flags().String(flagOutputPath, "./genesis_new.json", "export file path (print to stdout if empty)")

	return cmd
}

func processOutput(docBz []byte) error {
	outputPath := viper.GetString(flagOutputPath)
	if outputPath == "" {
		fmt.Println(string(docBz))
		return nil
	}

	if err := ioutil.WriteFile(outputPath, docBz, 0644); err != nil {
		return fmt.Errorf("write to file (%s): %w", outputPath, err)
	}

	return nil
}
