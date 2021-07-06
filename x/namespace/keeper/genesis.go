package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dfinance/dstation/x/namespace/types"
)

// InitGenesis inits module genesis state: creates Whois, the LastWhoisId.
func (k Keeper) InitGenesis(ctx sdk.Context, state *types.GenesisState) {
	if state.LastWhoisId != nil {
		k.setLastWhoisID(ctx, *state.LastWhoisId)
	}

	for _, call := range state.Whois {
		k.setWhois(ctx, call)
	}
}

// ExportGenesis exports module genesis state using current params state.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	state := types.GenesisState{
		LastWhoisId: k.getLastWhoisID(ctx),
		Whois:      k.GetAllWhois(ctx),
	}

	return &state
}
