package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	dvmTypes "github.com/dfinance/dstation/pkg/types/dvm"
	"github.com/dfinance/dstation/x/vm/types"
)

// HasValue implements types.VMStorage interface.
func (k Keeper) HasValue(ctx sdk.Context, accessPath *dvmTypes.VMAccessPath) bool {
	store := ctx.KVStore(k.storeKey)
	vmStore := prefix.NewStore(store, types.VMKeyPrefix)
	key := types.GetVMStorageKey(accessPath)

	return vmStore.Has(key)
}

// GetValue implements types.VMStorage interface.
func (k Keeper) GetValue(ctx sdk.Context, accessPath *dvmTypes.VMAccessPath) []byte {
	store := ctx.KVStore(k.storeKey)
	vmStore := prefix.NewStore(store, types.VMKeyPrefix)
	key := types.GetVMStorageKey(accessPath)

	return vmStore.Get(key)
}

// SetValue implements types.VMStorage interface.
func (k Keeper) SetValue(ctx sdk.Context, accessPath *dvmTypes.VMAccessPath, value []byte) {
	store := ctx.KVStore(k.storeKey)
	vmStore := prefix.NewStore(store, types.VMKeyPrefix)
	key := types.GetVMStorageKey(accessPath)

	vmStore.Set(key, value)
}

// DelValue implements types.VMStorage interface.
func (k Keeper) DelValue(ctx sdk.Context, accessPath *dvmTypes.VMAccessPath) {
	store := ctx.KVStore(k.storeKey)
	vmStore := prefix.NewStore(store, types.VMKeyPrefix)
	key := types.GetVMStorageKey(accessPath)

	vmStore.Delete(key)
}

// iterateVMStorageValues iterates over all VMStorage values and processes them with handler (stop when handler returns true).
func (k Keeper) iterateVMStorageValues(ctx sdk.Context, handler func(accessPath *dvmTypes.VMAccessPath, value []byte) bool) {
	store := ctx.KVStore(k.storeKey)
	vmStore := prefix.NewStore(store, types.VMKeyPrefix)

	iterator := vmStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		accessPath := types.MustParseVMStorageKey(iterator.Key())
		value := iterator.Value()

		if handler(accessPath, value) {
			break
		}
	}
}
