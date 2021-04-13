package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dfinance/dstation/x/oracle/types"
)

// InitGenesis inits module genesis state: creates currencies.
func (k Keeper) InitGenesis(ctx sdk.Context, state *types.GenesisState) {
	k.SetParams(ctx, state.Params)

	for _, oracle := range state.Oracles {
		k.setOracle(ctx, oracle)
	}

	for _, asset := range state.Assets {
		k.setAsset(ctx, asset)
	}

	for _, price := range state.CurrentPrices {
		k.setCurrentPrice(ctx, price)
	}
}

// ExportGenesis exports module genesis state using current params state.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	state := types.GenesisState{
		Params:        k.GetParams(ctx),
		Oracles:       k.GetOracles(ctx),
		Assets:        k.GetAssets(ctx),
		CurrentPrices: k.GetCurrentPrices(ctx),
	}

	return &state
}
