package keeper

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"

	dnTypes "github.com/dfinance/dstation/pkg/types"
	dvmTypes "github.com/dfinance/dstation/pkg/types/dvm"
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
	execList := make([]*dvmTypes.VMExecuteResponse, 0, len(msg.Modules))
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
func (k Keeper) processVMExecution(ctx sdk.Context, exec *dvmTypes.VMExecuteResponse) {
	// Consume gas (if execution took too much gas - panic and mark transaction as out of gas)
	ctx.GasMeter().ConsumeGas(exec.GasUsed, "vm script/module execution")

	// Emit execution events
	ctx.EventManager().EmitEvent(dnTypes.NewModuleNameEvent(types.ModuleName))
	ctx.EventManager().EmitEvents(types.NewContractEvents(exec))

	// Process success status
	if exec.GetStatus().GetError() == nil {
		k.processVMBalanceChangeSet(ctx, exec.BalanceChangeSet)
		k.processVMWriteSet(ctx, exec.WriteSet)

		// Emit VM events (panic on "out of gas", emitted events stays in the EventManager)
		for _, vmEvent := range exec.Events {
			ctx.EventManager().EmitEvent(types.NewMoveEvent(ctx.GasMeter(), vmEvent))
		}

		// Update DS context with the current changes to be visible for other Tx messages (if any)
		// DS context is updated within BeginBlock as well
		k.SetDSContext(ctx)
	}
}

// processVMBalanceChangeSet processes VM native balance changes (delegate / undelegate coins to / from VM).
func (k Keeper) processVMBalanceChangeSet(ctx sdk.Context, changeSet []*dvmTypes.VMBalanceChange) {
	for _, value := range changeSet {
		if value == nil {
			panic(fmt.Errorf("processing account balance changes: nil value received"))
		}

		accAddr, err := types.LibraToBech32(value.Address)
		if err != nil {
			panic(fmt.Errorf("processing account balance changes: converting Libra address (%v) to Bech32: %w", value.Address, err))
		}
		denom := strings.ToLower(value.Ticker)

		switch op := value.Op.(type) {
		case *dvmTypes.VMBalanceChange_Deposit:
			coin := sdk.Coin{Denom: denom, Amount: types.VmU128ToSdkInt(op.Deposit)}
			if err := k.DelegateCoinsToPool(ctx, accAddr, sdk.Coins{coin}); err != nil {
				panic(fmt.Errorf("processing account balance changes: delegating coin (%s) to account (%s): %w", coin, accAddr, err))
			}
		case *dvmTypes.VMBalanceChange_Withdraw:
			coin := sdk.Coin{Denom: denom, Amount: types.VmU128ToSdkInt(op.Withdraw)}
			if err := k.UndelegateCoinsFromPool(ctx, accAddr, sdk.Coins{coin}); err != nil {
				panic(fmt.Errorf("processing account balance changes: undelegating coin (%s) to account (%s): %w", coin, accAddr, err))
			}
		default:
			panic(fmt.Errorf("processing account balance changes: unsupported change.Op: %T", value.Op))
		}
	}
}

// processVMWriteSet processes VM execution writeSets (set/delete).
func (k Keeper) processVMWriteSet(ctx sdk.Context, writeSet []*dvmTypes.VMValue) {
	for _, value := range writeSet {
		if value == nil {
			panic(fmt.Errorf("processing writeSets: nil value received"))
		}

		switch value.Type {
		case dvmTypes.VmWriteOp_Value:
			k.SetValue(ctx, value.Path, value.Value)
		case dvmTypes.VmWriteOp_Deletion:
			k.DelValue(ctx, value.Path)
		default:
			panic(fmt.Errorf("processing writeSets: unsupported writeOp.Type: %d", value.Type))
		}
	}
}
