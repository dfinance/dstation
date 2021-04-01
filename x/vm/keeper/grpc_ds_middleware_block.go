package keeper

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dfinance/dvm-proto/go/vm_grpc"
	"github.com/dfinance/glav"
	"github.com/dfinance/lcs"

	"github.com/dfinance/dstation/x/vm/types"
)

type BlockMetadata struct {
	Height uint64
}

// NewBlockMiddleware creates DS server middleware which return current blockHeight.
func NewBlockMiddleware() types.DSDataMiddleware {
	blockHeaderPath := vm_grpc.VMAccessPath{
		Address: types.StdLibAddress,
		Path:    glav.BlockMetadataVector(),
	}

	return func(ctx sdk.Context, path *vm_grpc.VMAccessPath) (data []byte, err error) {
		if bytes.Equal(blockHeaderPath.Address, path.Address) && bytes.Equal(blockHeaderPath.Path, path.Path) {
			return lcs.Marshal(BlockMetadata{Height: uint64(ctx.BlockHeader().Height)})
		}

		return
	}
}
