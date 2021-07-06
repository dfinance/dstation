package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dfinance/dstation/x/namespace/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	Keeper
}

func (k msgServer) Buy(c context.Context, call *types.MsgBuyCall) (*types.MsgBuyResponse, error) {
	if call == nil {
		return nil, status.Error(codes.InvalidArgument, "empty msg")
	}

	ctx := sdk.UnwrapSDKContext(c)
	wh, err := k.Keeper.Buy(ctx, call.Value, call.Address, call.Amount)
	if err != nil {
		return nil, err
	}

	return &types.MsgBuyResponse{Id: wh.ID}, nil
}

func (k msgServer) Delete(c context.Context, call *types.MsgDeleteCall) (*types.MsgDeleteResponse, error) {
	if call == nil {
		return nil, status.Error(codes.InvalidArgument, "empty msg")
	}

	ctx := sdk.UnwrapSDKContext(c)
	err := k.Keeper.Delete(ctx, call.Value, call.Address)
	if err != nil {
		return nil, err
	}

	return &types.MsgDeleteResponse{}, nil
}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}