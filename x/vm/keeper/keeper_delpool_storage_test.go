package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/dfinance/dstation/pkg/tests"
	"github.com/dfinance/dstation/x/vm/types"
)

func (s *KeeperMockVmTestSuite) TestDelPoolStorage() {
	ctx, accKeeper, bankKeeper, keeper := s.ctx, s.app.DnApp.AccountKeeper, s.app.DnApp.BankKeeper, s.keeper

	accCoins := sdk.NewCoins(
		sdk.NewCoin("testa", sdk.NewInt(1000)),
		sdk.NewCoin("testb", sdk.NewInt(500)),
	)
	acc, _ := s.app.AddAccount(ctx, accCoins...)

	delPoolAcc := accKeeper.GetModuleAccount(ctx, types.DelPoolName)
	s.Require().NotNil(delPoolAcc)
	delPoolCoins := keeper.GetDelegatedPoolSupply(ctx)
	s.Require().True(delPoolCoins.Empty())

	// fail: DelegateCoinsToPool: validation
	{
		s.Require().Error(
			keeper.DelegateCoinsToPool(ctx, acc.GetAddress(),
				sdk.Coins{
					sdk.Coin{
						Denom:  "",
						Amount: sdk.ZeroInt(),
					},
				},
			),
		)
		s.Require().Error(
			keeper.DelegateCoinsToPool(ctx, acc.GetAddress(),
				sdk.Coins{
					sdk.Coin{
						Denom:  "xfi",
						Amount: sdk.NewInt(-1),
					},
				},
			),
		)
	}

	// fail: DelegateCoinsToPool: insufficient funds
	{
		s.Require().Error(
			keeper.DelegateCoinsToPool(ctx, acc.GetAddress(),
				sdk.NewCoins(
					sdk.Coin{Denom: accCoins[0].Denom, Amount: accCoins[0].Amount.MulRaw(2)},
				),
			),
		)
	}

	// fail: DelegateCoinsToPool: non-existing account
	{
		accAddr, _, _ := tests.GenAccAddress()
		s.Require().Error(keeper.DelegateCoinsToPool(ctx, accAddr, sdk.NewCoins(sdk.NewCoin("xfi", sdk.NewInt(1)))))
	}

	// ok: DelegateCoinsToPool (500testa, 500testb)
	{
		opCoins := sdk.NewCoins(
			sdk.NewCoin(accCoins[0].Denom, accCoins[0].Amount.QuoRaw(2)),
			accCoins[1],
		)
		s.Require().NoError(keeper.DelegateCoinsToPool(ctx, acc.GetAddress(), opCoins))

		curDelPoolCoins := bankKeeper.GetAllBalances(ctx, delPoolAcc.GetAddress())
		s.Require().True(delPoolCoins.Add(opCoins...).IsEqual(curDelPoolCoins))
		delPoolCoins = curDelPoolCoins

		curAccCoins := bankKeeper.GetAllBalances(ctx, acc.GetAddress())
		s.Require().True(accCoins.Sub(opCoins).IsEqual(curAccCoins))
		accCoins = curAccCoins
	}

	// fail: UndelegateCoinsFromPool: validation
	{
		s.Require().Error(
			keeper.UndelegateCoinsFromPool(ctx, acc.GetAddress(),
				sdk.Coins{
					sdk.Coin{
						Denom:  "",
						Amount: sdk.ZeroInt(),
					},
				},
			),
		)
		s.Require().Error(
			keeper.UndelegateCoinsFromPool(ctx, acc.GetAddress(),
				sdk.Coins{
					sdk.Coin{
						Denom:  "xfi",
						Amount: sdk.NewInt(-1),
					},
				},
			),
		)
	}

	// fail: UndelegateCoinsFromPool: insufficient funds
	{
		s.Require().Error(
			keeper.UndelegateCoinsFromPool(ctx, acc.GetAddress(),
				sdk.NewCoins(
					sdk.Coin{Denom: delPoolCoins[0].Denom, Amount: delPoolCoins[0].Amount.MulRaw(2)},
				),
			),
		)
	}

	// fail: UndelegateCoinsFromPool: non-existing account
	{
		accAddr, _, _ := tests.GenAccAddress()
		s.Require().Error(keeper.UndelegateCoinsFromPool(ctx, accAddr, sdk.NewCoins(sdk.NewCoin("xfi", sdk.NewInt(1)))))
	}

	// ok: UndelegateCoinsFromPool  (500testa, 250testb)
	{
		opCoins := sdk.NewCoins(
			delPoolCoins[0],
			sdk.NewCoin(delPoolCoins[1].Denom, delPoolCoins[1].Amount.QuoRaw(2)),
		)
		s.Require().NoError(keeper.UndelegateCoinsFromPool(ctx, acc.GetAddress(), opCoins))

		curDelPoolCoins := bankKeeper.GetAllBalances(ctx, delPoolAcc.GetAddress())
		s.Require().True(delPoolCoins.Sub(opCoins).IsEqual(curDelPoolCoins))
		delPoolCoins = curDelPoolCoins

		curAccCoins := bankKeeper.GetAllBalances(ctx, acc.GetAddress())
		s.Require().True(accCoins.Add(opCoins...).IsEqual(curAccCoins))
		accCoins = curAccCoins
	}

	// manual check
	{
		s.Require().True(delPoolCoins.IsEqual(
			sdk.NewCoins(
				sdk.NewCoin("testb", sdk.NewInt(250)),
			),
		))

		s.Require().True(accCoins.IsEqual(
			sdk.NewCoins(
				sdk.NewCoin("testa", sdk.NewInt(1000)),
				sdk.NewCoin("testb", sdk.NewInt(250)),
			),
		))
	}
}
