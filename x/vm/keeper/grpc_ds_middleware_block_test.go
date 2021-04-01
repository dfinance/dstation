package keeper_test

import (
	"context"
	"testing"

	"github.com/dfinance/dvm-proto/go/ds_grpc"
	"github.com/dfinance/dvm-proto/go/vm_grpc"
	"github.com/dfinance/glav"
	"github.com/dfinance/lcs"
	"github.com/stretchr/testify/require"

	"github.com/dfinance/dstation/pkg/tests"
	"github.com/dfinance/dstation/x/vm/keeper"
	"github.com/dfinance/dstation/x/vm/types"
)

func TestVM_DSServer_BlockMiddleware(t *testing.T) {
	app := tests.SetupDSimApp(tests.WithMockVM())
	defer app.TearDown()

	vmBlockPath := vm_grpc.VMAccessPath{
		Address: types.StdLibAddress,
		Path:    glav.BlockMetadataVector(),
	}

	// Skip genesis block
	app.BeginBlock()
	app.EndBlock()

	// ok
	{
		DoDSClientRequest(t, app, func(ctx context.Context, client ds_grpc.DSServiceClient) {
			resp, err := client.GetRaw(ctx, &ds_grpc.DSAccessPath{Address: vmBlockPath.Address, Path: vmBlockPath.Path})
			require.NoError(t, err)
			require.NotNil(t, resp)

			blockMeta := keeper.BlockMetadata{}
			require.NoError(t, lcs.Unmarshal(resp.Blob, &blockMeta))

			require.EqualValues(t, 1, blockMeta.Height)
		})
	}
}
