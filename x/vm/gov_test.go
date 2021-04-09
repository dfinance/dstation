package vm_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	dnConfig "github.com/dfinance/dstation/cmd/dstation/config"
	"github.com/dfinance/dstation/x/vm"
	"github.com/dfinance/dstation/x/vm/client"
	"github.com/dfinance/dstation/x/vm/types"
)

func (s *ModuleDVVTestSuite) TestUpdateStdlibGovProposal() {
	ctx, keeper := s.ctx, s.keeper
	acc, accPrivKey := s.app.AddAccount(ctx, sdk.NewCoin(dnConfig.MainDenom, sdk.NewInt(1000)))

	// Deploy module v1 to the StdLib
	{
		// Compile
		byteCode := s.CompileMoveFile(acc.GetAddress(), "gov/module_v1.move")

		// Build Gov proposal
		curBlock, _ := s.app.GetCurrentBlockHeightTime()
		updateProposal := types.NewStdLibUpdateProposal("http://github.com/myStdLib", "v1: initial version", byteCode[0])
		pProposal, err := types.NewPlannedProposal(curBlock+1, updateProposal)
		s.Require().NoError(err)
		s.Require().NoError(pProposal.ValidateBasic())

		// Emulate Gov proposal approved and routed to the VM module
		s.Require().NoError(vm.NewGovHandler(keeper)(ctx, pProposal))

		// Skip block (planned proposal apply)
		s.app.BeginBlock()
		s.app.EndBlock()
	}

	// Check module v1 deployed
	{
		_, scriptEvents := s.CheckContractExecuted(
			s.ExecuteScript(acc.GetAddress(), accPrivKey,
				"gov/script.move", nil,
				s.BuildScriptArg("255", client.NewU64ScriptArg),
			),
		)

		s.CheckABCIEventsContain(scriptEvents, []sdk.Event{
			sdk.NewEvent(
				types.EventTypeMoveEvent,
				sdk.NewAttribute(types.AttributeVmEventSender, acc.GetAddress().String()),
				sdk.NewAttribute(types.AttributeVmEventSource, types.AttributeValueSourceScript),
				sdk.NewAttribute(types.AttributeVmEventType, "u64"),
				sdk.NewAttribute(types.AttributeVmEventData, "ff00000000000000"),
			),
		})
	}

	// Update module v1 to v2
	{
		// Compile
		byteCode := s.CompileMoveFile(acc.GetAddress(), "gov/module_v2.move")

		// Build Gov proposal
		curBlock, _ := s.app.GetCurrentBlockHeightTime()
		updateProposal := types.NewStdLibUpdateProposal("http://github.com/myStdLib", "v2: fix for inc() func", byteCode[0])
		pProposal, err := types.NewPlannedProposal(curBlock+1, updateProposal)
		s.Require().NoError(err)
		s.Require().NoError(pProposal.ValidateBasic())

		// Emulate Gov proposal approved and routed to the VM module
		s.Require().NoError(vm.NewGovHandler(keeper)(ctx, pProposal))

		// Skip block (planned proposal apply)
		s.app.BeginBlock()
		s.app.EndBlock()
	}

	// Check module was updated to v2
	{
		_, scriptEvents := s.CheckContractExecuted(
			s.ExecuteScript(acc.GetAddress(), accPrivKey,
				"gov/script.move", nil,
				s.BuildScriptArg("255", client.NewU64ScriptArg),
			),
		)

		s.CheckABCIEventsContain(scriptEvents, []sdk.Event{
			sdk.NewEvent(
				types.EventTypeMoveEvent,
				sdk.NewAttribute(types.AttributeVmEventSender, acc.GetAddress().String()),
				sdk.NewAttribute(types.AttributeVmEventSource, types.AttributeValueSourceScript),
				sdk.NewAttribute(types.AttributeVmEventType, "u64"),
				sdk.NewAttribute(types.AttributeVmEventData, "0001000000000000"),
			),
		})
	}

	// Check Gov proposal queue is empty
	{
		proposalsCount := 0
		keeper.IterateProposalsQueue(ctx, func(id uint64, pProposal *types.PlannedProposal) {
			proposalsCount++
		})
		s.Require().Zero(proposalsCount)
	}
}
