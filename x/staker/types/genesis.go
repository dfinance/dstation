package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Validate checks that genesis state is valid.
func (m GenesisState) Validate() error {
	if err := m.Params.ValidateBasic(); err != nil {
		return fmt.Errorf("params: %w", err)
	}

	if m.LastCallId != nil {
		uintNil := sdk.Uint{}
		if *m.LastCallId == uintNil {
			return fmt.Errorf("last_call_id: nil sdk.Uint")
		}
	}

	callsSet := make(map[string]struct{})
	for i, call := range m.Calls {
		if _, found := callsSet[call.Id.String()]; found {
			return fmt.Errorf("calls [%d]: duplicated", i)
		}
		callsSet[call.Id.String()] = struct{}{}

		if err := call.Validate(); err != nil {
			return fmt.Errorf("calls [%d]: %w", i, err)
		}
	}

	return nil
}

// DefaultGenesisState returns default genesis state (validation is done on module init).
func DefaultGenesisState() GenesisState {
	return GenesisState{
		Params: DefaultParams(),
	}
}
