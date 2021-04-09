package vm

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"
	govTypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/dfinance/dstation/x/vm/keeper"
	"github.com/dfinance/dstation/x/vm/types"
)

// NewGovHandler creates a governance handler to manage new proposal types.
func NewGovHandler(k keeper.Keeper) govTypes.Handler {
	return func(ctx sdk.Context, content govTypes.Content) error {
		if content.ProposalRoute() != types.RouterKey {
			return fmt.Errorf("invalid proposal route %q for module %q", content.ProposalRoute(), types.ModuleName)
		}

		switch c := content.(type) {
		case *types.PlannedProposal:
			return handlePlannedProposal(ctx, k, c)
		default:
			return sdkErrors.Wrapf(sdkErrors.ErrUnknownRequest, "unsupported proposal content type: %T", c)
		}
	}
}

// handlePlannedProposal handles all types.PlannedProposal Gov proposals.
func handlePlannedProposal(ctx sdk.Context, k keeper.Keeper, pProposal *types.PlannedProposal) error {
	c := pProposal.GetContent()
	switch c.ProposalType() {
	case types.ProposalTypeStdlibUpdate:
		return handleStdlibUpdatePlannedProposal(ctx, k, pProposal)
	default:
		return sdkErrors.Wrapf(sdkErrors.ErrUnknownRequest, "unsupported proposal type for PlannedProposal: %s", c.ProposalType())
	}
}

// handleStdlibUpdatePlannedProposal handles types.StdLibUpdateProposal Gov proposal.
func handleStdlibUpdatePlannedProposal(ctx sdk.Context, k keeper.Keeper, pProposal *types.PlannedProposal) error {
	updateProposal, ok := pProposal.GetContent().(*types.StdLibUpdateProposal)
	if !ok {
		return fmt.Errorf("type assert to %s proposal for PlannedProposal failed: %T", types.ProposalTypeStdlibUpdate, pProposal.GetContent())
	}

	// DVM check (dry-run deploy)
	msg := types.NewMsgDeployModule(types.StdLibAddress, updateProposal.Code...)
	if err := msg.ValidateBasic(); err != nil {
		return fmt.Errorf("module deploy (dry run): validation: %w", err)
	}

	if err := k.DeployContractDryRun(ctx, msg); err != nil {
		return fmt.Errorf("module deploy (dry run): %w", err)
	}

	// Add proposal to the gov proposal queue
	if err := k.ScheduleProposal(ctx, pProposal); err != nil {
		return fmt.Errorf("proposal scheduling: %w", err)
	}

	k.Logger(ctx).Info(fmt.Sprintf("Proposal scheduled:\n%s", pProposal.String()))

	return nil
}
