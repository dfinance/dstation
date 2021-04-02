package keeper

import (
	"encoding/hex"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/dfinance/dstation/pkg/types/dvm"
	"github.com/dfinance/dstation/x/vm/types"
)

// InitGenesis inits module genesis state: creates currencies.
func (k Keeper) InitGenesis(ctx sdk.Context, state *types.GenesisState) {
	for genWOIdx, genWriteOp := range state.WriteSet {
		accessPath, value, err := genWriteOp.ToBytes()
		if err != nil {
			panic(fmt.Errorf("writeSet [%d]: %w", genWOIdx, err))
		}

		k.SetValue(ctx, accessPath, value)
	}

	// Edge-case: set storage context for DS before the BeginBlock occurs
	k.dsServer.SetContext(ctx)
}

// ExportGenesis exports module genesis state using current params state.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	state := types.GenesisState{}
	k.iterateVMStorageValues(ctx, func(accessPath *dvm.VMAccessPath, value []byte) bool {
		writeSetOp := types.GenesisState_WriteOp{
			Address: hex.EncodeToString(accessPath.Address),
			Path:    hex.EncodeToString(accessPath.Path),
			Value:   hex.EncodeToString(value),
		}
		state.WriteSet = append(state.WriteSet, writeSetOp)

		return true
	})

	return &state
}
