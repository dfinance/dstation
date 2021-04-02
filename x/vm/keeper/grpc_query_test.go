package keeper_test

import (
	"context"
	"encoding/hex"

	"github.com/dfinance/dstation/pkg/mock"
	"github.com/dfinance/dstation/x/vm/types"
)

func (s *KeeperMockVmTestSuite) TestQuerier() {
	ctx, keeper, client := s.ctx, s.keeper, s.queryClient
	vmPath, writeSetData := mock.GetRandomVMAccessPath(), mock.GetRandomBytes(12)

	// fail: invalid address
	{
		resp, err := client.Data(context.Background(), &types.QueryDataRequest{
			Address: "abc",
			Path:    hex.EncodeToString(vmPath.Path),
		})
		s.Require().Error(err)
		s.Require().Nil(resp)
	}

	// fail: invalid path
	{
		resp, err := client.Data(context.Background(), &types.QueryDataRequest{
			Address: hex.EncodeToString(vmPath.Address),
			Path:    "@$@",
		})
		s.Require().Error(err)
		s.Require().Nil(resp)
	}

	// ok: non-existing
	{
		resp, err := client.Data(context.Background(), &types.QueryDataRequest{
			Address: hex.EncodeToString(vmPath.Address),
			Path:    hex.EncodeToString(vmPath.Path),
		})
		s.Require().NoError(err)
		s.Require().NotNil(resp)
		s.Require().Empty(resp.Value)
	}

	// ok: existing
	{
		keeper.SetValue(ctx, vmPath, writeSetData)

		resp, err := client.Data(context.Background(), &types.QueryDataRequest{
			Address: hex.EncodeToString(vmPath.Address),
			Path:    hex.EncodeToString(vmPath.Path),
		})
		s.Require().NoError(err)
		s.Require().NotNil(resp)
		s.Require().Equal(hex.EncodeToString(writeSetData), resp.Value)
	}
}
