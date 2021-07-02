package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dfinance/dstation/x/namespace/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Querier{}

// Querier implements gRPC query RPCs.
type Querier struct {
	Keeper
}


func (k Querier) DomainsAccount(c context.Context, request *types.DomainsAccountRequest) (*types.DomainsAccountResponse, error) {
	if request == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	return &types.DomainsAccountResponse{
		Whois: k.Keeper.GetWhoisByAccount(ctx, request.Address),
	}, nil
}
