package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/dfinance/dstation/x/oracle/types"
)

// Keeper is a module keeper object.
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        codec.BinaryMarshaler
	paramSpace paramTypes.Subspace
	//
	cache *keeperCache
}

// keeperCache optimized Gas usage for frequent PostPrice operations.
type keeperCache struct {
	oracles map[string]types.Oracle // key: AccAddress
	assets  map[string]types.Asset  // key: AssetCode
}

// IsNominee checks if nominee exist within keeper parameters and returns sdk wrapped error.
func (k Keeper) IsNominee(ctx sdk.Context, accAddr string) error {
	for _, nominee := range k.GetParams(ctx).Nominees {
		if nominee == accAddr {
			return nil
		}
	}

	return sdkErrors.Wrapf(types.ErrNotAuthorized, "address (%s)", accAddr)
}

// Logger returns logger with keeper context.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// NewKeeper create a new Keeper.
func NewKeeper(
	cdc codec.BinaryMarshaler, storeKey sdk.StoreKey, paramSpace paramTypes.Subspace,
) Keeper {

	// Set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		paramSpace: paramSpace,
		cache: &keeperCache{
			oracles: make(map[string]types.Oracle),
			assets:  make(map[string]types.Asset),
		},
	}
}
