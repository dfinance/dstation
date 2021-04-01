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

func TestVM_DSServer_TimeMiddleware(t *testing.T) {
	app := tests.SetupDSimApp(tests.WithMockVM())
	defer app.TearDown()

	vmTimePath := vm_grpc.VMAccessPath{
		Address: types.StdLibAddress,
		Path:    glav.TimeMetadataVector(),
	}

	// ok
	{
		DoDSClientRequest(t, app, func(ctx context.Context, client ds_grpc.DSServiceClient) {
			resp, err := client.GetRaw(ctx, &ds_grpc.DSAccessPath{Address: vmTimePath.Address, Path: vmTimePath.Path})
			require.NoError(t, err)
			require.NotNil(t, resp)

			timeMeta := keeper.TimeMetadata{}
			require.NoError(t, lcs.Unmarshal(resp.Blob, &timeMeta))

			require.NotZero(t, timeMeta.Seconds)
		})
	}
}
