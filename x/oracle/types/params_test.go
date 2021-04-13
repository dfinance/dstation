package types_test

import (
	"testing"

	"github.com/dfinance/dstation/pkg/tests"
	"github.com/dfinance/dstation/x/oracle/types"
	"github.com/stretchr/testify/require"
)

func TestOracle_Params(t *testing.T) {
	mockAddr, _, _ := tests.GenAccAddress()

	// fail: invalid nominee
	{
		params := types.Params{
			Nominees:  []string{"abc"},
			PostPrice: types.Params_PostPriceParams{},
		}
		require.Error(t, params.ValidateBasic())
	}

	// fail: duplicated nominee
	{
		params := types.Params{
			Nominees:  []string{mockAddr.String(), mockAddr.String()},
			PostPrice: types.Params_PostPriceParams{},
		}
		require.Error(t, params.ValidateBasic())
	}

	// ok
	{
		params := types.Params{
			Nominees: []string{mockAddr.String()},
			PostPrice: types.Params_PostPriceParams{
				ReceivedAtDiffInS: 0,
			},
		}
		require.NoError(t, params.ValidateBasic())
	}

	// ok
	{
		params := types.Params{
			Nominees:  nil,
			PostPrice: types.Params_PostPriceParams{},
		}
		require.NoError(t, params.ValidateBasic())
	}
}
