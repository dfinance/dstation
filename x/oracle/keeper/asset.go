package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
	dnTypes "github.com/dfinance/dstation/pkg/types"
	"github.com/dfinance/dstation/x/oracle/types"
)

// GetAsset returns an types.Asset if exists.
func (k Keeper) GetAsset(ctx sdk.Context, assetCode dnTypes.AssetCode) *types.Asset {
	if cachedValue, found := k.cache.assets[assetCode.String()]; found {
		return &cachedValue
	}

	store := ctx.KVStore(k.storeKey)
	assetStore := prefix.NewStore(store, types.AssetsPrefix)
	key, _ := assetCode.Marshal()

	bz := assetStore.Get(key)
	if bz == nil {
		return nil
	}

	asset := &types.Asset{}
	k.cdc.MustUnmarshalBinaryBare(bz, asset)

	k.cache.assets[asset.AssetCode.String()] = *asset

	return asset
}

// GetAssets returns all registered types.Asset objects.
func (k Keeper) GetAssets(ctx sdk.Context) []types.Asset {
	store := ctx.KVStore(k.storeKey)
	assetStore := prefix.NewStore(store, types.AssetsPrefix)

	iterator := assetStore.Iterator(nil, nil)
	defer iterator.Close()

	assets := make([]types.Asset, 0)
	for ; iterator.Valid(); iterator.Next() {
		asset := types.Asset{}
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &asset)
		assets = append(assets, asset)
	}

	return assets
}

// SetAsset wraps setAsset with checking if requester is a registered nominee.
func (k Keeper) SetAsset(ctx sdk.Context, msg types.MsgSetAsset) error {
	if err := k.IsNominee(ctx, msg.Nominee); err != nil {
		return err
	}

	for _, oracle := range msg.Asset.Oracles {
		if !k.hasOracle(ctx, oracle) {
			return sdkErrors.Wrapf(types.ErrOracleNotFound, "oracle (%s)", oracle)
		}
	}

	k.setAsset(ctx, msg.Asset)

	return nil
}

// setAsset sets an types.Asset.
func (k Keeper) setAsset(ctx sdk.Context, asset types.Asset) {
	store := ctx.KVStore(k.storeKey)
	assetStore := prefix.NewStore(store, types.AssetsPrefix)
	key, _ := asset.AssetCode.Marshal()

	bz := k.cdc.MustMarshalBinaryBare(&asset)
	assetStore.Set(key, bz)

	k.cache.assets[asset.AssetCode.String()] = asset
}
