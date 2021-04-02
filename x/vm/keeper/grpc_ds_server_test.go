package keeper_test

import (
	"context"

	"github.com/dfinance/dvm-proto/go/ds_grpc"

	"github.com/dfinance/dstation/pkg/mock"
)

func (s *KeeperMockVmTestSuite) TestDSServer_GetRaw() {
	ctx, keeper := s.ctx, s.keeper

	// ok: non-existing
	{
		vmPath := mock.GetRandomVMAccessPath()

		s.DoDSClientRequest(func(ctx context.Context, client ds_grpc.DSServiceClient) {
			resp, err := client.GetRaw(ctx, &ds_grpc.DSAccessPath{Address: vmPath.Address, Path: vmPath.Path})
			s.Require().NoError(err)
			s.Require().NotNil(resp)

			s.Require().Equal(ds_grpc.DSRawResponse_NO_DATA, resp.ErrorCode)
			s.Require().NotEmpty(resp.ErrorMessage)
			s.Require().Nil(resp.Blob)
		})
	}

	// ok
	{
		vmPath, writeSetData := mock.GetRandomVMAccessPath(), mock.GetRandomBytes(8)
		keeper.SetValue(ctx, vmPath, writeSetData)

		s.DoDSClientRequest(func(ctx context.Context, client ds_grpc.DSServiceClient) {
			resp, err := client.GetRaw(ctx, &ds_grpc.DSAccessPath{Address: vmPath.Address, Path: vmPath.Path})
			s.Require().NoError(err)
			s.Require().NotNil(resp)

			s.Require().Equal(ds_grpc.DSRawResponse_NONE, resp.ErrorCode)
			s.Require().Empty(resp.ErrorMessage)
			s.Require().Equal(writeSetData, resp.Blob)
		})
	}
}
