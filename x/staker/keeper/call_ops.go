package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dfinance/dstation/x/staker/types"
)

// Deposit creates a new types.Call, mints coins and sends them to the target account.
func (k Keeper) Deposit(ctx sdk.Context, msg types.MsgDepositCall) (types.Call, error) {
	call, err := k.createNewCall(ctx, types.Call_DEPOSIT, msg.Nominee, msg.Address, msg.Amount)
	if err != nil {
		return types.Call{}, err
	}

	targetAccAddr, _ := sdk.AccAddressFromBech32(call.Address)
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, call.Amount); err != nil {
		return types.Call{}, fmt.Errorf("minting coins (%s) for module: %w", call.Amount, err)
	}
	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, targetAccAddr, call.Amount); err != nil {
		return types.Call{}, fmt.Errorf("sending coins (%s) from module to account (%s): %w", call.Amount, targetAccAddr.String(), err)
	}

	return call, nil
}

// Withdraw creates a new types.Call, withdraws coins from the target account and burns them.
func (k Keeper) Withdraw(ctx sdk.Context, msg types.MsgWithdrawCall) (types.Call, error) {
	call, err := k.createNewCall(ctx, types.Call_WITHDRAW, msg.Nominee, msg.Address, msg.Amount)
	if err != nil {
		return types.Call{}, err
	}

	targetAccAddr, _ := sdk.AccAddressFromBech32(call.Address)
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, targetAccAddr, types.ModuleName, call.Amount); err != nil {
		return types.Call{}, fmt.Errorf("sending coins (%s) from account (%s) to module: %w", call.Amount, targetAccAddr.String(), err)
	}
	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, call.Amount); err != nil {
		return types.Call{}, fmt.Errorf("burning coins (%s) from module: %w", call.Amount, err)
	}

	return call, nil
}

// createNewCall creates and saves a new types.Call.
func (k Keeper) createNewCall(ctx sdk.Context, callType types.Call_CallType, nomineeAddr, accAddr string, amount sdk.Coins) (types.Call, error) {
	if err := k.IsNominee(ctx, nomineeAddr); err != nil {
		return types.Call{}, err
	}

	call := types.Call{
		Id:        k.getNextCallID(ctx),
		Nominee:   nomineeAddr,
		Address:   accAddr,
		Type:      callType,
		Amount:    amount,
		Timestamp: ctx.BlockTime(),
	}
	if err := call.Validate(); err != nil {
		return types.Call{}, fmt.Errorf("call validation: %w", err)
	}

	k.setCall(ctx, call)
	k.setLastCallID(ctx, call.Id)

	return call, nil
}
