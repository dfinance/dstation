package keeper

import (
	"context"
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dfinance/dvm-proto/go/vm_grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/dfinance/dstation/pkg"
	"github.com/dfinance/dstation/x/vm/types"
)

var _ types.QueryServer = Querier{}

// Querier implements gRPC query RPCs.
type Querier struct {
	Keeper
}

// Data queries VMStorage value.
func (k Querier) Data(c context.Context, req *types.QueryDataRequest) (*types.QueryDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	// Validate and parse request
	addr, err := pkg.ParseSdkAddressParam("address", req.Address, pkg.ParamTypeRequest)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	_, path, err := pkg.ParseHexStringParam("path", req.Path, pkg.ParamTypeRequest)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	ctx := sdk.UnwrapSDKContext(c)
	value := k.GetValue(ctx, &vm_grpc.VMAccessPath{Address: types.Bech32ToLibra(addr), Path: path})

	return &types.QueryDataResponse{Value: hex.EncodeToString(value)}, nil
}

// TxVmStatus queries Tx VM status based on Tx ABCI logs.
func (k Querier) TxVmStatus(c context.Context, req *types.QueryTxVmStatusRequest) (*types.QueryTxVmStatusResponse, error) {
	txVmStatus := types.NewVmStatusFromABCILogs(req.TxMeta)
	return &types.QueryTxVmStatusResponse{VmStatus: txVmStatus}, nil
}
