package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	assetCodeDelimiter = '_'
)

// AssetCode is a wrapper type for denom pair.
type AssetCode string

// Validate validates AssetCode.
func (a AssetCode) Validate() error {
	if err := AssetCodeFilter(a.String()); err != nil {
		return err
	}

	leftDenom, rightDenom := a.Split()
	if err := sdk.ValidateDenom(leftDenom); err != nil {
		return fmt.Errorf("leftDenom: %w", err)
	}
	if err := sdk.ValidateDenom(rightDenom); err != nil {
		return fmt.Errorf("rightDenom: %w", err)
	}

	return nil
}

// String implements fmt.Stringer interface.
func (a AssetCode) String() string {
	return string(a)
}

// Split returns AssetCode denoms.
func (a AssetCode) Split() (string, string) {
	parts := strings.Split(a.String(), string(assetCodeDelimiter))
	if len(parts) != 2 {
		panic(fmt.Errorf("wrong asset code format: %s", a))
	}

	return parts[0], parts[1]
}

// ReverseCode reverses asset code.
// Will panic if use it with the wrong format asset code.
func (a AssetCode) ReverseCode() AssetCode {
	leftDenom, rightDenom := a.Split()

	return AssetCode(rightDenom + string(assetCodeDelimiter) + leftDenom)
}

// Marshal implements the gogo proto custom type interface.
func (a AssetCode) Marshal() ([]byte, error) {
	return []byte(a), nil
}

// MarshalTo implements the gogo proto custom type interface.
func (a *AssetCode) MarshalTo(data []byte) (n int, err error) {
	if a == nil {
		*a = ""
	}

	bz, err := a.Marshal()
	if err != nil {
		return 0, err
	}
	copy(data, bz)

	return len(bz), nil
}

// Unmarshal implements the gogo proto custom type interface.
func (a *AssetCode) Unmarshal(data []byte) error {
	*a = AssetCode(data)

	return nil
}

// Size implements the gogo proto custom type interface.
func (a *AssetCode) Size() int {
	bz, _ := a.Marshal()

	return len(bz)
}

// NewAssetCodeByDenoms creates a new AssetCode.
func NewAssetCodeByDenoms(leftDenom, rightDenom string) (AssetCode, error) {
	a := AssetCode(leftDenom + string(assetCodeDelimiter) + rightDenom)
	if err := a.Validate(); err != nil {
		return "", err
	}

	return a, nil
}

// NewAssetCode creates a new AssetCode.
func NewAssetCode(assetCode string) (AssetCode, error) {
	a := AssetCode(assetCode)
	if err := a.Validate(); err != nil {
		return "", err
	}

	return a, nil
}
