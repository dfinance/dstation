package app

import (
	"fmt"
	"io"
	"math/big"
	"net"
	"net/http"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/dfinance/dstation/x/gravity"
	gravityTypes "github.com/dfinance/dstation/x/gravity/types"
	"github.com/dfinance/dstation/x/oracle"
	oracleTypes "github.com/dfinance/dstation/x/oracle/types"
	"github.com/gorilla/mux"
	"github.com/rakyll/statik/fs"
	"google.golang.org/grpc"

	abci "github.com/tendermint/tendermint/abci/types"
	tmJson "github.com/tendermint/tendermint/libs/json"
	"github.com/tendermint/tendermint/libs/log"
	tmOs "github.com/tendermint/tendermint/libs/os"
	tmProto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmDb "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server/api"
	"github.com/cosmos/cosmos-sdk/server/config"
	serverTypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authRest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	authKeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authSims "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authTx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankKeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	capabilityKeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilityTypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisisKeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisisTypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrClient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	distrKeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrTypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidenceKeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidenceTypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutilTypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govKeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govTypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	transfer "github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer"
	ibcTransferKeeper "github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer/keeper"
	ibcTransferTypes "github.com/cosmos/cosmos-sdk/x/ibc/applications/transfer/types"
	ibc "github.com/cosmos/cosmos-sdk/x/ibc/core"
	ibcClient "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client"
	ibcPortTypes "github.com/cosmos/cosmos-sdk/x/ibc/core/05-port/types"
	ibcHost "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
	ibcKeeper "github.com/cosmos/cosmos-sdk/x/ibc/core/keeper"
	"github.com/cosmos/cosmos-sdk/x/mint"
	mintKeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	mintTypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsClient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramsKeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramsTypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramsProposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingKeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingTypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingKeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeClient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	upgradeKeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradeTypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	"github.com/dfinance/dstation/pkg"
	gravityKeeper "github.com/dfinance/dstation/x/gravity/keeper"
	oracleKeeper "github.com/dfinance/dstation/x/oracle/keeper"
	"github.com/dfinance/dstation/x/vm"
	vmConfig "github.com/dfinance/dstation/x/vm/config"
	vmKeeper "github.com/dfinance/dstation/x/vm/keeper"
	vmTypes "github.com/dfinance/dstation/x/vm/types"
	// unnamed import of statik for swagger UI support
	_ "github.com/cosmos/cosmos-sdk/client/docs/statik"
)

const appName = "DfinanceNode"

var (
	// ModuleBasics defines the module BasicManager is in charge of setting up basic,
	// non-dependant module elements, such as codec registration
	// and genesis verification.
	ModuleBasics = module.NewBasicManager(
		auth.AppModuleBasic{},
		genutil.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distr.AppModuleBasic{},
		gov.NewAppModuleBasic(
			paramsClient.ProposalHandler, distrClient.ProposalHandler, upgradeClient.ProposalHandler, upgradeClient.CancelProposalHandler,
		),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		ibc.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		evidence.AppModuleBasic{},
		transfer.AppModuleBasic{},
		vesting.AppModuleBasic{},
		// DN modules
		oracle.AppModuleBasic{},
		vm.AppModuleBasic{},
		gravity.AppModuleBasic{},
	)

	// module account permissions
	maccPerms = map[string][]string{
		authTypes.FeeCollectorName:     nil,
		distrTypes.ModuleName:          nil,
		mintTypes.ModuleName:           {authTypes.Minter},
		stakingTypes.BondedPoolName:    {authTypes.Burner, authTypes.Staking},
		stakingTypes.NotBondedPoolName: {authTypes.Burner, authTypes.Staking},
		govTypes.ModuleName:            {authTypes.Burner},
		ibcTransferTypes.ModuleName:    {authTypes.Minter, authTypes.Burner},
		vmTypes.DelPoolName:            {authTypes.Staking},
		gravityTypes.ModuleName:        {authTypes.Minter, authTypes.Burner},
	}
)

var (
	_ simapp.App              = (*DnApp)(nil)
	_ serverTypes.Application = (*DnApp)(nil)
)

