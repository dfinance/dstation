package keeper

import (
	"fmt"
	"net"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
	"google.golang.org/grpc"

	dvmTypes "github.com/dfinance/dstation/pkg/types/dvm"
	"github.com/dfinance/dstation/x/vm/config"
	"github.com/dfinance/dstation/x/vm/types"
)

var _ types.VMStorage = (*Keeper)(nil)
var _ types.CurrencyInfoResProvider = (*Keeper)(nil)
var _ types.AccountBalanceResProvider = (*Keeper)(nil)

// Keeper is a module keeper object.
type Keeper struct {
	storeKey sdk.StoreKey
	cdc      codec.BinaryMarshaler
	config   *config.VMConfig
	// Dependency keepers
	accKeeper  types.AccountKeeper
	bankKeeper types.BankKeeper
	// VM connection
	vmClient VMClient
	vmConn   *grpc.ClientConn
	// DataSource server
	dsServer   *DSServer
	dsListener net.Listener
	//
	cache *keeperCache
}

// keeperCache keeps Keeper cache (Keeper is created by value, so cache should be defined as a pointer)
type keeperCache struct {
	currencyInfo map[string]dvmTypes.CurrencyInfo // key: denom
}

// Logger returns logger with keeper context.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// StartDSServer starts DataSource server.
func (k Keeper) StartDSServer() {
	k.dsServer.Start(k.dsListener)
}

// StopDSServer stops DataSource server.
func (k Keeper) StopDSServer() {
	k.dsServer.Stop()
}

// SetDSContext sets DataSource server context (storage context should be updated periodically to provide actual data).
func (k Keeper) SetDSContext(ctx sdk.Context) {
	k.dsServer.SetContext(ctx.WithGasMeter(types.NewDumbGasMeter()))
}

// NewKeeper create a new Keeper.
func NewKeeper(
	cdc codec.BinaryMarshaler, storeKey sdk.StoreKey,
	vmConn *grpc.ClientConn, dsListener net.Listener,
	config *config.VMConfig,
	accKeeper types.AccountKeeper, bankKeeper types.BankKeeper,
) Keeper {

	k := Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		accKeeper:  accKeeper,
		bankKeeper: bankKeeper,
		vmConn:     vmConn,
		vmClient:   NewVMClient(vmConn),
		dsListener: dsListener,
		config:     config,
		cache: &keeperCache{
			currencyInfo: make(map[string]dvmTypes.CurrencyInfo),
		},
	}

	k.dsServer = NewDSServer(&k, &k, &k)

	return k
}
