package mock

import (
	"context"
	"math/rand"
	"net"
	"sync"
	"time"

	"github.com/dfinance/dvm-proto/go/vm_grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	grpcStatus "google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
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
	response      *vm_grpc.VMExecuteResponse
	responseDelay time.Duration
	failCountdown uint
	//
	vmListener   *bufconn.Listener
	dsListener   *bufconn.Listener
	vmGrpcServer *grpc.Server
}

func (s *VMServer) PublishModule(context.Context, *vm_grpc.VMPublishModule) (*vm_grpc.VMExecuteResponse, error) {
	return s.buildResponse()
}

func (s *VMServer) ExecuteScript(context.Context, *vm_grpc.VMExecuteScript) (*vm_grpc.VMExecuteResponse, error) {
	return s.buildResponse()
}

func (s *VMServer) SetResponse(resp *vm_grpc.VMExecuteResponse) {
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

func (s *VMServer) buildResponse() (*vm_grpc.VMExecuteResponse, error) {
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

	return s.response, nil
}

func (s *VMServer) getDefaultResponse() *vm_grpc.VMExecuteResponse {
	values := []*vm_grpc.VMValue{
		{
			Type:  vm_grpc.VmWriteOp_Value,
			Value: GetRandomBytes(8),
			Path:  GetRandomVMAccessPath(),
		},
		{
			Type:  vm_grpc.VmWriteOp_Value,
			Value: GetRandomBytes(32),
			Path:  GetRandomVMAccessPath(),
		},
	}

	events := []*vm_grpc.VMEvent{
		{
			SenderAddress: VMStdLibAddress,
			EventType: &vm_grpc.LcsTag{
				TypeTag: vm_grpc.LcsType_LcsVector,
				VectorType: &vm_grpc.LcsTag{
					TypeTag: vm_grpc.LcsType_LcsU8,
				},
			},
			EventData: GetRandomBytes(32),
		},
	}

	return &vm_grpc.VMExecuteResponse{
		WriteSet: values,
		Events:   events,
		GasUsed:  10000,
		Status:   &vm_grpc.VMStatus{},
	}
}

func NewVMServer() *VMServer {
	vmServer := &VMServer{
		dsListener:   bufconn.Listen(bufferedListenerSize),
		vmListener:   bufconn.Listen(bufferedListenerSize),
		vmGrpcServer: grpc.NewServer(),
	}
	vm_grpc.RegisterVMModulePublisherServer(vmServer.vmGrpcServer, vmServer)
	vm_grpc.RegisterVMScriptExecutorServer(vmServer.vmGrpcServer, vmServer)

	go func() {
		if err := vmServer.vmGrpcServer.Serve(vmServer.vmListener); err != nil {
			panic(err)
		}
	}()

	return vmServer
}

func GetRandomVMAccessPath() *vm_grpc.VMAccessPath {
	return &vm_grpc.VMAccessPath{
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
