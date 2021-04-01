package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/dfinance/dstation/pkg/mock"
	"github.com/dfinance/dstation/pkg/tests"
)

func TestVM_VMStorage(t *testing.T) {
	app := tests.SetupDSimApp(tests.WithMockVM())
	defer app.TearDown()

	ctx, keeper := app.GetContext(), app.DnApp.VmKeeper
	vmPath, writeSetData := mock.GetRandomVMAccessPath(), mock.GetRandomBytes(12)

	// ok: HasValue, GetValue: non-existing
	{
		require.False(t, keeper.HasValue(ctx, vmPath))
		require.Nil(t, keeper.GetValue(ctx, vmPath))
	}

	// ok: SetValue
	{
		keeper.SetValue(ctx, vmPath, writeSetData)

		require.True(t, keeper.HasValue(ctx, vmPath))
		require.NotNil(t, keeper.GetValue(ctx, vmPath))
		require.Equal(t, writeSetData, keeper.GetValue(ctx, vmPath))
	}

	// ok: DelValue
	{
		keeper.DelValue(ctx, vmPath)

		require.False(t, keeper.HasValue(ctx, vmPath))
		require.Nil(t, keeper.GetValue(ctx, vmPath))
	}
}
