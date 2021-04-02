package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/dfinance/dstation/pkg/types/dvm"
	"github.com/dfinance/dstation/x/vm/types"
)

var _ types.QueryServer = Querier{}

// Querier implements gRPC query RPCs.
type Querier struct {
	Keeper
}

// Data queries VMStorage value.
func (k Querier) Data(c context.Context, req *types.QueryDataRequest) (*types.QueryDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	// Validate and parse request
	if len(req.Address) != types.VMAddressLength {
		return nil, status.Errorf(codes.InvalidArgument, "address: invalid length (should be %d)", types.VMAddressLength)
	}
	if len(req.Path) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "path: empty")
	}

	ctx := sdk.UnwrapSDKContext(c)
	value := k.GetValue(ctx, &dvm.VMAccessPath{Address: req.Address, Path: req.Path})

	return &types.QueryDataResponse{Value: value}, nil
}

// TxVmStatus queries Tx VM status based on Tx ABCI logs.
func (k Querier) TxVmStatus(c context.Context, req *types.QueryTxVmStatusRequest) (*types.QueryTxVmStatusResponse, error) {
	txVmStatus := types.NewVmStatusFromABCILogs(req.TxMeta)
	return &types.QueryTxVmStatusResponse{VmStatus: txVmStatus}, nil
}

// Compile queries Move code compilation.
func (k Querier) Compile(c context.Context, req *types.QueryCompileRequest) (*types.QueryCompileResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	// Validate and parse request
	if len(req.Address) != types.VMAddressLength {
		return nil, status.Errorf(codes.InvalidArgument, "address: invalid length (should be %d)", types.VMAddressLength)
	}
	if len(req.Code) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "code: empty")
	}

	// Build compile request
	compReq := &dvm.SourceFiles{
		Units: []*dvm.CompilationUnit{
			{
				Text: string(req.Code),
				Name: "CompilationUnit",
			},
		},
		Address: req.Address,
	}

	// Compile request
	compResp, err := k.vmClient.Compile(c, compReq)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "compilation failed: VM connection (compilation): %v", err)
	}

	// Check for compilation errors
	if len(compResp.Errors) > 0 {
		return nil, status.Errorf(codes.InvalidArgument, "compilation failed: compiler errors: [%s]", strings.Join(compResp.Errors, ", "))
	}

	// Build response
	resp := make([]types.CompiledItem, 0, len(compResp.Units))
	for _, unit := range compResp.Units {
		compItem := types.CompiledItem{
			ByteCode: unit.Bytecode,
			Name:     unit.Name,
		}

		meta, err := k.vmClient.GetMetadata(c, &dvm.Bytecode{Code: unit.Bytecode})
		if err != nil {
			return nil, status.Errorf(codes.Unavailable, "compilation failed: VM connection (getting meta information): %v", err)
		}

		if ok := meta.GetScript(); ok != nil {
			compItem.CodeType = types.CompiledItem_SCRIPT
		}

		if moduleMeta := meta.GetModule(); moduleMeta != nil {
			compItem.CodeType = types.CompiledItem_MODULE
			compItem.Types = moduleMeta.GetTypes()
			compItem.Methods = moduleMeta.GetFunctions()
		}

		resp = append(resp, compItem)
	}

	return &types.QueryCompileResponse{CompiledItems: resp}, nil
}

// Metadata queries VM for byteCode metadata.
func (k Querier) Metadata(c context.Context, req *types.QueryMetadataRequest) (*types.QueryMetadataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	res, err := k.vmClient.GetMetadata(c, &dvm.Bytecode{
		Code: req.Code,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "getting meta information: %v", err)
	}

	return &types.QueryMetadataResponse{Metadata: res}, nil
}
