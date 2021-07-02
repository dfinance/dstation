package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = (*MsgBuyCall)(nil)
	_ sdk.Msg = (*MsgDeleteCall)(nil)
)

const (
	TypeMsgBuyCall  = "buy_call"
	TypeMsgDeleteCall = "delete_call"
)

// Route implements sdk.Msg interface.
func (MsgBuyCall) Route() string {
	return RouterKey
}

// Type implements sdk.Msg interface.
func (MsgBuyCall) Type() string {
	return TypeMsgBuyCall
}

func (m MsgBuyCall) ValidateBasic() error {
	if strings.Trim(m.Value, " ") == "" {
		return fmt.Errorf("value: empty")
	}

	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return fmt.Errorf("address: %w", err)
	}

	if m.Amount.IsZero() {
		return fmt.Errorf("amount: empty coins")
	}
	if err := m.Amount.Validate(); err != nil {
		return fmt.Errorf("amount: coins invalid: %w", err)
	}

	return nil
}

// GetSignBytes implements sdk.Msg interface.
func (m MsgBuyCall) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg interface.
func (m MsgBuyCall) GetSigners() []sdk.AccAddress {
	// signer pays fees
	signerAddr, err := sdk.AccAddressFromBech32(m.Address)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signerAddr}
}

// NewMsgDepositCall creates a new MsgDepositCall message.
func NewMsgBuyCall(address sdk.AccAddress, value string, coins sdk.Coins) MsgBuyCall {
	return MsgBuyCall{
		Address:  address.String(),
		Value: value,
		Amount: coins,
	}
}

// Route implements sdk.Msg interface.
func (MsgDeleteCall) Route() string {
	return RouterKey
}

// Type implements sdk.Msg interface.
func (MsgDeleteCall) Type() string {
	return TypeMsgDeleteCall
}

func (m MsgDeleteCall) ValidateBasic() error {
	if strings.Trim(m.Value, " ") == "" {
		return fmt.Errorf("value: empty")
	}

	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return fmt.Errorf("address: %w", err)
	}

	return nil
}

// GetSignBytes implements sdk.Msg interface.
func (m MsgDeleteCall) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg interface.
func (m MsgDeleteCall) GetSigners() []sdk.AccAddress {
	signerAddr, err := sdk.AccAddressFromBech32(m.Address)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signerAddr}
}

// MsgDeleteCall creates a new MsgDepositCall message.
func NewMsgDeleteCall(address sdk.AccAddress, value string) MsgDeleteCall {
	return MsgDeleteCall{
		Address:  address.String(),
		Value: value,
	}
}
