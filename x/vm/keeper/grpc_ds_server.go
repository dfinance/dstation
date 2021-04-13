package keeper

import (
	"context"
	"encoding/hex"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	dnTypes "github.com/dfinance/dstation/pkg/types"
	"google.golang.org/grpc"

	dvmTypes "github.com/dfinance/dstation/pkg/types/dvm"
	"github.com/dfinance/dstation/x/vm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ dvmTypes.DSServiceServer = &DSServer{}

// DSServer is a DataSource server that catches VM client data requests.
type DSServer struct {
	sync.Mutex
	// Current storage context and implementation
	ctx              sdk.Context
	storage          types.VMStorage
	ccInfoProvider   types.CurrencyInfoResProvider
	nBalanceProvider types.AccountBalanceResProvider
	oPriceProvider   types.OracleKeeper
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
	dvmTypes.RegisterDSServiceServer(srv.server, srv)

	go func() {
		if err := srv.server.Serve(listener); err != nil {
			panic(err) // should not happen
		}
	}()
	time.Sleep(10 * time.Millisecond) // force context switch for server to start
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

// GetOraclePrice implements gRPC service handler: returns Oracle price for the specified currency pair.
func (srv *DSServer) GetOraclePrice(_ context.Context, req *dvmTypes.OraclePriceRequest) (*dvmTypes.OraclePriceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	// Input check
	assetCode, err := dnTypes.NewAssetCodeByDenoms(strings.ToLower(req.Currency_1), strings.ToLower(req.Currency_2))
	if err != nil {
		return &dvmTypes.OraclePriceResponse{
			ErrorCode:    dvmTypes.ErrorCode_BAD_REQUEST,
			ErrorMessage: fmt.Sprintf("assetCode: %v", err),
		}, nil
	}

	// Get price
	price := srv.oPriceProvider.GetCurrentPrice(srv.ctx, assetCode)
	if price == nil {
		return &dvmTypes.OraclePriceResponse{
			ErrorCode:    dvmTypes.ErrorCode_NO_DATA,
			ErrorMessage: fmt.Sprintf("current price for assetCode (%s): not found", assetCode),
		}, nil
	}

	// Build response
	exchangeRateU128, err := types.SdkIntToVmU128(price.AskPrice)
	if err != nil {
		panic(fmt.Errorf("converting askPrice (%s) to U128: %w", price.AskPrice, err))
	}

	return &dvmTypes.OraclePriceResponse{Price: exchangeRateU128}, nil
}

// GetNativeBalance implements gRPC service handler: returns account native balance resource.
func (srv *DSServer) GetNativeBalance(_ context.Context, req *dvmTypes.NativeBalanceRequest) (*dvmTypes.NativeBalanceResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	// Input check
	accAddr, err := types.LibraToBech32(req.Address)
	if err != nil {
		return &dvmTypes.NativeBalanceResponse{
			ErrorCode:    dvmTypes.ErrorCode_BAD_REQUEST,
			ErrorMessage: fmt.Sprintf("address: %v", err),
		}, nil
	}

	denom := strings.ToLower(req.Ticker)
	if err := sdk.ValidateDenom(denom); err != nil {
		return &dvmTypes.NativeBalanceResponse{
			ErrorCode:    dvmTypes.ErrorCode_BAD_REQUEST,
			ErrorMessage: fmt.Sprintf("ticker: %v", err),
		}, nil
	}

	// Build response
	balance := srv.nBalanceProvider.GetVmAccountBalance(srv.ctx, accAddr, denom)

	return &dvmTypes.NativeBalanceResponse{
		Balance: balance,
	}, nil
}

// GetCurrencyInfo implements gRPC service handler: returns CurrencyInfo resource.
func (srv *DSServer) GetCurrencyInfo(_ context.Context, req *dvmTypes.CurrencyInfoRequest) (*dvmTypes.CurrencyInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	// Input check
	denom := strings.ToLower(req.Ticker)
	if err := sdk.ValidateDenom(denom); err != nil {
		return &dvmTypes.CurrencyInfoResponse{
			ErrorCode:    dvmTypes.ErrorCode_BAD_REQUEST,
			ErrorMessage: fmt.Sprintf("ticker: %v", err),
		}, nil
	}

	// Build response
	ccInfo := srv.ccInfoProvider.GetVmCurrencyInfo(srv.ctx, denom)
	if ccInfo == nil {
		return &dvmTypes.CurrencyInfoResponse{
			Info:         nil,
			ErrorCode:    dvmTypes.ErrorCode_NO_DATA,
			ErrorMessage: "denom not registered",
		}, nil
	}

	return &dvmTypes.CurrencyInfoResponse{Info: ccInfo}, nil
}

// GetRaw implements gRPC service handler: returns value from the storage.
func (srv *DSServer) GetRaw(_ context.Context, req *dvmTypes.DSAccessPath) (*dvmTypes.DSRawResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	noDataErr := func(path *dvmTypes.DSAccessPath) *dvmTypes.DSRawResponse {
		return &dvmTypes.DSRawResponse{
			ErrorCode:    dvmTypes.ErrorCode_NO_DATA,
			ErrorMessage: fmt.Sprintf("data not found for access path: %s", path.String()),
		}
	}

	path := &dvmTypes.VMAccessPath{
		Address: req.Address,
		Path:    req.Path,
	}
	srv.Logger().Info(fmt.Sprintf("Get path: %s", types.StringifyVMAccessPath(path)))

	// check middlewares
	blob, err := srv.processMiddlewares(path)
	if err != nil {
		srv.Logger().Error(fmt.Sprintf("Error processing middlewares for path %s: %v", types.StringifyVMAccessPath(path), err))
		return noDataErr(req), nil
	}
	if blob != nil {
		return &dvmTypes.DSRawResponse{Blob: blob}, nil
	}

	// check storage
	if !srv.storage.HasValue(srv.ctx, path) {
		srv.Logger().Debug(fmt.Sprintf("Can't find path: %s", types.StringifyVMAccessPath(path)))
		return noDataErr(req), nil
	}

	srv.Logger().Debug(fmt.Sprintf("Get path: %s", types.StringifyVMAccessPath(path)))
	blob = srv.storage.GetValue(srv.ctx, path)
	srv.Logger().Debug(fmt.Sprintf("Return values: %s\n", hex.EncodeToString(blob)))

	return &dvmTypes.DSRawResponse{Blob: blob}, nil
}

// MultiGetRaw implements gRPC service handler: returns multiple values from the storage.
func (srv *DSServer) MultiGetRaw(_ context.Context, req *dvmTypes.DSAccessPaths) (*dvmTypes.DSRawResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "MultiGetRaw unimplemented")
}

// processMiddlewares checks that accessPath can be processed by any registered middleware.
// Contract: if {data} != nil, middleware was found.
func (srv *DSServer) processMiddlewares(path *dvmTypes.VMAccessPath) (data []byte, err error) {
	for _, f := range srv.dataMiddlewares {
		data, err = f(srv.ctx, path)
		if err != nil || data != nil {
			return
		}
	}

	return
}

// NewDSServer creates a new DS server.
func NewDSServer(
	storage types.VMStorage,
	ccInfoProvider types.CurrencyInfoResProvider,
	nBalanceProvider types.AccountBalanceResProvider,
	oraclePriceProvider types.OracleKeeper,
) *DSServer {

	return &DSServer{
		storage:          storage,
		ccInfoProvider:   ccInfoProvider,
		nBalanceProvider: nBalanceProvider,
		oPriceProvider:   oraclePriceProvider,
	}
}
