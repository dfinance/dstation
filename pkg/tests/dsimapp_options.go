package tests

import (
	"encoding/json"

	dnApp "github.com/dfinance/dstation/app"
	"github.com/dfinance/dstation/pkg/mock"
	vmConfig "github.com/dfinance/dstation/x/vm/config"
)

// DSimAppOption defines functional arguments for DSimApp constructor.
type DSimAppOption func(app *DSimApp)

// DSimAppGenesisSetter defines setter callback for WithCustomGenesisState DSimApp option.
type DSimAppGenesisSetter func(oldState json.RawMessage) (newState json.RawMessage)

// WithMockVM adds the mock VMServer to DSimApp environment.
func WithMockVM() DSimAppOption {
	return func(app *DSimApp) {
		app.MockVMServer = mock.NewVMServer()

		app.appOptions.Set(dnApp.FlagCustomVMConnection, app.MockVMServer.GetVMClientConnection())
		app.appOptions.Set(dnApp.FlagCustomDSListener, app.MockVMServer.GetDSListener())
	}
}

// WithDVMConfig sets DVM config.
func WithDVMConfig(cfg vmConfig.VMConfig) DSimAppOption {
	return func(app *DSimApp) {
		app.vmConfig = cfg
	}
}

// WithCustomGenesisState sets custom GenesisState for module.
func WithCustomGenesisState(moduleName string, setter DSimAppGenesisSetter) DSimAppOption {
	return func(app *DSimApp) {
		app.genesisState[moduleName] = setter(app.genesisState[moduleName])
	}
}
