package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/dfinance/dstation/pkg/types/dvm"
)

// NewVMPublishModuleRequest builds a new dvm.VMPublishModule VM request.
func NewVMPublishModuleRequests(ctx sdk.Context, signerAddrRaw string, code []byte) *dvm.VMPublishModule {
	return &dvm.VMPublishModule{
		Sender:       MustBech32ToLibra(signerAddrRaw),
		MaxGasAmount: getVMLimitedGas(ctx),
		GasUnitPrice: VmGasPrice,
		Code:         code,
	}
}

// NewVMExecuteScriptRequest builds a new dvm.VMExecuteScript VM request.
func NewVMExecuteScriptRequest(ctx sdk.Context, signerAddrRaw string, code []byte, args ...MsgExecuteScript_ScriptArg) *dvm.VMExecuteScript {
	vmArgs := make([]*dvm.VMArgs, 0, len(args))
	for _, arg := range args {
		vmArgs = append(vmArgs, &dvm.VMArgs{
			Type:  arg.Type,
			Value: arg.Value,
		})
	}

	return &dvm.VMExecuteScript{
		Senders:      [][]byte{MustBech32ToLibra(signerAddrRaw)},
		MaxGasAmount: getVMLimitedGas(ctx),
		GasUnitPrice: VmGasPrice,
		Code:         code,
		TypeParams:   nil,
		Args:         vmArgs,
	}
}

// getVMLimitedGas returns gas limit which is LTE VM max limit.
func getVMLimitedGas(ctx sdk.Context) sdk.Gas {
	gas := ctx.GasMeter().Limit() - ctx.GasMeter().GasConsumed()
	if gas > VmGasLimit {
		return VmGasLimit
	}

	return gas
}
