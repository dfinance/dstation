package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = (*MsgDepositCall)(nil)
	_ sdk.Msg = (*MsgWithdrawCall)(nil)
)

const (
	TypeMsgDepositCall  = "deposit_call"
	TypeMsgWithdrawCall = "withdraw_call"
)

// Route implements sdk.Msg interface.
func (MsgDepositCall) Route() string {
	return RouterKey
}

// Type implements sdk.Msg interface.
func (MsgDepositCall) Type() string {
	return TypeMsgDepositCall
}

func (m MsgDepositCall) ValidateBasic() error {
	if m.UniqueId == "" {
		return fmt.Errorf("unique_id: empty")
	}

	if _, err := sdk.AccAddressFromBech32(m.Nominee); err != nil {
		return fmt.Errorf("nominee: %w", err)
	}

	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return fmt.Errorf("address: %w", err)
	}

	if err := m.SourceMeta.Validate(); err != nil {
		return fmt.Errorf("source_meta: %w", err)
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
func (m MsgDepositCall) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg interface.
func (m MsgDepositCall) GetSigners() []sdk.AccAddress {
	// signer pays fees
	signerAddr, err := sdk.AccAddressFromBech32(m.Nominee)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signerAddr}
}

// NewMsgDepositCall creates a new MsgDepositCall message.
func NewMsgDepositCall(uniqueId string, nomineeAddress, targetAccAddress sdk.AccAddress, srcEthAddr, srcChainId string, coins sdk.Coins) MsgDepositCall {
	return MsgDepositCall{
		UniqueId: uniqueId,
		Nominee:  nomineeAddress.String(),
		Address:  targetAccAddress.String(),
		SourceMeta: CallSourceMeta{
			EthAddress: srcEthAddr,
			ChainId:    srcChainId,
		},
		Amount: coins,
	}
}

// Route implements sdk.Msg interface.
func (MsgWithdrawCall) Route() string {
	return RouterKey
}

// Type implements sdk.Msg interface.
func (MsgWithdrawCall) Type() string {
	return TypeMsgWithdrawCall
}

func (m MsgWithdrawCall) ValidateBasic() error {
	if m.UniqueId == "" {
		return fmt.Errorf("unique_id: empty")
	}

	if _, err := sdk.AccAddressFromBech32(m.Nominee); err != nil {
		return fmt.Errorf("nominee: %w", err)
	}

	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return fmt.Errorf("address: %w", err)
	}

	if err := m.SourceMeta.Validate(); err != nil {
		return fmt.Errorf("source_meta: %w", err)
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
func (m MsgWithdrawCall) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg interface.
func (m MsgWithdrawCall) GetSigners() []sdk.AccAddress {
	// signer pays fees
	signerAddr, err := sdk.AccAddressFromBech32(m.Nominee)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signerAddr}
}

// NewMsgWithdrawCall creates a new MsgDepositCall message.
func NewMsgWithdrawCall(uniqueId string, nomineeAddress, targetAccAddress sdk.AccAddress, srcEthAddr, srcChainId string, coins sdk.Coins) MsgWithdrawCall {
	return MsgWithdrawCall{
		UniqueId: uniqueId,
		Nominee:  nomineeAddress.String(),
		Address:  targetAccAddress.String(),
		SourceMeta: CallSourceMeta{
			EthAddress: srcEthAddr,
			ChainId:    srcChainId,
		},
		Amount: coins,
	}
}
