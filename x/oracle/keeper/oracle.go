package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dfinance/dstation/x/oracle/types"
)

// GetOracle returns an types.Oracle if exists.
func (k Keeper) GetOracle(ctx sdk.Context, addr string) *types.Oracle {
	if cachedValue, found := k.cache.oracles[addr]; found {
		return &cachedValue
	}

	store := ctx.KVStore(k.storeKey)
	oracleStore := prefix.NewStore(store, types.OraclesPrefix)
	key := types.GetOraclesKey(addr)

	bz := oracleStore.Get(key)
	if bz == nil {
		return nil
	}

	oracle := &types.Oracle{}
	k.cdc.MustUnmarshalBinaryBare(bz, oracle)

	k.cache.oracles[addr] = *oracle

	return oracle
}

// GetOracles returns all registered types.Oracle objects.
func (k Keeper) GetOracles(ctx sdk.Context) []types.Oracle {
	store := ctx.KVStore(k.storeKey)
	oracleStore := prefix.NewStore(store, types.OraclesPrefix)

	iterator := oracleStore.Iterator(nil, nil)
	defer iterator.Close()

	oracles := make([]types.Oracle, 0)
	for ; iterator.Valid(); iterator.Next() {
		oracle := types.Oracle{}
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &oracle)
		oracles = append(oracles, oracle)
	}

	return oracles
}

// SetOracle wraps setOracle with checking if requester is a registered nominee.
func (k Keeper) SetOracle(ctx sdk.Context, msg types.MsgSetOracle) error {
	if err := k.IsNominee(ctx, msg.Nominee); err != nil {
		return err
	}

	k.setOracle(ctx, msg.Oracle)

	return nil
}

// hasOracle checks if types.Oracle with addr is registered.
func (k Keeper) hasOracle(ctx sdk.Context, addr string) bool {
	if _, found := k.cache.oracles[addr]; found {
		return found
	}

	store := ctx.KVStore(k.storeKey)
	oracleStore := prefix.NewStore(store, types.OraclesPrefix)
	key := types.GetOraclesKey(addr)

	return oracleStore.Has(key)
}

// setOracle sets an types.Oracle.
func (k Keeper) setOracle(ctx sdk.Context, oracle types.Oracle) {
	store := ctx.KVStore(k.storeKey)
	oracleStore := prefix.NewStore(store, types.OraclesPrefix)
	key := types.GetOraclesKey(oracle.AccAddress)

	bz := k.cdc.MustMarshalBinaryBare(&oracle)
	oracleStore.Set(key, bz)

	k.cache.oracles[oracle.AccAddress] = oracle
}
