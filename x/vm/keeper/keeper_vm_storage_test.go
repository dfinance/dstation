package keeper_test

import (
	"github.com/dfinance/dstation/pkg/mock"
)

func (s *KeeperMockVmTestSuite) TestVMStorage() {
	ctx, keeper := s.ctx, s.keeper
	vmPath, writeSetData := mock.GetRandomVMAccessPath(), mock.GetRandomBytes(12)

	// ok: HasValue, GetValue: non-existing
	{
		s.Require().False(keeper.HasValue(ctx, vmPath))
		s.Require().Nil(keeper.GetValue(ctx, vmPath))
	}

	// ok: SetValue
	{
		keeper.SetValue(ctx, vmPath, writeSetData)

		s.Require().True(keeper.HasValue(ctx, vmPath))
		s.Require().NotNil(keeper.GetValue(ctx, vmPath))
		s.Require().Equal(writeSetData, keeper.GetValue(ctx, vmPath))
	}

	// ok: DelValue
	{
		keeper.DelValue(ctx, vmPath)

		s.Require().False(keeper.HasValue(ctx, vmPath))
		s.Require().Nil(keeper.GetValue(ctx, vmPath))
	}
}
