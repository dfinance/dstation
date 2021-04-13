package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dfinance/dstation/x/oracle/types"
)

func (s *KeeperTestSuite) TestCurrentPrice() {
	ctx, keeper := s.ctx, s.keeper
	var lastAsset0Price, lastAsset1Price types.CurrentPrice

	// ok: GetCurrentPrice / GetCurrentPrices: empty
	{
		s.Require().Nil(keeper.GetCurrentPrice(ctx, s.assets[0].AssetCode))
		s.Require().Nil(keeper.GetCurrentPrice(ctx, s.assets[1].AssetCode))
		s.Require().Empty(keeper.GetCurrentPrices(ctx))
	}

	// ok: UpdateCurrentPrices: even
	{
		s.app.BeginBlock()
		ctx = s.app.GetContext(false)
		_, curBlockTime := s.app.GetCurrentBlockHeightTime()

		expAskPrice := sdk.NewInt(5050000000)
		expBidPrice := sdk.NewInt(5525000000)
		expReceivedAt := curBlockTime.Add(15 * time.Millisecond)
		postPrices := []types.MsgPostPrice{
			{
				AssetCode:     s.assets[1].AssetCode,
				OracleAddress: s.oracles[1].AccAddress,
				AskPrice:      sdk.NewInt(5000000000), // 50.0
				BidPrice:      sdk.NewInt(5500000000), // 55.0
				ReceivedAt:    curBlockTime.Add(10 * time.Millisecond),
			},
			{
				AssetCode:     s.assets[1].AssetCode,
				OracleAddress: s.oracles[2].AccAddress,
				AskPrice:      sdk.NewInt(5100000000), // 51.0
				BidPrice:      sdk.NewInt(5550000000), // 55.5
				ReceivedAt:    expReceivedAt,          // 15 ms
			},
		}
		for _, msg := range postPrices {
			s.Require().NoError(keeper.PostPrice(ctx, msg))
		}
		s.app.EndBlock()
		ctx = s.app.GetContext(false)

		rcvPrice := keeper.GetCurrentPrice(ctx, s.assets[1].AssetCode)
		s.Require().NotNil(rcvPrice)
		s.Require().Equal(s.assets[1].AssetCode, rcvPrice.AssetCode)
		s.Require().Equal(expAskPrice.String(), rcvPrice.AskPrice.String())
		s.Require().Equal(expBidPrice.String(), rcvPrice.BidPrice.String())
		s.Require().True(rcvPrice.ReceivedAt.Equal(expReceivedAt))
		lastAsset1Price = *rcvPrice

		s.Require().ElementsMatch([]types.CurrentPrice{lastAsset1Price}, keeper.GetCurrentPrices(ctx))
	}

	// ok: UpdateCurrentPrices: odd
	{
		s.app.BeginBlock()
		ctx = s.app.GetContext(false)
		_, curBlockTime := s.app.GetCurrentBlockHeightTime()

		expAskPrice := sdk.NewInt(5500000000)
		expBidPrice := sdk.NewInt(6500000000)
		expReceivedAt := curBlockTime.Add(15 * time.Millisecond)
		postPrices := []types.MsgPostPrice{
			{
				AssetCode:     s.assets[1].AssetCode,
				OracleAddress: s.oracles[1].AccAddress,
				AskPrice:      sdk.NewInt(5000000000), // 50.0
				BidPrice:      sdk.NewInt(5500000000), // 55.0
				ReceivedAt:    expReceivedAt,          // 15 ms
			},
			{
				AssetCode:     s.assets[1].AssetCode,
				OracleAddress: s.oracles[2].AccAddress,
				AskPrice:      sdk.NewInt(6000000000), // 60.0
				BidPrice:      expBidPrice,            // 65.0
				ReceivedAt:    curBlockTime.Add(20 * time.Millisecond),
			},
			{
				AssetCode:     s.assets[1].AssetCode,
				OracleAddress: s.oracles[3].AccAddress,
				AskPrice:      expAskPrice,            // 55.0
				BidPrice:      sdk.NewInt(7000000000), // 70.0
				ReceivedAt:    curBlockTime.Add(5 * time.Millisecond),
			},
		}
		for _, msg := range postPrices {
			s.Require().NoError(keeper.PostPrice(ctx, msg))
		}
		s.app.EndBlock()
		ctx = s.app.GetContext(false)

		rcvPrice := keeper.GetCurrentPrice(ctx, s.assets[1].AssetCode)
		s.Require().NotNil(rcvPrice)
		s.Require().Equal(s.assets[1].AssetCode, rcvPrice.AssetCode)
		s.Require().Equal(expAskPrice.String(), rcvPrice.AskPrice.String())
		s.Require().Equal(expBidPrice.String(), rcvPrice.BidPrice.String())
		s.Require().True(rcvPrice.ReceivedAt.Equal(expReceivedAt))
		lastAsset1Price = *rcvPrice

		s.Require().ElementsMatch([]types.CurrentPrice{lastAsset1Price}, keeper.GetCurrentPrices(ctx))
	}

	// ok: UpdateCurrentPrices: 1 element
	{
		s.app.BeginBlock()
		ctx = s.app.GetContext(false)
		_, curBlockTime := s.app.GetCurrentBlockHeightTime()

		expAskPrice := sdk.NewInt(20000000)
		expBidPrice := sdk.NewInt(30000000)
		expReceivedAt := curBlockTime.Add(5 * time.Millisecond)
		postPrices := []types.MsgPostPrice{
			{
				AssetCode:     s.assets[0].AssetCode,
				OracleAddress: s.oracles[0].AccAddress,
				AskPrice:      sdk.NewInt(2000000000), // 20.0
				BidPrice:      sdk.NewInt(3000000000), // 30.0
				ReceivedAt:    expReceivedAt,
			},
		}
		for _, msg := range postPrices {
			s.Require().NoError(keeper.PostPrice(ctx, msg))
		}
		s.app.EndBlock()
		ctx = s.app.GetContext(false)

		rcvPrice := keeper.GetCurrentPrice(ctx, s.assets[0].AssetCode)
		s.Require().NotNil(rcvPrice)
		s.Require().Equal(s.assets[0].AssetCode, rcvPrice.AssetCode)
		s.Require().Equal(expAskPrice.String(), rcvPrice.AskPrice.String())
		s.Require().Equal(expBidPrice.String(), rcvPrice.BidPrice.String())
		s.Require().True(rcvPrice.ReceivedAt.Equal(expReceivedAt))
		lastAsset0Price = *rcvPrice

		s.Require().ElementsMatch([]types.CurrentPrice{lastAsset0Price, lastAsset1Price}, keeper.GetCurrentPrices(ctx))
	}

	// ok: UpdateCurrentPrices: no change (Ask / Bid not changed)
	{
		s.app.BeginBlock()
		ctx = s.app.GetContext(false)
		_, curBlockTime := s.app.GetCurrentBlockHeightTime()

		postPrices := []types.MsgPostPrice{
			{
				AssetCode:     s.assets[0].AssetCode,
				OracleAddress: s.oracles[0].AccAddress,
				AskPrice:      sdk.NewInt(2000000000), // 20.0
				BidPrice:      sdk.NewInt(3000000000), // 30.0
				ReceivedAt:    curBlockTime.Add(5 * time.Millisecond),
			},
		}
		for _, msg := range postPrices {
			s.Require().NoError(keeper.PostPrice(ctx, msg))
		}
		s.app.EndBlock()
		ctx = s.app.GetContext(false)

		s.Require().ElementsMatch([]types.CurrentPrice{lastAsset0Price, lastAsset1Price}, keeper.GetCurrentPrices(ctx))
	}

	// ok: GetCurrentPrice: reversed
	{
		price := keeper.GetCurrentPrice(ctx, s.assets[0].AssetCode.ReverseCode())
		s.Require().NotNil(price)
		s.Require().Equal(s.assets[0].AssetCode.ReverseCode(), price.AssetCode)
		s.Require().True(price.IsReversed)
	}
}
