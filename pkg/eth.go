package pkg

import (
	"encoding/hex"
)

const (
	EthAddressLength = 20
)

// Check it's hex
func isHex(s string) bool {
	_, err := hex.DecodeString(s)
	return err == nil
}

// IsEthereumAddress check if address is Ethereum address.
func IsEthereumAddress(address string) bool {
	if len(address) < 2 {
		return false
	}

	s := address[2:]
	return len(s) == 2*EthAddressLength && isHex(s)
}
