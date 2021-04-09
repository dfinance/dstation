package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/dfinance/dstation/x/vm/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (k msgServer) DeployModule(c context.Context, msg *types.MsgDeployModule) (*types.MsgDeployModuleResponse, error) {
	if msg == nil {
		return nil, status.Error(codes.InvalidArgument, "empty msg")
	}

	ctx := sdk.UnwrapSDKContext(c)
	if err := k.DeployContract(ctx, *msg); err != nil {
		return nil, err
	}

	return &types.MsgDeployModuleResponse{}, nil
}

func (k msgServer) ExecuteScript(c context.Context, msg *types.MsgExecuteScript) (*types.MsgExecuteScriptResponse, error) {
	if msg == nil {
		return nil, status.Error(codes.InvalidArgument, "empty msg")
	}

	ctx := sdk.UnwrapSDKContext(c)
	if err := k.ExecuteContract(ctx, *msg); err != nil {
		return nil, err
	}

	return &types.MsgExecuteScriptResponse{}, nil
}