// DnApp extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type DnApp struct { // nolint: golint
	*baseapp.BaseApp
	legacyAmino       *codec.LegacyAmino
	appCodec          codec.Marshaler
	interfaceRegistry types.InterfaceRegistry

	invCheckPeriod uint

	// keys to access the substores
	keys    map[string]*sdk.KVStoreKey
	tkeys   map[string]*sdk.TransientStoreKey
	memKeys map[string]*sdk.MemoryStoreKey

	// keepers
	AccountKeeper    authKeeper.AccountKeeper
	BankKeeper       bankKeeper.Keeper
	CapabilityKeeper *capabilityKeeper.Keeper
	StakingKeeper    stakingKeeper.Keeper
	SlashingKeeper   slashingKeeper.Keeper
	MintKeeper       mintKeeper.Keeper
	DistrKeeper      distrKeeper.Keeper
	GovKeeper        govKeeper.Keeper
	CrisisKeeper     crisisKeeper.Keeper
	UpgradeKeeper    upgradeKeeper.Keeper
	ParamsKeeper     paramsKeeper.Keeper
	IBCKeeper        *ibcKeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	EvidenceKeeper   evidenceKeeper.Keeper
	TransferKeeper   ibcTransferKeeper.Keeper
	// DN keepers
	OracleKeeper  oracleKeeper.Keeper
	VmKeeper      vmKeeper.Keeper
	GravityKeeper gravityKeeper.Keeper

	// make scoped keepers public for test purposes
	ScopedIBCKeeper      capabilityKeeper.ScopedKeeper
	ScopedTransferKeeper capabilityKeeper.ScopedKeeper

	// the module manager
	mm *module.Manager

	// simulation manager
	sm *module.SimulationManager

	// vm connections
	vmConfig   *vmConfig.VMConfig
	vmConn     *grpc.ClientConn
	dsListener net.Listener
}

// Name returns the name of the App.
func (app *DnApp) Name() string { return app.BaseApp.Name() }

