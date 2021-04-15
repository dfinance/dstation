package keeper

import (
	"fmt"
	"sort"
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	dnTypes "github.com/dfinance/dstation/pkg/types"
	"github.com/dfinance/dstation/x/oracle/types"
)

// GetCurrentPrice returns an types.CurrentPrice for assetCode / reversed assetCode (if exists).
func (k Keeper) GetCurrentPrice(ctx sdk.Context, assetCode dnTypes.AssetCode) *types.CurrentPrice {
	if price := k.getCurrentPrice(ctx, assetCode); price != nil {
		return price
	}

	if priceDirect := k.getCurrentPrice(ctx, assetCode.ReverseCode()); priceDirect != nil {
		asset := k.GetAsset(ctx, priceDirect.AssetCode)
		if asset == nil {
			panic(fmt.Errorf("GetAsset(%s) for reversed assetCode (%s): not found", priceDirect.AssetCode, assetCode))
		}
		priceReversed := priceDirect.ReversedPrice(asset.Decimals)

		return &priceReversed
	}

	return nil
}

// GetCurrentPrices returns all types.CurrentPrice objects in the storage.
func (k Keeper) GetCurrentPrices(ctx sdk.Context) []types.CurrentPrice {
	store := ctx.KVStore(k.storeKey)
	priceStore := prefix.NewStore(store, types.CurPricesPrefix)

	iterator := priceStore.Iterator(nil, nil)
	defer iterator.Close()

	prices := make([]types.CurrentPrice, 0)
	for ; iterator.Valid(); iterator.Next() {
		price := types.CurrentPrice{}
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &price)
		prices = append(prices, price)
	}

	return prices
}

// UpdateCurrentPrices updates types.CurrentPrice values for each registered types.Asset.
func (k Keeper) UpdateCurrentPrices(ctx sdk.Context) {
	updatesCnt := 0
	for _, asset := range k.GetAssets(ctx) {
		medianAskPrice, medianBidPrice, medianReceivedAt := k.getMedianAssetPrices(ctx, asset.AssetCode)

		// Check if there is no rawPrices or medianPrice is invalid
		if medianAskPrice.IsZero() || medianBidPrice.IsZero() {
			continue
		}

		// Check if a new price appeared (no need to update it every block)
		oldPrice := k.GetCurrentPrice(ctx, asset.AssetCode)
		if oldPrice != nil && oldPrice.AskPrice.Equal(medianAskPrice) && oldPrice.BidPrice.Equal(medianBidPrice) {
			continue
		}

		// Set the new price and set the reversed price
		newPrice := types.CurrentPrice{
			AssetCode:  asset.AssetCode,
			AskPrice:   medianAskPrice,
			BidPrice:   medianBidPrice,
			ReceivedAt: medianReceivedAt,
		}
		k.setCurrentPrice(ctx, newPrice)

		// Emit event
		updatesCnt++
		ctx.EventManager().EmitEvent(types.NewPriceEvent(newPrice))
	}

	if updatesCnt > 0 {
		ctx.EventManager().EmitEvent(dnTypes.NewModuleNameEvent(types.ModuleName))
	}
}

// getMedianAssetPrices returns median Ask/Bid/ReceivedAt values for an asset based on current block types.RawPrice objects.
func (k Keeper) getMedianAssetPrices(ctx sdk.Context, assetCode dnTypes.AssetCode) (medianAskPrice, medianBidPrice sdk.Int, medianReceivedAt time.Time) {
	rawPrices := k.GetRawPrices(ctx, assetCode, ctx.BlockHeight())
	pricesCnt := len(rawPrices)
	if pricesCnt == 0 {
		medianAskPrice, medianBidPrice = sdk.ZeroInt(), sdk.ZeroInt()
		return
	}

	if pricesCnt == 1 {
		medianAskPrice, medianBidPrice, medianReceivedAt = rawPrices[0].AskPrice, rawPrices[0].BidPrice, rawPrices[0].ReceivedAt
		return
	}

	// Sort prices
	askPrices, bidPrices := make([]types.RawPrice, pricesCnt), make([]types.RawPrice, pricesCnt)
	copy(askPrices, rawPrices)
	copy(bidPrices, rawPrices)

	sort.Slice(askPrices, func(i, j int) bool {
		return askPrices[i].AskPrice.LT(askPrices[j].AskPrice)
	})
	sort.Slice(bidPrices, func(i, j int) bool {
		return bidPrices[i].BidPrice.LT(bidPrices[j].BidPrice)
	})
	sort.Slice(rawPrices, func(i, j int) bool {
		return rawPrices[i].ReceivedAt.Before(rawPrices[j].ReceivedAt)
	})

	// Odd number of prices case
	if pricesCnt%2 != 0 {
		medianAskPrice, medianBidPrice = askPrices[pricesCnt/2].AskPrice, bidPrices[pricesCnt/2].BidPrice
		medianReceivedAt = rawPrices[pricesCnt/2].ReceivedAt
		return
	}

	// Even number of prices case
	a1 := askPrices[pricesCnt/2-1].AskPrice
	a2 := askPrices[pricesCnt/2].AskPrice
	sumA := a1.Add(a2)

	b1 := bidPrices[pricesCnt/2-1].BidPrice
	b2 := bidPrices[pricesCnt/2].BidPrice
	sumB := b1.Add(b2)

	medianAskPrice, medianBidPrice = sumA.QuoRaw(2), sumB.QuoRaw(2) // since it's a price and not a balance, division with precision loss is OK
	medianReceivedAt = rawPrices[pricesCnt/2].ReceivedAt

	return
}

// getCurrentPrice returns an types.CurrentPrice if exists.
func (k Keeper) getCurrentPrice(ctx sdk.Context, assetCode dnTypes.AssetCode) *types.CurrentPrice {
	store := ctx.KVStore(k.storeKey)
	priceStore := prefix.NewStore(store, types.CurPricesPrefix)
	key, _ := assetCode.Marshal()

	bz := priceStore.Get(key)
	if bz == nil {
		return nil
	}

	price := &types.CurrentPrice{}
	k.cdc.MustUnmarshalBinaryBare(bz, price)

	return price
}

// setCurrentPrice sets a types.CurrentPrice.
func (k Keeper) setCurrentPrice(ctx sdk.Context, price types.CurrentPrice) {
	store := ctx.KVStore(k.storeKey)
	priceStore := prefix.NewStore(store, types.CurPricesPrefix)
	key, _ := price.AssetCode.Marshal()

	bz := k.cdc.MustMarshalBinaryBare(&price)
	priceStore.Set(key, bz)
}
