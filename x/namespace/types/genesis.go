package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Validate checks that genesis state is valid.
func (m GenesisState) Validate() error {

	if m.LastWhoisId != nil {
		uintNil := sdk.Uint{}
		if *m.LastWhoisId == uintNil {
			return fmt.Errorf("last_call_id: nil sdk.Uint")
		}
	}

	callsSet := make(map[string]struct{})
	for i, whois := range m.Whois {
		if _, found := callsSet[whois.ID.String()]; found {
			return fmt.Errorf("calls [%d]: duplicated", i)
		}

		callsSet[whois.ID.String()] = struct{}{}

		if err := whois.Validate(); err != nil {
			return fmt.Errorf("calls [%d]: %w", i, err)
		}
	}

	return nil
}

// DefaultGenesisState returns default genesis state (validation is done on module init).
func DefaultGenesisState() GenesisState {
	return GenesisState{
	}
}
