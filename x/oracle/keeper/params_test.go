package keeper_test

import "github.com/dfinance/dstation/pkg/tests"

func (s *KeeperTestSuite) TestParams() {
	ctx, keeper := s.ctx, s.keeper

	mockNomineeAddr, _, _ := tests.GenAccAddress()
	initParams := keeper.GetParams(ctx)

	// Modify params
	updParams := initParams
	updParams.Nominees = append(updParams.Nominees, mockNomineeAddr.String())
	updParams.PostPrice.ReceivedAtDiffInS += 1
	keeper.SetParams(ctx, updParams)

	// Check params updated
	s.Require().Equal(updParams, keeper.GetParams(ctx))

	// Rollback values
	keeper.SetParams(ctx, initParams)
}
