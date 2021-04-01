package keeper_test

import (
	"testing"
	"time"

	"github.com/dfinance/dvm-proto/go/vm_grpc"
	"github.com/stretchr/testify/require"

	"github.com/dfinance/dstation/pkg/tests"
	"github.com/dfinance/dstation/x/vm/types"
)

func TestVM_retryMechanism(t *testing.T) {
	accAddr, _, _ := tests.GenAccAddress()

	// ok
	t.Log("OK: in one attempt (infinite settings)")
	{
		app := tests.SetupDSimApp(tests.WithMockVM(), tests.WithCustomVMRetryParams(0, 0))
		defer app.TearDown()

		ctx, keeper := app.GetContext(), app.DnApp.VmKeeper

		// Build msg
		vmResp := &vm_grpc.VMExecuteResponse{GasUsed: 1, Status: &vm_grpc.VMStatus{}}
		msg := types.NewMsgDeployModule(accAddr, [][]byte{{0x1}})

		// Request
		app.MockVMServer.SetResponseDelay(50 * time.Millisecond)
		app.MockVMServer.SetFailCountdown(0)
		app.MockVMServer.SetResponse(vmResp)
		require.NoError(t, keeper.DeployContract(ctx, msg))
	}

	// ok
	t.Log("OK: in one attempt (settings with limit)")
	{
		app := tests.SetupDSimApp(tests.WithMockVM(), tests.WithCustomVMRetryParams(1, 5000))
		defer app.TearDown()

		ctx, keeper := app.GetContext(), app.DnApp.VmKeeper

		// Build msg
		vmResp := &vm_grpc.VMExecuteResponse{GasUsed: 1, Status: &vm_grpc.VMStatus{}}
		msg := types.NewMsgDeployModule(accAddr, [][]byte{{0x1}})

		// Request
		app.MockVMServer.SetResponseDelay(10 * time.Millisecond)
		app.MockVMServer.SetFailCountdown(0)
		app.MockVMServer.SetResponse(vmResp)
		require.NoError(t, keeper.DeployContract(ctx, msg))
	}

	// ok
	t.Log("OK: in one attempt (without request timeout)")
	{
		app := tests.SetupDSimApp(tests.WithMockVM(), tests.WithCustomVMRetryParams(1, 0))
		defer app.TearDown()

		ctx, keeper := app.GetContext(), app.DnApp.VmKeeper

		// Build msg
		vmResp := &vm_grpc.VMExecuteResponse{GasUsed: 1, Status: &vm_grpc.VMStatus{}}
		msg := types.NewMsgDeployModule(accAddr, [][]byte{{0x1}})

		// Request
		app.MockVMServer.SetResponseDelay(500 * time.Millisecond)
		app.MockVMServer.SetFailCountdown(0)
		app.MockVMServer.SetResponse(vmResp)
		require.NoError(t, keeper.DeployContract(ctx, msg))
	}

	// ok
	t.Log("OK: in multiple attempts (with request timeout)")
	{
		app := tests.SetupDSimApp(tests.WithMockVM(), tests.WithCustomVMRetryParams(10, 50))
		defer app.TearDown()

		ctx, keeper := app.GetContext(), app.DnApp.VmKeeper

		// Build msg
		vmResp := &vm_grpc.VMExecuteResponse{GasUsed: 1, Status: &vm_grpc.VMStatus{}}
		msg := types.NewMsgDeployModule(accAddr, [][]byte{{0x1}})

		// Request
		app.MockVMServer.SetResponseDelay(10 * time.Millisecond)
		app.MockVMServer.SetFailCountdown(5)
		app.MockVMServer.SetResponse(vmResp)
		require.NoError(t, keeper.DeployContract(ctx, msg))
	}

	// ok
	t.Log("OK: in multiple attempts (without request timeout)")
	{
		app := tests.SetupDSimApp(tests.WithMockVM(), tests.WithCustomVMRetryParams(10, 0))
		defer app.TearDown()

		ctx, keeper := app.GetContext(), app.DnApp.VmKeeper

		// Build msg
		vmResp := &vm_grpc.VMExecuteResponse{GasUsed: 1, Status: &vm_grpc.VMStatus{}}
		msg := types.NewMsgDeployModule(accAddr, [][]byte{{0x1}})

		// Request
		app.MockVMServer.SetResponseDelay(100 * time.Millisecond)
		app.MockVMServer.SetFailCountdown(5)
		app.MockVMServer.SetResponse(vmResp)
		require.NoError(t, keeper.DeployContract(ctx, msg))
	}

	// ok
	t.Log("OK: in one attempt with long response (without limits)")
	{
		app := tests.SetupDSimApp(tests.WithMockVM(), tests.WithCustomVMRetryParams(0, 0))
		defer app.TearDown()

		ctx, keeper := app.GetContext(), app.DnApp.VmKeeper

		// Build msg
		vmResp := &vm_grpc.VMExecuteResponse{GasUsed: 1, Status: &vm_grpc.VMStatus{}}
		msg := types.NewMsgDeployModule(accAddr, [][]byte{{0x1}})

		// Request
		app.MockVMServer.SetResponseDelay(3000 * time.Millisecond)
		app.MockVMServer.SetFailCountdown(0)
		app.MockVMServer.SetResponse(vmResp)
		require.NoError(t, keeper.DeployContract(ctx, msg))
	}

	// fail
	t.Log("FAIL: by timeout (deadline)")
	{
		app := tests.SetupDSimApp(tests.WithMockVM(), tests.WithCustomVMRetryParams(5, 30))
		defer app.TearDown()

		ctx, keeper := app.GetContext(), app.DnApp.VmKeeper

		// Build msg
		vmResp := &vm_grpc.VMExecuteResponse{GasUsed: 1, Status: &vm_grpc.VMStatus{}}
		msg := types.NewMsgDeployModule(accAddr, [][]byte{{0x1}})

		// Request
		app.MockVMServer.SetResponseDelay(300 * time.Millisecond)
		app.MockVMServer.SetFailCountdown(0)
		app.MockVMServer.SetResponse(vmResp)
		tests.CheckPanicErrorContains(t,
			func() {
				keeper.DeployContract(ctx, msg)
			},
			"context deadline exceeded",
		)
	}

	// fail
	t.Log("FAIL: by attempts")
	{
		app := tests.SetupDSimApp(tests.WithMockVM(), tests.WithCustomVMRetryParams(5, 0))
		defer app.TearDown()

		ctx, keeper := app.GetContext(), app.DnApp.VmKeeper

		// Build msg
		vmResp := &vm_grpc.VMExecuteResponse{GasUsed: 1, Status: &vm_grpc.VMStatus{}}
		msg := types.NewMsgDeployModule(accAddr, [][]byte{{0x1}})

		// Request
		app.MockVMServer.SetResponseDelay(50 * time.Millisecond)
		app.MockVMServer.SetFailCountdown(10)
		app.MockVMServer.SetResponse(vmResp)
		tests.CheckPanicErrorContains(t,
			func() {
				keeper.DeployContract(ctx, msg)
			},
			"failing gRPC execution",
		)
	}
}
