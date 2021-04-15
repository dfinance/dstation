package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	dnTypes "github.com/dfinance/dstation/pkg/types"
	"github.com/dfinance/dstation/x/oracle/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Querier{}

// Querier implements gRPC query RPCs.
type Querier struct {
	Keeper
}

func (k Querier) Oracles(c context.Context, req *types.QueryOraclesRequest) (*types.QueryOraclesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	oracles := k.GetOracles(ctx)

	return &types.QueryOraclesResponse{Oracles: oracles}, nil
}

func (k Querier) Assets(c context.Context, req *types.QueryAssetsRequest) (*types.QueryAssetsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	assets := k.GetAssets(ctx)

	return &types.QueryAssetsResponse{Assets: assets}, nil
}

func (k Querier) CurrentPrice(c context.Context, req *types.QueryCurrentPriceRequest) (*types.QueryCurrentPriceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	assetCode, err := dnTypes.NewAssetCodeByDenoms(req.LeftDenom, req.RightDenom)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "assetCode: invalid: %v", err)
	}
	price := k.GetCurrentPrice(ctx, assetCode)

	return &types.QueryCurrentPriceResponse{Price: price}, nil
}

func (k Querier) CurrentPrices(c context.Context, req *types.QueryCurrentPricesRequest) (*types.QueryCurrentPricesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	prices := k.GetCurrentPrices(ctx)

	return &types.QueryCurrentPricesResponse{Prices: prices}, nil
}
