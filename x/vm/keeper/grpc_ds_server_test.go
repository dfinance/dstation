package keeper_test

import (
	"context"

	"github.com/dfinance/dstation/pkg/mock"
	"github.com/dfinance/dstation/pkg/types/dvm"
)

func (s *KeeperMockVmTestSuite) TestDSServer_GetRaw() {
	ctx, keeper := s.ctx, s.keeper

	// ok: non-existing
	{
		vmPath := mock.GetRandomVMAccessPath()

		s.DoDSClientRequest(func(ctx context.Context, client dvm.DSServiceClient) {
			resp, err := client.GetRaw(ctx, &dvm.DSAccessPath{Address: vmPath.Address, Path: vmPath.Path})
			s.Require().NoError(err)
			s.Require().NotNil(resp)

			s.Require().Equal(dvm.ErrorCode_NO_DATA, resp.ErrorCode)
			s.Require().NotEmpty(resp.ErrorMessage)
			s.Require().Nil(resp.Blob)
		})
	}

	// ok
	{
		vmPath, writeSetData := mock.GetRandomVMAccessPath(), mock.GetRandomBytes(8)
		keeper.SetValue(ctx, vmPath, writeSetData)

		s.DoDSClientRequest(func(ctx context.Context, client dvm.DSServiceClient) {
			resp, err := client.GetRaw(ctx, &dvm.DSAccessPath{Address: vmPath.Address, Path: vmPath.Path})
			s.Require().NoError(err)
			s.Require().NotNil(resp)

			s.Require().Equal(dvm.ErrorCode_NONE, resp.ErrorCode)
			s.Require().Empty(resp.ErrorMessage)
			s.Require().Equal(writeSetData, resp.Blob)
		})
	}
}
