package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Validate validates Asset object.
func (m Asset) Validate() error {
	if err := m.AssetCode.Validate(); err != nil {
		return fmt.Errorf("asset_code: %w", err)
	}

	for i, oracleAddr := range m.Oracles {
		if _, err := sdk.AccAddressFromBech32(oracleAddr); err != nil {
			return fmt.Errorf("oracles [%d]: Bech32 address convert: %w", i, err)
		}
	}

	return nil
}

// NormalizePriceValue normalizes sdk.Int with decimals to sdk.Int with asset decimals (truncate).
func (m Asset) NormalizePriceValue(value sdk.Int, valueDecimals uint32) sdk.Int {
	assetDec := sdk.NewDecWithPrec(1, int64(m.Decimals))
	valueDec := sdk.NewDecFromIntWithPrec(value, int64(valueDecimals))

	return valueDec.Quo(assetDec).TruncateInt()
}
