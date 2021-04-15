package keeper

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
	dnTypes "github.com/dfinance/dstation/pkg/types"
	"github.com/dfinance/dstation/x/oracle/types"
)

// PostPrice validates types.PostPrice, converts it to types.RawPrice and stores.
func (k Keeper) PostPrice(ctx sdk.Context, msg types.MsgPostPrice) error {
	// Input checks
	oracle := k.GetOracle(ctx, msg.OracleAddress)
	if oracle == nil {
		return sdkErrors.Wrapf(types.ErrInvalidPostPrice, "oracle (%s) not registered", msg.OracleAddress)
	}

	asset := k.GetAsset(ctx, msg.AssetCode)
	if asset == nil {
		return sdkErrors.Wrapf(types.ErrInvalidPostPrice, "asset (%s) not registered", msg.AssetCode)
	}

	found := false
	for _, assetSrc := range asset.Oracles {
		if assetSrc == msg.OracleAddress {
			found = true
			break
		}
	}
	if !found {
		return sdkErrors.Wrapf(types.ErrInvalidPostPrice, "asset (%s) has no source Oracle (%s)", msg.AssetCode, msg.OracleAddress)
	}

	if err := k.checkPostPriceReceivedAt(ctx, msg.ReceivedAt); err != nil {
		return sdkErrors.Wrapf(types.ErrInvalidPostPrice, "receivedAt validation: %v", err)
	}

	oracleBits := oracle.PriceMaxBytes*8
	if priceBits := uint32(msg.AskPrice.BigInt().BitLen()); priceBits > oracleBits {
		return sdkErrors.Wrapf(types.ErrInvalidPostPrice, "askPrice: invalid bit length %d (%d expected for %s oracle)", priceBits, oracleBits, oracle.AccAddress)
	}
	if priceBits := uint32(msg.BidPrice.BigInt().BitLen()); priceBits > oracleBits {
		return sdkErrors.Wrapf(types.ErrInvalidPostPrice, "bidPrice: invalid bit length %d (%d expected for %s oracle)", priceBits, oracleBits, oracle.AccAddress)
	}

	// Normalize prices, build RawPrice and store (overwrite if Oracle has already posted at this block)
	rawPrice := types.RawPrice{
		AskPrice:   asset.NormalizePriceValue(msg.AskPrice, oracle.PriceDecimals),
		BidPrice:   asset.NormalizePriceValue(msg.BidPrice, oracle.PriceDecimals),
		ReceivedAt: msg.ReceivedAt,
	}
	rawPriceBz := k.cdc.MustMarshalBinaryBare(&rawPrice)

	store := ctx.KVStore(k.storeKey)
	priceStore := prefix.NewStore(store, types.RawPricesPrefix)
	key := types.GetRawPricesKey(msg.AssetCode, ctx.BlockHeight(), msg.OracleAddress)
	priceStore.Set(key, rawPriceBz)

	return nil
}

// GetRawPrices returns all registered types.RawPrice for assetCode and blockHeight.
func (k Keeper) GetRawPrices(ctx sdk.Context, assetCode dnTypes.AssetCode, blockHeight int64) []types.RawPrice {
	store := ctx.KVStore(k.storeKey)
	priceStore := prefix.NewStore(store, types.RawPricesPrefix)

	keyPrefix := types.GetRawPricesKeyPrefix(assetCode, blockHeight)
	iterator := priceStore.Iterator(keyPrefix, sdk.PrefixEndBytes(keyPrefix))
	defer iterator.Close()

	prices := make([]types.RawPrice, 0)
	for ; iterator.Valid(); iterator.Next() {
		var price types.RawPrice
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &price)
		prices = append(prices, price)
	}

	return prices
}

// checkPostPriceReceivedAt checks PostPrice's ReceivedAt timestamp (algorithm depends on module params).
func (k Keeper) checkPostPriceReceivedAt(ctx sdk.Context, receivedAt time.Time) error {
	cfg := k.GetParams(ctx).PostPrice

	if cfg.ReceivedAtDiffInS > 0 {
		thresholdDur := time.Duration(cfg.ReceivedAtDiffInS) * time.Second

		absDuration := func(dur time.Duration) time.Duration {
			if dur < 0 {
				return -dur
			}
			return dur
		}

		blockTime := ctx.BlockTime()
		diffDur := blockTime.Sub(receivedAt)
		if absDuration(diffDur) > thresholdDur {
			return fmt.Errorf("timestamp difference %v should be LT %v", diffDur, thresholdDur)
		}
	}

	return nil
}
