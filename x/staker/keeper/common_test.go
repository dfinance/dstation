package keeper_test

import (
	"encoding/json"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dfinance/dstation/pkg/tests"
	"github.com/dfinance/dstation/x/staker/keeper"
	"github.com/dfinance/dstation/x/staker/types"
	"github.com/stretchr/testify/suite"
)

type KeeperTestSuite struct {
	suite.Suite

	app    *tests.DSimApp
	ctx    sdk.Context
	keeper keeper.Keeper
	//
	nominee sdk.AccAddress
}

func (s *KeeperTestSuite) SetupSuite() {
	// Genesis state params and values
	nomineeAddr, _, _ := tests.GenAccAddress()

	genStateSetter := func(prevStateBz json.RawMessage) json.RawMessage {
		var curState types.GenesisState
		types.ModuleCdc.MustUnmarshalJSON(prevStateBz, &curState)

		curState.Params.Nominees = append(curState.Params.Nominees, nomineeAddr.String())

		curStateBz := types.ModuleCdc.MustMarshalJSON(&curState)

		return curStateBz
	}

	// Init the SimApp
	s.app = tests.SetupDSimApp(tests.WithCustomGenesisState(types.ModuleName, genStateSetter))
	s.ctx = s.app.GetContext(false)
	s.keeper = s.app.DnApp.StakerKeeper
	s.nominee = nomineeAddr
}

func (s *KeeperTestSuite) TearDownSuite() {
	s.app.TearDown()
}

func (s *KeeperTestSuite) SetupTest() {
	s.ctx = s.app.GetContext(false)
}

func (s *KeeperTestSuite) CheckCall(ctx sdk.Context, call types.Call, expId uint64, expAccAddr sdk.AccAddress, expType types.Call_CallType, expAmt sdk.Coins) {
	s.Require().Equal(expId, call.Id.Uint64(), "id")
	s.Require().Equal(s.nominee.String(), call.Nominee, "nominee")
	s.Require().Equal(expAccAddr.String(), call.Address, "address")
	s.Require().Equal(expType, call.Type, "type")
	s.Require().True(expAmt.IsEqual(call.Amount), "amount")
	s.Require().True(ctx.BlockTime().Equal(call.Timestamp), "timestamp")
}

func TestStakerKeeper(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(KeeperTestSuite))
}
