package keeper_test

import (
	"encoding/hex"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dfinance/dvm-proto/go/vm_grpc"
	"github.com/stretchr/testify/require"

	"github.com/dfinance/dstation/pkg/mock"
	"github.com/dfinance/dstation/pkg/tests"
	dnTypes "github.com/dfinance/dstation/pkg/types"
	"github.com/dfinance/dstation/x/vm/types"
)

func TestVM_DeployContract(t *testing.T) {
	app := tests.SetupDSimApp(tests.WithMockVM())
	defer app.TearDown()

	ctx, keeper := app.GetContext(), app.DnApp.VmKeeper

	// Build msg
	accAddr, _, _ := tests.GenAccAddress()
	vmResp := &vm_grpc.VMExecuteResponse{
		WriteSet: []*vm_grpc.VMValue{
			{
				Type: vm_grpc.VmWriteOp_Value,
				Path: &vm_grpc.VMAccessPath{
					Address: types.Bech32ToLibra(accAddr),
					Path:    mock.GetRandomBytes(mock.VMAddressLength),
				},
				Value: mock.GetRandomBytes(512),
			},
		},
		Events:  nil,
		GasUsed: 10000,
		Status:  &vm_grpc.VMStatus{},
	}
	msg := types.NewMsgDeployModule(accAddr, [][]byte{{0x1}})

	// Request
	{
		app.MockVMServer.SetResponse(vmResp)
		require.NoError(t, keeper.DeployContract(ctx, msg))
	}

	// Check events
	{
		events := ctx.EventManager().Events()
		require.Len(t, events, 2)

		require.EqualValues(t, sdk.EventTypeMessage, events[0].Type)
		require.Len(t, events[0].Attributes, 1)
		require.EqualValues(t, sdk.AttributeKeyModule, events[0].Attributes[0].Key)
		require.EqualValues(t, types.ModuleName, events[0].Attributes[0].Value)

		require.EqualValues(t, types.EventTypeContractStatus, events[1].Type)
		require.Len(t, events[1].Attributes, 1)
		require.EqualValues(t, types.AttributeStatus, events[1].Attributes[0].Key)
		require.EqualValues(t, types.AttributeValueStatusKeep, events[1].Attributes[0].Value)
	}

	// Check writeSets
	{
		rcvValue := keeper.GetValue(ctx, vmResp.WriteSet[0].Path)
		require.EqualValues(t, vmResp.WriteSet[0].Value, rcvValue)
	}
}

func TestVM_ExecuteScript(t *testing.T) {
	app := tests.SetupDSimApp(tests.WithMockVM())
	defer app.TearDown()

	ctx, keeper := app.GetContext(), app.DnApp.VmKeeper

	// Build msg
	accAddr, _, _ := tests.GenAccAddress()
	vmResp := &vm_grpc.VMExecuteResponse{
		WriteSet: []*vm_grpc.VMValue{
			{
				Type: vm_grpc.VmWriteOp_Value,
				Path: &vm_grpc.VMAccessPath{
					Address: types.Bech32ToLibra(accAddr),
					Path:    mock.GetRandomBytes(mock.VMAddressLength),
				},
				Value: mock.GetRandomBytes(512),
			},
			{
				Type: vm_grpc.VmWriteOp_Deletion,
				Path: &vm_grpc.VMAccessPath{
					Address: types.Bech32ToLibra(accAddr),
					Path:    mock.GetRandomBytes(mock.VMAddressLength),
				},
				Value: mock.GetRandomBytes(256),
			},
		},
		Events: []*vm_grpc.VMEvent{
			{
				SenderAddress: mock.VMStdLibAddress,
				EventType: &vm_grpc.LcsTag{
					TypeTag: vm_grpc.LcsType_LcsU64,
				},
				EventData: []byte{0x10},
			},
		},
		GasUsed: 10000,
		Status:  &vm_grpc.VMStatus{},
	}
	msg := types.NewMsgExecuteScript(accAddr, []byte{0x1})

	// Set writeSet for the next delete writeOp
	{
		keeper.SetValue(ctx, vmResp.WriteSet[1].Path, vmResp.WriteSet[1].Value)
		require.True(t, keeper.HasValue(ctx, vmResp.WriteSet[1].Path))
	}

	// Request
	{
		app.MockVMServer.SetResponse(vmResp)
		require.NoError(t, keeper.ExecuteScript(ctx, msg))
	}

	// Check events
	{
		events := ctx.EventManager().Events()
		require.Len(t, events, 3)

		// Module message
		require.EqualValues(t, sdk.EventTypeMessage, events[0].Type)
		require.Len(t, events[0].Attributes, 1)
		require.EqualValues(t, sdk.AttributeKeyModule, events[0].Attributes[0].Key)
		require.EqualValues(t, types.ModuleName, events[0].Attributes[0].Value)

		require.EqualValues(t, types.EventTypeContractStatus, events[1].Type)
		require.Len(t, events[1].Attributes, 1)
		require.EqualValues(t, types.AttributeStatus, events[1].Attributes[0].Key)
		require.EqualValues(t, types.AttributeValueStatusKeep, events[1].Attributes[0].Value)

		require.EqualValues(t, types.EventTypeMoveEvent, events[2].Type)
		require.Len(t, events[2].Attributes, 4)

		require.EqualValues(t,
			types.AttributeVmEventSender,
			events[2].Attributes[0].Key,
		)
		require.EqualValues(t,
			types.StdLibAddressShortStr,
			events[2].Attributes[0].Value,
		)

		require.EqualValues(t,
			types.AttributeVmEventSource,
			events[2].Attributes[1].Key,
		)
		require.EqualValues(t,
			types.AttributeValueSourceScript,
			events[2].Attributes[1].Value,
		)

		require.EqualValues(t,
			types.AttributeVmEventType,
			events[2].Attributes[2].Key,
		)
		require.EqualValues(t,
			[]byte("u64"),
			events[2].Attributes[2].Value,
		)

		require.EqualValues(t,
			types.AttributeVmEventData,
			events[2].Attributes[3].Key,
		)
		require.EqualValues(t,
			hex.EncodeToString(vmResp.Events[0].EventData),
			events[2].Attributes[3].Value,
		)
	}

	// Check writeSets
	{
		rcvValue := keeper.GetValue(ctx, vmResp.WriteSet[0].Path)
		require.EqualValues(t, vmResp.WriteSet[0].Value, rcvValue)

		require.False(t, keeper.HasValue(ctx, vmResp.WriteSet[1].Path))
	}
}

