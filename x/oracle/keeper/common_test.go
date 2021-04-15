package keeper_test

import (
	"encoding/json"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dfinance/dstation/pkg/tests"
	"github.com/dfinance/dstation/x/oracle/types"
	"github.com/stretchr/testify/suite"

	"github.com/dfinance/dstation/x/oracle/keeper"
)

type KeeperTestSuite struct {
	suite.Suite

	app    *tests.DSimApp
	ctx    sdk.Context
	keeper keeper.Keeper
	//
	nominee sdk.AccAddress
	oracles []types.Oracle
	assets  []types.Asset
}

func (s *KeeperTestSuite) SetupSuite() {
	// Genesis state params and values
	nomineeAddr, _, _ := tests.GenAccAddress()
	oracle1Addr, _, _ := tests.GenAccAddress()
	oracle2Addr, _, _ := tests.GenAccAddress()
	oracle3Addr, _, _ := tests.GenAccAddress()
	oracle4Addr, _, _ := tests.GenAccAddress()
	oracles := []types.Oracle{
		{
			AccAddress:    oracle1Addr.String(),
			Description:   "mock_oracle1",
			PriceMaxBytes: 8,
			PriceDecimals: 8,
		},
		{
			AccAddress:    oracle2Addr.String(),
			Description:   "mock_oracle2",
			PriceMaxBytes: 8,
			PriceDecimals: 8,
		},
		{
			AccAddress:    oracle3Addr.String(),
			Description:   "mock_oracle3",
			PriceMaxBytes: 8,
			PriceDecimals: 8,
		},
		{
			AccAddress:    oracle4Addr.String(),
			Description:   "mock_oracle4",
			PriceMaxBytes: 8,
			PriceDecimals: 8,
		},
	}
	assets := []types.Asset{
		{
			AssetCode: "btc_usdt",
			Oracles:   []string{oracle1Addr.String()},
			Decimals:  6,
		},
		{
			AssetCode: "eth_usdt",
			Oracles:   []string{oracle2Addr.String(), oracle3Addr.String(), oracle4Addr.String()},
			Decimals:  8,
		},
	}

	genStateSetter := func(prevStateBz json.RawMessage) json.RawMessage {
		var curState types.GenesisState
		types.ModuleCdc.MustUnmarshalJSON(prevStateBz, &curState)

		curState.Params.Nominees = append(curState.Params.Nominees, nomineeAddr.String())
		curState.Oracles = oracles
		curState.Assets = assets

		curStateBz := types.ModuleCdc.MustMarshalJSON(&curState)

		return curStateBz
	}

	// Init the SimApp
	s.app = tests.SetupDSimApp(tests.WithCustomGenesisState(types.ModuleName, genStateSetter))
	s.ctx = s.app.GetContext(false)
	s.keeper = s.app.DnApp.OracleKeeper
	s.nominee = nomineeAddr
	s.oracles = oracles
	s.assets = assets
}

func (s *KeeperTestSuite) TearDownSuite() {
	s.app.TearDown()
}

func (s *KeeperTestSuite) SetupTest() {
	s.ctx = s.app.GetContext(false)
}

func TestOracleKeeper(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(KeeperTestSuite))
}
