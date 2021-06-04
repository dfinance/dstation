package types_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dfinance/dstation/pkg/tests"
	"github.com/dfinance/dstation/x/staker/types"
	"github.com/stretchr/testify/require"
)

func TestStaker_Genesis(t *testing.T) {
	mockAddr, _, _ := tests.GenAccAddress()

	lastCallId := sdk.NewUint(1)
	gen := types.GenesisState{
		Params:     types.Params{
			Nominees: []string{mockAddr.String()},
		},
		LastCallId: &lastCallId,
		Calls:      []types.Call{
			{
				Id:        sdk.NewUint(0),
				Nominee:   mockAddr.String(),
				Address:   mockAddr.String(),
				Type:      types.Call_DEPOSIT,
				Amount:    sdk.Coins{sdk.Coin{Denom: "test", Amount: sdk.OneInt()}},
				Timestamp: time.Now(),
			},
			{
				Id:        sdk.NewUint(1),
				Nominee:   mockAddr.String(),
				Address:   mockAddr.String(),
				Type:      types.Call_WITHDRAW,
				Amount:    sdk.Coins{sdk.Coin{Denom: "test", Amount: sdk.OneInt()}},
				Timestamp: time.Now(),
			},
		},
	}

	// ok
	{
		require.NoError(t, gen.Validate())
	}

	// fail: LastCallId: nil
	{
		g := gen
		nilUint := sdk.Uint{}
		g.LastCallId = &nilUint
		require.Error(t, g.Validate())
	}

	// fail: Calls: duplicate
	{
		g := gen
		g.Calls[1].Id = sdk.NewUint(0)
		require.Error(t, g.Validate())
	}

	// fail: Calls: invalid
	{
		g := gen
		g.Calls[1].Amount = sdk.Coins{}
		require.Error(t, g.Validate())
	}
}
