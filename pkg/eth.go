package pkg

import (
	"encoding/hex"
	"fmt"
	"strings"
)

const (
	EthAddressLength = 20
)

// ValidateEthereumAddress validates Ethereum chain address.
func ValidateEthereumAddress(address string) error {
	if address == "" {
		return fmt.Errorf("empty")
	}

	if !strings.HasPrefix(address, "0x") {
		return fmt.Errorf("should be prefixed with 0x (HEX string)")
	}
	address = address[2:]

	addrBytes, err := hex.DecodeString(address)
	if err != nil {
		return fmt.Errorf("HEX decode: %w", err)
	}

	if len(addrBytes) != EthAddressLength {
		return fmt.Errorf("length mismatch, expected / actualaddress: %d / %d", EthAddressLength, len(addrBytes))
	}

	return nil
}
