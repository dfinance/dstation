package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = (*MsgExecuteScript)(nil)
	_ sdk.Msg = (*MsgDeployModule)(nil)
)

const (
	TypeMsgExecuteScript = "execute_script"
	TypeMsgDeployModule  = "deploy_module"
)

// Route implements sdk.Msg interface.
func (MsgExecuteScript) Route() string {
	return RouterKey
}

// Type implements sdk.Msg interface.
func (MsgExecuteScript) Type() string {
	return TypeMsgExecuteScript
}

// ValidateBasic implements sdk.Msg interface.
func (m MsgExecuteScript) ValidateBasic() error {
	signerAddr, err := sdk.AccAddressFromBech32(m.Signer)
	if err != nil {
		return sdkErrors.Wrapf(sdkErrors.ErrInvalidAddress, "signer address: invalid")
	}
	if signerAddr.Empty() {
		return sdkErrors.Wrapf(sdkErrors.ErrInvalidAddress, "signer address: empty")
	}

	if len(m.Script) == 0 {
		return sdkErrors.Wrapf(ErrEmptyContract, "script: empty")
	}

	for i, arg := range m.Args {
		if _, err := StringifyVMTypeTag(arg.Type); err != nil {
			return sdkErrors.Wrapf(ErrWrongArgTypeTag, "args [%d]: type: %v", i, err)
		}
		if len(arg.Value) == 0 {
			return sdkErrors.Wrapf(ErrWrongArgTypeTag, "args [%d]: value: empty", i)
		}
	}

	return nil
}

// GetSignBytes implements sdk.Msg interface.
func (m MsgExecuteScript) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg interface.
func (m MsgExecuteScript) GetSigners() []sdk.AccAddress {
	// signer pays fees
	signerAddr, err := sdk.AccAddressFromBech32(m.Signer)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signerAddr}
}

// NewMsgExecuteScript creates a new MsgExecuteScript message.
func NewMsgExecuteScript(signer sdk.AccAddress, script []byte, args ...MsgExecuteScript_ScriptArg) MsgExecuteScript {
	if len(args) == 0 {
		args = nil
	}

	return MsgExecuteScript{
		Signer: signer.String(),
		Script: script,
		Args:   args,
	}
}

// Route implements sdk.Msg interface.
func (MsgDeployModule) Route() string {
	return RouterKey
}

// Type implements sdk.Msg interface.
func (MsgDeployModule) Type() string {
	return TypeMsgDeployModule
}

// ValidateBasic implements sdk.Msg interface.
func (m MsgDeployModule) ValidateBasic() error {
	signerAddr, err := sdk.AccAddressFromBech32(m.Signer)
	if err != nil {
		return sdkErrors.Wrapf(sdkErrors.ErrInvalidAddress, "signer address: invalid")
	}
	if signerAddr.Empty() {
		return sdkErrors.Wrapf(sdkErrors.ErrInvalidAddress, "signer address: empty")
	}

	if len(m.Modules) == 0 {
		return sdkErrors.Wrapf(ErrEmptyContract, "modules: empty")
	}
	for i, module := range m.Modules {
		if len(module) == 0 {
			return sdkErrors.Wrapf(ErrEmptyContract, "modules [%d]: empty", i)
		}
	}

	return nil
}

// GetSignBytes implements sdk.Msg interface.
func (m MsgDeployModule) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg interface.
func (m MsgDeployModule) GetSigners() []sdk.AccAddress {
	// signer pays fees
	signerAddr, err := sdk.AccAddressFromBech32(m.Signer)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signerAddr}
}

// NewMsgDeployModule creates a new MsgDeployModule message.
func NewMsgDeployModule(signer sdk.AccAddress, modules ...[]byte) MsgDeployModule {
	return MsgDeployModule{
		Signer:  signer.String(),
		Modules: modules,
	}
}
