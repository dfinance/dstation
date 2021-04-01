package app

import (
	"github.com/cosmos/cosmos-sdk/baseapp"

	vmTypes "github.com/dfinance/dstation/x/vm/types"
)

const (
	FlagCustomVMConnection = "custom-vm-connection"
	FlagCustomDSListener   = "custom-ds-listener"
)

func VMCrashHandleBaseAppOption() func(*baseapp.BaseApp) {
	return func(app *baseapp.BaseApp) {
		app.AddRunTxRecoveryHandler(func(recoveryObj interface{}) error {
			if err, ok := recoveryObj.(error); ok {
				if vmTypes.ErrVMCrashed.Is(err) {
					panic(recoveryObj)
				}
			}

			return nil
		})
	}
}
