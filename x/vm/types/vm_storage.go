package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dfinance/dvm-proto/go/vm_grpc"
)

// DSDataMiddleware defines prototype for DSServer middleware.
type DSDataMiddleware func(ctx sdk.Context, path *vm_grpc.VMAccessPath) ([]byte, error)

// VMStorage interface defines VM storage IO operations.
type VMStorage interface {
	HasValue(ctx sdk.Context, accessPath *vm_grpc.VMAccessPath) bool
	GetValue(ctx sdk.Context, accessPath *vm_grpc.VMAccessPath) []byte
	SetValue(ctx sdk.Context, accessPath *vm_grpc.VMAccessPath, value []byte)
	DelValue(ctx sdk.Context, accessPath *vm_grpc.VMAccessPath)
}
