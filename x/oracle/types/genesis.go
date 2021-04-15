package types

import "fmt"

// Validate checks that genesis state is valid.
func (m GenesisState) Validate() error {
	if err := m.Params.ValidateBasic(); err != nil {
		return fmt.Errorf("params: %w", err)
	}

	oracleSet := make(map[string]struct{})
	for i, oracle := range m.Oracles {
		if _, found := oracleSet[oracle.AccAddress]; found {
			return fmt.Errorf("oracles [%d]: duplicated", i)
		}
		oracleSet[oracle.AccAddress] = struct{}{}

		if err := oracle.Validate(); err != nil {
			return fmt.Errorf("oracles [%d]: %w", i, err)
		}
	}

	assetsSet := make(map[string]struct{})
	for i, asset := range m.Assets {
		if _, found := assetsSet[asset.String()]; found {
			return fmt.Errorf("assets [%d]: duplicated", i)
		}
		assetsSet[asset.String()] = struct{}{}

		if err := asset.Validate(); err != nil {
			return fmt.Errorf("assets [%d]: %w", i, err)
		}
	}

	for i, price := range m.CurrentPrices {
		if err := price.Validate(); err != nil {
			return fmt.Errorf("current_prices [%d]: %w", i, err)
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
