package pkg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ValidateEthereumAddress(t *testing.T) {
	ethAddr := "0x89205A3A3b2A69De6Dbf7f01ED13B2108B2c43e7"

	// fail: empty
	require.Error(t, ValidateEthereumAddress(""))

	// fail: no prefix
	require.Error(t, ValidateEthereumAddress(ethAddr[2:]))

	// fail: not a HEX string
	require.Error(t, ValidateEthereumAddress("0x@1234"))

	// fail: invalid length
	require.Error(t, ValidateEthereumAddress("0x001122AABBCC"))

	require.NoError(t, ValidateEthereumAddress(ethAddr))
}
