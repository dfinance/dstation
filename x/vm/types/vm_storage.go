package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/dfinance/dstation/pkg/types/dvm"
)

// DSDataMiddleware defines prototype for DSServer middleware.
type DSDataMiddleware func(ctx sdk.Context, path *dvm.VMAccessPath) ([]byte, error)

// VMStorage interface defines VM storage IO operations.
type VMStorage interface {
	HasValue(ctx sdk.Context, accessPath *dvm.VMAccessPath) bool
	GetValue(ctx sdk.Context, accessPath *dvm.VMAccessPath) []byte
	SetValue(ctx sdk.Context, accessPath *dvm.VMAccessPath, value []byte)
	DelValue(ctx sdk.Context, accessPath *dvm.VMAccessPath)
}
