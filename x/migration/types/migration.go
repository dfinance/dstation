package types

import (
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/dfinance/dstation/x/migration/migrations/v10"
	tmTypes "github.com/tendermint/tendermint/types"
)

// GenDocMigrationHandler converts an tmTypes.GenesisDoc (appState, consensus params, etc.) from the previous version to the targeted one.
type GenDocMigrationHandler func(initialGenDoc *tmTypes.GenesisDoc, newChainId string, newGenTime time.Time, clientCtx client.Context) (migratedGenDoc *tmTypes.GenesisDoc, retErr error)

// TargetMigrationMap defines a mapping from a migration target version to a GenDocMigrationHandler.
type TargetMigrationMap map[string]GenDocMigrationHandler

// MigrationMap is a registered migrations map.
var MigrationMap = TargetMigrationMap{
	"v1.0.0": v10.Migrate,
}
