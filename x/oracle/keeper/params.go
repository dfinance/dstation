package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dfinance/dstation/x/oracle/types"
)

// GetParams returns the keeper parameters.
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	params := types.Params{}
	k.paramSpace.GetParamSet(ctx, &params)

	return params
}

// SetParams sets the keeper parameters to the param space.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}
