package keeper_test

import (
	"encoding/hex"

	"github.com/dfinance/dstation/pkg/mock"
	"github.com/dfinance/dstation/x/vm/types"
)

func (s *KeeperMockVmTestSuite) TestGenesis() {
	ctx, keeper := s.ctx, s.keeper

	genStateExpected := keeper.ExportGenesis(ctx)
	genStateAppendix := types.GenesisState{
		WriteSet: make([]types.GenesisState_WriteOp, 0),
	}

	for i := 0; i < 5; i++ {
		vmPath, writeSetData := mock.GetRandomVMAccessPath(), mock.GetRandomBytes(20)
		writeOp := types.GenesisState_WriteOp{
			Address: hex.EncodeToString(vmPath.Address),
			Path:    hex.EncodeToString(vmPath.Path),
			Value:   hex.EncodeToString(writeSetData),
		}

		genStateAppendix.WriteSet = append(genStateAppendix.WriteSet, writeOp)
		genStateExpected.WriteSet = append(genStateExpected.WriteSet, writeOp)
	}

	// ok: Add writeSets to an existing default genesis
	{
		keeper.InitGenesis(ctx, &genStateAppendix)
		s.Require().ElementsMatch(genStateExpected.WriteSet, keeper.ExportGenesis(ctx).WriteSet)
	}
}
