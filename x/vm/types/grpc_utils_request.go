package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dfinance/dvm-proto/go/vm_grpc"
)

// NewVMPublishModuleRequest builds a new vm_grpc.VMPublishModule VM request.
func NewVMPublishModuleRequests(ctx sdk.Context, signerAddrRaw string, code []byte) *vm_grpc.VMPublishModule {
	return &vm_grpc.VMPublishModule{
		Sender:       MustBech32ToLibra(signerAddrRaw),
		MaxGasAmount: getVMLimitedGas(ctx),
		GasUnitPrice: VmGasPrice,
		Code:         code,
	}
}

// NewVMExecuteScriptRequest builds a new vm_grpc.VMExecuteScript VM request.
func NewVMExecuteScriptRequest(ctx sdk.Context, signerAddrRaw string, code []byte, args ...MsgExecuteScript_ScriptArg) *vm_grpc.VMExecuteScript {
	vmArgs := make([]*vm_grpc.VMArgs, 0, len(args))
	for _, arg := range args {
		vmArgs = append(vmArgs, &vm_grpc.VMArgs{
			Type:  arg.Type,
			Value: arg.Value,
		})
	}

	return &vm_grpc.VMExecuteScript{
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
