package types_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dfinance/dstation/x/oracle/types"
	"github.com/stretchr/testify/require"
)

func TestOracle_CurrentPrice(t *testing.T) {
	// fail: invalid assetCode
	{
		price := types.CurrentPrice{
			AssetCode:  "btcusdt",
			AskPrice:   sdk.NewInt(1),
			BidPrice:   sdk.NewInt(2),
			ReceivedAt: time.Now(),
		}
		require.Error(t, price.Validate())
	}

	// fail: invalid bid/ask
	{
		{
			price := types.CurrentPrice{
				AssetCode:  "btc_usdt",
				AskPrice:   sdk.ZeroInt(),
				BidPrice:   sdk.NewInt(2),
				ReceivedAt: time.Now(),
			}
			require.Error(t, price.Validate())
		}
		{
			price := types.CurrentPrice{
				AssetCode:  "btc_usdt",
				AskPrice:   sdk.NewInt(-1),
				BidPrice:   sdk.NewInt(2),
				ReceivedAt: time.Now(),
			}
			require.Error(t, price.Validate())
		}
	}

	// fail: invalid receivedAt
	{
		price := types.CurrentPrice{
			AssetCode:  "btc_usdt",
			AskPrice:   sdk.NewInt(1),
			BidPrice:   sdk.NewInt(2),
			ReceivedAt: time.Time{},
		}
		require.Error(t, price.Validate())
	}

	// ok
	{
		price := types.CurrentPrice{
			AssetCode:  "btc_usdt",
			AskPrice:   sdk.NewInt(1),
			BidPrice:   sdk.NewInt(2),
			ReceivedAt: time.Now(),
		}
		require.NoError(t, price.Validate())
	}
}

func TestOracle_CurrentPrice_ReversedPrice(t *testing.T) {
	// Ask / Bid / Decimals: 10900.55, 10889.95, 8
	{
		price := types.CurrentPrice{
			AssetCode:  "eth_usdt",
			AskPrice:   sdk.NewInt(213500000000), // 2135.0
			BidPrice:   sdk.NewInt(213494000000), // 2134.94
			ReceivedAt: time.Time{},
		}
		expAskReversed := sdk.NewInt(46839) // 0.00046839
		expBidReversed := sdk.NewInt(46838) // 0.00046838

		priceReversed := price.ReversedPrice(8)
		require.Equal(t, "usdt_eth", priceReversed.AssetCode.String())
		require.Equal(t, expAskReversed.String(), priceReversed.AskPrice.String())
		require.Equal(t, expBidReversed.String(), priceReversed.BidPrice.String())
		require.True(t, priceReversed.IsReversed)
	}
}
