package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	dvmTypes "github.com/dfinance/dstation/pkg/types/dvm"
)

// DSDataMiddleware defines prototype for DSServer middleware.
type DSDataMiddleware func(ctx sdk.Context, path *dvmTypes.VMAccessPath) ([]byte, error)

// VMStorage interface defines VM storage IO operations: WriteSet ops.
type VMStorage interface {
	// HasValue checks if VMStorage has a writeSet data by dvmTypes.VMAccessPath.
	HasValue(ctx sdk.Context, accessPath *dvmTypes.VMAccessPath) bool
	// GetValue returns a VMStorage writeSet data by dvmTypes.VMAccessPath.
	GetValue(ctx sdk.Context, accessPath *dvmTypes.VMAccessPath) []byte
	// SetValue sets the VMStorage writeSet data by dvmTypes.VMAccessPath.
	SetValue(ctx sdk.Context, accessPath *dvmTypes.VMAccessPath, value []byte)
	// DelValue removes the VMStorage writeSet data by dvmTypes.VMAccessPath.
	DelValue(ctx sdk.Context, accessPath *dvmTypes.VMAccessPath)
}

// CurrencyInfoResProvider interfaces defines DVM CurrencyInfo resource provider.
type CurrencyInfoResProvider interface {
	// GetVmCurrencyInfo returns dvmTypes.CurrencyInfo resource if found.
	GetVmCurrencyInfo(ctx sdk.Context, denom string) *dvmTypes.CurrencyInfo
}

// AccountBalanceResProvider interface defines DVM native account balance resource provider.
type AccountBalanceResProvider interface {
	// GetVmAccountBalance returns VM account native balance resource (SDK account balance).
	// Contract: always returns non-nil value.
	GetVmAccountBalance(ctx sdk.Context, accAddress sdk.AccAddress, denom string) *dvmTypes.U128
}
