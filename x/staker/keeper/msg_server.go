package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	dnTypes "github.com/dfinance/dstation/pkg/types"
	"github.com/dfinance/dstation/x/staker/types"
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

func (k msgServer) Deposit(c context.Context, msg *types.MsgDepositCall) (*types.MsgDepositCallResponse, error) {
	if msg == nil {
		return nil, status.Error(codes.InvalidArgument, "empty msg")
	}

	ctx := sdk.UnwrapSDKContext(c)
	call, err := k.Keeper.Deposit(ctx, *msg)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(dnTypes.NewModuleNameEvent(types.ModuleName))
	ctx.EventManager().EmitEvent(types.NewDepositEvent(call))

	return &types.MsgDepositCallResponse{
		Id: call.Id,
	}, nil
}

func (k msgServer) Withdraw(c context.Context, msg *types.MsgWithdrawCall) (*types.MsgWithdrawCallResponse, error) {
	if msg == nil {
		return nil, status.Error(codes.InvalidArgument, "empty msg")
	}

	ctx := sdk.UnwrapSDKContext(c)
	call, err := k.Keeper.Withdraw(ctx, *msg)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(dnTypes.NewModuleNameEvent(types.ModuleName))
	ctx.EventManager().EmitEvent(types.NewWithdrawEvent(call))

	return &types.MsgWithdrawCallResponse{
		Id: call.Id,
	}, nil
}
