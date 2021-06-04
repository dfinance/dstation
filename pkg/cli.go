package pkg

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/spf13/cobra"
)

// UpdateModuleGenesis read, updates module state and saves application genesis state to default genesis.json location.
// {genStatePtr} - pointer to module's genesis object.
// {genStateUpdater} - update handler ({genStatePtr} is passed through to the handler after unmarshal).
func UpdateModuleGenesis(cmd *cobra.Command, moduleName string, genStatePtr interface{}, genStateUpdater func(genStatePtr interface{}) error) error {
	if genStatePtr == nil {
		return fmt.Errorf("genStatePtr: nil")
	}
	if genStateUpdater == nil {
		return fmt.Errorf("genStateUpdater: nil")
	}

	serverCtx := server.GetServerContextFromCmd(cmd)
	clientCtx, err := client.GetClientTxContext(cmd)
	if err != nil {
		return err
	}

	config := serverCtx.Config
	config.SetRoot(clientCtx.HomeDir)

	genFile := config.GenesisFile()
	appState, genDoc, err := genutiltypes.GenesisStateFromGenFile(genFile)
	if err != nil {
		return fmt.Errorf("application genesis state: unmarshal failed (%s): %w", genFile, err)
	}

	if err := json.Unmarshal(appState[moduleName], genStatePtr); err != nil {
		return fmt.Errorf("module (%s) genesis state: unmarshal failed: %w", moduleName, err)
	}

	if err := genStateUpdater(genStatePtr); err != nil {
		return fmt.Errorf("module (%s) genesis state: update failed: %w", moduleName, err)
	}

	genesisStateBz, err := json.Marshal(genStatePtr)
	if err != nil {
		return fmt.Errorf("module (%s) genesis state: marshal failed: %w", moduleName, err)
	}
	appState[moduleName] = genesisStateBz

	appStateBz, err := json.Marshal(appState)
	if err != nil {
		return fmt.Errorf("application genesis state: marshal failed: %w", err)
	}
	genDoc.AppState = appStateBz

	return genutil.ExportGenesisFile(genDoc, genFile)
}
