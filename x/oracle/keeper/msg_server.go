package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	dnTypes "github.com/dfinance/dstation/pkg/types"
	"github.com/dfinance/dstation/x/oracle/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (k msgServer) SetOracle(c context.Context, msg *types.MsgSetOracle) (*types.MsgSetOracleResponse, error) {
	if msg == nil {
		return nil, status.Error(codes.InvalidArgument, "empty msg")
	}

	ctx := sdk.UnwrapSDKContext(c)
	if err := k.Keeper.SetOracle(ctx, *msg); err != nil {
		return nil, err
	}

	return &types.MsgSetOracleResponse{}, nil
}

func (k msgServer) SetAsset(c context.Context, msg *types.MsgSetAsset) (*types.MsgSetAssetResponse, error) {
	if msg == nil {
		return nil, status.Error(codes.InvalidArgument, "empty msg")
	}

	ctx := sdk.UnwrapSDKContext(c)
	if err := k.Keeper.SetAsset(ctx, *msg); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(dnTypes.NewModuleNameEvent(types.ModuleName))
	ctx.EventManager().EmitEvent(types.NewAssetAddedEvent(msg.Asset))

	return &types.MsgSetAssetResponse{}, nil
}

func (k msgServer) PostPrice(c context.Context, msg *types.MsgPostPrice) (*types.MsgPostPriceResponse, error) {
	if msg == nil {
		return nil, status.Error(codes.InvalidArgument, "empty msg")
	}

	ctx := sdk.UnwrapSDKContext(c)
	if err := k.Keeper.PostPrice(ctx, *msg); err != nil {
		return nil, err
	}

	return &types.MsgPostPriceResponse{}, nil
}
