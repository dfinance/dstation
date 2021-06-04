package keeper_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dfinance/dstation/pkg/tests"
	"github.com/dfinance/dstation/x/staker/types"
	"github.com/stretchr/testify/require"
)

func TestStakerKeeper_Genesis(t *testing.T) {
	app := tests.SetupDSimApp()
	defer app.TearDown()

	ctx, keeper := app.GetContext(false), app.DnApp.StakerKeeper
	mockAddr, _, _ := tests.GenAccAddress()

	lastCallId := sdk.NewUint(0)
	stateInit := &types.GenesisState{
		Params: types.Params{
			Nominees: []string{mockAddr.String()},
		},
		LastCallId: &lastCallId,
		Calls: []types.Call{
			{
				Id:        sdk.NewUint(0),
				Nominee:   mockAddr.String(),
				Address:   mockAddr.String(),
				Type:      types.Call_DEPOSIT,
				Amount:    sdk.Coins{sdk.Coin{Denom: "test", Amount: sdk.OneInt()}},
				Timestamp: time.Now(),
			},
		},
	}

	// import
	keeper.InitGenesis(ctx, stateInit)

	// export
	{
		stateExport := keeper.ExportGenesis(ctx)
		require.Equal(t, stateInit.Params, stateExport.Params)
		require.True(t, stateInit.LastCallId.Equal(*stateExport.LastCallId))
		require.Len(t, stateExport.Calls, len(stateInit.Calls))
	}
}
