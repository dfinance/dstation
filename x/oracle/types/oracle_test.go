package types_test

import (
	"testing"

	"github.com/dfinance/dstation/pkg/tests"
	"github.com/dfinance/dstation/x/oracle/types"
	"github.com/stretchr/testify/require"
)

func TestOracle_Oracle(t *testing.T) {
	mockAddr, _, _ := tests.GenAccAddress()

	// fail: invalid address
	{
		oracle := types.Oracle{
			AccAddress:    "abc",
			PriceMaxBytes: 1,
		}
		require.Error(t, oracle.Validate())
	}

	// fail: invalid bytes number
	{
		oracle := types.Oracle{
			AccAddress:    mockAddr.String(),
			PriceMaxBytes: 0,
		}
		require.Error(t, oracle.Validate())
	}

	// ok
	{
		oracle := types.Oracle{
			AccAddress:    mockAddr.String(),
			PriceMaxBytes: 1,
		}
		require.NoError(t, oracle.Validate())
	}
}
