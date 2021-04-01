package keeper_test

import (
	"context"
	"testing"

	"github.com/dfinance/dvm-proto/go/ds_grpc"
	"github.com/stretchr/testify/require"

	"github.com/dfinance/dstation/pkg/mock"
	"github.com/dfinance/dstation/pkg/tests"
)

func TestVM_DSServer_GetRaw(t *testing.T) {
	app := tests.SetupDSimApp(tests.WithMockVM())
	defer app.TearDown()

	ctx, keeper := app.GetContext(), app.DnApp.VmKeeper

	// ok: non-existing
	{
		vmPath := mock.GetRandomVMAccessPath()

		DoDSClientRequest(t, app, func(ctx context.Context, client ds_grpc.DSServiceClient) {
			resp, err := client.GetRaw(ctx, &ds_grpc.DSAccessPath{Address: vmPath.Address, Path: vmPath.Path})
			require.NoError(t, err)
			require.NotNil(t, resp)

			require.Equal(t, ds_grpc.DSRawResponse_NO_DATA, resp.ErrorCode)
			require.NotEmpty(t, resp.ErrorMessage)
			require.Nil(t, resp.Blob)
		})
	}

	// ok
	{
		vmPath, writeSetData := mock.GetRandomVMAccessPath(), mock.GetRandomBytes(8)
		keeper.SetValue(ctx, vmPath, writeSetData)

		DoDSClientRequest(t, app, func(ctx context.Context, client ds_grpc.DSServiceClient) {
			resp, err := client.GetRaw(ctx, &ds_grpc.DSAccessPath{Address: vmPath.Address, Path: vmPath.Path})
			require.NoError(t, err)
			require.NotNil(t, resp)

			require.Equal(t, ds_grpc.DSRawResponse_NONE, resp.ErrorCode)
			require.Empty(t, resp.ErrorMessage)
			require.Equal(t, writeSetData, resp.Blob)
		})
	}
}
