package migrations

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/gogo/protobuf/proto"
)

// ModuleMigrationHandler migrates an old module GenesisState to a new one.
type ModuleMigrationHandler func(cdc codec.Marshaler, oldStateBz, newStateBz json.RawMessage) (newState proto.Message, retErr error)
