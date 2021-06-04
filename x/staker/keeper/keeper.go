package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/dfinance/dstation/x/staker/types"
	"github.com/tendermint/tendermint/libs/log"
)

// Keeper is a module keeper object.
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        codec.BinaryMarshaler
	paramSpace paramTypes.Subspace
	// Dependency keepers
	bankKeeper types.BankKeeper
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
	bankKeeper types.BankKeeper,
) Keeper {

	// Set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		paramSpace: paramSpace,
		bankKeeper: bankKeeper,
	}
}
