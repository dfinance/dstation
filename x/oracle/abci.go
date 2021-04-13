package oracle

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dfinance/dstation/x/oracle/keeper"
	"github.com/dfinance/dstation/x/oracle/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// EndBlocker is called every block and updates types.CurrentPrice for each registered types.Asset.
func EndBlocker(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	k.UpdateCurrentPrices(ctx)

	return []abci.ValidatorUpdate{}
}
