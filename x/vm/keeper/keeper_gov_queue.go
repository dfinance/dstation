package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/dfinance/dstation/x/vm/types"
)

// ScheduleProposal checks if types.PlannedProposal can be added to the gov proposal queue and adds it.
func (k Keeper) ScheduleProposal(ctx sdk.Context, pProposal *types.PlannedProposal) error {
	if pProposal == nil {
		return fmt.Errorf("pProposal: nil")
	}

	if pProposal.Height <= ctx.BlockHeight() {
		return sdkErrors.Wrapf(
			sdkErrors.ErrInvalidRequest,
			"schedule failed: planned blockHeight (%d) should be LT current (%d)",
			pProposal.Height, ctx.BlockHeight())
	}

	k.addProposalToQueue(ctx, pProposal)

	return nil
}

// RemoveProposalFromQueue removes types.PlannedProposal from the gov proposal queue.
func (k Keeper) RemoveProposalFromQueue(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)
	queueStore := prefix.NewStore(store, types.GovQueuePrefix)
	key := types.GetGovQueueStorageKey(id)

	queueStore.Delete(key)
}

// IterateProposalsQueue iterates over gov proposal queue.
func (k Keeper) IterateProposalsQueue(ctx sdk.Context, handler func(id uint64, pProposal *types.PlannedProposal)) {
	store := ctx.KVStore(k.storeKey)
	queueStore := prefix.NewStore(store, types.GovQueuePrefix)

	iterator := queueStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		id := types.ParseGovQueueStorageKey(iterator.Key())

		p, err := k.unmarshalPlannedProposal(iterator.Value())
		if err != nil {
			panic(fmt.Errorf("unmarshalPlannedProposal: %w", err))
		}

		handler(id, p)
	}
}

// getNextProposalID returns next gov proposal queue ID.
func (k Keeper) getNextProposalID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(types.GovQueueLastIdKey) {
		return 0
	}

	bz := store.Get(types.GovQueueLastIdKey)
	id := types.ParseGovQueueStorageKey(bz)

	return id + 1
}

// setProposalID updates gov proposal queue last ID.
func (k Keeper) setProposalID(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(k.storeKey)

	bz := types.GetGovQueueStorageKey(id)
	store.Set(types.GovQueueLastIdKey, bz)
}

// addProposalToQueue adds types.PlannedProposal to the gov proposal queue.
func (k Keeper) addProposalToQueue(ctx sdk.Context, pProposal *types.PlannedProposal) {
	store := ctx.KVStore(k.storeKey)
	queueStore := prefix.NewStore(store, types.GovQueuePrefix)

	id := k.getNextProposalID(ctx)
	key := types.GetGovQueueStorageKey(id)
	bz, err := k.marshalPlannedProposal(pProposal)
	if err != nil {
		panic(fmt.Errorf("marshalPlannedProposal: %w", err))
	}

	queueStore.Set(key, bz)
	k.setProposalID(ctx, id)
}

// marshalPlannedProposal returns binary serialization of types.PlannedProposal.
func (k Keeper) marshalPlannedProposal(pProposal *types.PlannedProposal) ([]byte, error) {
	bz, err := k.cdc.MarshalBinaryLengthPrefixed(pProposal)
	if err != nil {
		return nil, err
	}

	return bz, nil
}

// unmarshalPlannedProposal return deserialized binary serialization of types.PlannedProposal.
func (k Keeper) unmarshalPlannedProposal(bz []byte) (*types.PlannedProposal, error) {
	pProposal := &types.PlannedProposal{}
	if err := k.cdc.UnmarshalBinaryLengthPrefixed(bz, pProposal); err != nil {
		return nil, err
	}

	return pProposal, nil
}
