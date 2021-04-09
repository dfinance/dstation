package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/dfinance/dstation/x/vm/types"
)

// DelegateCoinsToPool transfers coins from an account to the module account (DecPool).
func (k Keeper) DelegateCoinsToPool(ctx sdk.Context, accAddr sdk.AccAddress, coins sdk.Coins) error {
	// Validate input
	if err := coins.Validate(); err != nil {
		return sdkErrors.Wrapf(sdkErrors.ErrInvalidCoins, "coins (%s): invalid: %v", coins, err)
	}

	// Transfer
	if err := k.bankKeeper.DelegateCoinsFromAccountToModule(ctx, accAddr, types.DelPoolName, coins); err != nil {
		return fmt.Errorf("bank operation failed: %w", err)
	}

	return nil
}

// UndelegateCoinsFromPool transfers coins from the module account (DecPool) to an account.
func (k Keeper) UndelegateCoinsFromPool(ctx sdk.Context, accAddr sdk.AccAddress, coins sdk.Coins) error {
	// Validate input
	if err := coins.Validate(); err != nil {
		return sdkErrors.Wrapf(sdkErrors.ErrInvalidCoins, "coins (%s): invalid: %v", coins, err)
	}

	// Transfer
	if err := k.bankKeeper.UndelegateCoinsFromModuleToAccount(ctx, types.DelPoolName, accAddr, coins); err != nil {
		return fmt.Errorf("bank operation failed: %w", err)
	}

	return nil
}

// GetDelegatedPoolSupply returns module account (DecPool) current balance.
func (k Keeper) GetDelegatedPoolSupply(ctx sdk.Context) sdk.Coins {
	moduleAcc := k.accKeeper.GetModuleAccount(ctx, types.DelPoolName)

	return k.bankKeeper.GetAllBalances(ctx, moduleAcc.GetAddress())
}
