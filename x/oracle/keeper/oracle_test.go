package keeper_test

import (
	"errors"

	"github.com/dfinance/dstation/pkg/tests"
	"github.com/dfinance/dstation/x/oracle/types"
)

func (s *KeeperTestSuite) TestOracle() {
	ctx, keeper := s.ctx, s.keeper
	mockOracleAddr, _, _ := tests.GenAccAddress()

	initOracles := keeper.GetOracles(ctx)

	// fail: SetOracle: not a nominee
	{
		oracle := types.Oracle{
			AccAddress:    mockOracleAddr.String(),
			PriceMaxBytes: 1,
		}

		err := keeper.SetOracle(ctx, types.NewMsgSetOracle(mockOracleAddr, mockOracleAddr, oracle.Description, oracle.PriceMaxBytes, oracle.PriceDecimals))
		s.Require().Error(err)
		s.Require().True(errors.Is(err, types.ErrNotAuthorized))
	}

	// ok: SetOracle / GetOracle / GetOracles
	{
		oracle := types.Oracle{
			AccAddress:    mockOracleAddr.String(),
			Description:   "some_desc",
			PriceMaxBytes: 1,
		}

		s.Require().Nil(keeper.GetOracle(ctx, oracle.AccAddress))

		s.Require().NoError(keeper.SetOracle(ctx, types.NewMsgSetOracle(s.nominee, mockOracleAddr, oracle.Description, oracle.PriceMaxBytes, oracle.PriceDecimals)))

		s.Require().NotNil(keeper.GetOracle(ctx, oracle.AccAddress))
		s.Require().Equal(oracle, *keeper.GetOracle(ctx, oracle.AccAddress))

		s.Require().ElementsMatch(append(initOracles, oracle), keeper.GetOracles(ctx))
	}
}
