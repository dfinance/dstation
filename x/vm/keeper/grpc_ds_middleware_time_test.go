package keeper_test

import (
	"context"

	"github.com/dfinance/glav"
	"github.com/dfinance/lcs"

	"github.com/dfinance/dstation/pkg/types/dvm"
	"github.com/dfinance/dstation/x/vm/keeper"
	"github.com/dfinance/dstation/x/vm/types"
)

func (s *KeeperMockVmTestSuite) TestDSServer_TimeMiddleware() {
	vmTimePath := dvm.VMAccessPath{
		Address: types.StdLibAddress,
		Path:    glav.TimeMetadataVector(),
	}

	// ok
	{
		s.DoDSClientRequest(func(ctx context.Context, client dvm.DSServiceClient) {
			resp, err := client.GetRaw(ctx, &dvm.DSAccessPath{Address: vmTimePath.Address, Path: vmTimePath.Path})
			s.Require().NoError(err)
			s.Require().NotNil(resp)

			timeMeta := keeper.TimeMetadata{}
			s.Require().NoError(lcs.Unmarshal(resp.Blob, &timeMeta))

			sdkCtx := s.app.GetContext(true)
			s.Require().Equal(uint64(sdkCtx.BlockTime().Unix()), timeMeta.Seconds)
		})
	}
}
