package tests

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cryptoTypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func GenAccAddress() (sdk.AccAddress, cryptoTypes.PubKey, cryptoTypes.PrivKey) {
	pk := ed25519.GenPrivKey()

	return sdk.AccAddress(pk.PubKey().Address()), pk.PubKey(), pk
}

func CheckPanicErrorContains(t *testing.T, handler func(), errSubStr string) {
	defer func() {
		r := recover()
		require.NotNil(t, r, "handler did not panic")

		err, ok := r.(error)
		require.True(t, ok, "panic obj is not an error")

		require.Contains(t, err.Error(), errSubStr)
	}()

	handler()
}
