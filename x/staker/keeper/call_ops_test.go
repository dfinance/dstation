package keeper_test

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	dnConfig "github.com/dfinance/dstation/cmd/dstation/config"
	"github.com/dfinance/dstation/x/staker/types"
)

func (s *KeeperTestSuite) TestOperations() {
	ctx, keeper, bankKeeper, accKeeper := s.ctx, s.keeper, s.app.DnApp.BankKeeper, s.app.DnApp.AccountKeeper
	moduleAcc := accKeeper.GetModuleAddress(types.ModuleName)
	acc1, _ := s.app.AddAccount(ctx)
	acc2, _ := s.app.AddAccount(ctx)

	srcMeta := types.CallSourceMeta{
		EthAddress: "0x89205A3A3b2A69De6Dbf7f01ED13B2108B2c43e7",
		ChainId:    "Ethereum",
	}

	// ok: GetCallById, GetAddressCalls: non-existing
	{
		s.Require().Nil(keeper.GetCallById(ctx, sdk.NewUint(1)))
		s.Require().Nil(keeper.GetCallByUniqueId(ctx, "1"))
		s.Require().Empty(keeper.GetAddressCalls(ctx, acc1.GetAddress()))
	}

	// fail: not a nominee
	{
		msg := types.NewMsgDepositCall("0", acc1.GetAddress(), acc1.GetAddress(), srcMeta.EthAddress, srcMeta.ChainId, sdk.Coins{sdk.NewCoin(dnConfig.MainDenom, sdk.OneInt())})
		s.Require().NoError(msg.ValidateBasic())

		_, err := keeper.Deposit(ctx, msg)
		s.Require().Error(err)
		s.Require().True(types.ErrNotAuthorized.Is(err))
	}

	// ok: Deposit: acc1
	depositAcc1_1 := sdk.NewCoins(sdk.NewCoin(dnConfig.MainDenom, sdk.NewInt(100)))
	depositAcc1_2 := sdk.NewCoins(sdk.NewCoin(dnConfig.MainDenom, sdk.NewInt(50)))
	{
		msg := types.NewMsgDepositCall("0", s.nominee, acc1.GetAddress(), srcMeta.EthAddress, srcMeta.ChainId, depositAcc1_1)
		s.Require().NoError(msg.ValidateBasic())

		call, err := keeper.Deposit(ctx, msg)
		s.Require().NoError(err)
		s.CheckCall(ctx, call, 0, "0", acc1.GetAddress(), types.Call_DEPOSIT, srcMeta, depositAcc1_1)

		expBalance := depositAcc1_1
		s.Require().True(expBalance.IsEqual(bankKeeper.GetAllBalances(ctx, acc1.GetAddress())))
		s.Require().True(bankKeeper.GetAllBalances(ctx, moduleAcc).IsZero())
	}
	{
		msg := types.NewMsgDepositCall("1", s.nominee, acc1.GetAddress(), srcMeta.EthAddress, srcMeta.ChainId, depositAcc1_2)
		s.Require().NoError(msg.ValidateBasic())

		call, err := keeper.Deposit(ctx, msg)
		s.Require().NoError(err)
		s.CheckCall(ctx, call, 1, "1", acc1.GetAddress(), types.Call_DEPOSIT, srcMeta, depositAcc1_2)

		expBalance := depositAcc1_1.Add(depositAcc1_2...)
		s.Require().True(expBalance.IsEqual(bankKeeper.GetAllBalances(ctx, acc1.GetAddress())))
		s.Require().True(bankKeeper.GetAllBalances(ctx, moduleAcc).IsZero())
	}

	// ok: Withdraw: acc1
	withdrawAcc1_1 := sdk.NewCoins(sdk.NewCoin(dnConfig.MainDenom, sdk.NewInt(25)))
	{
		msg := types.NewMsgWithdrawCall("2", s.nominee, acc1.GetAddress(), srcMeta.EthAddress, srcMeta.ChainId, withdrawAcc1_1)
		s.Require().NoError(msg.ValidateBasic())

		call, err := keeper.Withdraw(ctx, msg)
		s.Require().NoError(err)
		s.CheckCall(ctx, call, 2, "2", acc1.GetAddress(), types.Call_WITHDRAW, srcMeta, withdrawAcc1_1)

		expBalance := depositAcc1_1.Add(depositAcc1_2...).Sub(withdrawAcc1_1)
		s.Require().True(expBalance.IsEqual(bankKeeper.GetAllBalances(ctx, acc1.GetAddress())))
		s.Require().True(bankKeeper.GetAllBalances(ctx, moduleAcc).IsZero())
	}

	// ok: GetCallById, GetCallByUniqueId, GetAddressCalls: acc1
	{
		call1ById := keeper.GetCallById(ctx, sdk.NewUint(0))
		call2ById := keeper.GetCallById(ctx, sdk.NewUint(1))
		call3ById := keeper.GetCallById(ctx, sdk.NewUint(2))
		s.Require().NotNil(call1ById)
		s.Require().NotNil(call2ById)
		s.Require().NotNil(call3ById)

		call1ByUniqueId := keeper.GetCallByUniqueId(ctx, "0")
		call2ByUniqueId := keeper.GetCallByUniqueId(ctx, "1")
		call3ByUniqueId := keeper.GetCallByUniqueId(ctx, "2")
		s.Require().NotNil(call1ByUniqueId)
		s.Require().NotNil(call2ByUniqueId)
		s.Require().NotNil(call3ByUniqueId)

		s.CheckCall(ctx, *call1ById, 0, "0", acc1.GetAddress(), types.Call_DEPOSIT, srcMeta, depositAcc1_1)
		s.CheckCall(ctx, *call1ByUniqueId, 0, "0", acc1.GetAddress(), types.Call_DEPOSIT, srcMeta, depositAcc1_1)
		s.CheckCall(ctx, *call2ById, 1, "1", acc1.GetAddress(), types.Call_DEPOSIT, srcMeta, depositAcc1_2)
		s.CheckCall(ctx, *call2ByUniqueId, 1, "1", acc1.GetAddress(), types.Call_DEPOSIT, srcMeta, depositAcc1_2)
		s.CheckCall(ctx, *call3ById, 2, "2", acc1.GetAddress(), types.Call_WITHDRAW, srcMeta, withdrawAcc1_1)
		s.CheckCall(ctx, *call3ByUniqueId, 2, "2", acc1.GetAddress(), types.Call_WITHDRAW, srcMeta, withdrawAcc1_1)

		calls := keeper.GetAddressCalls(ctx, acc1.GetAddress())
		s.Require().ElementsMatch([]types.Call{*call1ById, *call2ById, *call3ById}, calls)
	}

	// fail: Withdraw: acc2: no funds
	{
		msg := types.NewMsgWithdrawCall("3", s.nominee, acc2.GetAddress(), srcMeta.EthAddress, srcMeta.ChainId, sdk.Coins{sdk.Coin{Denom: "test", Amount: sdk.OneInt()}})
		s.Require().NoError(msg.ValidateBasic())

		_, err := keeper.Withdraw(ctx, msg)
		s.Require().Error(err)
	}

	// fail: Deposit: uniqueId already exists
	{
		msg := types.NewMsgDepositCall("2", s.nominee, acc1.GetAddress(), srcMeta.EthAddress, srcMeta.ChainId, depositAcc1_2)
		s.Require().NoError(msg.ValidateBasic())

		_, err := keeper.Deposit(ctx, msg)
		s.Require().Error(err)
		s.Require().True(errors.Is(err, types.ErrUniqueIdExists))
	}
}
