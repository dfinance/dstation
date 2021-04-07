package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	dvmTypes "github.com/dfinance/dstation/pkg/types/dvm"
	"github.com/dfinance/dstation/x/vm/types"
)

// GetVmAccountBalance implements types.AccountBalanceResProvider interface.
func (k Keeper) GetVmAccountBalance(ctx sdk.Context, accAddress sdk.AccAddress, denom string) *dvmTypes.U128 {
	coinBalance := k.bankKeeper.GetBalance(ctx, accAddress, denom)

	value, err := types.SdkIntToVmU128(coinBalance.Amount)
	if err != nil {
		panic(fmt.Errorf("converting balance (%s) for denom (%s) to VM U128 мфдгу: %w", coinBalance, denom, err))
	}

	return value
}
