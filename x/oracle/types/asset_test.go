package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dfinance/dstation/pkg/tests"
	"github.com/dfinance/dstation/x/oracle/types"
	"github.com/stretchr/testify/require"
)

func TestOracle_Asset(t *testing.T) {
	mockAddr, _, _ := tests.GenAccAddress()

	// fail: invalid assetCode
	{
		asset := types.Asset{
			AssetCode: "btcusdt",
			Oracles:   []string{mockAddr.String()},
		}
		require.Error(t, asset.Validate())
	}

	// fail: invalid oracle
	{
		asset := types.Asset{
			AssetCode: "btc_usdt",
			Oracles:   []string{""},
		}
		require.Error(t, asset.Validate())
	}

	// ok
	{
		asset := types.Asset{
			AssetCode: "btc_usdt",
			Oracles:   []string{mockAddr.String()},
		}
		require.NoError(t, asset.Validate())
	}

	// ok
	{
		asset := types.Asset{
			AssetCode: "btc_usdt",
			Oracles:   nil,
		}
		require.NoError(t, asset.Validate())
	}
}

func TestOracle_AssetNormalizePriceValue(t *testing.T) {
	asset := types.Asset{
		AssetCode: "a_b",
		Decimals:  6,
	}

	// Price with more decimals
	{
		value := sdk.NewInt(5012345678) // 50.12345678
		decimals := uint32(8)

		price := asset.NormalizePriceValue(value, decimals)
		require.EqualValues(t, 50123456, price.Int64()) // 50.123456
	}

	// Price with less decimals
	{
		value := sdk.NewInt(2010) // 20.10
		decimals := uint32(2)

		price := asset.NormalizePriceValue(value, decimals)
		require.EqualValues(t, 20100000, price.Int64()) // 20.100000
	}

	// Price with the same number of decimals
	{
		value := sdk.NewInt(99987654) // 99.987654
		decimals := uint32(6)

		price := asset.NormalizePriceValue(value, decimals)
		require.EqualValues(t, 99987654, price.Int64()) // 99.987654
	}
}
