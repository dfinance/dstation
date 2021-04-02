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

func (s *KeeperMockVmTestSuite) TestDSServer_BlockMiddleware() {
	vmBlockPath := vm_grpc.VMAccessPath{
		Address: types.StdLibAddress,
		Path:    glav.BlockMetadataVector(),
	}

	// Skip genesis block
	s.app.BeginBlock()
	s.app.EndBlock()

	// ok
	{
		s.DoDSClientRequest(func(ctx context.Context, client ds_grpc.DSServiceClient) {
			resp, err := client.GetRaw(ctx, &ds_grpc.DSAccessPath{Address: vmBlockPath.Address, Path: vmBlockPath.Path})
			s.Require().NoError(err)
			s.Require().NotNil(resp)

			blockMeta := keeper.BlockMetadata{}
			s.Require().NoError(lcs.Unmarshal(resp.Blob, &blockMeta))

			sdkCtx := s.app.GetContext(true)
			s.Require().EqualValues(uint64(sdkCtx.BlockHeight()), blockMeta.Height)
		})
	}
}
