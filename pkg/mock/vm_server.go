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

	dvmTypes "github.com/dfinance/dstation/pkg/types/dvm"
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
	response      *dvmTypes.VMExecuteResponse
	responseDelay time.Duration
	failCountdown uint
	//
	vmListener   *bufconn.Listener
	dsListener   *bufconn.Listener
	vmGrpcServer *grpc.Server
}

func (s *VMServer) PublishModule(context.Context, *dvmTypes.VMPublishModule) (*dvmTypes.VMExecuteResponse, error) {
	return s.buildResponse()
}

func (s *VMServer) ExecuteScript(context.Context, *dvmTypes.VMExecuteScript) (*dvmTypes.VMExecuteResponse, error) {
	return s.buildResponse()
}

func (s *VMServer) SetResponse(resp *dvmTypes.VMExecuteResponse) {
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

func (s *VMServer) buildResponse() (*dvmTypes.VMExecuteResponse, error) {
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

func (s *VMServer) getDefaultResponse() *dvmTypes.VMExecuteResponse {
	values := []*dvmTypes.VMValue{
		{
			Type:  dvmTypes.VmWriteOp_Value,
			Value: GetRandomBytes(8),
			Path:  GetRandomVMAccessPath(),
		},
		{
			Type:  dvmTypes.VmWriteOp_Value,
			Value: GetRandomBytes(32),
			Path:  GetRandomVMAccessPath(),
		},
	}

	events := []*dvmTypes.VMEvent{
		{
			SenderAddress: VMStdLibAddress,
			EventType: &dvmTypes.LcsTag{
				TypeTag: dvmTypes.LcsType_LcsVector,
				VectorType: &dvmTypes.LcsTag{
					TypeTag: dvmTypes.LcsType_LcsU8,
				},
			},
			EventData: GetRandomBytes(32),
		},
	}

	return &dvmTypes.VMExecuteResponse{
		WriteSet: values,
		Events:   events,
		GasUsed:  10000,
		Status:   &dvmTypes.VMStatus{},
	}
}

func NewVMServer() *VMServer {
	vmServer := &VMServer{
		dsListener:   bufconn.Listen(bufferedListenerSize),
		vmListener:   bufconn.Listen(bufferedListenerSize),
		vmGrpcServer: grpc.NewServer(),
	}
	dvmTypes.RegisterVMModulePublisherServer(vmServer.vmGrpcServer, vmServer)
	dvmTypes.RegisterVMScriptExecutorServer(vmServer.vmGrpcServer, vmServer)

	go func() {
		if err := vmServer.vmGrpcServer.Serve(vmServer.vmListener); err != nil {
			panic(err)
		}
	}()

	return vmServer
}

func GetRandomVMAccessPath() *dvmTypes.VMAccessPath {
	return &dvmTypes.VMAccessPath{
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
