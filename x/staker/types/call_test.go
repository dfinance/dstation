package types_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dfinance/dstation/pkg/tests"
	"github.com/dfinance/dstation/x/staker/types"
	"github.com/stretchr/testify/require"
)

func TestStaker_Call(t *testing.T) {
	mockAddr, _, _ := tests.GenAccAddress()

	call := types.Call{
		Id:        sdk.ZeroUint(),
		UniqueId: "unique_id",
		Nominee:   mockAddr.String(),
		Address:   mockAddr.String(),
		Type:      types.Call_DEPOSIT,
		Amount:    sdk.NewCoins(sdk.NewCoin("test", sdk.OneInt())),
		Timestamp: time.Now(),
	}

	// ok
	{
		require.NoError(t, call.Validate())
	}

	// fail: id: nil
	{
		c := call
		c.Id = sdk.Uint{}
		require.Error(t, c.Validate())
	}

	// fail: unique_id: empty
	{
		c := call
		c.UniqueId = ""
		require.Error(t, c.Validate())
	}

	// fail: nominee: invalid
	{
		c := call
		c.Nominee = "abc"
		require.Error(t, c.Validate())
	}

	// fail: address: invalid
	{
		c := call
		c.Address = ""
		require.Error(t, c.Validate())
	}

	// fail: type: invalid
	{
		c := call
		c.Type = 100
		require.Error(t, c.Validate())
	}

	// fail: amount: empty
	{
		c := call
		c.Amount = sdk.Coins{}
		require.Error(t, c.Validate())
	}

	// fail: amount: invalid
	{
		c := call
		c.Amount = sdk.Coins{sdk.Coin{
			Denom:  "@",
			Amount: sdk.OneInt(),
		}}
		require.Error(t, c.Validate())
	}

	// fail: timestamp: empty
	{
		c := call
		c.Timestamp = time.Time{}
		require.Error(t, c.Validate())
	}
}
