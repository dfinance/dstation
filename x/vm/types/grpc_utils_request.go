package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	dvmTypes "github.com/dfinance/dstation/pkg/types/dvm"
)

// NewVMPublishModuleRequest builds a new dvmTypes.VMPublishModule VM request.
func NewVMPublishModuleRequests(ctx sdk.Context, signerAddrRaw string, code []byte) *dvmTypes.VMPublishModule {
	return &dvmTypes.VMPublishModule{
		Sender:       MustBech32ToLibra(signerAddrRaw),
		MaxGasAmount: getVMLimitedGas(ctx),
		GasUnitPrice: VmGasPrice,
		Code:         code,
	}
}

// NewVMExecuteScriptRequest builds a new dvmTypes.VMExecuteScript VM request.
func NewVMExecuteScriptRequest(ctx sdk.Context, signerAddrRaw string, code []byte, args ...MsgExecuteScript_ScriptArg) *dvmTypes.VMExecuteScript {
	vmArgs := make([]*dvmTypes.VMArgs, 0, len(args))
	for _, arg := range args {
		vmArgs = append(vmArgs, &dvmTypes.VMArgs{
			Type:  arg.Type,
			Value: arg.Value,
		})
	}

	return &dvmTypes.VMExecuteScript{
		Senders:      [][]byte{MustBech32ToLibra(signerAddrRaw)},
		MaxGasAmount: getVMLimitedGas(ctx),
		GasUnitPrice: VmGasPrice,
		Block:        uint64(ctx.BlockHeight()),
		Timestamp:    uint64(ctx.BlockTime().Unix()),
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
