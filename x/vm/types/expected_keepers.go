package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankExported "github.com/cosmos/cosmos-sdk/x/bank/exported"
)

// AccountKeeper defines the expected account keeper.
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authTypes.AccountI

	GetModuleAddress(name string) sdk.AccAddress
	GetModuleAccount(ctx sdk.Context, moduleName string) authTypes.ModuleAccountI

	SetModuleAccount(sdk.Context, authTypes.ModuleAccountI)
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins

	GetSupply(ctx sdk.Context) bankExported.SupplyI

	UndelegateCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	DelegateCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}