// BeginBlocker application updates every begin block.
func (app *DnApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

// EndBlocker application updates every end block.
func (app *DnApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

// InitChainer application update at chain initialization.
func (app *DnApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var genesisState GenesisState
	if err := tmJson.Unmarshal(req.AppStateBytes, &genesisState); err != nil {
		panic(err)
	}
	return app.mm.InitGenesis(ctx, app.appCodec, genesisState)
}

// LoadHeight loads a particular height.
func (app *DnApp) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (app *DnApp) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[authTypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// LegacyAmino returns DnApp's amino codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *DnApp) LegacyAmino() *codec.LegacyAmino {
	return app.legacyAmino
}

// AppCodec returns Gaia's app codec.
//
// NOTE: This is solely to be used for testing purposes as it may be desirable
// for modules to register their own custom testing types.
func (app *DnApp) AppCodec() codec.Marshaler {
	return app.appCodec
}

// InterfaceRegistry returns Gaia's InterfaceRegistry.
func (app *DnApp) InterfaceRegistry() types.InterfaceRegistry {
	return app.interfaceRegistry
}

// GetKey returns the KVStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *DnApp) GetKey(storeKey string) *sdk.KVStoreKey {
	return app.keys[storeKey]
}

// GetTKey returns the TransientStoreKey for the provided store key.
//
// NOTE: This is solely to be used for testing purposes.
func (app *DnApp) GetTKey(storeKey string) *sdk.TransientStoreKey {
	return app.tkeys[storeKey]
}

// GetMemKey returns the MemStoreKey for the provided mem key.
//
// NOTE: This is solely used for testing purposes.
func (app *DnApp) GetMemKey(storeKey string) *sdk.MemoryStoreKey {
	return app.memKeys[storeKey]
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (app *DnApp) GetSubspace(moduleName string) paramsTypes.Subspace {
	subspace, _ := app.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

// SimulationManager implements the SimulationApp interface.
func (app *DnApp) SimulationManager() *module.SimulationManager {
	return app.sm
}

// RegisterAPIRoutes registers all application module routes with the provided API server.
func (app *DnApp) RegisterAPIRoutes(apiSvr *api.Server, apiConfig config.APIConfig) {
	clientCtx := apiSvr.ClientCtx
	rpc.RegisterRoutes(clientCtx, apiSvr.Router)
	// Register legacy tx routes.
	authRest.RegisterTxRoutes(clientCtx, apiSvr.Router)
	// Register new tx routes from grpc-gateway.
	authTx.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)
	// Register new tendermint queries routes from grpc-gateway.
	tmservice.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// Register legacy and grpc-gateway routes for all modules.
	ModuleBasics.RegisterRESTRoutes(clientCtx, apiSvr.Router)
	ModuleBasics.RegisterGRPCGatewayRoutes(clientCtx, apiSvr.GRPCGatewayRouter)

	// register swagger API from root so that other applications can override easily
	if apiConfig.Swagger {
		RegisterSwaggerAPI(apiSvr.Router)
	}
}

// RegisterTxService implements the Application.RegisterTxService method.
func (app *DnApp) RegisterTxService(clientCtx client.Context) {
	authTx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.BaseApp.Simulate, app.interfaceRegistry)
}

// RegisterTendermintService implements the Application.RegisterTendermintService method.
func (app *DnApp) RegisterTendermintService(clientCtx client.Context) {
	tmservice.RegisterTendermintService(app.BaseApp.GRPCQueryRouter(), clientCtx, app.interfaceRegistry)
}

// Initialize connection to VM server.
func (app *DnApp) InitializeVMConnection(addr string, appOpts serverTypes.AppOptions) {
	// Custom (used for mock connection)
	if obj := appOpts.Get(FlagCustomVMConnection); obj != nil {
		conn, ok := obj.(*grpc.ClientConn)
		if !ok {
			panic(fmt.Errorf("%s appOpt: type assertion failed: %T", FlagCustomVMConnection, obj))
		}
		app.vmConn = conn
		return
	}

	// gRPC connection
	app.Logger().Info(fmt.Sprintf("Creating VM connection, address: %s", addr))
	conn, err := pkg.GetGRpcClientConnection(addr, 1*time.Second)
	if err != nil {
		panic(err)
	}
	app.vmConn = conn

	app.Logger().Info(fmt.Sprintf("Non-blocking connection initialized, status: %s", app.vmConn.GetState()))
}

// Initialize listener to listen for connections from VM for data server.
func (app *DnApp) InitializeVMDataServer(addr string, appOpts serverTypes.AppOptions) {
	// Custom (used for mock connection)
	if obj := appOpts.Get(FlagCustomDSListener); obj != nil {
		listener, ok := obj.(net.Listener)
		if !ok {
			panic(fmt.Errorf("%s appOpt: type assertion failed: %T", FlagCustomDSListener, obj))
		}
		app.dsListener = listener
		return
	}

	app.Logger().Info(fmt.Sprintf("Starting VM data server listener, address: %s", addr))
	listener, err := pkg.GetGRpcNetListener(addr)
	if err != nil {
		panic(err)
	}
	app.dsListener = listener

	app.Logger().Info("VM data server is running")
}

// CloseConnections closes VM connection and stops DS server.
func (app DnApp) CloseConnections() {
	app.VmKeeper.StopDSServer()
	if app.dsListener != nil {
		app.dsListener.Close()
	}
	if app.vmConn != nil {
		app.vmConn.Close()
	}
}

// NewDnApp returns a reference to an initialized Dnode service.
func NewDnApp(
	logger log.Logger, db tmDb.DB, traceStore io.Writer, loadLatest bool, skipUpgradeHeights map[int64]bool,
	homePath string, invCheckPeriod uint, encodingConfig EncodingConfig, vmConfig *vmConfig.VMConfig, appOpts serverTypes.AppOptions, baseAppOptions ...func(*baseapp.BaseApp),
) *DnApp {

	appCodec := encodingConfig.Marshaler
	legacyAmino := encodingConfig.Amino
	interfaceRegistry := encodingConfig.InterfaceRegistry

	bApp := baseapp.NewBaseApp(appName, logger, db, encodingConfig.TxConfig.TxDecoder(), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetAppVersion(version.Version)
	bApp.SetInterfaceRegistry(interfaceRegistry)

	keys := sdk.NewKVStoreKeys(
		authTypes.StoreKey, bankTypes.StoreKey, stakingTypes.StoreKey,
		mintTypes.StoreKey, distrTypes.StoreKey, slashingTypes.StoreKey,
		govTypes.StoreKey, paramsTypes.StoreKey, ibcHost.StoreKey, upgradeTypes.StoreKey,
		evidenceTypes.StoreKey, ibcTransferTypes.StoreKey, capabilityTypes.StoreKey,
		oracleTypes.StoreKey, vmTypes.StoreKey,
		gravityTypes.StoreKey,
	)
	tkeys := sdk.NewTransientStoreKeys(paramsTypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilityTypes.MemStoreKey)

	app := &DnApp{
		BaseApp:           bApp,
		legacyAmino:       legacyAmino,
		appCodec:          appCodec,
		interfaceRegistry: interfaceRegistry,
		invCheckPeriod:    invCheckPeriod,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
		vmConfig:          vmConfig,
	}

	// Initialize VM connections
	app.InitializeVMDataServer(vmConfig.DataListen, appOpts)
	app.InitializeVMConnection(vmConfig.Address, appOpts)

	// Reduce ConsensusPower reduction coefficient (1 xfi == 1 power unit)
	// 1 xfi == 1000000000000000000 (exp == 18)
	sdk.PowerReduction = sdk.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))

	app.ParamsKeeper = initParamsKeeper(appCodec, legacyAmino, keys[paramsTypes.StoreKey], tkeys[paramsTypes.TStoreKey])
	// set the BaseApp's parameter store
	bApp.SetParamStore(app.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramsKeeper.ConsensusParamsKeyTable()))

	// add capability keeper and ScopeToModule for ibc module
	app.CapabilityKeeper = capabilityKeeper.NewKeeper(appCodec, keys[capabilityTypes.StoreKey], memKeys[capabilityTypes.MemStoreKey])
	scopedIBCKeeper := app.CapabilityKeeper.ScopeToModule(ibcHost.ModuleName)
	scopedTransferKeeper := app.CapabilityKeeper.ScopeToModule(ibcTransferTypes.ModuleName)

	// Add keepers
	app.AccountKeeper = authKeeper.NewAccountKeeper(
		appCodec, keys[authTypes.StoreKey], app.GetSubspace(authTypes.ModuleName), authTypes.ProtoBaseAccount, maccPerms,
	)
	app.BankKeeper = bankKeeper.NewBaseKeeper(
		appCodec, keys[bankTypes.StoreKey], app.AccountKeeper, app.GetSubspace(bankTypes.ModuleName), app.ModuleAccountAddrs(),
	)
	stakingKeeper := stakingKeeper.NewKeeper(
		appCodec, keys[stakingTypes.StoreKey], app.AccountKeeper, app.BankKeeper, app.GetSubspace(stakingTypes.ModuleName),
	)
	app.MintKeeper = mintKeeper.NewKeeper(
		appCodec, keys[mintTypes.StoreKey], app.GetSubspace(mintTypes.ModuleName), &stakingKeeper,
		app.AccountKeeper, app.BankKeeper, authTypes.FeeCollectorName,
	)
	app.DistrKeeper = distrKeeper.NewKeeper(
		appCodec, keys[distrTypes.StoreKey], app.GetSubspace(distrTypes.ModuleName), app.AccountKeeper, app.BankKeeper,
		&stakingKeeper, authTypes.FeeCollectorName, app.ModuleAccountAddrs(),
	)
	app.SlashingKeeper = slashingKeeper.NewKeeper(
		appCodec, keys[slashingTypes.StoreKey], &stakingKeeper, app.GetSubspace(slashingTypes.ModuleName),
	)
	app.CrisisKeeper = crisisKeeper.NewKeeper(
		app.GetSubspace(crisisTypes.ModuleName), invCheckPeriod, app.BankKeeper, authTypes.FeeCollectorName,
	)
	app.UpgradeKeeper = upgradeKeeper.NewKeeper(skipUpgradeHeights, keys[upgradeTypes.StoreKey], appCodec, homePath)

	// Register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	app.StakingKeeper = *stakingKeeper.SetHooks(
		stakingTypes.NewMultiStakingHooks(
			app.DistrKeeper.Hooks(),
			app.SlashingKeeper.Hooks(),
			app.GravityKeeper.Hooks(),
		),
	)

	// Create IBC Keeper
	app.IBCKeeper = ibcKeeper.NewKeeper(
		appCodec, keys[ibcHost.StoreKey], app.GetSubspace(ibcHost.ModuleName), app.StakingKeeper, scopedIBCKeeper,
	)

	// DN keepers
	app.OracleKeeper = oracleKeeper.NewKeeper(
		appCodec, keys[oracleTypes.StoreKey], app.GetSubspace(oracleTypes.ModuleName),
	)

	app.VmKeeper = vmKeeper.NewKeeper(
		appCodec, keys[vmTypes.StoreKey],
		app.vmConn, app.dsListener, app.vmConfig,
		app.AccountKeeper, app.BankKeeper, app.OracleKeeper,
	)

	app.GravityKeeper = gravityKeeper.NewKeeper(
		appCodec,
		keys[gravityTypes.StoreKey],
		app.GetSubspace(gravityTypes.ModuleName),
		stakingKeeper,
		app.BankKeeper,
		app.SlashingKeeper,
	)

	// Register the proposal types
	govRouter := govTypes.NewRouter()
	govRouter.AddRoute(govTypes.RouterKey, govTypes.ProposalHandler).
		AddRoute(paramsProposal.RouterKey, params.NewParamChangeProposalHandler(app.ParamsKeeper)).
		AddRoute(distrTypes.RouterKey, distr.NewCommunityPoolSpendProposalHandler(app.DistrKeeper)).
		AddRoute(upgradeTypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(app.UpgradeKeeper)).
		AddRoute(ibcHost.RouterKey, ibcClient.NewClientUpdateProposalHandler(app.IBCKeeper.ClientKeeper)).
		AddRoute(vmTypes.RouterKey, vm.NewGovHandler(app.VmKeeper))
	app.GovKeeper = govKeeper.NewKeeper(
		appCodec, keys[govTypes.StoreKey], app.GetSubspace(govTypes.ModuleName), app.AccountKeeper, app.BankKeeper,
		&stakingKeeper, govRouter,
	)

	// Create Transfer Keepers
	app.TransferKeeper = ibcTransferKeeper.NewKeeper(
		appCodec, keys[ibcTransferTypes.StoreKey], app.GetSubspace(ibcTransferTypes.ModuleName),
		app.IBCKeeper.ChannelKeeper, &app.IBCKeeper.PortKeeper,
		app.AccountKeeper, app.BankKeeper, scopedTransferKeeper,
	)
	transferModule := transfer.NewAppModule(app.TransferKeeper)

	// Create static IBC router, add transfer route, then set and seal it
	ibcRouter := ibcPortTypes.NewRouter()
	ibcRouter.AddRoute(ibcTransferTypes.ModuleName, transferModule)
	app.IBCKeeper.SetRouter(ibcRouter)

	// create evidence keeper with router
	evidenceKeeper := evidenceKeeper.NewKeeper(
		appCodec, keys[evidenceTypes.StoreKey], &app.StakingKeeper, app.SlashingKeeper,
	)
	// If evidence needs to be handled for the app, set routes in router here and seal
	app.EvidenceKeeper = *evidenceKeeper
	/****  Module Options ****/

	/****  Module Options ****/
	var skipGenesisInvariants = false
	opt := appOpts.Get(crisis.FlagSkipGenesisInvariants)
	if opt, ok := opt.(bool); ok {
		skipGenesisInvariants = opt
	}

	// NOTE: Any module instantiated in the module manager that is later modified
	// must be passed by reference here.
	app.mm = module.NewManager(
		genutil.NewAppModule(
			app.AccountKeeper, app.StakingKeeper, app.BaseApp.DeliverTx,
			encodingConfig.TxConfig,
		),
		auth.NewAppModule(appCodec, app.AccountKeeper, nil),
		vesting.NewAppModule(app.AccountKeeper, app.BankKeeper),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper),
		crisis.NewAppModule(&app.CrisisKeeper, skipGenesisInvariants),
		gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper),
		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper),
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		upgrade.NewAppModule(app.UpgradeKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		ibc.NewAppModule(app.IBCKeeper),
		params.NewAppModule(app.ParamsKeeper),
		transferModule,
		// DN modules
		oracle.NewAppModule(appCodec, app.OracleKeeper),
		vm.NewAppModule(appCodec, app.VmKeeper, app.AccountKeeper, app.BankKeeper),
		// Gravity bridge
		gravity.NewAppModule(app.GravityKeeper, app.BankKeeper),
	)

	// During begin block slashing happens after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool, so as to keep the
	// CanWithdrawInvariant invariant.
	// NOTE: staking module is required if HistoricalEntries param > 0
	app.mm.SetOrderBeginBlockers(
		upgradeTypes.ModuleName,
		mintTypes.ModuleName,
		distrTypes.ModuleName,
		slashingTypes.ModuleName,
		evidenceTypes.ModuleName,
		stakingTypes.ModuleName,
		ibcHost.ModuleName,
		// DN modules
		oracleTypes.ModuleName,
		vmTypes.ModuleName,
	)
	app.mm.SetOrderEndBlockers(
		crisisTypes.ModuleName,
		govTypes.ModuleName,
		stakingTypes.ModuleName,
		// DN modules
		oracleTypes.ModuleName,
		// Gravity bridge
		gravityTypes.ModuleName,
	)

	// NOTE: The genutils module must occur after staking so that pools are
	// properly initialized with tokens from genesis accounts.
	// NOTE: Capability module must occur first so that it can initialize any capabilities
	// so that other modules that want to create or claim capabilities afterwards in InitChain
	// can do so safely.
	app.mm.SetOrderInitGenesis(
		capabilityTypes.ModuleName,
		authTypes.ModuleName,
		bankTypes.ModuleName,
		distrTypes.ModuleName,
		stakingTypes.ModuleName,
		slashingTypes.ModuleName,
		govTypes.ModuleName,
		mintTypes.ModuleName,
		crisisTypes.ModuleName,
		ibcHost.ModuleName,
		genutilTypes.ModuleName,
		evidenceTypes.ModuleName,
		ibcTransferTypes.ModuleName,
		// DN modules
		oracleTypes.ModuleName,
		vmTypes.ModuleName,
		// Gravity bridge
		gravityTypes.ModuleName,
	)

	app.mm.RegisterInvariants(&app.CrisisKeeper)
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter(), encodingConfig.Amino)
	app.mm.RegisterServices(module.NewConfigurator(app.MsgServiceRouter(), app.GRPCQueryRouter()))

	// create the simulation manager and define the order of the modules for deterministic simulations
	//
	// NOTE: this is not required apps that don't use the simulator for fuzz testing
	// transactions
	app.sm = module.NewSimulationManager(
		auth.NewAppModule(appCodec, app.AccountKeeper, authSims.RandomGenesisAccounts),
		bank.NewAppModule(appCodec, app.BankKeeper, app.AccountKeeper),
		capability.NewAppModule(appCodec, *app.CapabilityKeeper),
		gov.NewAppModule(appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper),
		mint.NewAppModule(appCodec, app.MintKeeper, app.AccountKeeper),
		staking.NewAppModule(appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		slashing.NewAppModule(appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		params.NewAppModule(app.ParamsKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		ibc.NewAppModule(app.IBCKeeper),
		transferModule,
	)

	app.sm.RegisterStoreDecoders()

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	// initialize BaseApp
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetAnteHandler(
		ante.NewAnteHandler(
			app.AccountKeeper, app.BankKeeper, ante.DefaultSigVerificationGasConsumer,
			encodingConfig.TxConfig.SignModeHandler(),
		),
	)
	app.SetEndBlocker(app.EndBlocker)

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmOs.Exit(err.Error())
		}

		// Initialize and seal the capability keeper so all persistent capabilities
		// are loaded in-memory and prevent any further modules from creating scoped
		// sub-keepers.
		// This must be done during creation of baseapp rather than in InitChain so
		// that in-memory capabilities get regenerated on app restart.
		// Note that since this reads from the store, we can only perform it when
		// `loadLatest` is set to true.
		ctx := app.BaseApp.NewUncachedContext(true, tmProto.Header{})
		app.CapabilityKeeper.InitializeAndSeal(ctx)
	}
	app.ScopedIBCKeeper = scopedIBCKeeper
	app.ScopedTransferKeeper = scopedTransferKeeper

	// Start the VM data source server after all the initialization is done
	app.VmKeeper.StartDSServer()

	return app
}

