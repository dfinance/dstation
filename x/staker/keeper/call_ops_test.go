package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	dnConfig "github.com/dfinance/dstation/cmd/dstation/config"
	"github.com/dfinance/dstation/x/staker/types"
)

func (s *KeeperTestSuite) TestOperations() {
	ctx, keeper, bankKeeper, accKeeper := s.ctx, s.keeper, s.app.DnApp.BankKeeper, s.app.DnApp.AccountKeeper
	moduleAcc := accKeeper.GetModuleAddress(types.ModuleName)
	acc1, _ := s.app.AddAccount(ctx)
	acc2, _ := s.app.AddAccount(ctx)

	// ok: GetCall, GetAddressCalls: non-existing
	{
		s.Require().Nil(keeper.GetCall(ctx, sdk.NewUint(1)))
		s.Require().Empty(keeper.GetAddressCalls(ctx, acc1.GetAddress()))
	}

	// fail: not a nominee
	{
		msg := types.NewMsgDepositCall(acc1.GetAddress(), acc1.GetAddress(), sdk.Coins{sdk.NewCoin(dnConfig.MainDenom, sdk.OneInt())})
		s.Require().NoError(msg.ValidateBasic())

		_, err := keeper.Deposit(ctx, msg)
		s.Require().Error(err)
		s.Require().True(types.ErrNotAuthorized.Is(err))
	}

	// ok: Deposit: acc1
	depositAcc1_1 := sdk.NewCoins(sdk.NewCoin(dnConfig.MainDenom, sdk.NewInt(100)))
	depositAcc1_2 := sdk.NewCoins(sdk.NewCoin(dnConfig.MainDenom, sdk.NewInt(50)))
	{
		msg := types.NewMsgDepositCall(s.nominee, acc1.GetAddress(), depositAcc1_1)
		s.Require().NoError(msg.ValidateBasic())

		call, err := keeper.Deposit(ctx, msg)
		s.Require().NoError(err)
		s.CheckCall(ctx, call, 0, acc1.GetAddress(), types.Call_DEPOSIT, depositAcc1_1)

		expBalance := depositAcc1_1
		s.Require().True(expBalance.IsEqual(bankKeeper.GetAllBalances(ctx, acc1.GetAddress())))
		s.Require().True(bankKeeper.GetAllBalances(ctx, moduleAcc).IsZero())
	}
	{
		msg := types.NewMsgDepositCall(s.nominee, acc1.GetAddress(), depositAcc1_2)
		s.Require().NoError(msg.ValidateBasic())

		call, err := keeper.Deposit(ctx, msg)
		s.Require().NoError(err)
		s.CheckCall(ctx, call, 1, acc1.GetAddress(), types.Call_DEPOSIT, depositAcc1_2)

		expBalance := depositAcc1_1.Add(depositAcc1_2...)
		s.Require().True(expBalance.IsEqual(bankKeeper.GetAllBalances(ctx, acc1.GetAddress())))
		s.Require().True(bankKeeper.GetAllBalances(ctx, moduleAcc).IsZero())
	}

	// ok: Withdraw: acc1
	withdrawAcc1_1 := sdk.NewCoins(sdk.NewCoin(dnConfig.MainDenom, sdk.NewInt(25)))
	{
		msg := types.NewMsgWithdrawCall(s.nominee, acc1.GetAddress(), withdrawAcc1_1)
		s.Require().NoError(msg.ValidateBasic())

		call, err := keeper.Withdraw(ctx, msg)
		s.Require().NoError(err)
		s.CheckCall(ctx, call, 2, acc1.GetAddress(), types.Call_WITHDRAW, withdrawAcc1_1)

		expBalance := depositAcc1_1.Add(depositAcc1_2...).Sub(withdrawAcc1_1)
		s.Require().True(expBalance.IsEqual(bankKeeper.GetAllBalances(ctx, acc1.GetAddress())))
		s.Require().True(bankKeeper.GetAllBalances(ctx, moduleAcc).IsZero())
	}

	// ok: GetCall, GetAddressCalls: acc1
	{
		call1 := keeper.GetCall(ctx, sdk.NewUint(0))
		call2 := keeper.GetCall(ctx, sdk.NewUint(1))
		call3 := keeper.GetCall(ctx, sdk.NewUint(2))
		s.Require().NotNil(call1)
		s.Require().NotNil(call2)
		s.Require().NotNil(call3)
		s.CheckCall(ctx, *call1, 0, acc1.GetAddress(), types.Call_DEPOSIT, depositAcc1_1)
		s.CheckCall(ctx, *call2, 1, acc1.GetAddress(), types.Call_DEPOSIT, depositAcc1_2)
		s.CheckCall(ctx, *call3, 2, acc1.GetAddress(), types.Call_WITHDRAW, withdrawAcc1_1)

		calls := keeper.GetAddressCalls(ctx, acc1.GetAddress())
		s.Require().ElementsMatch([]types.Call{*call1, *call2, *call3}, calls)
	}

	// fail: Withdraw: acc2: no funds
	{
		msg := types.NewMsgWithdrawCall(s.nominee, acc2.GetAddress(), sdk.Coins{sdk.Coin{Denom: "test", Amount: sdk.OneInt()}})
		s.Require().NoError(msg.ValidateBasic())

		_, err := keeper.Withdraw(ctx, msg)
		s.Require().Error(err)
	}
}
