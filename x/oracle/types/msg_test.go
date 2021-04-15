package types_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dfinance/dstation/pkg/tests"
	"github.com/dfinance/dstation/x/oracle/types"
	"github.com/stretchr/testify/require"
)

func TestOracle_MsgSetOracle(t *testing.T) {
	mockOracleAddr, _, _ := tests.GenAccAddress()

	// fail: invalid nominee
	{
		msg := types.NewMsgSetOracle([]byte("abc"), mockOracleAddr, "desc", 8, 8)
		require.Error(t, msg.ValidateBasic())
	}

	// fail: invalid oracle
	{
		msg := types.NewMsgSetOracle(mockOracleAddr, mockOracleAddr, "desc", 0, 8)
		require.Error(t, msg.ValidateBasic())
	}

	// ok
	{
		msg := types.NewMsgSetOracle(mockOracleAddr, mockOracleAddr, "desc", 8, 6)

		require.Equal(t, mockOracleAddr.String(), msg.Nominee)
		require.Equal(t, mockOracleAddr.String(), msg.Oracle.AccAddress)
		require.Equal(t, "desc", msg.Oracle.Description)
		require.EqualValues(t, 8, msg.Oracle.PriceMaxBytes)
		require.EqualValues(t, 6, msg.Oracle.PriceDecimals)
		//
		require.Equal(t, types.RouterKey, msg.Route())
		require.Equal(t, types.TypeMsgSetOracle, msg.Type())
		require.Equal(t, []sdk.AccAddress{mockOracleAddr}, msg.GetSigners())
		require.Equal(t, getMsgSignBytes(&msg), msg.GetSignBytes())

		require.NoError(t, msg.ValidateBasic())
	}
}

func TestOracle_MsgSetAsset(t *testing.T) {
	mockOracleAddr, _, _ := tests.GenAccAddress()

	// fail: invalid nominee
	{
		msg := types.NewMsgSetAsset([]byte("abc"), "btc_usdt", 8, mockOracleAddr)
		require.Error(t, msg.ValidateBasic())
	}

	// fail: invalid asset
	{
		msg := types.NewMsgSetAsset(mockOracleAddr, "btcusdt", 8, mockOracleAddr)
		require.Error(t, msg.ValidateBasic())
	}

	// ok
	{
		msg := types.NewMsgSetAsset(mockOracleAddr, "btc_usdt", 8, mockOracleAddr)

		require.Equal(t, mockOracleAddr.String(), msg.Nominee)
		require.Equal(t, "btc_usdt", msg.Asset.AssetCode.String())
		require.Len(t, msg.Asset.Oracles, 1)
		require.Equal(t, mockOracleAddr.String(), msg.Asset.Oracles[0])
		require.EqualValues(t, 8, msg.Asset.Decimals)
		//
		require.Equal(t, types.RouterKey, msg.Route())
		require.Equal(t, types.TypeMsgSetAsset, msg.Type())
		require.Equal(t, []sdk.AccAddress{mockOracleAddr}, msg.GetSigners())
		require.Equal(t, getMsgSignBytes(&msg), msg.GetSignBytes())

		require.NoError(t, msg.ValidateBasic())
	}
}

func TestOracle_MsgPostPrice(t *testing.T) {
	mockOracleAddr, _, _ := tests.GenAccAddress()

	// fail: invalid assetCode
	{
		msg := types.NewMsgPostPrice("btcusdt", mockOracleAddr, sdk.NewInt(1), sdk.NewInt(2), time.Now())
		require.Error(t, msg.ValidateBasic())
	}

	// fail: invalid bid/ask
	{
		{
			msg := types.NewMsgPostPrice("btc_usdt", mockOracleAddr, sdk.ZeroInt(), sdk.NewInt(2), time.Now())
			require.Error(t, msg.ValidateBasic())
		}
		{
			msg := types.NewMsgPostPrice("btc_usdt", mockOracleAddr, sdk.NewInt(-1), sdk.NewInt(2), time.Now())
			require.Error(t, msg.ValidateBasic())
		}
	}

	// fail: invalid oracleAddress
	{
		msg := types.NewMsgPostPrice("btc_usdt", []byte("abc"), sdk.NewInt(1), sdk.NewInt(2), time.Time{})
		require.Error(t, msg.ValidateBasic())
	}

	// fail: invalid receivedAt
	{
		msg := types.NewMsgPostPrice("btc_usdt", mockOracleAddr, sdk.NewInt(1), sdk.NewInt(2), time.Time{})
		require.Error(t, msg.ValidateBasic())
	}

	// ok
	{
		now := time.Now()
		msg := types.NewMsgPostPrice("btc_usdt", mockOracleAddr, sdk.NewInt(1), sdk.NewInt(2), now)

		require.Equal(t, "btc_usdt", msg.AssetCode.String())
		require.Equal(t, mockOracleAddr.String(), msg.OracleAddress)
		require.Equal(t, "1", msg.AskPrice.String())
		require.Equal(t, "2", msg.BidPrice.String())
		require.True(t, now.Equal(msg.ReceivedAt))
		//
		require.Equal(t, types.RouterKey, msg.Route())
		require.Equal(t, types.TypeMsgPostPrice, msg.Type())
		require.Equal(t, []sdk.AccAddress{mockOracleAddr}, msg.GetSigners())
		require.Equal(t, getMsgSignBytes(&msg), msg.GetSignBytes())

		require.NoError(t, msg.ValidateBasic())
	}
}

func getMsgSignBytes(msg sdk.Msg) []byte {
	return sdk.MustSortJSON(types.ModuleCdc.MustMarshalJSON(msg))
}
