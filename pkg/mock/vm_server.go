package mock

import (
	"context"
	"math/rand"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	grpcStatus "google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"

	"github.com/dfinance/dstation/pkg/types/dvm"
)

const (
	VMAddressLength = 20
	//
	bufferedListenerSize = 1024 * 1024
)

var (
	VMStdLibAddress = make([]byte, VMAddressLength)
)

type VMServer struct {
	sync.Mutex
	//
	response      *dvm.VMExecuteResponse
	responseDelay time.Duration
	failCountdown uint
	//
	vmListener   *bufconn.Listener
	dsListener   *bufconn.Listener
	vmGrpcServer *grpc.Server
}

func (s *VMServer) PublishModule(context.Context, *dvm.VMPublishModule) (*dvm.VMExecuteResponse, error) {
	return s.buildResponse()
}

func (s *VMServer) ExecuteScript(context.Context, *dvm.VMExecuteScript) (*dvm.VMExecuteResponse, error) {
	return s.buildResponse()
}

func (s *VMServer) SetResponse(resp *dvm.VMExecuteResponse) {
	s.Lock()
	defer s.Unlock()

	s.response = resp
}

func (s *VMServer) SetFailCountdown(value uint) {
	s.Lock()
	defer s.Unlock()

	s.failCountdown = value
}

func (s *VMServer) SetResponseDelay(dur time.Duration) {
	s.Lock()
	defer s.Unlock()

	s.responseDelay = dur
}

func (s *VMServer) Stop() {
	s.Lock()
	defer s.Unlock()

	if s.vmGrpcServer != nil {
		s.vmGrpcServer.Stop()
	}
	if s.vmListener != nil {
		s.vmListener.Close()
	}
}

func (s *VMServer) GetVMListener() net.Listener { return s.vmListener }

func (s *VMServer) GetDSListener() net.Listener { return s.dsListener }

func (s *VMServer) GetVMClientConnection() *grpc.ClientConn {
	bufDialer := func(ctx context.Context, url string) (net.Conn, error) {
		return s.vmListener.Dial()
	}

	ctx := context.TODO()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return conn
}

func (s *VMServer) GetDSClientConnection() *grpc.ClientConn {
	bufDialer := func(ctx context.Context, url string) (net.Conn, error) {
		return s.dsListener.Dial()
	}

	ctx := context.TODO()
	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return conn
}

func (s *VMServer) buildResponse() (*dvm.VMExecuteResponse, error) {
	s.Lock()
	defer s.Unlock()

	time.Sleep(s.responseDelay)

	if s.failCountdown > 0 {
		s.failCountdown--
		return nil, grpcStatus.Errorf(codes.Internal, "failing gRPC execution")
	}

	resp := s.response
	if resp == nil {
		resp = s.getDefaultResponse()
	}

	return resp, nil
}

func (s *VMServer) getDefaultResponse() *dvm.VMExecuteResponse {
	values := []*dvm.VMValue{
		{
			Type:  dvm.VmWriteOp_Value,
			Value: GetRandomBytes(8),
			Path:  GetRandomVMAccessPath(),
		},
		{
			Type:  dvm.VmWriteOp_Value,
			Value: GetRandomBytes(32),
			Path:  GetRandomVMAccessPath(),
		},
	}

	events := []*dvm.VMEvent{
		{
			SenderAddress: VMStdLibAddress,
			EventType: &dvm.LcsTag{
				TypeTag: dvm.LcsType_LcsVector,
				VectorType: &dvm.LcsTag{
					TypeTag: dvm.LcsType_LcsU8,
				},
			},
			EventData: GetRandomBytes(32),
		},
	}

	return &dvm.VMExecuteResponse{
		WriteSet: values,
		Events:   events,
		GasUsed:  10000,
		Status:   &dvm.VMStatus{},
	}
}

func NewVMServer() *VMServer {
	vmServer := &VMServer{
		dsListener:   bufconn.Listen(bufferedListenerSize),
		vmListener:   bufconn.Listen(bufferedListenerSize),
		vmGrpcServer: grpc.NewServer(),
	}
	dvm.RegisterVMModulePublisherServer(vmServer.vmGrpcServer, vmServer)
	dvm.RegisterVMScriptExecutorServer(vmServer.vmGrpcServer, vmServer)

	go func() {
		if err := vmServer.vmGrpcServer.Serve(vmServer.vmListener); err != nil {
			panic(err)
		}
	}()

	return vmServer
}

func GetRandomVMAccessPath() *dvm.VMAccessPath {
	return &dvm.VMAccessPath{
		Address: GetRandomBytes(VMAddressLength),
		Path:    GetRandomBytes(VMAddressLength),
	}
}

func GetRandomBytes(len int) []byte {
	rndBytes := make([]byte, len)

	_, err := rand.Read(rndBytes)
	if err != nil {
		panic(err)
	}

	return rndBytes
}

func init() {
	VMStdLibAddress[VMAddressLength-1] = 1
}
