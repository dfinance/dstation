package keeper_test

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/dfinance/dstation/pkg/mock"
	"github.com/dfinance/dstation/pkg/tests"
	"github.com/dfinance/dstation/x/vm/types"
)

func TestVMKeeper_Genesis(t *testing.T) {
	app := tests.SetupDSimApp()
	defer app.TearDown()

	ctx, keeper := app.GetContext(false), app.DnApp.VmKeeper

	genStateExpected, genStateAppendix := types.DefaultGenesisState(), types.GenesisState{}

	// Add genesis data
	{
		// VM writeSets
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
	}


	// ok: Init / Export check
	{
		keeper.InitGenesis(ctx, &genStateAppendix)
		genStateExported := keeper.ExportGenesis(ctx)

		require.ElementsMatch(t, genStateExpected.WriteSet, genStateExported.WriteSet)
	}

	// ok: DelPool module account check
	{
		delPoolAcc := app.DnApp.AccountKeeper.GetModuleAccount(ctx, types.DelPoolName)
		require.NotNil(t, delPoolAcc)

		delPoolCoins := app.DnApp.BankKeeper.GetAllBalances(ctx, delPoolAcc.GetAddress())
		require.True(t, delPoolCoins.IsZero())
	}
}
