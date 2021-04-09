package keeper_test

import (
	"context"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"

	"github.com/dfinance/dstation/pkg/mock"
	"github.com/dfinance/dstation/pkg/tests"
	dvmTypes "github.com/dfinance/dstation/pkg/types/dvm"
	"github.com/dfinance/dstation/x/vm/keeper"
	"github.com/dfinance/dstation/x/vm/types"
)

type KeeperMockVmTestSuite struct {
	suite.Suite

	app         *tests.DSimApp
	ctx         sdk.Context
	keeper      keeper.Keeper
	vmServer    *mock.VMServer
	queryClient types.QueryClient
}

func (s *KeeperMockVmTestSuite) SetupSuite() {
	// Init the SimApp
	s.app = tests.SetupDSimApp(tests.WithMockVM())
	s.ctx = s.app.GetContext(false)
	s.keeper = s.app.DnApp.VmKeeper
	s.vmServer = s.app.MockVMServer

	// Init the querier client
	querier := keeper.Querier{Keeper: s.keeper}
	queryHelper := baseapp.NewQueryServerTestHelper(s.ctx, s.app.DnApp.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, querier)
	s.queryClient = types.NewQueryClient(queryHelper)
}

func (s *KeeperMockVmTestSuite) TearDownSuite() {
	s.app.TearDown()
}

func (s *KeeperMockVmTestSuite) SetupTest() {
	s.ctx = s.ctx.WithEventManager(sdk.NewEventManager())
}

func (s *KeeperMockVmTestSuite) DoDSClientRequest(handler func(ctx context.Context, client dvmTypes.DSServiceClient)) {
	client := dvmTypes.NewDSServiceClient(s.app.MockVMServer.GetDSClientConnection())
	ctx, ctxCancel := context.WithDeadline(context.Background(), time.Now().Add(100*time.Millisecond))
	defer ctxCancel()

	handler(ctx, client)
}

func TestVMKeeper_MockVM(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(KeeperMockVmTestSuite))
}
