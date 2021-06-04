package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dfinance/dstation/x/staker/types"
)

// InitGenesis inits module genesis state: creates Calls, the latest call ID.
func (k Keeper) InitGenesis(ctx sdk.Context, state *types.GenesisState) {
	k.SetParams(ctx, state.Params)

	if state.LastCallId != nil {
		k.setLastCallID(ctx, *state.LastCallId)
	}

	for _, call := range state.Calls {
		k.setCall(ctx, call)
	}
}

// ExportGenesis exports module genesis state using current params state.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	state := types.GenesisState{
		Params:     k.GetParams(ctx),
		LastCallId: k.getLastCallID(ctx),
		Calls:      k.GetAllCalls(ctx),
	}

	return &state
}