// RegisterSwaggerAPI registers swagger route with API Server.
func RegisterSwaggerAPI(rtr *mux.Router) {
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}

	staticServer := http.FileServer(statikFS)
	rtr.PathPrefix("/swagger/").Handler(http.StripPrefix("/swagger/", staticServer))
}

// GetMaccPerms returns a copy of the module account permissions.
func GetMaccPerms() map[string][]string {
	dupMaccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		dupMaccPerms[k] = v
	}
	return dupMaccPerms
}

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(appCodec codec.BinaryMarshaler, legacyAmino *codec.LegacyAmino, key, tkey sdk.StoreKey) paramsKeeper.Keeper {
	paramsKeeper := paramsKeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	paramsKeeper.Subspace(authTypes.ModuleName)
	paramsKeeper.Subspace(bankTypes.ModuleName)
	paramsKeeper.Subspace(stakingTypes.ModuleName)
	paramsKeeper.Subspace(mintTypes.ModuleName)
	paramsKeeper.Subspace(distrTypes.ModuleName)
	paramsKeeper.Subspace(slashingTypes.ModuleName)
	paramsKeeper.Subspace(govTypes.ModuleName).WithKeyTable(govTypes.ParamKeyTable())
	paramsKeeper.Subspace(crisisTypes.ModuleName)
	paramsKeeper.Subspace(ibcTransferTypes.ModuleName)
	paramsKeeper.Subspace(ibcHost.ModuleName)
	//
	paramsKeeper.Subspace(oracleTypes.ModuleName)
	//
	paramsKeeper.Subspace(gravityTypes.ModuleName)

	return paramsKeeper
}
