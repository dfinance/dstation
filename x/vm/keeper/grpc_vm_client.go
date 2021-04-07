package keeper

import (
	"google.golang.org/grpc"

	dvmTypes "github.com/dfinance/dstation/pkg/types/dvm"
)

// VMClient is an aggregated gRPC VM services client.
type VMClient struct {
	dvmTypes.DvmCompilerClient
	dvmTypes.DVMBytecodeMetadataClient
	dvmTypes.VMModulePublisherClient
	dvmTypes.VMScriptExecutorClient
}

// NewVMClient creates VMClient using gRPC connection.
func NewVMClient(conn *grpc.ClientConn) VMClient {
	return VMClient{
		DvmCompilerClient:         dvmTypes.NewDvmCompilerClient(conn),
		DVMBytecodeMetadataClient: dvmTypes.NewDVMBytecodeMetadataClient(conn),
		VMModulePublisherClient:   dvmTypes.NewVMModulePublisherClient(conn),
		VMScriptExecutorClient:    dvmTypes.NewVMScriptExecutorClient(conn),
	}
}
