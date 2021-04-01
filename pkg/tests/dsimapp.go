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
	"github.com/dfinance/dstation/cmd/dstation/config"
	"github.com/dfinance/dstation/pkg/mock"
	vmConfig "github.com/dfinance/dstation/x/vm/config"
)

// DSimAppOption defines functional arguments for DSimApp constructor.
type DSimAppOption func(app *DSimApp)

// DSimApp wraps DnApp, provides VM environment and helper functions.
type DSimApp struct {
	DnApp *dnApp.DnApp
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
func (app *DSimApp) GetContext() sdk.Context {
	return app.DnApp.NewContext(false, tmProto.Header{})
}

// TearDown stops VM, servers and closes all connections.
func (app *DSimApp) TearDown() {
	app.DnApp.CloseConnections()
	if app.MockVMServer != nil {
		app.MockVMServer.Stop()
	}
}

// BeginBlock starts a new block.
func (app *DSimApp) BeginBlock() {
	app.curBlockHeight++
	app.curBlockTime = app.curBlockTime.Add(5 * time.Second)

	app.DnApp.BeginBlock(
		abci.RequestBeginBlock{
			Header: tmProto.Header{
				Height: app.curBlockHeight,
				Time:   app.curBlockTime,
			},
		},
	)
}

// EndBlock ends the current block.
func (app *DSimApp) EndBlock() {
	app.DnApp.EndBlock(abci.RequestEndBlock{})
	app.DnApp.Commit()
}

// SetupDSimApp creates a new DSimApp, setups VM environment.
func SetupDSimApp(opts ...DSimAppOption) *DSimApp {
	encCfg := dnApp.MakeEncodingConfig()

	// Create simApp with defaults
	app := &DSimApp{
		appOptions:     NewAppOptionsProvider(),
		vmConfig:       vmConfig.DefaultVMConfig(),
		curBlockHeight: 0,
		curBlockTime:   time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
	}

	genState, err := config.SetGenesisDefaults(encCfg.Marshaler, dnApp.NewDefaultGenesisState())
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
		config.DefaultNodeHome,
		1,
		dnApp.MakeEncodingConfig(),
		app.vmConfig,
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

	return app
}
