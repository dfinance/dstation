package keeper_test

import (
	"context"

	"github.com/dfinance/dvm-proto/go/ds_grpc"
	"github.com/dfinance/dvm-proto/go/vm_grpc"
	"github.com/dfinance/glav"
	"github.com/dfinance/lcs"

	"github.com/dfinance/dstation/x/vm/keeper"
	"github.com/dfinance/dstation/x/vm/types"
)

func (s *KeeperMockVmTestSuite) TestDSServer_TimeMiddleware() {
	vmTimePath := vm_grpc.VMAccessPath{
		Address: types.StdLibAddress,
		Path:    glav.TimeMetadataVector(),
	}

	// ok
	{
		s.DoDSClientRequest(func(ctx context.Context, client ds_grpc.DSServiceClient) {
			resp, err := client.GetRaw(ctx, &ds_grpc.DSAccessPath{Address: vmTimePath.Address, Path: vmTimePath.Path})
			s.Require().NoError(err)
			s.Require().NotNil(resp)

			timeMeta := keeper.TimeMetadata{}
			s.Require().NoError(lcs.Unmarshal(resp.Blob, &timeMeta))

			sdkCtx := s.app.GetContext(true)
			s.Require().Equal(uint64(sdkCtx.BlockTime().Unix()), timeMeta.Seconds)
		})
	}
}
