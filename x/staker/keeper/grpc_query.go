package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dfinance/dstation/x/staker/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Querier{}

// Querier implements gRPC query RPCs.
type Querier struct {
	Keeper
}

func (k Querier) CallById(c context.Context, req *types.QueryCallByIdRequest) (*types.QueryCallByIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryCallByIdResponse{
		Call: k.Keeper.GetCallById(ctx, req.Id),
	}, nil
}

func (k Querier) CallByUniqueId(c context.Context, req *types.QueryCallByUniqueIdRequest) (*types.QueryCallByUniqueIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryCallByUniqueIdResponse{
		Call: k.Keeper.GetCallByUniqueId(ctx, req.UniqueId),
	}, nil
}

func (k Querier) CallsByAccount(c context.Context, req *types.QueryCallsByAccountRequest) (*types.QueryCallsByAccountResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	accAddr, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "address: invalid: %v", err)
	}

	return &types.QueryCallsByAccountResponse{
		Calls: k.Keeper.GetAddressCalls(ctx, accAddr),
	}, nil
}

func (k Querier) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetParams(ctx)

	return &types.QueryParamsResponse{
		Params: params,
	}, nil
}
