package keeper

import (
	"context"
	"encoding/hex"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/dfinance/dvm-proto/go/ds_grpc"
	"github.com/dfinance/dvm-proto/go/vm_grpc"
	"google.golang.org/grpc"

	"github.com/dfinance/dstation/x/vm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ ds_grpc.DSServiceServer = &DSServer{}

// DSServer is a DataSource server that catches VM client data requests.
type DSServer struct {
	sync.Mutex
	// Current storage context and implementation
	ctx     sdk.Context
	storage types.VMStorage
	// Data middleware handlers
	dataMiddlewares []types.DSDataMiddleware
	// Started gRPC server instance
	server *grpc.Server
}

// Logger returns logger with DS server context.
func (srv *DSServer) Logger() log.Logger {
	return srv.ctx.Logger().With("module", fmt.Sprintf("x/%s/dsserver", types.ModuleName))
}

// RegisterDataMiddleware registers new data middleware.
func (srv *DSServer) RegisterDataMiddleware(md types.DSDataMiddleware) {
	srv.dataMiddlewares = append(srv.dataMiddlewares, md)
}

// SetContext updates server storage context.
func (srv *DSServer) SetContext(ctx sdk.Context) {
	srv.Lock()
	defer srv.Unlock()

	srv.ctx = ctx
}

// IsStarted checks if gRPC server has been started.
func (srv *DSServer) IsStarted() bool {
	srv.Lock()
	defer srv.Unlock()

	return srv.server != nil
}

// Start starts gRPC DS server in the go routine.
func (srv *DSServer) Start(listener net.Listener) {
	srv.Lock()
	defer srv.Unlock()

	if srv.server != nil {
		return
	}

	srv.server = grpc.NewServer()
	ds_grpc.RegisterDSServiceServer(srv.server, srv)

	go func() {
		if err := srv.server.Serve(listener); err != nil {
			panic(err) // should not happen
		}
	}()
	time.Sleep(10 * time.Millisecond) // force context switch for server to start

	return
}

// Stop stops gRPC DS server.
func (srv *DSServer) Stop() {
	srv.Lock()
	defer srv.Unlock()

	if srv.server == nil {
		return
	}

	srv.server.Stop()
}

// GetRaw implements gRPC service handler: returns value from the storage.
func (srv *DSServer) GetRaw(_ context.Context, req *ds_grpc.DSAccessPath) (*ds_grpc.DSRawResponse, error) {
	path := &vm_grpc.VMAccessPath{
		Address: req.Address,
		Path:    req.Path,
	}
	srv.Logger().Info(fmt.Sprintf("Get path: %s", types.StringifyVMAccessPath(path)))

	// check middlewares
	blob, err := srv.processMiddlewares(path)
	if err != nil {
		srv.Logger().Error(fmt.Sprintf("Error processing middlewares for path %s: %v", types.StringifyVMAccessPath(path), err))
		return ErrNoData(req), nil
	}
	if blob != nil {
		return &ds_grpc.DSRawResponse{Blob: blob}, nil
	}

	// check storage
	if !srv.storage.HasValue(srv.ctx, path) {
		srv.Logger().Debug(fmt.Sprintf("Can't find path: %s", types.StringifyVMAccessPath(path)))
		return ErrNoData(req), nil
	}

	srv.Logger().Debug(fmt.Sprintf("Get path: %s", types.StringifyVMAccessPath(path)))
	blob = srv.storage.GetValue(srv.ctx, path)
	srv.Logger().Debug(fmt.Sprintf("Return values: %s\n", hex.EncodeToString(blob)))

	return &ds_grpc.DSRawResponse{Blob: blob}, nil
}

// MultiGetRaw implements gRPC service handler: returns multiple values from the storage.
func (srv *DSServer) MultiGetRaw(_ context.Context, req *ds_grpc.DSAccessPaths) (*ds_grpc.DSRawResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "MultiGetRaw unimplemented")
}

// processMiddlewares checks that accessPath can be processed by any registered middleware.
// Contract: if {data} != nil, middleware was found.
func (srv *DSServer) processMiddlewares(path *vm_grpc.VMAccessPath) (data []byte, err error) {
	for _, f := range srv.dataMiddlewares {
		data, err = f(srv.ctx, path)
		if err != nil || data != nil {
			return
		}
	}

	return
}

// NewDSServer creates a new DS server.
func NewDSServer(storage types.VMStorage) *DSServer {
	return &DSServer{
		storage: storage,
	}
}

// ErrNoData builds gRPC error response when data wasn't found.
func ErrNoData(path *ds_grpc.DSAccessPath) *ds_grpc.DSRawResponse {
	return &ds_grpc.DSRawResponse{
		ErrorCode:    ds_grpc.DSRawResponse_NO_DATA,
		ErrorMessage: fmt.Sprintf("data not found for access path: %s", path.String()),
	}
}
