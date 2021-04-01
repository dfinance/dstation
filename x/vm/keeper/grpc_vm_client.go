package keeper

import (
	"github.com/dfinance/dvm-proto/go/compiler_grpc"
	"github.com/dfinance/dvm-proto/go/metadata_grpc"
	"github.com/dfinance/dvm-proto/go/vm_grpc"
	"google.golang.org/grpc"
)

// VMClient is an aggregated gRPC VM services client.
type VMClient struct {
	compiler_grpc.DvmCompilerClient
	metadata_grpc.DVMBytecodeMetadataClient
	vm_grpc.VMModulePublisherClient
	vm_grpc.VMScriptExecutorClient
}

// NewVMClient creates VMClient using gRPC connection.
func NewVMClient(conn *grpc.ClientConn) VMClient {
	return VMClient{
		DvmCompilerClient:         compiler_grpc.NewDvmCompilerClient(conn),
		DVMBytecodeMetadataClient: metadata_grpc.NewDVMBytecodeMetadataClient(conn),
		VMModulePublisherClient:   vm_grpc.NewVMModulePublisherClient(conn),
		VMScriptExecutorClient:    vm_grpc.NewVMScriptExecutorClient(conn),
	}
}
