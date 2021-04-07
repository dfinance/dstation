package tests

import (
	"fmt"

	cryptoTypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	tmProto "github.com/tendermint/tendermint/proto/tendermint/types"

	dnConfig "github.com/dfinance/dstation/cmd/dstation/config"
)

// DeliverTxOption defines DeliverTx functional arguments.
type DeliverTxOption func(gasLimit *uint64)

// BeginBlock starts a new block.
func (app *DSimApp) BeginBlock() {
	app.curBlockHeight, app.curBlockTime = app.GetNextBlockHeightTime()

	app.DnApp.BeginBlock(
		abci.RequestBeginBlock{
			Header: tmProto.Header{
				Height: app.curBlockHeight,
				Time:   app.curBlockTime,
			},
		},
	)
}

// EndBlock ends the current block.
func (app *DSimApp) EndBlock() {
	app.DnApp.EndBlock(abci.RequestEndBlock{})
	app.DnApp.Commit()
}

// DeliverTx generates, signs and sends Tx containing messages.
func (app *DSimApp) DeliverTx(ctx sdk.Context, accAddr sdk.AccAddress, accPrivKey cryptoTypes.PrivKey, msgs []sdk.Msg, options ...DeliverTxOption) (sdk.GasInfo, *sdk.Result, error) {
	txCfg := app.encCfg.TxConfig

	// Get the latest account data
	acc := app.DnApp.AccountKeeper.GetAccount(ctx, accAddr)
	if acc == nil {
		panic(fmt.Errorf("account (%s): not found", accAddr))
	}

	// Apply options
	gasLimit := uint64(helpers.DefaultGenTxGas)
	for _, opt := range options {
		opt(&gasLimit)
	}

	// Generate and sign Tx
	feeCoin := sdk.Coins{sdk.NewInt64Coin(dnConfig.MainDenom, 1)}
	tx, err := helpers.GenTx(txCfg, msgs, feeCoin, gasLimit, "", []uint64{acc.GetAccountNumber()}, []uint64{acc.GetSequence()}, accPrivKey)
	if err != nil {
		panic(fmt.Errorf("generating Tx: %w", err))
	}

	// Simulate Tx
	txBytes, err := txCfg.TxEncoder()(tx)
	if err != nil {
		panic(fmt.Errorf("encoding Tx: %w", err))
	}

	gasInfo, res, err := app.DnApp.Simulate(txBytes)
	if err != nil {
		return gasInfo, res, err
	}

	// Deliver Tx
	app.BeginBlock()
	gasInfo, res, err = app.DnApp.Deliver(txCfg.TxEncoder(), tx)
	app.EndBlock()

	return gasInfo, res, err
}

func TxWithGasLimit(limit uint64) DeliverTxOption {
	return func(gasLimit *uint64) {
		*gasLimit = limit
	}
}
