package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dfinance/dstation/x/staker/types"
)

// GetCallById returns a types.Call by ID if exists.
func (k Keeper) GetCallById(ctx sdk.Context, id sdk.Uint) *types.Call {
	store := ctx.KVStore(k.storeKey)
	callStore := prefix.NewStore(store, types.CallsPrefix)
	keyCallId, _ := id.Marshal()

	callBz := callStore.Get(keyCallId)
	if callBz == nil {
		return nil
	}

	call := &types.Call{}
	k.cdc.MustUnmarshalBinaryBare(callBz, call)

	return call
}

// GetCallByUniqueId returns a types.Call by unique ID if match exists.
func (k Keeper) GetCallByUniqueId(ctx sdk.Context, uniqueId string) *types.Call {
	callId := k.getCallIdByUniqueId(ctx, uniqueId)
	if callId == nil {
		return nil
	}

	return k.GetCallById(ctx, *callId)
}

// GetAllCalls returns all stored types.Call entries.
func (k Keeper) GetAllCalls(ctx sdk.Context) (retList []types.Call) {
	k.IterateAllCalls(ctx, func(call types.Call) (stop bool) {
		retList = append(retList, call)
		return false
	})

	return
}

// GetAddressCalls returns all stored types.Call entries for target account address.
func (k Keeper) GetAddressCalls(ctx sdk.Context, accAddr sdk.Address) (retList []types.Call) {
	accAddrStr := accAddr.String()
	k.IterateAllCalls(ctx, func(call types.Call) (stop bool) {
		if call.Address != accAddrStr {
			return false
		}
		retList = append(retList, call)

		return false
	})

	return
}

// IterateAllCalls iterates through all stored types.Call entries.
func (k Keeper) IterateAllCalls(ctx sdk.Context, handler func(call types.Call) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	callStore := prefix.NewStore(store, types.CallsPrefix)

	iterator := callStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		call := types.Call{}
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &call)
		if handler(call) {
			break
		}
	}
}

// setCall sets a types.Call and call.uniqueId <-> call.Id match.
func (k Keeper) setCall(ctx sdk.Context, call types.Call) {
	store := ctx.KVStore(k.storeKey)
	callStore := prefix.NewStore(store, types.CallsPrefix)
	uniqueIdStore := prefix.NewStore(store, types.UniqueIdsPrefix)

	keyCallId, _ := call.Id.Marshal()
	keyCallUniqueId := []byte(call.UniqueId)
	callBz := k.cdc.MustMarshalBinaryBare(&call)
	callIdBz, _ := call.Id.Marshal()

	callStore.Set(keyCallId, callBz)
	uniqueIdStore.Set(keyCallUniqueId, callIdBz)
}

// getCallIdByUniqueId returns types.Call ID by its uniqueID if found.
func (k Keeper) getCallIdByUniqueId(ctx sdk.Context, uniqueId string) *sdk.Uint {
	store := ctx.KVStore(k.storeKey)
	uniqueIdStore := prefix.NewStore(store, types.UniqueIdsPrefix)

	keyUniqueId := []byte(uniqueId)
	callIdBz := uniqueIdStore.Get(keyUniqueId)
	if callIdBz == nil {
		return nil
	}

	var callId sdk.Uint
	callId.Unmarshal(callIdBz)

	return &callId
}

// getLastCallID returns the latest stored unique types.Call ID.
func (k Keeper) getLastCallID(ctx sdk.Context) *sdk.Uint {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.LastCallId)
	if bz == nil {
		return nil
	}

	id := sdk.Uint{}
	if err := id.Unmarshal(bz); err != nil {
		panic(fmt.Errorf("unmarshal dnTypes.ID (%v): %w", bz, err))
	}

	return &id
}

// setLastCallID sets the latest used unique types.Call ID.
func (k Keeper) setLastCallID(ctx sdk.Context, id sdk.Uint) {
	store := ctx.KVStore(k.storeKey)
	bz, _ := id.Marshal()
	store.Set(types.LastCallId, bz)
}

// getNextCallID returns the next unique types.Call ID (0 if not exists).
func (k Keeper) getNextCallID(ctx sdk.Context) sdk.Uint {
	id := k.getLastCallID(ctx)
	if id == nil {
		return sdk.ZeroUint()
	}

	return id.Incr()
}
