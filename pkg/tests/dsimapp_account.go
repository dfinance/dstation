package tests

import (
	"fmt"

	cryptoTypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authTypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// AddAccount creates a new account, sets account initial balance and updates totalSupply (if defined).
func (app *DSimApp) AddAccount(ctx sdk.Context, coins ...sdk.Coin) (authTypes.AccountI, cryptoTypes.PrivKey) {
	addr, _, privKey := GenAccAddress()

	acc := app.DnApp.AccountKeeper.NewAccountWithAddress(ctx, addr)
	app.DnApp.AccountKeeper.SetAccount(ctx, acc)

	if len(coins) > 0 {
		if err := app.DnApp.BankKeeper.AddCoins(ctx, addr, coins); err != nil {
			panic(fmt.Errorf("generating test account with coins (%s): %w", sdk.Coins(coins).String(), err))
		}

		prevSupply := app.DnApp.BankKeeper.GetSupply(ctx)
		app.DnApp.BankKeeper.SetSupply(ctx, banktypes.NewSupply(prevSupply.GetTotal().Add(coins...)))
	}

	return acc, privKey
}