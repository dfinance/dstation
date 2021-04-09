package tests

import (
	"encoding/json"
	"time"

	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmProto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	dnApp "github.com/dfinance/dstation/app"
	dnConfig "github.com/dfinance/dstation/cmd/dstation/config"
	"github.com/dfinance/dstation/pkg/mock"
	vmConfig "github.com/dfinance/dstation/x/vm/config"
)

// DSimAppOption defines functional arguments for DSimApp constructor.
type DSimAppOption func(app *DSimApp)

// DSimApp wraps DnApp, provides VM environment and helper functions.
type DSimApp struct {
	DnApp  *dnApp.DnApp
	encCfg dnApp.EncodingConfig
	//
	MockVMServer *mock.VMServer
	//
	appOptions   AppOptionsProvider
	genesisState dnApp.GenesisState
	vmConfig     vmConfig.VMConfig
	//
	curBlockHeight int64
	curBlockTime   time.Time
}

// GetContext creates a new sdk.Context.
func (app *DSimApp) GetContext(checkTx bool) sdk.Context {
	return app.DnApp.NewUncachedContext(checkTx, tmProto.Header{
		Height: app.curBlockHeight,
		Time:   app.curBlockTime,
	})
}

// GetCurrentBlockHeightTime returns current block height and time (set at the BeginBlock).
func (app *DSimApp) GetCurrentBlockHeightTime() (int64, time.Time) {
	return app.curBlockHeight, app.curBlockTime
}

// GetNextBlockHeightTime returns next block height and time (as DeliverTx begins / ends a new block, that func peeks a new values).
func (app *DSimApp) GetNextBlockHeightTime() (int64, time.Time) {
	return app.curBlockHeight + 1, app.curBlockTime.Add(5 * time.Second)
}

// SetCustomVMRetryParams alters VM config params.
func (app *DSimApp) SetCustomVMRetryParams(maxAttempts, reqTimeoutInMs uint) {
	app.vmConfig.MaxAttempts = maxAttempts
	app.vmConfig.ReqTimeoutInMs = reqTimeoutInMs
}

// TearDown stops VM, servers and closes all connections.
func (app *DSimApp) TearDown() {
	app.DnApp.CloseConnections()
	if app.MockVMServer != nil {
		app.MockVMServer.Stop()
	}
}

// SetupDSimApp creates a new DSimApp, setups VM environment.
func SetupDSimApp(opts ...DSimAppOption) *DSimApp {
	// Create simApp with defaults
	app := &DSimApp{
		encCfg:         dnApp.MakeEncodingConfig(),
		appOptions:     NewAppOptionsProvider(),
		vmConfig:       vmConfig.DefaultVMConfig(),
		curBlockHeight: 0,
		curBlockTime:   time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	genState, err := dnConfig.SetGenesisDefaults(app.encCfg.Marshaler, dnApp.NewDefaultGenesisState())
	if err != nil {
		panic(err)
	}
	app.genesisState = genState

	// Apply simApp options
	for _, opt := range opts {
		opt(app)
	}

	// Init the main app
	app.DnApp = dnApp.NewDnApp(
		log.NewNopLogger(),
		dbm.NewMemDB(),
		nil,
		true,
		map[int64]bool{},
		dnConfig.DefaultNodeHome,
		1,
		app.encCfg,
		&app.vmConfig,
		app.appOptions,
		dnApp.VMCrashHandleBaseAppOption(),
	)

	// Init chain
	appStateBz, err := json.MarshalIndent(app.genesisState, "", " ")
	if err != nil {
		panic(err)
	}

	app.DnApp.InitChain(
		abci.RequestInitChain{
			Validators:      []abci.ValidatorUpdate{},
			ConsensusParams: simapp.DefaultConsensusParams,
			AppStateBytes:   appStateBz,
		},
	)

	// Skip the genesis block
	app.BeginBlock()
	app.EndBlock()

	return app
}
