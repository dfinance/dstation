package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	dnTypes "github.com/dfinance/dstation/pkg/types"
)

var (
	_ sdk.Msg = (*MsgSetOracle)(nil)
	_ sdk.Msg = (*MsgSetAsset)(nil)
	_ sdk.Msg = (*MsgPostPrice)(nil)
)

const (
	TypeMsgSetOracle = "set_oracle"
	TypeMsgSetAsset  = "set_asset"
	TypeMsgPostPrice = "post_price"
)

// Route implements sdk.Msg interface.
func (MsgSetOracle) Route() string {
	return RouterKey
}

// Type implements sdk.Msg interface.
func (MsgSetOracle) Type() string {
	return TypeMsgSetOracle
}

func (m MsgSetOracle) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Nominee); err != nil {
		return fmt.Errorf("nominee: %w", err)
	}

	if err := m.Oracle.Validate(); err != nil {
		return fmt.Errorf("oracle: %w", err)
	}

	return nil
}

// GetSignBytes implements sdk.Msg interface.
func (m MsgSetOracle) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg interface.
func (m MsgSetOracle) GetSigners() []sdk.AccAddress {
	// signer pays fees
	signerAddr, err := sdk.AccAddressFromBech32(m.Nominee)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signerAddr}
}

// NewMsgSetOracle creates a new MsgSetOracle message.
func NewMsgSetOracle(nomineeAddress, accAddress sdk.AccAddress, description string, priceMaxBytes, priceDecimals uint32) MsgSetOracle {
	return MsgSetOracle{
		Nominee: nomineeAddress.String(),
		Oracle: Oracle{
			AccAddress:    accAddress.String(),
			Description:   description,
			PriceMaxBytes: priceMaxBytes,
			PriceDecimals: priceDecimals,
		},
	}
}

// Route implements sdk.Msg interface.
func (MsgSetAsset) Route() string {
	return RouterKey
}

// Type implements sdk.Msg interface.
func (MsgSetAsset) Type() string {
	return TypeMsgSetAsset
}

func (m MsgSetAsset) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Nominee); err != nil {
		return fmt.Errorf("nominee: %w", err)
	}

	if err := m.Asset.Validate(); err != nil {
		return fmt.Errorf("asset: %w", err)
	}

	return nil
}

// GetSignBytes implements sdk.Msg interface.
func (m MsgSetAsset) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg interface.
func (m MsgSetAsset) GetSigners() []sdk.AccAddress {
	// signer pays fees
	signerAddr, err := sdk.AccAddressFromBech32(m.Nominee)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signerAddr}
}

// NewMsgSetAsset creates a new MsgSetAsset message.
func NewMsgSetAsset(nomineeAddress sdk.AccAddress, assetCode dnTypes.AssetCode, decimals uint32, oracleAddresses ...sdk.AccAddress) MsgSetAsset {
	m := MsgSetAsset{
		Nominee: nomineeAddress.String(),
		Asset: Asset{
			AssetCode: assetCode,
			Decimals:  decimals,
		},
	}
	for _, addr := range oracleAddresses {
		m.Asset.Oracles = append(m.Asset.Oracles, addr.String())
	}

	return m
}

// Route implements sdk.Msg interface.
func (MsgPostPrice) Route() string {
	return RouterKey
}

// Type implements sdk.Msg interface.
func (MsgPostPrice) Type() string {
	return TypeMsgPostPrice
}

func (m MsgPostPrice) ValidateBasic() error {
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

	if _, err := sdk.AccAddressFromBech32(m.OracleAddress); err != nil {
		return fmt.Errorf("oracle_address: %w", err)
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

// GetSignBytes implements sdk.Msg interface.
func (m MsgPostPrice) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&m)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg interface.
func (m MsgPostPrice) GetSigners() []sdk.AccAddress {
	// signer pays fees
	signerAddr, err := sdk.AccAddressFromBech32(m.OracleAddress)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{signerAddr}
}

// NewMsgPostPrice creates a new MsgPostPrice message.
func NewMsgPostPrice(assetCode dnTypes.AssetCode, oracleAddress sdk.AccAddress, askPrice, bidPrice sdk.Int, receivedAt time.Time) MsgPostPrice {
	return MsgPostPrice{
		AssetCode:     assetCode,
		OracleAddress: oracleAddress.String(),
		AskPrice:      askPrice,
		BidPrice:      bidPrice,
		ReceivedAt:    receivedAt,
	}
}
