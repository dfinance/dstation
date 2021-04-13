package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Validate validates CurrentPrice object.
func (m CurrentPrice) Validate() error {
	validatePrice := func(v sdk.Int) error {
		if v.IsNil() {
			return fmt.Errorf("nil")
		}
		if v.IsZero() {
			return fmt.Errorf("zero")
		}
		if v.IsNegative() {
			return fmt.Errorf("negative")
		}

		return nil
	}

	if err := m.AssetCode.Validate(); err != nil {
		return fmt.Errorf("asset_code: %w", err)
	}

	if err := validatePrice(m.BidPrice); err != nil {
		return fmt.Errorf("bid_price: %w", err)
	}
	if err := validatePrice(m.AskPrice); err != nil {
		return fmt.Errorf("ask_price: %w", err)
	}

	if m.ReceivedAt.IsZero() {
		return fmt.Errorf("received_at: zero")
	}

	return nil
}

// ReversedPrice returns a new CurrentPrice with reversed exchange rate.
func (m CurrentPrice) ReversedPrice(assetDecimals uint32) CurrentPrice {
	reverseInt := func(price sdk.Int) sdk.Int {
		priceDec := sdk.NewDecFromBigIntWithPrec(price.BigInt(), int64(assetDecimals))
		priceDec = sdk.NewDec(1).Quo(priceDec)
		priceDec = priceDec.Mul(sdk.NewDec(10).Power(uint64(assetDecimals)))

		return priceDec.TruncateInt()
	}

	return CurrentPrice{
		AssetCode:  m.AssetCode.ReverseCode(),
		AskPrice:   reverseInt(m.BidPrice),
		BidPrice:   reverseInt(m.AskPrice),
		ReceivedAt: m.ReceivedAt,
		IsReversed: true,
	}
}
