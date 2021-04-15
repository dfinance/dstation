package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Validate validates Oracle object.
func (m Oracle) Validate() error {
	if _, err := sdk.AccAddressFromBech32(m.AccAddress); err != nil {
		return fmt.Errorf("acc_address: invalid Bech32 address: %w", err)
	}

	if m.PriceMaxBytes == 0 {
		return fmt.Errorf("price_bytes: zero")
	}

	return nil
}

// MustGetAccAddress converts Oracle AccAddress to sdk.AccAddress and panics of failure.
func (m Oracle) MustGetAccAddress() sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(m.AccAddress)
	if err != nil {
		panic(fmt.Errorf("converting Oracle AccAddress (%s) to sdk.AccAddress: %w", m.AccAddress, err))
	}

	return accAddr
}
