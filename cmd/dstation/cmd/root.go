package cmd

import (
	"io"
	"os"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/debug"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	serverTypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/snapshots"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authClient "github.com/cosmos/cosmos-sdk/x/auth/client"
	authCmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	genutilCli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	tmCli "github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"
	tmDb "github.com/tendermint/tm-db"

	"github.com/dfinance/dstation/app"
	"github.com/dfinance/dstation/cmd/dstation/config"
	vmConfig "github.com/dfinance/dstation/x/vm/config"
)

// NewRootCmd creates a new root command for simd. It is called once in the main function.
func NewRootCmd() (*cobra.Command, app.EncodingConfig) {
	sdkConfig := sdk.GetConfig()
	config.SetConfigBech32Prefixes(sdkConfig)
	sdkConfig.Seal()

	encodingConfig := app.MakeEncodingConfig()
	authClient.Codec = encodingConfig.Marshaler

	initClientCtx := client.Context{}.
		WithJSONMarshaler(encodingConfig.Marshaler).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithInput(os.Stdin).
		WithAccountRetriever(types.AccountRetriever{}).
		WithBroadcastMode(flags.BroadcastBlock).
		WithHomeDir(config.DefaultNodeHome)

	rootCmd := &cobra.Command{
		Use:   "dstation",
		Short: "Dfinance Cosmos SDK Stargate app",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			if err := client.SetCmdClientContextHandler(initClientCtx, cmd); err != nil {
				return err
			}

			return server.InterceptConfigsPreRunHandler(cmd)
		},
	}

	rootCmd.AddCommand(
		genutilCli.InitCmd(app.ModuleBasics, config.DefaultNodeHome),
		genutilCli.CollectGenTxsCmd(bankTypes.GenesisBalancesIterator{}, config.DefaultNodeHome),
		genutilCli.GenTxCmd(app.ModuleBasics, encodingConfig.TxConfig, bankTypes.GenesisBalancesIterator{}, config.DefaultNodeHome),
		genutilCli.ValidateGenesisCmd(app.ModuleBasics),

		SetGenesisDefaultsCmd(config.DefaultNodeHome),
		AddGenesisAccountCmd(config.DefaultNodeHome),

		rpc.StatusCommand(),
		keys.Commands(config.DefaultNodeHome),
		debug.Cmd(),
		tmCli.NewCompletionCmd(rootCmd, true),

		queryCommand(),
		txCommand(),
	)
	server.AddCommands(rootCmd, config.DefaultNodeHome, newApp, appExporter, serverFlags)

	return rootCmd, encodingConfig
}

// appExporter returns an serverTypes.AppExporter.
func appExporter(
	logger log.Logger, db tmDb.DB, traceStore io.Writer, height int64, forZeroHeight bool, jailAllowedAddrs []string, appOpts serverTypes.AppOptions,
) (serverTypes.ExportedApp, error) {

	vmConfig := vmConfig.ReadVMConfig(viper.GetString(flags.FlagHome))

	encCfg := app.MakeEncodingConfig() // Ideally, we would reuse the one created by NewRootCmd.
	encCfg.Marshaler = codec.NewProtoCodec(encCfg.InterfaceRegistry)
	var dnApp *app.DnApp
	if height != -1 {
		dnApp = app.NewDnApp(logger, db, traceStore, false, map[int64]bool{}, "", cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod)), encCfg, vmConfig, appOpts)

		if err := dnApp.LoadHeight(height); err != nil {
			return serverTypes.ExportedApp{}, err
		}
	} else {
		dnApp = app.NewDnApp(logger, db, traceStore, true, map[int64]bool{}, "", cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod)), encCfg, vmConfig, appOpts)
	}

	return dnApp.ExportAppStateAndValidators(forZeroHeight, jailAllowedAddrs)
}

// serverFlags returns flags for start command.
func serverFlags(startCmd *cobra.Command) {
	crisis.AddModuleInitFlags(startCmd)
}

// queryCommand returns query command.
func queryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "query",
		Aliases:                    []string{"q"},
		Short:                      "Querying subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authCmd.GetAccountCmd(),
		rpc.ValidatorCommand(),
		rpc.BlockCommand(),
		authCmd.QueryTxsByEventsCmd(),
		authCmd.QueryTxCmd(),
	)

	app.ModuleBasics.AddQueryCommands(cmd)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

// txCommand returns tx command.
func txCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "tx",
		Short:                      "Transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authCmd.GetSignCommand(),
		authCmd.GetSignBatchCommand(),
		authCmd.GetMultiSignCommand(),
		authCmd.GetMultiSignBatchCmd(),
		authCmd.GetValidateSignaturesCommand(),
		flags.LineBreak,
		authCmd.GetBroadcastCommand(),
		authCmd.GetEncodeCommand(),
		authCmd.GetDecodeCommand(),
	)

	app.ModuleBasics.AddTxCommands(cmd)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

// newApp returns an AppCreator.
func newApp(logger log.Logger, db tmDb.DB, traceStore io.Writer, appOpts serverTypes.AppOptions) serverTypes.Application {
	var cache sdk.MultiStorePersistentCache

	vmConfig := vmConfig.ReadVMConfig(viper.GetString(flags.FlagHome))

	if cast.ToBool(appOpts.Get(server.FlagInterBlockCache)) {
		cache = store.NewCommitKVStoreCacheManager()
	}

	skipUpgradeHeights := make(map[int64]bool)
	for _, h := range cast.ToIntSlice(appOpts.Get(server.FlagUnsafeSkipUpgrades)) {
		skipUpgradeHeights[int64(h)] = true
	}

	pruningOpts, err := server.GetPruningOptionsFromFlags(appOpts)
	if err != nil {
		panic(err)
	}

	snapshotDir := filepath.Join(cast.ToString(appOpts.Get(flags.FlagHome)), "data", "snapshots")
	snapshotDB, err := sdk.NewLevelDB("metadata", snapshotDir)
	if err != nil {
		panic(err)
	}
	snapshotStore, err := snapshots.NewStore(snapshotDB, snapshotDir)
	if err != nil {
		panic(err)
	}

	return app.NewDnApp(
		logger, db, traceStore, true, skipUpgradeHeights,
		cast.ToString(appOpts.Get(flags.FlagHome)),
		cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod)),
		app.MakeEncodingConfig(), // Ideally, we would reuse the one created by NewRootCmd.
		vmConfig,
		appOpts,
		baseapp.SetPruning(pruningOpts),
		baseapp.SetMinGasPrices(cast.ToString(appOpts.Get(server.FlagMinGasPrices))),
		baseapp.SetHaltHeight(cast.ToUint64(appOpts.Get(server.FlagHaltHeight))),
		baseapp.SetHaltTime(cast.ToUint64(appOpts.Get(server.FlagHaltTime))),
		baseapp.SetMinRetainBlocks(cast.ToUint64(appOpts.Get(server.FlagMinRetainBlocks))),
		baseapp.SetInterBlockCache(cache),
		baseapp.SetTrace(cast.ToBool(appOpts.Get(server.FlagTrace))),
		baseapp.SetIndexEvents(cast.ToStringSlice(appOpts.Get(server.FlagIndexEvents))),
		baseapp.SetSnapshotStore(snapshotStore),
		baseapp.SetSnapshotInterval(cast.ToUint64(appOpts.Get(server.FlagStateSyncSnapshotInterval))),
		baseapp.SetSnapshotKeepRecent(cast.ToUint32(appOpts.Get(server.FlagStateSyncSnapshotKeepRecent))),
		app.VMCrashHandleBaseAppOption(),
	)
}
