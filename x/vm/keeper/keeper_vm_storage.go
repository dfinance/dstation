package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/dfinance/dstation/pkg/types/dvm"
	"github.com/dfinance/dstation/x/vm/types"
)

// HasValue checks if VMStorage has a writeSet data by dvm.VMAccessPath.
func (k Keeper) HasValue(ctx sdk.Context, accessPath *dvm.VMAccessPath) bool {
	store := ctx.KVStore(k.storeKey)
	key := types.GetVMStorageKey(accessPath)

	return store.Has(key)
}

// GetValue returns a VMStorage writeSet data by dvm.VMAccessPath.
func (k Keeper) GetValue(ctx sdk.Context, accessPath *dvm.VMAccessPath) []byte {
	store := ctx.KVStore(k.storeKey)
	key := types.GetVMStorageKey(accessPath)

	return store.Get(key)
}

// SetValue sets the VMStorage writeSet data by dvm.VMAccessPath.
func (k Keeper) SetValue(ctx sdk.Context, accessPath *dvm.VMAccessPath, value []byte) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetVMStorageKey(accessPath)

	store.Set(key, value)
}

// DelValue removes the VMStorage writeSet data by dvm.VMAccessPath.
func (k Keeper) DelValue(ctx sdk.Context, accessPath *dvm.VMAccessPath) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetVMStorageKey(accessPath)

	store.Delete(key)
}

// iterateVMStorageValues iterates over all VMStorage values and processes them with handler (stop when handler returns false).
func (k Keeper) iterateVMStorageValues(ctx sdk.Context, handler func(accessPath *dvm.VMAccessPath, value []byte) bool) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.GetVMStorageKeyPrefix())
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		accessPath := types.MustParseVMStorageKey(iterator.Key())
		value := iterator.Value()

		if !handler(accessPath, value) {
			break
		}
	}
}
