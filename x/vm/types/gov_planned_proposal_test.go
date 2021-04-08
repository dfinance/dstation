package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVM_PlannedProposal(t *testing.T) {
	// ok
	{
		content := NewStdLibUpdateProposal("http://github.com/repo", "test", []byte{1})
		p, err := NewPlannedProposal(1, content)
		require.NoError(t, err)
		require.NoError(t, p.ValidateBasic())
	}

	// fail: height: < 0
	{
		content := NewStdLibUpdateProposal("http://github.com/repo", "test", []byte{1})
		p, err := NewPlannedProposal(-1, content)
		require.NoError(t, err)
		require.Error(t, p.ValidateBasic())
	}

	// fail: content: nil
	{
		_, err := NewPlannedProposal(1, nil)
		require.Error(t, err)
	}

	// fail: content: non-ProtoMessage
	{
		_, err := NewPlannedProposal(1, StdLibUpdateProposal{})
		require.Error(t, err)
	}

	// fail: content: invalid
	{
		content := NewStdLibUpdateProposal("http://github.com/repo", "", []byte{1})
		p, err := NewPlannedProposal(1, content)
		require.NoError(t, err)
		require.Error(t, p.ValidateBasic())
	}
}
