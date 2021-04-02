package keeper_test

import (
	"encoding/hex"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/dfinance/dstation/pkg/mock"
	"github.com/dfinance/dstation/pkg/tests"
	dnTypes "github.com/dfinance/dstation/pkg/types"
	"github.com/dfinance/dstation/pkg/types/dvm"
	"github.com/dfinance/dstation/x/vm/types"
)

func (s *KeeperMockVmTestSuite) TestDeployContract() {
	ctx, keeper, vmServer := s.ctx, s.keeper, s.vmServer

	// Build msg
	accAddr, _, _ := tests.GenAccAddress()
	vmResp := &dvm.VMExecuteResponse{
		WriteSet: []*dvm.VMValue{
			{
				Type: dvm.VmWriteOp_Value,
				Path: &dvm.VMAccessPath{
					Address: types.Bech32ToLibra(accAddr),
					Path:    mock.GetRandomBytes(mock.VMAddressLength),
				},
				Value: mock.GetRandomBytes(512),
			},
		},
		Events:  nil,
		GasUsed: 10000,
		Status:  &dvm.VMStatus{},
	}
	msg := types.NewMsgDeployModule(accAddr, []byte{0x1})

	// Request
	{
		vmServer.SetResponse(vmResp)
		s.Require().NoError(keeper.DeployContract(ctx, msg))
	}

	// Check events
	{
		events := ctx.EventManager().Events()
		s.Require().Len(events, 2)

		s.Require().EqualValues(sdk.EventTypeMessage, events[0].Type)
		s.Require().Len(events[0].Attributes, 1)
		s.Require().EqualValues(sdk.AttributeKeyModule, events[0].Attributes[0].Key)
		s.Require().EqualValues(types.ModuleName, events[0].Attributes[0].Value)

		s.Require().EqualValues(types.EventTypeContractStatus, events[1].Type)
		s.Require().Len(events[1].Attributes, 1)
		s.Require().EqualValues(types.AttributeStatus, events[1].Attributes[0].Key)
		s.Require().EqualValues(types.AttributeValueStatusKeep, events[1].Attributes[0].Value)
	}

	// Check writeSets
	{
		rcvValue := keeper.GetValue(ctx, vmResp.WriteSet[0].Path)
		s.Require().EqualValues(vmResp.WriteSet[0].Value, rcvValue)
	}
}

func (s *KeeperMockVmTestSuite) TestExecuteContract() {
	ctx, keeper, vmServer := s.ctx, s.keeper, s.vmServer

	// Build msg
	accAddr, _, _ := tests.GenAccAddress()
	vmResp := &dvm.VMExecuteResponse{
		WriteSet: []*dvm.VMValue{
			{
				Type: dvm.VmWriteOp_Value,
				Path: &dvm.VMAccessPath{
					Address: types.Bech32ToLibra(accAddr),
					Path:    mock.GetRandomBytes(mock.VMAddressLength),
				},
				Value: mock.GetRandomBytes(512),
			},
			{
				Type: dvm.VmWriteOp_Deletion,
				Path: &dvm.VMAccessPath{
					Address: types.Bech32ToLibra(accAddr),
					Path:    mock.GetRandomBytes(mock.VMAddressLength),
				},
				Value: mock.GetRandomBytes(256),
			},
		},
		Events: []*dvm.VMEvent{
			{
				SenderAddress: mock.VMStdLibAddress,
				EventType: &dvm.LcsTag{
					TypeTag: dvm.LcsType_LcsU64,
				},
				EventData: []byte{0x10},
			},
		},
		GasUsed: 10000,
		Status:  &dvm.VMStatus{},
	}
	msg := types.NewMsgExecuteScript(accAddr, []byte{0x1})

	// Set writeSet for the next delete writeOp
	{
		keeper.SetValue(ctx, vmResp.WriteSet[1].Path, vmResp.WriteSet[1].Value)
		s.Require().True(keeper.HasValue(ctx, vmResp.WriteSet[1].Path))
	}

	// Request
	{
		vmServer.SetResponse(vmResp)
		s.Require().NoError(keeper.ExecuteContract(ctx, msg))
	}

	// Check events
	{
		events := ctx.EventManager().Events()
		s.Require().Len(events, 3)

		// Module message
		s.Require().EqualValues(sdk.EventTypeMessage, events[0].Type)
		s.Require().Len(events[0].Attributes, 1)
		s.Require().EqualValues(sdk.AttributeKeyModule, events[0].Attributes[0].Key)
		s.Require().EqualValues(types.ModuleName, events[0].Attributes[0].Value)

		s.Require().EqualValues(types.EventTypeContractStatus, events[1].Type)
		s.Require().Len(events[1].Attributes, 1)
		s.Require().EqualValues(types.AttributeStatus, events[1].Attributes[0].Key)
		s.Require().EqualValues(types.AttributeValueStatusKeep, events[1].Attributes[0].Value)

		s.Require().EqualValues(types.EventTypeMoveEvent, events[2].Type)
		s.Require().Len(events[2].Attributes, 4)

		s.Require().EqualValues(types.AttributeVmEventSender,
			events[2].Attributes[0].Key,
		)
		s.Require().EqualValues(types.StdLibAddressShortStr,
			events[2].Attributes[0].Value,
		)

		s.Require().EqualValues(types.AttributeVmEventSource,
			events[2].Attributes[1].Key,
		)
		s.Require().EqualValues(types.AttributeValueSourceScript,
			events[2].Attributes[1].Value,
		)

		s.Require().EqualValues(types.AttributeVmEventType,
			events[2].Attributes[2].Key,
		)
		s.Require().EqualValues([]byte("u64"),
			events[2].Attributes[2].Value,
		)

		s.Require().EqualValues(types.AttributeVmEventData,
			events[2].Attributes[3].Key,
		)
		s.Require().EqualValues(hex.EncodeToString(vmResp.Events[0].EventData),
			events[2].Attributes[3].Value,
		)
	}

	// Check writeSets
	{
		rcvValue := keeper.GetValue(ctx, vmResp.WriteSet[0].Path)
		s.Require().EqualValues(vmResp.WriteSet[0].Value, rcvValue)

		s.Require().False(keeper.HasValue(ctx, vmResp.WriteSet[1].Path))
	}
}