func TestVM_FailedExecution(t *testing.T) {
	app := tests.SetupDSimApp(tests.WithMockVM())
	defer app.TearDown()

	ctx, keeper := app.GetContext(), app.DnApp.VmKeeper

	// Status: nil
	{
		// Build msg
		accAddr, _, _ := tests.GenAccAddress()
		vmResp := &vm_grpc.VMExecuteResponse{
			WriteSet: nil,
			Events:   nil,
			GasUsed:  10000,
			Status:   &vm_grpc.VMStatus{},
		}
		msg := types.NewMsgExecuteScript(accAddr, []byte{0x1})

		// Request
		ctx := ctx.WithEventManager(sdk.NewEventManager())
		app.MockVMServer.SetResponse(vmResp)
		require.NoError(t, keeper.ExecuteScript(ctx, msg))

		// Check events
		events := ctx.EventManager().Events()
		require.Len(t, events, 2)

		require.Equal(t, dnTypes.NewModuleNameEvent(types.ModuleName), events[0])

		require.Equal(t, types.EventTypeContractStatus, events[1].Type)
		require.Len(t, events[1].Attributes, 1)

		require.Equal(t,
			[]byte(types.AttributeStatus),
			events[1].Attributes[0].Key,
		)
		require.Equal(t,
			[]byte(types.AttributeValueStatusKeep),
			events[1].Attributes[0].Value,
		)
	}

	// Status: failure
	{
		// Build msg
		accAddr, _, _ := tests.GenAccAddress()
		vmResp := &vm_grpc.VMExecuteResponse{
			WriteSet: nil,
			Events:   nil,
			GasUsed:  10000,
			Status: &vm_grpc.VMStatus{
				Error: &vm_grpc.VMStatus_ExecutionFailure{
					ExecutionFailure: &vm_grpc.Failure{
						StatusCode: 100,
					},
				},
				Message: &vm_grpc.Message{
					Text: "something went wrong",
				},
			},
		}
		msg := types.NewMsgExecuteScript(accAddr, []byte{0x1})

		// Request
		ctx := ctx.WithEventManager(sdk.NewEventManager())
		app.MockVMServer.SetResponse(vmResp)
		require.NoError(t, keeper.ExecuteScript(ctx, msg))

		// Check events
		events := ctx.EventManager().Events()
		require.Len(t, events, 2)

		require.Equal(t, dnTypes.NewModuleNameEvent(types.ModuleName), events[0])

		require.Equal(t, types.EventTypeContractStatus, events[1].Type)
		require.Len(t, events[1].Attributes, 4)

		require.Equal(t,
			[]byte(types.AttributeStatus),
			events[1].Attributes[0].Key,
		)
		require.Equal(t,
			[]byte(types.AttributeValueStatusDiscard),
			events[1].Attributes[0].Value,
		)

		require.Equal(t,
			[]byte(types.AttributeErrMajorStatus),
			events[1].Attributes[1].Key,
		)
		require.Equal(t,
			[]byte("100"),
			events[1].Attributes[1].Value,
		)

		require.Equal(t,
			[]byte(types.AttributeErrSubStatus),
			events[1].Attributes[2].Key,
		)
		require.Equal(t,
			[]byte("0"),
			events[1].Attributes[2].Value,
		)

		require.Equal(t,
			[]byte(types.AttributeErrMessage),
			events[1].Attributes[3].Key,
		)
		require.Equal(t,
			[]byte(vmResp.Status.Message.Text),
			events[1].Attributes[3].Value,
		)
	}

	// Status: abort
	{
		// Build msg
		accAddr, _, _ := tests.GenAccAddress()
		vmResp := &vm_grpc.VMExecuteResponse{
			WriteSet: nil,
			Events:   nil,
			GasUsed:  10000,
			Status: &vm_grpc.VMStatus{
				Error: &vm_grpc.VMStatus_Abort{
					Abort: &vm_grpc.Abort{
						AbortCode: 100,
					},
				},
			},
		}
		msg := types.NewMsgExecuteScript(accAddr, []byte{0x1})

		// Request
		ctx := ctx.WithEventManager(sdk.NewEventManager())
		app.MockVMServer.SetResponse(vmResp)
		require.NoError(t, keeper.ExecuteScript(ctx, msg))

		// Check events
		events := ctx.EventManager().Events()
		require.Len(t, events, 2)

		require.Equal(t, dnTypes.NewModuleNameEvent(types.ModuleName), events[0])

		require.Equal(t, types.EventTypeContractStatus, events[1].Type)
		require.Len(t, events[1].Attributes, 3)

		require.Equal(t,
			[]byte(types.AttributeStatus),
			events[1].Attributes[0].Key,
		)
		require.Equal(t,
			[]byte(types.AttributeValueStatusDiscard),
			events[1].Attributes[0].Value,
		)

		require.Equal(t,
			[]byte(types.AttributeErrMajorStatus),
			events[1].Attributes[1].Key,
		)
		require.Equal(t,
			[]byte("4016"),
			events[1].Attributes[1].Value,
		)

		require.Equal(t,
			[]byte(types.AttributeErrSubStatus),
			events[1].Attributes[2].Key,
		)
		require.Equal(t,
			[]byte("100"),
			events[1].Attributes[2].Value,
		)
	}

	// Status: move error
	{
		// Build msg
		accAddr, _, _ := tests.GenAccAddress()
		vmResp := &vm_grpc.VMExecuteResponse{
			WriteSet: nil,
			Events:   nil,
			GasUsed:  10000,
			Status: &vm_grpc.VMStatus{
				Error: &vm_grpc.VMStatus_MoveError{
					MoveError: &vm_grpc.MoveError{
						StatusCode: 100,
					},
				},
			},
		}
		msg := types.NewMsgExecuteScript(accAddr, []byte{0x1})

		// Request
		ctx := ctx.WithEventManager(sdk.NewEventManager())
		app.MockVMServer.SetResponse(vmResp)
		require.NoError(t, keeper.ExecuteScript(ctx, msg))

		// Check events
		events := ctx.EventManager().Events()
		require.Len(t, events, 2)

		require.Equal(t, dnTypes.NewModuleNameEvent(types.ModuleName), events[0])

		require.Equal(t, types.EventTypeContractStatus, events[1].Type)
		require.Len(t, events[1].Attributes, 3)

		require.Equal(t,
			[]byte(types.AttributeStatus),
			events[1].Attributes[0].Key,
		)
		require.Equal(t,
			[]byte(types.AttributeValueStatusDiscard),
			events[1].Attributes[0].Value,
		)

		require.Equal(t,
			[]byte(types.AttributeErrMajorStatus),
			events[1].Attributes[1].Key,
		)
		require.Equal(t,
			[]byte("100"),
			events[1].Attributes[1].Value,
		)

		require.Equal(t,
			[]byte(types.AttributeErrSubStatus),
			events[1].Attributes[2].Key,
		)
		require.Equal(t,
			[]byte("0"),
			events[1].Attributes[2].Value,
		)
	}
}
