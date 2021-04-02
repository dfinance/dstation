package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/dfinance/dvm-proto/go/vm_grpc"

	dnTypes "github.com/dfinance/dstation/pkg/types"
	"github.com/dfinance/dstation/x/vm/types"
)

// ExecuteContract executes Move script and processes execution results (events, writeSets).
func (k Keeper) ExecuteContract(ctx sdk.Context, msg types.MsgExecuteScript) error {
	req := types.NewVMExecuteScriptRequest(ctx, msg.Signer, msg.Script, msg.Args...)

	exec, err := k.sendExecuteReq(ctx, nil, req)
	if err != nil {
		sdkErr := sdkErrors.Wrapf(types.ErrVMCrashed, "grpc error: %v", err)

		k.Logger(ctx).Error(sdkErr.Error())
		panic(sdkErr)
	}

	k.processVMExecution(ctx, exec)

	return nil
}

// DeployContract deploys Move module (contract) and processes execution results (events, writeSets).
func (k Keeper) DeployContract(ctx sdk.Context, msg types.MsgDeployModule) error {
	execList := make([]*vm_grpc.VMExecuteResponse, 0, len(msg.Modules))
	for i, code := range msg.Modules {
		req := types.NewVMPublishModuleRequests(ctx, msg.Signer, code)

		exec, err := k.sendExecuteReq(ctx, req, nil)
		if err != nil {
			sdkErr := sdkErrors.Wrapf(types.ErrVMCrashed, "contract [%d]: grpc error: %v", i, err)

			k.Logger(ctx).Error(sdkErr.Error())
			panic(sdkErr)
		}
		execList = append(execList, exec)
	}

	for _, exec := range execList {
		k.processVMExecution(ctx, exec)
	}

	return nil
}

// DeployContractDryRun checks that contract can be deployed (returned writeSets are not persisted to store).
func (k Keeper) DeployContractDryRun(ctx sdk.Context, msg types.MsgDeployModule) error {
	for i, code := range msg.Modules {
		req := types.NewVMPublishModuleRequests(ctx, msg.Signer, code)

		exec, err := k.sendExecuteReq(ctx, req, nil)
		if err != nil {
			sdkErr := sdkErrors.Wrapf(types.ErrVMCrashed, "contract [%d]: grpc error: %v", i, err)
			return sdkErr
		}

		if exec.GetStatus().GetError() != nil {
			statusMsg := types.StringifyVMStatus(exec.Status)
			sdkErr := sdkErrors.Wrapf(types.ErrWrongExecutionResponse, "contract [%d]: %s", i, statusMsg)
			return sdkErr
		}
	}

	return nil
}

// processVMExecution processes VM execution result: emits events, converts VM events, updates writeSets.
func (k Keeper) processVMExecution(ctx sdk.Context, exec *vm_grpc.VMExecuteResponse) {
	// Consume gas (if execution took too much gas - panic and mark transaction as out of gas)
	ctx.GasMeter().ConsumeGas(exec.GasUsed, "vm script/module execution")

	// Emit execution events
	ctx.EventManager().EmitEvent(dnTypes.NewModuleNameEvent(types.ModuleName))
	ctx.EventManager().EmitEvents(types.NewContractEvents(exec))

	// Process success status
	if exec.GetStatus().GetError() == nil {
		k.processVMWriteSet(ctx, exec.WriteSet)

		// Emit VM events (panic on "out of gas", emitted events stays in the EventManager)
		for _, vmEvent := range exec.Events {
			ctx.EventManager().EmitEvent(types.NewMoveEvent(ctx.GasMeter(), vmEvent))
		}
	}
}

// processVMWriteSet processes VM execution writeSets (set/delete).
func (k Keeper) processVMWriteSet(ctx sdk.Context, writeSet []*vm_grpc.VMValue) {
	for _, value := range writeSet {
		if value.Type == vm_grpc.VmWriteOp_Deletion {
			k.DelValue(ctx, value.Path)
		} else if value.Type == vm_grpc.VmWriteOp_Value {
			k.SetValue(ctx, value.Path, value.Value)
		} else {
			// must not happens
			panic(fmt.Errorf("unknown writeOp (should not happen): %d", value.Type))
		}
	}
}
