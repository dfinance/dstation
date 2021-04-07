package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// Default address length (Move address length)
	VMAddressLength = 20
)

var (
	// Move stdlib addresses
	StdLibAddress         = make([]byte, VMAddressLength)
	StdLibAddressShortStr = "0x1"
)

// Bech32ToLibra converts Bech32 to Libra hex.
func Bech32ToLibra(addr sdk.AccAddress) []byte {
	return addr.Bytes()
}

// MustBech32ToLibra converts Bech32 address string to Libra hex and panics on failure.
func MustBech32ToLibra(addrRaw string) []byte {
	addr, err := sdk.AccAddressFromBech32(addrRaw)
	if err != nil {
		panic(fmt.Errorf("raw bech32 string convert: %v", err))
	}

	return Bech32ToLibra(addr)
}

// LibraToBech32 converts Libra bytes to sdk.AccAddress.
func LibraToBech32(addrRaw []byte) (sdk.AccAddress, error) {
	if len(addrRaw) != VMAddressLength {
		return sdk.AccAddress{}, fmt.Errorf("invalid length (must be %d)", VMAddressLength)
	}

	return sdk.AccAddress(addrRaw), nil
}

func init() {
	StdLibAddress[VMAddressLength-1] = 1
}
