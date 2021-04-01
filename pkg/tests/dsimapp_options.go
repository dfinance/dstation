package tests

import (
	dnApp "github.com/dfinance/dstation/app"
	"github.com/dfinance/dstation/pkg/mock"
)

// WithMockVM adds the mock VMServer to DSimApp environment.
func WithMockVM() DSimAppOption {
	return func(app *DSimApp) {
		app.MockVMServer = mock.NewVMServer()

		app.appOptions.Set(dnApp.FlagCustomVMConnection, app.MockVMServer.GetVMClientConnection())
		app.appOptions.Set(dnApp.FlagCustomDSListener, app.MockVMServer.GetDSListener())
	}
}

// WithCustomVMRetryParams alters VM config params.
func WithCustomVMRetryParams(maxAttempts, reqTimeoutInMs uint) DSimAppOption {
	return func(app *DSimApp) {
		app.vmConfig.MaxAttempts = maxAttempts
		app.vmConfig.ReqTimeoutInMs = reqTimeoutInMs
	}
}
