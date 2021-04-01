package types

import (
	"errors"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/dfinance/dvm-proto/go/types_grpc"
	"github.com/stretchr/testify/require"
)

func TestVM_MsgDeployModule(t *testing.T) {
	t.Parallel()

	acc, err := sdk.AccAddressFromBech32("cosmos18557aj0s0dxdd24elwmf6tgv4x6vas3h3vzg5u")
	require.NoError(t, err)
	code := make([]byte, 128)

	// ok
	{
		msg := NewMsgDeployModule(acc, [][]byte{code})

		require.Equal(t, acc.String(), msg.Signer)
		require.Equal(t, code, msg.Modules[0])
		//
		require.Equal(t, RouterKey, msg.Route())
		require.Equal(t, TypeMsgDeployModule, msg.Type())
		require.Equal(t, []sdk.AccAddress{acc}, msg.GetSigners())
		require.Equal(t, getMsgSignBytes(&msg), msg.GetSignBytes())

		require.NoError(t, msg.ValidateBasic())
	}

	// fail
	{
		msg := NewMsgDeployModule([]byte{}, [][]byte{code})
		require.True(t, errors.Is(msg.ValidateBasic(), sdkErrors.ErrInvalidAddress))
	}

	// fail
	{
		msg := NewMsgDeployModule(acc, [][]byte{})
		require.True(t, errors.Is(msg.ValidateBasic(), ErrEmptyContract))
	}
}

func TestVM_MsgExecuteScript(t *testing.T) {
	t.Parallel()

	acc, err := sdk.AccAddressFromBech32("cosmos18557aj0s0dxdd24elwmf6tgv4x6vas3h3vzg5u")
	require.NoError(t, err)
	code := make([]byte, 128)

	args := []MsgExecuteScript_ScriptArg{
		{Type: types_grpc.VMTypeTag_U64, Value: []byte{0x1, 0x2, 0x3, 0x4}},
		{Type: types_grpc.VMTypeTag_Vector, Value: []byte{0x0}},
		{Type: types_grpc.VMTypeTag_Address, Value: Bech32ToLibra(acc)},
	}

	// ok
	{
		msg := NewMsgExecuteScript(acc, code, args...)

		require.Equal(t, acc.String(), msg.Signer)
		require.Equal(t, code, msg.Script)
		require.EqualValues(t, args, msg.Args)

		require.Equal(t, RouterKey, msg.Route())
		require.Equal(t, TypeMsgExecuteScript, msg.Type())
		require.Equal(t, []sdk.AccAddress{acc}, msg.GetSigners())
		require.Equal(t, getMsgSignBytes(&msg), msg.GetSignBytes())

		require.NoError(t, msg.ValidateBasic())
	}

	// ok
	{
		msg := NewMsgExecuteScript(acc, code)
		require.NoError(t, msg.ValidateBasic())
	}

	// fail
	{
		msg := NewMsgExecuteScript([]byte{}, code)
		require.True(t, errors.Is(msg.ValidateBasic(), sdkErrors.ErrInvalidAddress))
	}

	// fail
	{
		msg := NewMsgExecuteScript(acc, nil)
		require.True(t, errors.Is(msg.ValidateBasic(), ErrEmptyContract))
	}
}

func getMsgSignBytes(msg sdk.Msg) []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}
