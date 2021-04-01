package keeper_test

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/dfinance/dstation/pkg/mock"
	"github.com/dfinance/dstation/pkg/tests"
	"github.com/dfinance/dstation/x/vm/types"
)

func TestVM_Genesis(t *testing.T) {
	app := tests.SetupDSimApp(tests.WithMockVM())
	defer app.TearDown()

	ctx, keeper := app.GetContext(), app.DnApp.VmKeeper

	genStateExpected := types.DefaultGenesisState()
	genStateAppendix := types.GenesisState{
		WriteSet: make([]types.GenesisState_WriteOp, 0),
	}

	for i := 0; i < 5; i++ {
		vmPath, writeSetData := mock.GetRandomVMAccessPath(), mock.GetRandomBytes(20)
		writeOp := types.GenesisState_WriteOp{
			Address: hex.EncodeToString(vmPath.Address),
			Path:    hex.EncodeToString(vmPath.Path),
			Value:   hex.EncodeToString(writeSetData),
		}

		genStateAppendix.WriteSet = append(genStateAppendix.WriteSet, writeOp)
		genStateExpected.WriteSet = append(genStateExpected.WriteSet, writeOp)
	}

	// ok: Add writeSets to an existing default genesis
	{
		keeper.InitGenesis(ctx, &genStateAppendix)
		require.ElementsMatch(t, genStateExpected.WriteSet, keeper.ExportGenesis(ctx).WriteSet)
	}
}
