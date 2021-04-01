package vm

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/dfinance/dstation/x/vm/keeper"
	"github.com/dfinance/dstation/x/vm/types"
)

// BeginBlocker set dataSource server storage context and handles gov proposal scheduler
// iterating over plannedProposals and checking if it is time to execute.
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)

	k.SetDSContext(ctx)
}
