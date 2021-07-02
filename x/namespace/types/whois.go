package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Validate validates Call object.
func (w Whois) Validate() error {
	uintNil := sdk.Uint{}
	if w.ID == uintNil {
		return fmt.Errorf("id: nil sdk.Uint")
	}

	if _, err := sdk.AccAddressFromBech32(w.Creator); err != nil {
		return fmt.Errorf("address: invalid AccAddress: %w", err)
	}

	if w.Price.IsZero() {
		return fmt.Errorf("price: empty")
	}
	if err := w.Price.Validate(); err != nil {
		return fmt.Errorf("price: invalid sdk.Coins: %w", err)
	}

	return nil
}