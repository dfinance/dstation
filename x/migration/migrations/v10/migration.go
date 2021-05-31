package v10

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	genutilTypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/dfinance/dstation/x/migration/migrations"
	v076 "github.com/dfinance/dstation/x/migration/migrations/v076"
	tmTypes "github.com/tendermint/tendermint/types"
)

// Migrate migrates exported state from dnode v0.7.6 (v0.7.5) to the dstation v1.0.0 genesis state.
// We only pick data we need, the old genesis is no longer needed (the default one is build instead).
func Migrate(oldGenDoc *tmTypes.GenesisDoc, chainId string, genTime time.Time, clientCtx client.Context) (*tmTypes.GenesisDoc, error) {
	cdc := clientCtx.JSONMarshaler.(codec.Marshaler)

	// Input checks
	if oldGenDoc == nil {
		return nil, fmt.Errorf("oldGenDoc: nil")
	}
	if chainId == "" {
		return nil, fmt.Errorf("chainId: empty")
	}
	if genTime.IsZero() {
		return nil, fmt.Errorf("genTime: empty")
	}

	// Build a new GenDoc with default and custom Dfinance param overwrites
	newGenDoc, err := buildNewDefaultGenDoc(cdc, chainId, genTime)
	if err != nil {
		return nil, fmt.Errorf("building default GenesisDoc: %w", err)
	}

	// Migrate modules
	var oldAppState, newAppState genutilTypes.AppMap
	if err := json.Unmarshal(oldGenDoc.AppState, &oldAppState); err != nil {
		return nil, fmt.Errorf("old GenesisDoc: AppState JSON unmarshal: %w", err)
	}
	if err := json.Unmarshal(newGenDoc.AppState, &newAppState); err != nil {
		return nil, fmt.Errorf("new GenesisDoc: AppState JSON unmarshal: %w", err)
	}

	if err := migrateModule(cdc, v076.AuthModuleName, authTypes.ModuleName, oldAppState, newAppState, migrateAuthModule); err != nil {
		return nil, err
	}
	if err := migrateModule(cdc, v076.AuthModuleName, bankTypes.ModuleName, oldAppState, newAppState, migrateBankModule); err != nil {
		return nil, err
	}

	// Update new genDoc
	newGenDoc.AppState, err = json.Marshal(newAppState)
	if err != nil {
		return nil, fmt.Errorf("new GenesisDoc: AppState JSON marshal: %w", err)
	}

	return newGenDoc, nil
}

// migrateModule migrates a single module GenesisState.
func migrateModule(cdc codec.Marshaler, oldModuleName, newModuleName string, oldAppState, newAppState genutilTypes.AppMap, handler migrations.ModuleMigrationHandler) error {
	newGenState, err := handler(cdc, oldAppState[oldModuleName], newAppState[newModuleName])
	if err != nil {
		return fmt.Errorf("migrating module (%s -> %s): %w", oldModuleName, newModuleName, err)
	}

	newGenStateBz, err := cdc.MarshalJSON(newGenState)
	if err != nil {
		return fmt.Errorf("migrating module (%s -> %s): new state proto JSON marshal: %w", oldModuleName, newModuleName, err)
	}
	newAppState[newModuleName] = newGenStateBz

	return nil
}
