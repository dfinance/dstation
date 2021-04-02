package keeper

import (
	"google.golang.org/grpc"

	"github.com/dfinance/dstation/pkg/types/dvm"
)

// VMClient is an aggregated gRPC VM services client.
type VMClient struct {
	dvm.DvmCompilerClient
	dvm.DVMBytecodeMetadataClient
	dvm.VMModulePublisherClient
	dvm.VMScriptExecutorClient
}

// NewVMClient creates VMClient using gRPC connection.
func NewVMClient(conn *grpc.ClientConn) VMClient {
	return VMClient{
		DvmCompilerClient:         dvm.NewDvmCompilerClient(conn),
		DVMBytecodeMetadataClient: dvm.NewDVMBytecodeMetadataClient(conn),
		VMModulePublisherClient:   dvm.NewVMModulePublisherClient(conn),
		VMScriptExecutorClient:    dvm.NewVMScriptExecutorClient(conn),
	}
}