func (s *KeeperMockVmTestSuite) TestFailedExecution() {
	ctx, keeper, vmServer := s.ctx, s.keeper, s.vmServer

	// Status: nil
	{
		// Build msg
		accAddr, _, _ := tests.GenAccAddress()
		vmResp := &dvm.VMExecuteResponse{
			WriteSet: nil,
			Events:   nil,
			GasUsed:  10000,
			Status:   &dvm.VMStatus{},
		}
		msg := types.NewMsgExecuteScript(accAddr, []byte{0x1})

		// Request
		ctx := ctx.WithEventManager(sdk.NewEventManager())
		vmServer.SetResponse(vmResp)
		s.Require().NoError(keeper.ExecuteContract(ctx, msg))

		// Check events
		events := ctx.EventManager().Events()
		s.Require().Len(events, 2)

		s.Require().Equal(dnTypes.NewModuleNameEvent(types.ModuleName), events[0])

		s.Require().Equal(types.EventTypeContractStatus, events[1].Type)
		s.Require().Len(events[1].Attributes, 1)

		s.Require().Equal(
			[]byte(types.AttributeStatus),
			events[1].Attributes[0].Key,
		)
		s.Require().Equal(
			[]byte(types.AttributeValueStatusKeep),
			events[1].Attributes[0].Value,
		)
	}

	// Status: failure
	{
		// Build msg
		accAddr, _, _ := tests.GenAccAddress()
		vmResp := &dvm.VMExecuteResponse{
			WriteSet: nil,
			Events:   nil,
			GasUsed:  10000,
			Status: &dvm.VMStatus{
				Error: &dvm.VMStatus_ExecutionFailure{
					ExecutionFailure: &dvm.Failure{
						StatusCode: 100,
					},
				},
				Message: &dvm.Message{
					Text: "something went wrong",
				},
			},
		}
		msg := types.NewMsgExecuteScript(accAddr, []byte{0x1})

		// Request
		ctx := ctx.WithEventManager(sdk.NewEventManager())
		vmServer.SetResponse(vmResp)
		s.Require().NoError(keeper.ExecuteContract(ctx, msg))

		// Check events
		events := ctx.EventManager().Events()
		s.Require().Len(events, 2)

		s.Require().Equal(dnTypes.NewModuleNameEvent(types.ModuleName), events[0])

		s.Require().Equal(types.EventTypeContractStatus, events[1].Type)
		s.Require().Len(events[1].Attributes, 4)

		s.Require().Equal(
			[]byte(types.AttributeStatus),
			events[1].Attributes[0].Key,
		)
		s.Require().Equal(
			[]byte(types.AttributeValueStatusDiscard),
			events[1].Attributes[0].Value,
		)

		s.Require().Equal(
			[]byte(types.AttributeErrMajorStatus),
			events[1].Attributes[1].Key,
		)
		s.Require().Equal(
			[]byte("100"),
			events[1].Attributes[1].Value,
		)

		s.Require().Equal(
			[]byte(types.AttributeErrSubStatus),
			events[1].Attributes[2].Key,
		)
		s.Require().Equal(
			[]byte("0"),
			events[1].Attributes[2].Value,
		)

		s.Require().Equal(
			[]byte(types.AttributeErrMessage),
			events[1].Attributes[3].Key,
		)
		s.Require().Equal(
			[]byte(vmResp.Status.Message.Text),
			events[1].Attributes[3].Value,
		)
	}

	// Status: abort
	{
		// Build msg
		accAddr, _, _ := tests.GenAccAddress()
		vmResp := &dvm.VMExecuteResponse{
			WriteSet: nil,
			Events:   nil,
			GasUsed:  10000,
			Status: &dvm.VMStatus{
				Error: &dvm.VMStatus_Abort{
					Abort: &dvm.Abort{
						AbortCode: 100,
					},
				},
			},
		}
		msg := types.NewMsgExecuteScript(accAddr, []byte{0x1})

		// Request
		ctx := ctx.WithEventManager(sdk.NewEventManager())
		vmServer.SetResponse(vmResp)
		s.Require().NoError(keeper.ExecuteContract(ctx, msg))

		// Check events
		events := ctx.EventManager().Events()
		s.Require().Len(events, 2)

		s.Require().Equal(dnTypes.NewModuleNameEvent(types.ModuleName), events[0])

		s.Require().Equal(types.EventTypeContractStatus, events[1].Type)
		s.Require().Len(events[1].Attributes, 3)

		s.Require().Equal(
			[]byte(types.AttributeStatus),
			events[1].Attributes[0].Key,
		)
		s.Require().Equal(
			[]byte(types.AttributeValueStatusDiscard),
			events[1].Attributes[0].Value,
		)

		s.Require().Equal(
			[]byte(types.AttributeErrMajorStatus),
			events[1].Attributes[1].Key,
		)
		s.Require().Equal(
			[]byte("4016"),
			events[1].Attributes[1].Value,
		)

		s.Require().Equal(
			[]byte(types.AttributeErrSubStatus),
			events[1].Attributes[2].Key,
		)
		s.Require().Equal(
			[]byte("100"),
			events[1].Attributes[2].Value,
		)
	}

	// Status: move error
	{
		// Build msg
		accAddr, _, _ := tests.GenAccAddress()
		vmResp := &dvm.VMExecuteResponse{
			WriteSet: nil,
			Events:   nil,
			GasUsed:  10000,
			Status: &dvm.VMStatus{
				Error: &dvm.VMStatus_MoveError{
					MoveError: &dvm.MoveError{
						StatusCode: 100,
					},
				},
			},
		}
		msg := types.NewMsgExecuteScript(accAddr, []byte{0x1})

		// Request
		ctx := ctx.WithEventManager(sdk.NewEventManager())
		vmServer.SetResponse(vmResp)
		s.Require().NoError(keeper.ExecuteContract(ctx, msg))

		// Check events
		events := ctx.EventManager().Events()
		s.Require().Len(events, 2)

		s.Require().Equal(dnTypes.NewModuleNameEvent(types.ModuleName), events[0])

		s.Require().Equal(types.EventTypeContractStatus, events[1].Type)
		s.Require().Len(events[1].Attributes, 3)

		s.Require().Equal(
			[]byte(types.AttributeStatus),
			events[1].Attributes[0].Key,
		)
		s.Require().Equal(
			[]byte(types.AttributeValueStatusDiscard),
			events[1].Attributes[0].Value,
		)

		s.Require().Equal(
			[]byte(types.AttributeErrMajorStatus),
			events[1].Attributes[1].Key,
		)
		s.Require().Equal(
			[]byte("100"),
			events[1].Attributes[1].Value,
		)

		s.Require().Equal(
			[]byte(types.AttributeErrSubStatus),
			events[1].Attributes[2].Key,
		)
		s.Require().Equal(
			[]byte("0"),
			events[1].Attributes[2].Value,
		)
	}
}
