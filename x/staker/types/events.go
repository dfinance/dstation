package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	EventTypeDeposit  = ModuleName + ".deposit"
	EventTypeWithdraw = ModuleName + ".withdraw"
	//
	AttributeTargetAddress = "target_address"
	AttributeAmount        = "amount"
)

// NewDepositEvent creates an Event on Deposit operation.
func NewDepositEvent(call Call) sdk.Event {
	return sdk.NewEvent(EventTypeDeposit,
		sdk.NewAttribute(AttributeTargetAddress, call.Address),
		sdk.NewAttribute(AttributeAmount, call.Amount.String()),
	)
}

// NewWithdrawEvent creates an Event on Withdraw operation.
func NewWithdrawEvent(call Call) sdk.Event {
	return sdk.NewEvent(EventTypeWithdraw,
		sdk.NewAttribute(AttributeTargetAddress, call.Address),
		sdk.NewAttribute(AttributeAmount, call.Amount.String()),
	)
}
