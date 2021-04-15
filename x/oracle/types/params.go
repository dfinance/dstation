package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramTypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter keys
var (
	ParamStoreKeyNominees  = []byte("oraclenominees")
	ParamStoreKeyPostPrice = []byte("oraclepostprice")
)

// ParamKeyTable returns parameters key table.
func ParamKeyTable() paramTypes.KeyTable {
	return paramTypes.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs returns the parameter set pairs.
func (m *Params) ParamSetPairs() paramTypes.ParamSetPairs {
	return paramTypes.ParamSetPairs{
		paramTypes.NewParamSetPair(ParamStoreKeyNominees, &m.Nominees, validateNominees),
		paramTypes.NewParamSetPair(ParamStoreKeyPostPrice, &m.PostPrice, validatePostPrice),
	}
}

// ValidateBasic validates keeper parameters.
func (m Params) ValidateBasic() error {
	if err := validateNominees(m.Nominees); err != nil {
		return err
	}
	if err := validatePostPrice(m.PostPrice); err != nil {
		return err
	}

	return nil
}

func (m Params) String() string {
	out := strings.Builder{}
	out.WriteString("Params:\n")

	out.WriteString("  Nominees:\n")
	for _, n := range m.Nominees {
		out.WriteString(fmt.Sprintf("    - %s\n", n))
	}

	out.WriteString("  PostPrice:\n")
	out.WriteString(fmt.Sprintf("    ReceivedAtDiff: %d [s]", m.PostPrice.ReceivedAtDiffInS))

	return out.String()
}

// validateNominees performs validation for ParamStoreKeyNominees params pair.
func validateNominees(i interface{}) error {
	const paramName = "nominees"

	v, ok := i.([]string)
	if !ok {
		return fmt.Errorf("%s param: invalid type: %T", paramName, i)
	}

	nomineeSet := make(map[string]struct{})
	for _, nominee := range v {
		if _, found := nomineeSet[nominee]; found {
			return fmt.Errorf("%s param: nominee (%s): duplicated", paramName, nominee)
		}
		nomineeSet[nominee] = struct{}{}

		if _, err := sdk.AccAddressFromBech32(nominee); err != nil {
			return fmt.Errorf("%s param: nominee (%s): invalid Bech32 accAddress: %w", paramName, nominee, err)
		}
	}

	return nil
}

// validatePostPrice performs validation for ParamStoreKeyPostPrice params pair.
func validatePostPrice(i interface{}) error {
	const paramName = "post_price"

	_, ok := i.(Params_PostPriceParams)
	if !ok {
		return fmt.Errorf("%s param: invalid type: %T", paramName, i)
	}

	return nil
}

// DefaultParams returns default keeper parameters.
func DefaultParams() Params {
	return Params{
		PostPrice: Params_PostPriceParams{
			ReceivedAtDiffInS: 60 * 60,
		},
	}
}
