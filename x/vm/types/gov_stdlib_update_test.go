package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVM_StdLibUpdateProposal(t *testing.T) {
	// ok
	{
		p := NewStdLibUpdateProposal("http://github.com/repo", "tst", []byte{1})
		require.NoError(t, p.ValidateBasic())
	}

	// fail: url: empty
	{
		p := NewStdLibUpdateProposal("", "tst", []byte{1})
		require.Error(t, p.ValidateBasic())
	}

	// fail: url: invalid
	{
		p := NewStdLibUpdateProposal("1://repo", "tst", []byte{1})
		require.Error(t, p.ValidateBasic())
	}

	// fail: description: empty
	{
		p := NewStdLibUpdateProposal("http://github.com/repo", "", []byte{1})
		require.Error(t, p.ValidateBasic())
	}

	// fail: byteCode: empty
	{
		p := NewStdLibUpdateProposal("http://github.com/repo", "tst", nil)
		require.Error(t, p.ValidateBasic())
	}

	// fail: byteCode: empty
	{
		p := NewStdLibUpdateProposal("http://github.com/repo", "tst", []byte{1}, nil)
		require.Error(t, p.ValidateBasic())
	}
}
