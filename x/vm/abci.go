package vm

import (
	"fmt"
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

	// Update DS server storage context
	k.SetDSContext(ctx)

	// Gov proposals processing
	k.IterateProposalsQueue(ctx, func(id uint64, pProposal *types.PlannedProposal) {
		if !pProposal.ShouldExecute(ctx) {
			return
		}

		var err error
		c := pProposal.GetContent()
		switch c.ProposalType() {
		case types.ProposalTypeStdlibUpdate:
			err = handleStdlibUpdateProposalExecution(ctx, k, pProposal)
		default:
			panic(fmt.Errorf("unsupported proposal type for PlannedProposal: %s", c.ProposalType()))
		}

		if err != nil {
			k.Logger(ctx).Error(fmt.Sprintf("%s\n\nExecution status: failed: %v", pProposal.String(), err))
		} else {
			k.SetDSContext(ctx)
			k.Logger(ctx).Info(fmt.Sprintf("%s\n\nExecution status: done", pProposal.String()))
		}

		k.RemoveProposalFromQueue(ctx, id)
	})
}

// handleStdlibUpdateProposalExecution requests DVM to update stdlib.
func handleStdlibUpdateProposalExecution(ctx sdk.Context, k keeper.Keeper, pProposal *types.PlannedProposal) error {
	updateProposal, ok := pProposal.GetContent().(*types.StdLibUpdateProposal)
	if !ok {
		return fmt.Errorf("type assert to %s proposal for PlannedProposal failed: %T", types.ProposalTypeStdlibUpdate, pProposal.GetContent())
	}

	msg := types.NewMsgDeployModule(types.StdLibAddress, updateProposal.Code...)
	if err := msg.ValidateBasic(); err != nil {
		return fmt.Errorf("module deploy: validation: %w", err)
	}

	if err := k.DeployContract(ctx, msg); err != nil {
		return fmt.Errorf("module deploy: %w", err)
	}

	return nil
}
