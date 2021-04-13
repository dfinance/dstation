package keeper_test

import (
	"errors"

	"github.com/dfinance/dstation/pkg/tests"
	"github.com/dfinance/dstation/x/oracle/types"
)

func (s *KeeperTestSuite) TestAsset() {
	ctx, keeper := s.ctx, s.keeper
	mockOracleAddr, _, _ := tests.GenAccAddress()

	initAssets := keeper.GetAssets(ctx)

	// fail: SetAsset: not a nominee
	{
		asset := types.Asset{
			AssetCode: "testa_testb",
			Oracles:   []string{mockOracleAddr.String()},
		}

		err := keeper.SetAsset(ctx, types.NewMsgSetAsset(mockOracleAddr, asset.AssetCode, asset.Decimals, mockOracleAddr))
		s.Require().Error(err)
		s.Require().True(errors.Is(err, types.ErrNotAuthorized))
	}

	// fail: SetAsset: unknown Oracle
	{
		asset := types.Asset{
			AssetCode: "testa_testb",
			Oracles:   []string{mockOracleAddr.String()},
		}

		err := keeper.SetAsset(ctx, types.NewMsgSetAsset(s.nominee, asset.AssetCode, asset.Decimals, mockOracleAddr))
		s.Require().Error(err)
		s.Require().True(errors.Is(err, types.ErrOracleNotFound))
	}

	// ok: SetAsset / GetAsset / GetAssets
	{
		oracle := types.Oracle{
			AccAddress:    mockOracleAddr.String(),
			Description:   "",
			PriceMaxBytes: 1,
		}
		s.Require().NoError(keeper.SetOracle(ctx, types.NewMsgSetOracle(s.nominee, mockOracleAddr, oracle.Description, oracle.PriceMaxBytes, oracle.PriceDecimals)))

		asset := types.Asset{
			AssetCode: "testa_testb",
			Decimals:  8,
			Oracles:   []string{mockOracleAddr.String()},
		}

		s.Require().Nil(keeper.GetAsset(ctx, asset.AssetCode))

		s.Require().NoError(keeper.SetAsset(ctx, types.NewMsgSetAsset(s.nominee, asset.AssetCode, asset.Decimals, mockOracleAddr)))

		s.Require().NotNil(keeper.GetAsset(ctx, asset.AssetCode))
		s.Require().Equal(asset, *keeper.GetAsset(ctx, asset.AssetCode))

		s.Require().ElementsMatch(append(initAssets, asset), keeper.GetAssets(ctx))
	}
}
