package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dfinance/dstation/pkg/tests"
	dnTypes "github.com/dfinance/dstation/pkg/types"
	"github.com/dfinance/dstation/x/oracle/types"
)

func (s *KeeperTestSuite) TestPostPrice() {
	ctx, keeper := s.ctx, s.keeper
	params := keeper.GetParams(ctx)

	// fail: invalid input
	{
		_, curBlockTime := s.app.GetCurrentBlockHeightTime()

		// non-existing oracle
		{
			mockOracleAddr, _, _ := tests.GenAccAddress()

			msg := types.MsgPostPrice{
				AssetCode:     s.assets[0].AssetCode,
				OracleAddress: mockOracleAddr.String(),
				AskPrice:      sdk.NewInt(2),
				BidPrice:      sdk.NewInt(1),
				ReceivedAt:    curBlockTime.Add(1 * time.Second),
			}
			s.Require().Error(keeper.PostPrice(ctx, msg))
		}

		// non-existing asset
		{
			msg := types.MsgPostPrice{
				AssetCode:     "nonexsa_nonexb",
				OracleAddress: s.oracles[0].AccAddress,
				AskPrice:      sdk.NewInt(2),
				BidPrice:      sdk.NewInt(1),
				ReceivedAt:    curBlockTime.Add(1 * time.Second),
			}
			s.Require().Error(keeper.PostPrice(ctx, msg))
		}

		// invalid asset oracle
		{
			msg := types.MsgPostPrice{
				AssetCode:     s.assets[0].AssetCode,
				OracleAddress: s.oracles[1].AccAddress,
				AskPrice:      sdk.NewInt(2),
				BidPrice:      sdk.NewInt(1),
				ReceivedAt:    curBlockTime.Add(1 * time.Second),
			}
			s.Require().Error(keeper.PostPrice(ctx, msg))
		}

		// invalid receivedAt
		{
			msg := types.MsgPostPrice{
				AssetCode:     s.assets[0].AssetCode,
				OracleAddress: s.oracles[0].AccAddress,
				AskPrice:      sdk.NewInt(2),
				BidPrice:      sdk.NewInt(1),
				ReceivedAt:    curBlockTime.Add(time.Duration(params.PostPrice.ReceivedAtDiffInS+1) * time.Second),
			}
			s.Require().Error(keeper.PostPrice(ctx, msg))
		}

		// invalid price bytes len
		bigPrice, ok := sdk.NewIntFromString("10000000000000000000000000000000")
		s.Require().True(ok)
		{
			msg1 := types.MsgPostPrice{
				AssetCode:     s.assets[0].AssetCode,
				OracleAddress: s.oracles[0].AccAddress,
				AskPrice:      bigPrice,
				BidPrice:      sdk.NewInt(1),
				ReceivedAt:    curBlockTime.Add(1 * time.Second),
			}
			s.Require().Error(keeper.PostPrice(ctx, msg1))

			msg2 := types.MsgPostPrice{
				AssetCode:     s.assets[0].AssetCode,
				OracleAddress: s.oracles[0].AccAddress,
				AskPrice:      sdk.NewInt(2),
				BidPrice:      bigPrice,
				ReceivedAt:    curBlockTime.Add(1 * time.Second),
			}
			s.Require().Error(keeper.PostPrice(ctx, msg2))
		}
	}

	// ok: "1st" block
	block1, curBlockTime := s.app.GetCurrentBlockHeightTime()
	postPrices1 := []types.MsgPostPrice{
		{
			AssetCode:     s.assets[0].AssetCode,
			OracleAddress: s.oracles[0].AccAddress,
			AskPrice:      sdk.NewInt(200000000),
			BidPrice:      sdk.NewInt(100000000),
			ReceivedAt:    curBlockTime.Add(10 * time.Millisecond),
		},
		{
			AssetCode:     s.assets[1].AssetCode,
			OracleAddress: s.oracles[1].AccAddress,
			AskPrice:      sdk.NewInt(50000000),
			BidPrice:      sdk.NewInt(45000000),
			ReceivedAt:    curBlockTime.Add(10 * time.Millisecond),
		},
		{
			AssetCode:     s.assets[0].AssetCode,
			OracleAddress: s.oracles[0].AccAddress,
			AskPrice:      sdk.NewInt(210000000),
			BidPrice:      sdk.NewInt(110000000),
			ReceivedAt:    curBlockTime.Add(20 * time.Millisecond),
		},
	}
	{
		for i, msg := range postPrices1 {
			s.Require().NoError(keeper.PostPrice(ctx, msg), "postPrice [%d]", i)
		}

		// Remove 1st as it should be overwritten
		postPrices1 = postPrices1[1:]

		// End block and update context
		s.app.BeginBlock()
		s.app.EndBlock()
		ctx = s.app.GetContext(false)
	}

	// ok: "2nd" block
	block2, curBlockTime := s.app.GetCurrentBlockHeightTime()
	postPrices2 := []types.MsgPostPrice{
		{
			AssetCode:     s.assets[0].AssetCode,
			OracleAddress: s.oracles[0].AccAddress,
			AskPrice:      sdk.NewInt(190000000),
			BidPrice:      sdk.NewInt(90000000),
			ReceivedAt:    curBlockTime.Add(5 * time.Millisecond),
		},
		{
			AssetCode:     s.assets[1].AssetCode,
			OracleAddress: s.oracles[1].AccAddress,
			AskPrice:      sdk.NewInt(50000001),
			BidPrice:      sdk.NewInt(44999999),
			ReceivedAt:    curBlockTime.Add(15 * time.Millisecond),
		},
		{
			AssetCode:     s.assets[1].AssetCode,
			OracleAddress: s.oracles[2].AccAddress,
			AskPrice:      sdk.NewInt(50000010),
			BidPrice:      sdk.NewInt(44999990),
			ReceivedAt:    curBlockTime.Add(20 * time.Millisecond),
		},
	}
	{
		for i, msg := range postPrices2 {
			s.Require().NoError(keeper.PostPrice(ctx, msg), "postPrice [%d]", i)
		}
	}

	// ok: GetRawPrices
	checkPostRawPricesEqual := func(assetCode dnTypes.AssetCode, post []types.MsgPostPrice, rawRcv []types.RawPrice) {
		rawExp := make([]types.RawPrice, 0, len(post))
		for _, p := range post {
			if p.AssetCode != assetCode {
				continue
			}

			asset := types.Asset{}
			for _, a := range s.assets {
				if p.AssetCode == a.AssetCode {
					asset = a
					break
				}
			}

			oracle := types.Oracle{}
			for _, o := range s.oracles {
				if p.OracleAddress == o.AccAddress {
					oracle = o
					break
				}
			}

			rawExp = append(rawExp, types.RawPrice{
				AskPrice:   asset.NormalizePriceValue(p.AskPrice, oracle.PriceDecimals),
				BidPrice:   asset.NormalizePriceValue(p.BidPrice, oracle.PriceDecimals),
				ReceivedAt: p.ReceivedAt,
			})
		}
		s.ElementsMatch(rawExp, rawRcv)
	}
	{
		// Non-existing assetCode
		rawPricesEmpty1 := keeper.GetRawPrices(ctx, dnTypes.AssetCode("nonex1_nonex2"), block1)
		s.Require().Empty(rawPricesEmpty1)

		// Empty block (the next one)
		rawPricesEmpty2 := keeper.GetRawPrices(ctx, s.assets[0].AssetCode, block2+1)
		s.Require().Empty(rawPricesEmpty2)

		// Block1
		checkPostRawPricesEqual(
			s.assets[0].AssetCode,
			postPrices1,
			keeper.GetRawPrices(ctx, s.assets[0].AssetCode, block1),
		)
		checkPostRawPricesEqual(
			s.assets[1].AssetCode,
			postPrices1,
			keeper.GetRawPrices(ctx, s.assets[1].AssetCode, block1),
		)

		// Block2
		checkPostRawPricesEqual(
			s.assets[0].AssetCode,
			postPrices2,
			keeper.GetRawPrices(ctx, s.assets[0].AssetCode, block2),
		)
		checkPostRawPricesEqual(
			s.assets[1].AssetCode,
			postPrices2,
			keeper.GetRawPrices(ctx, s.assets[1].AssetCode, block2),
		)
	}
}
