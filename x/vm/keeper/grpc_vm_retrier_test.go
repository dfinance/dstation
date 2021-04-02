package keeper_test

import (
	"time"

	"github.com/dfinance/dstation/pkg/tests"
	"github.com/dfinance/dstation/pkg/types/dvm"
	"github.com/dfinance/dstation/x/vm/types"
)

// nolint:errcheck
func (s *KeeperMockVmTestSuite) TestRetryMechanism() {
	// Rollback to defaults
	defer func() {
		s.app.SetCustomVMRetryParams(0, 0)
		s.vmServer.SetResponseDelay(0)
		s.vmServer.SetFailCountdown(0)
	}()

	accAddr, _, _ := tests.GenAccAddress()
	ctx, keeper, vmServer := s.ctx, s.keeper, s.vmServer

	// ok
	s.T().Log("OK: in one attempt (infinite settings)")
	{
		s.app.SetCustomVMRetryParams(0, 0)

		// Build msg
		vmResp := &dvm.VMExecuteResponse{GasUsed: 1, Status: &dvm.VMStatus{}}
		msg := types.NewMsgDeployModule(accAddr, []byte{0x1})

		// Request
		vmServer.SetResponseDelay(50 * time.Millisecond)
		vmServer.SetFailCountdown(0)
		vmServer.SetResponse(vmResp)
		s.Require().NoError(keeper.DeployContract(ctx, msg))
	}

	// ok
	s.T().Log("OK: in one attempt (settings with limit)")
	{
		s.app.SetCustomVMRetryParams(1, 5000)

		// Build msg
		vmResp := &dvm.VMExecuteResponse{GasUsed: 1, Status: &dvm.VMStatus{}}
		msg := types.NewMsgDeployModule(accAddr, []byte{0x1})

		// Request
		vmServer.SetResponseDelay(10 * time.Millisecond)
		vmServer.SetFailCountdown(0)
		vmServer.SetResponse(vmResp)
		s.Require().NoError(keeper.DeployContract(ctx, msg))
	}

	// ok
	s.T().Log("OK: in one attempt (without request timeout)")
	{
		s.app.SetCustomVMRetryParams(1, 0)

		// Build msg
		vmResp := &dvm.VMExecuteResponse{GasUsed: 1, Status: &dvm.VMStatus{}}
		msg := types.NewMsgDeployModule(accAddr, []byte{0x1})

		// Request
		vmServer.SetResponseDelay(500 * time.Millisecond)
		vmServer.SetFailCountdown(0)
		vmServer.SetResponse(vmResp)
		s.Require().NoError(keeper.DeployContract(ctx, msg))
	}

	// ok
	s.T().Log("OK: in multiple attempts (with request timeout)")
	{
		s.app.SetCustomVMRetryParams(10, 50)

		// Build msg
		vmResp := &dvm.VMExecuteResponse{GasUsed: 1, Status: &dvm.VMStatus{}}
		msg := types.NewMsgDeployModule(accAddr, []byte{0x1})

		// Request
		vmServer.SetResponseDelay(10 * time.Millisecond)
		vmServer.SetFailCountdown(5)
		vmServer.SetResponse(vmResp)
		s.Require().NoError(keeper.DeployContract(ctx, msg))
	}

	// ok
	s.T().Log("OK: in multiple attempts (without request timeout)")
	{
		s.app.SetCustomVMRetryParams(10, 0)

		// Build msg
		vmResp := &dvm.VMExecuteResponse{GasUsed: 1, Status: &dvm.VMStatus{}}
		msg := types.NewMsgDeployModule(accAddr, []byte{0x1})

		// Request
		vmServer.SetResponseDelay(100 * time.Millisecond)
		vmServer.SetFailCountdown(5)
		vmServer.SetResponse(vmResp)
		s.Require().NoError(keeper.DeployContract(ctx, msg))
	}

	// ok
	s.T().Log("OK: in one attempt with long response (without limits)")
	{
		s.app.SetCustomVMRetryParams(0, 0)

		// Build msg
		vmResp := &dvm.VMExecuteResponse{GasUsed: 1, Status: &dvm.VMStatus{}}
		msg := types.NewMsgDeployModule(accAddr, []byte{0x1})

		// Request
		vmServer.SetResponseDelay(3000 * time.Millisecond)
		vmServer.SetFailCountdown(0)
		vmServer.SetResponse(vmResp)
		s.Require().NoError(keeper.DeployContract(ctx, msg))
	}

	// fail
	s.T().Log("FAIL: by timeout (deadline)")
	{
		s.app.SetCustomVMRetryParams(5, 30)

		// Build msg
		vmResp := &dvm.VMExecuteResponse{GasUsed: 1, Status: &dvm.VMStatus{}}
		msg := types.NewMsgDeployModule(accAddr, []byte{0x1})

		// Request
		vmServer.SetResponseDelay(300 * time.Millisecond)
		vmServer.SetFailCountdown(0)
		vmServer.SetResponse(vmResp)
		tests.CheckPanicErrorContains(s.T(),
			func() {
				keeper.DeployContract(ctx, msg)
			},
			"context deadline exceeded",
		)
	}

	// fail
	s.T().Log("FAIL: by attempts")
	{
		s.app.SetCustomVMRetryParams(5, 0)

		// Build msg
		vmResp := &dvm.VMExecuteResponse{GasUsed: 1, Status: &dvm.VMStatus{}}
		msg := types.NewMsgDeployModule(accAddr, []byte{0x1})

		// Request
		vmServer.SetResponseDelay(50 * time.Millisecond)
		vmServer.SetFailCountdown(10)
		vmServer.SetResponse(vmResp)
		tests.CheckPanicErrorContains(s.T(),
			func() {
				keeper.DeployContract(ctx, msg)
			},
			"failing gRPC execution",
		)
	}
}
