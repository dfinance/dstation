package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/dfinance/dstation/x/vm/types"
)

type proposalInput struct {
	id       uint64
	proposal *types.PlannedProposal
}

func (i proposalInput) CheckEqual(t *testing.T, i2 proposalInput) {
	require.Equal(t, i.id, i2.id, "input Ids")
	require.Equal(t, i.proposal.Height, i2.proposal.Height, "proposal Heights")
	require.Equal(t, i.proposal.Content.Value, i2.proposal.Content.Value, "proposal Contents")
}

func (s *KeeperMockVmTestSuite) TestGovQueue() {
	ctx, keeper := s.ctx, s.keeper

	// Skip block (to be sure current is not 0)
	s.app.BeginBlock()
	s.app.EndBlock()
	startBlock, _ := s.app.GetCurrentBlockHeightTime()

	// Create mock proposals
	proposalInputs := make([]proposalInput, 0)
	{
		{
			proposal, err := types.NewPlannedProposal(
				startBlock+10,
				types.NewStdLibUpdateProposal("http://github.com/repo", "test_1", []byte{0x1}),
			)
			s.Require().NoError(err)
			proposalInputs = append(proposalInputs, proposalInput{
				id:       0,
				proposal: proposal,
			})
		}
		{
			proposal, err := types.NewPlannedProposal(
				startBlock+10,
				types.NewStdLibUpdateProposal("http://github.com/repo", "test_2", []byte{0x2}),
			)
			s.Require().NoError(err)
			proposalInputs = append(proposalInputs, proposalInput{
				id:       1,
				proposal: proposal,
			})
		}
		{
			proposal, err := types.NewPlannedProposal(
				startBlock+10,
				types.NewStdLibUpdateProposal("http://github.com/repo", "test_3", []byte{0x3}),
			)
			s.Require().NoError(err)
			proposalInputs = append(proposalInputs, proposalInput{
				id:       2,
				proposal: proposal,
			})
		}
	}

	// fail: ScheduleProposal: proposal from the past
	{
		proposal, err := types.NewPlannedProposal(
			startBlock-1,
			types.NewStdLibUpdateProposal("http://github.com/repo", "test_1", []byte{0x1}),
		)
		s.Require().NoError(err)

		s.Require().Error(keeper.ScheduleProposal(ctx, proposal))
	}

	// ok: ScheduleProposal
	{
		for _, input := range proposalInputs {
			s.Require().NoError(keeper.ScheduleProposal(ctx, input.proposal))
		}

		proposalOutputs := make([]proposalInput, 0)
		keeper.IterateProposalsQueue(ctx, func(id uint64, proposal *types.PlannedProposal) {
			proposalOutputs = append(proposalOutputs, proposalInput{id: id, proposal: proposal})
		})

		s.Require().Len(proposalOutputs, len(proposalInputs))
		for i := 0; i < len(proposalInputs); i++ {
			proposalInputs[i].CheckEqual(s.T(), proposalOutputs[i])
		}
	}

	// ok: RemoveProposalFromQueue
	{
		keeper.RemoveProposalFromQueue(ctx, 1)

		proposalInputs = append(proposalInputs[0:1], proposalInputs[2:]...)
		proposalOutputs := make([]proposalInput, 0)
		keeper.IterateProposalsQueue(ctx, func(id uint64, proposal *types.PlannedProposal) {
			proposalOutputs = append(proposalOutputs, proposalInput{id: id, proposal: proposal})
		})

		s.Require().Len(proposalOutputs, len(proposalInputs))
		for i := 0; i < len(proposalInputs); i++ {
			proposalInputs[i].CheckEqual(s.T(), proposalOutputs[i])
		}
	}
}
