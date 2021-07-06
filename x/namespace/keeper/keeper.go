package keeper

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/store/prefix"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dfinance/dstation/x/namespace/types"
	"github.com/tendermint/tendermint/libs/log"
)

// Keeper is a module keeper object.
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        codec.BinaryMarshaler
	// Dependency keepers
	bankKeeper types.BankKeeper
}

// Logger returns logger with keeper context.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// NewKeeper create a new Keeper.
func NewKeeper(
	cdc codec.BinaryMarshaler, storeKey sdk.StoreKey,
	bankKeeper types.BankKeeper,
) Keeper {
	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		bankKeeper: bankKeeper,
	}
}


func (k Keeper) Buy(ctx sdk.Context, value, address string, amount sdk.Coins) (types.Whois, error) {
	whois, err := k.newWhois(ctx, value, address, amount)
	if err != nil {
		return types.Whois{}, err
	}

	targetAccAddr, _ := sdk.AccAddressFromBech32(whois.Creator)

	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, targetAccAddr, types.ModuleName, whois.Price); err != nil {
		return types.Whois{}, fmt.Errorf("sending coins (%s) from account (%s) to module: %w", whois.Price, targetAccAddr.String(), err)
	}
	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, whois.Price); err != nil {
		return types.Whois{}, fmt.Errorf("burning coins (%s) from module: %w", whois.Price, err)
	}

	k.setWhois(ctx, whois)

	return whois, nil
}


// setWhois sets a types.Whois and call.uniqueId <-> call.Id match.
func (k Keeper) setWhois(ctx sdk.Context, whois types.Whois) {
	whoisStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.WhoisPrefix)

	keyWhoisId, _ := whois.ID.Marshal()
	whoisBz := k.cdc.MustMarshalBinaryBare(&whois)

	whoisStore.Set(keyWhoisId, whoisBz)

	k.setLastWhoisID(ctx , whois.ID)
}

func (k Keeper) Delete(ctx sdk.Context, value, adress string) error {
	wh, e := k.getWhoisByValue(ctx, value)
	if e != nil {
		return e
	}

	if wh.Creator != adress {
		return types.ErrDomainNotHandlesByAddress
	}

	k.dropWhois(ctx, wh)

	return nil
}


func (k Keeper) GetAllWhois(ctx sdk.Context) (res  []types.Whois) {
	k.IterateAllCalls(ctx, func(whois types.Whois) (stop bool) {
		res = append(res, whois)

		return false
	})

	return res
}

func (k Keeper) GetWhoisByAccount(ctx sdk.Context, address string) (res  []types.Whois) {
	k.IterateAllCalls(ctx, func(whois types.Whois) (stop bool) {
		if whois.Creator == address {
			res = append(res, whois)
		}

		return false
	})

	return res
}



// dropWhois sets a types.Whois and call.uniqueId <-> call.Id match.
func (k Keeper) dropWhois(ctx sdk.Context, whois types.Whois) {
	whoisStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.WhoisPrefix)

	keyWhoisId, _ := whois.ID.Marshal()

	whoisStore.Delete(keyWhoisId)
}

// newWhois creates and saves a new Whois.
func (k Keeper) newWhois(ctx sdk.Context, value, address string, amount sdk.Coins) (types.Whois, error) {
	wh := types.Whois{
		ID:         k.getNextCallID(ctx),
		Creator:   address,
		Value: value,
	}

	if err := wh.Validate(); err != nil {
		return types.Whois{}, fmt.Errorf("call validation: %w", err)
	}

	return wh, nil
}

func (k Keeper) getWhoisByValue(ctx sdk.Context, value string) (res types.Whois,e error) {
	k.IterateAllCalls(ctx, func(whois types.Whois) (stop bool) {
		if whois.Value == value {
			res = whois

			return true
		}

		return false
	})

	if res.Value == "" {
		return res, types.ErrNameDoesNotExist
	}

	return res, nil
}

// IterateAllWhois iterates through all stored types.Whois entries.
func (k Keeper) IterateAllCalls(ctx sdk.Context, handler func(call types.Whois) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	whoisStore := prefix.NewStore(store, types.WhoisPrefix)

	iterator := whoisStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		call := types.Whois{}
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &call)
		if handler(call) {
			break
		}
	}
}

// getLastWhoisID returns the latest stored unique types.Call ID.
func (k Keeper) getLastWhoisID(ctx sdk.Context) *sdk.Uint {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.LastWhoisId)
	if bz == nil {
		return nil
	}

	id := sdk.Uint{}
	if err := id.Unmarshal(bz); err != nil {
		panic(fmt.Errorf("unmarshal dnTypes.ID (%v): %w", bz, err))
	}

	return &id
}

// setLastWhoisID sets the latest used unique types.Whois ID.
func (k Keeper) setLastWhoisID(ctx sdk.Context, id sdk.Uint) {
	store := ctx.KVStore(k.storeKey)
	bz, _ := id.Marshal()
	store.Set(types.LastWhoisId, bz)
}

// getNextCallID returns the next unique types.Call ID (0 if not exists).
func (k Keeper) getNextCallID(ctx sdk.Context) sdk.Uint {
	id := k.getLastWhoisID(ctx)
	if id == nil {
		return sdk.ZeroUint()
	}

	return id.Incr()
}