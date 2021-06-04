package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Validate validates Call object.
func (m Call) Validate() error {
	uintNil := sdk.Uint{}
	if m.Id == uintNil {
		return fmt.Errorf("id: nil sdk.Uint")
	}

	if _, err := sdk.AccAddressFromBech32(m.Nominee); err != nil {
		return fmt.Errorf("nominee: invalid AccAddress: %w", err)
	}
	if _, err := sdk.AccAddressFromBech32(m.Address); err != nil {
		return fmt.Errorf("address: invalid AccAddress: %w", err)
	}

	if _, found := Call_CallType_name[int32(m.Type)]; !found {
		return fmt.Errorf("type: invalid Call_CallType")
	}

	if m.Amount.IsZero() {
		return fmt.Errorf("amount: empty")
	}
	if err := m.Amount.Validate(); err != nil {
		return fmt.Errorf("amount: invalid sdk.Coins: %w", err)
	}

	if m.Timestamp.IsZero() {
		return fmt.Errorf("timestamp: empty")
	}

	return nil
}
