package vm_test

import (
	"context"
	"encoding/hex"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkErrors "github.com/cosmos/cosmos-sdk/types/errors"

	dnConfig "github.com/dfinance/dstation/cmd/dstation/config"
	"github.com/dfinance/dstation/pkg/tests"
	dvmTypes "github.com/dfinance/dstation/pkg/types/dvm"
	"github.com/dfinance/dstation/x/vm/client"
	"github.com/dfinance/dstation/x/vm/types"
)

func (s *ModuleDVVTestSuite) TestScriptWithArgs() {
	ctx := s.ctx
	acc, accPrivKey := s.app.AddAccount(ctx, sdk.NewCoin(dnConfig.MainDenom, sdk.NewInt(1000)))

	s.CheckContractExecuted(
		s.ExecuteScript(acc.GetAddress(), accPrivKey,
			"scripts/arg_types.move", nil,
			s.BuildScriptArg("128", client.NewU8ScriptArg),
			s.BuildScriptArg("1000000", client.NewU64ScriptArg),
			s.BuildScriptArg("100000000000000000000000000000", client.NewU128ScriptArg),
			s.BuildScriptArg(acc.GetAddress().String(), client.NewAddressScriptArg),
			s.BuildScriptArg("true", client.NewBoolScriptArg),
			s.BuildScriptArg("false", client.NewBoolScriptArg),
			s.BuildScriptArg("0x0001", client.NewVectorScriptArg),
		),
	)
}

func (s *ModuleDVVTestSuite) TestScriptWithError() {
	ctx := s.ctx
	acc, accPrivKey := s.app.AddAccount(ctx, sdk.NewCoin(dnConfig.MainDenom, sdk.NewInt(1000)))

	s.CheckContractFailed(
		s.ExecuteScript(acc.GetAddress(), accPrivKey,
			"scripts/with_error.move", nil,
			s.BuildScriptArg("1", client.NewU64ScriptArg),
		),
	)
}

func (s *ModuleDVVTestSuite) TestEventTypeSerialization() {
	ctx := s.ctx
	acc, accPrivKey := s.app.AddAccount(ctx, sdk.NewCoin(dnConfig.MainDenom, sdk.NewInt(1000)))

	s.CheckContractExecuted(
		s.DeployModule(acc.GetAddress(), accPrivKey,
			"event_type_serialization/basic/module.move", nil,
		),
	)

	_, scriptEvents := s.CheckContractExecuted(
		s.ExecuteScript(acc.GetAddress(), accPrivKey,
			"event_type_serialization/basic/script.template.move", []string{
				s.GetLibraAccAddressString(acc.GetAddress()),
			},
		),
	)

	//s.PrintABCIEvents(scriptEvents)
	s.CheckABCIEventsContain(scriptEvents, []sdk.Event{
		sdk.NewEvent(
			types.EventTypeMoveEvent,
			sdk.NewAttribute(types.AttributeVmEventSender, acc.GetAddress().String()),
			sdk.NewAttribute(types.AttributeVmEventSource, types.AttributeValueSourceScript),
			sdk.NewAttribute(types.AttributeVmEventType, "u8"),
			sdk.NewAttribute(types.AttributeVmEventData, "80"),
		),
		sdk.NewEvent(
			types.EventTypeMoveEvent,
			sdk.NewAttribute(types.AttributeVmEventSender, acc.GetAddress().String()),
			sdk.NewAttribute(types.AttributeVmEventSource, types.AttributeValueSourceScript),
			sdk.NewAttribute(types.AttributeVmEventType, "vector<u8>"),
			sdk.NewAttribute(types.AttributeVmEventData, "020102"),
		),
		sdk.NewEvent(
			types.EventTypeMoveEvent,
			sdk.NewAttribute(types.AttributeVmEventSender, acc.GetAddress().String()),
			sdk.NewAttribute(types.AttributeVmEventSource, fmt.Sprintf("%s::Foo", acc.GetAddress())),
			sdk.NewAttribute(types.AttributeVmEventType, "bool"),
			sdk.NewAttribute(types.AttributeVmEventData, "01"),
		),
		sdk.NewEvent(
			types.EventTypeMoveEvent,
			sdk.NewAttribute(types.AttributeVmEventSender, acc.GetAddress().String()),
			sdk.NewAttribute(types.AttributeVmEventSource, types.AttributeValueSourceScript),
			sdk.NewAttribute(types.AttributeVmEventType, fmt.Sprintf("%s::Foo::FooEvent<u64, vector<u8>>", acc.GetAddress())),
			sdk.NewAttribute(types.AttributeVmEventData, "e803000000000000020102"),
		),
	})
}

func (s *ModuleDVVTestSuite) TestEventTypeSerializationGasCalculation() {
	ctx := s.ctx
	acc, accPrivKey := s.app.AddAccount(ctx, sdk.NewCoin(dnConfig.MainDenom, sdk.NewInt(1000)))

	s.CheckContractExecuted(
		s.DeployModule(acc.GetAddress(), accPrivKey,
			"event_type_serialization/gas_calculation/module.move", nil,
		),
	)

	gasInfo, scriptEvents := s.CheckContractExecuted(
		s.ExecuteScript(acc.GetAddress(), accPrivKey,
			"event_type_serialization/gas_calculation/script.template.move", []string{
				s.GetLibraAccAddressString(acc.GetAddress()),
			},
		),
	)

	//s.PrintABCIEvents(scriptEvents)
	s.CheckABCIEventsContain(scriptEvents, []sdk.Event{
		sdk.NewEvent(
			types.EventTypeMoveEvent,
			sdk.NewAttribute(types.AttributeVmEventSender, acc.GetAddress().String()),
			sdk.NewAttribute(types.AttributeVmEventSource, fmt.Sprintf("%s::GasEvent", acc.GetAddress())),
			sdk.NewAttribute(types.AttributeVmEventType, fmt.Sprintf("%s::GasEvent::D<%s::GasEvent::C<%s::GasEvent::B<%s::GasEvent::A>>>", acc.GetAddress(), acc.GetAddress(), acc.GetAddress(), acc.GetAddress())),
			sdk.NewAttribute(types.AttributeVmEventData, "0a00000000000000"),
		),
	})

	expMinGasUsed := uint64(0)
	for i := 1; i <= 4-types.EventTypeNoGasLevels; i++ {
		expMinGasUsed += uint64(i) * types.EventTypeProcessingGas
	}

	s.T().Logf("Consumed gas / expected min consumed gas: %d / %d", gasInfo.GasUsed, expMinGasUsed)
	s.Require().GreaterOrEqual(gasInfo.GasUsed, expMinGasUsed)
}

func (s *ModuleDVVTestSuite) TestEventTypeSerializationOutOfGas() {
	ctx := s.ctx
	acc, accPrivKey := s.app.AddAccount(ctx, sdk.NewCoin(dnConfig.MainDenom, sdk.NewInt(1000)))

	// Deploy module
	s.CheckContractExecuted(
		s.DeployModule(acc.GetAddress(), accPrivKey,
			"event_type_serialization/gas_limit/module.move", nil,
		),
	)

	// Estimate event serialization gas usage
	expMinGasUsed := uint64(0)
	for i := 1; i <= 6-types.EventTypeNoGasLevels; i++ {
		expMinGasUsed += uint64(i) * types.EventTypeProcessingGas
	}

	// Compile, execute script
	byteCodes := s.CompileMoveFile(acc.GetAddress(), "event_type_serialization/gas_limit/script.template.move", s.GetLibraAccAddressString(acc.GetAddress()))
	s.Require().Len(byteCodes, 1)
	msg := types.NewMsgExecuteScript(acc.GetAddress(), byteCodes[0])
	_, _, err := s.app.DeliverTx(ctx, acc.GetAddress(), accPrivKey, []sdk.Msg{&msg}, tests.TxWithGasLimit(expMinGasUsed))

	s.Require().True(sdkErrors.ErrOutOfGas.Is(err))
}

func (s *ModuleDVVTestSuite) TestDeployModule() {
	ctx := s.ctx
	acc, accPrivKey := s.app.AddAccount(ctx, sdk.NewCoin(dnConfig.MainDenom, sdk.NewInt(1000)))

	s.CheckContractExecuted(
		s.DeployModule(acc.GetAddress(), accPrivKey,
			"math/module.move", nil,
		),
	)

	_, scriptEvents := s.CheckContractExecuted(
		s.ExecuteScript(acc.GetAddress(), accPrivKey,
			"math/script.template.move", []string{
				s.GetLibraAccAddressString(acc.GetAddress()),
			},
			s.BuildScriptArg("10", client.NewU64ScriptArg),
			s.BuildScriptArg("100", client.NewU64ScriptArg),
		),
	)

	// uint64 -> Little-endian bytes -> HEX string
	expResult := uint64(110)
	expResultVM, err := client.NewU64ScriptArg(strconv.FormatUint(expResult, 10))
	s.Require().NoError(err)
	expResultVMEventAttr := hex.EncodeToString(expResultVM.Value)

	s.CheckABCIEventsContain(scriptEvents, []sdk.Event{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeySender, acc.GetAddress().String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyAction, types.TypeMsgExecuteScript),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		),
		sdk.NewEvent(
			types.EventTypeContractStatus,
			sdk.NewAttribute(types.AttributeStatus, types.AttributeValueStatusKeep),
		),
		sdk.NewEvent(
			types.EventTypeMoveEvent,
			sdk.NewAttribute(types.AttributeVmEventSender, acc.GetAddress().String()),
			sdk.NewAttribute(types.AttributeVmEventSource, types.AttributeValueSourceScript),
			sdk.NewAttribute(types.AttributeVmEventType, "u64"),
			sdk.NewAttribute(types.AttributeVmEventData, expResultVMEventAttr),
		),
	})
}

func (s *ModuleDVVTestSuite) TestCompileMetadata() {
	ctx, queryClient := s.ctx, s.queryClient
	acc, _ := s.app.AddAccount(ctx, sdk.NewCoin(dnConfig.MainDenom, sdk.NewInt(1000)))

	// Compile
	byteCode := s.GetMoveFileContent("modules/with_resources.move")
	resp, err := queryClient.Compile(context.Background(), &types.QueryCompileRequest{
		Address: types.Bech32ToLibra(acc.GetAddress()),
		Code:    string(byteCode),
	})
	s.Require().NoError(err)
	s.Require().NotNil(resp)

	// Check the compiled meta
	s.Require().Len(resp.CompiledItems, 1)
	s.Require().Equal(resp.CompiledItems[0].CodeType, types.CompiledItem_MODULE)
	s.Require().Equal(resp.CompiledItems[0].Name, "Foo")
	s.Require().NotEmpty(resp.CompiledItems[0].ByteCode)
	s.Require().ElementsMatch(resp.CompiledItems[0].Methods, []*dvmTypes.Function{
		{
			Name:           "add",
			IsPublic:       true,
			IsNative:       false,
			TypeParameters: nil,
			Arguments:      []string{"&signer", "u64", "u64"},
			Returns:        []string{"u64"},
		},
		{
			Name:           "build_obj",
			IsPublic:       true,
			IsNative:       false,
			TypeParameters: nil,
			Arguments:      []string{"&signer", "u64"},
			Returns:        nil,
		},
	})
	s.Require().ElementsMatch(resp.CompiledItems[0].Types, []*dvmTypes.Struct{
		{
			Name:           "Obj",
			IsResource:     true,
			TypeParameters: nil,
			Field: []*dvmTypes.Field{
				{
					Name: "val",
					Type: "u64",
				},
				{
					Name: "o",
					Type: "U64",
				},
			},
		},
		{
			Name:           "U64",
			IsResource:     true,
			TypeParameters: nil,
			Field: []*dvmTypes.Field{
				{
					Name: "val",
					Type: "u64",
				},
			},
		},
	})
}

func (s *ModuleDVVTestSuite) TestNativeBalanceRes() {
	ctx := s.ctx
	acc, accPrivKey := s.app.AddAccount(ctx, sdk.NewCoin(dnConfig.MainDenom, sdk.NewInt(1000)))

	// Pass the account balance - 1xfi (fee) to the script for assert
	s.CheckContractExecuted(
		s.ExecuteScript(acc.GetAddress(), accPrivKey,
			"scripts/native_balance.move", nil,
			s.BuildScriptArg("999", client.NewU128ScriptArg),
		),
	)
}

func (s *ModuleDVVTestSuite) TestNativeOracleRes() {
	ctx := s.ctx
	acc, accPrivKey := s.app.AddAccount(ctx, sdk.NewCoin(dnConfig.MainDenom, sdk.NewInt(1000)))

	// TODO: mock test without direct/reverse
	s.CheckContractExecuted(
		s.ExecuteScript(acc.GetAddress(), accPrivKey,
			"scripts/oracle_price_direct.move", nil,
			s.BuildScriptArg("100", client.NewU128ScriptArg),
		),
	)
}

func (s *ModuleDVVTestSuite) TestCurrencyInfoRes() {
	ctx := s.ctx
	acc, accPrivKey := s.app.AddAccount(ctx, sdk.NewCoin(dnConfig.MainDenom, sdk.NewInt(1000)))

	s.CheckContractExecuted(
		s.ExecuteScript(acc.GetAddress(), accPrivKey,
			"scripts/currency_infos.move", nil,
			s.BuildScriptArg("18", client.NewU8ScriptArg), // xfi
			s.BuildScriptArg("18", client.NewU8ScriptArg), // eth
			s.BuildScriptArg("8", client.NewU8ScriptArg),  // btc
			s.BuildScriptArg("6", client.NewU8ScriptArg),  // usdt
		),
	)
}

func (s *ModuleDVVTestSuite) TestBlockTimeTxMeta() {
	ctx := s.ctx
	acc, accPrivKey := s.app.AddAccount(ctx, sdk.NewCoin(dnConfig.MainDenom, sdk.NewInt(1000)))

	s.app.BeginBlock()
	s.app.EndBlock()
	blockHeight, blockTime := s.app.GetNextBlockHeightTime() // ExecuteScript starts a new block, so peek next values

	s.CheckContractExecuted(
		s.ExecuteScript(acc.GetAddress(), accPrivKey,
			"scripts/block_height_time.move", nil,
			s.BuildScriptArg(strconv.FormatInt(blockHeight, 10), client.NewU64ScriptArg),
			s.BuildScriptArg(strconv.FormatInt(blockTime.Unix(), 10), client.NewU64ScriptArg),
		),
	)
}

func (s *ModuleDVVTestSuite) TestNativeDepositWithdraw() {
	ctx, queryClient, bankKeeper := s.ctx, s.queryClient, s.app.DnApp.BankKeeper
	acc, accPrivKey := s.app.AddAccount(ctx, sdk.NewCoin(dnConfig.MainDenom, sdk.NewInt(1000)))

	delCoin := sdk.NewInt64Coin(dnConfig.MainDenom, 500)

	// Query initial DecPool supply
	getDecPoolSupply := func() sdk.Coins {
		resp, err := queryClient.DelegatedPoolSupply(context.Background(), &types.QueryDelegatedPoolSupplyRequest{})
		s.Require().NoError(err)
		s.Require().NotNil(resp)
		return resp.Coins
	}
	prevDecPoolCoins := getDecPoolSupply()

	// Deposit
	{
		s.CheckContractExecuted(
			s.ExecuteScript(acc.GetAddress(), accPrivKey,
				"scripts/native_deposit.move", nil,
				s.BuildScriptArg(delCoin.Amount.String(), client.NewU128ScriptArg),
			),
		)

		// Check bank balance (-1xfi for the fee)
		accBalance := bankKeeper.GetBalance(ctx, acc.GetAddress(), dnConfig.MainDenom)
		s.Require().EqualValues(delCoin.Amount.Int64()-1, accBalance.Amount.Int64())

		// Check DelPool supply
		curDecPoolCoins := getDecPoolSupply()
		s.Require().True(
			prevDecPoolCoins.
				Add(delCoin).
				IsEqual(curDecPoolCoins),
		)
		prevDecPoolCoins = curDecPoolCoins
	}

	// Withdraw
	{
		s.CheckContractExecuted(
			s.ExecuteScript(acc.GetAddress(), accPrivKey,
				"scripts/native_withdraw.move", nil,
				s.BuildScriptArg("500", client.NewU128ScriptArg),
			),
		)

		// Check bank balance (-1xfi for the fee)
		accBalance := bankKeeper.GetBalance(ctx, acc.GetAddress(), dnConfig.MainDenom)
		s.Require().EqualValues(998, accBalance.Amount.Int64())

		// Check DelPool supply
		curDecPoolCoins := getDecPoolSupply()
		s.Require().True(
			prevDecPoolCoins.
				Sub(sdk.NewCoins(delCoin)).
				IsEqual(curDecPoolCoins),
		)
		prevDecPoolCoins = curDecPoolCoins
	}
}

func (s *ModuleDVVTestSuite) TestTransferCoins() {
	ctx, queryClient, bankKeeper := s.ctx, s.queryClient, s.app.DnApp.BankKeeper

	// Query initial DecPool supply
	getDecPoolSupply := func() sdk.Coins {
		resp, err := queryClient.DelegatedPoolSupply(context.Background(), &types.QueryDelegatedPoolSupplyRequest{})
		s.Require().NoError(err)
		s.Require().NotNil(resp)
		return resp.Coins
	}
	prevDecPoolCoins := getDecPoolSupply()

	// Define initial accounts balances, transfer amounts and expected final balances
	acc1InitCoins := sdk.NewCoins(
		sdk.NewCoin(dnConfig.MainDenom, sdk.NewInt(1000)),
		sdk.NewCoin(dnConfig.EthDenom, sdk.NewInt(1000)),
		sdk.NewCoin(dnConfig.BtcDenom, sdk.NewInt(1000)),
		sdk.NewCoin(dnConfig.UsdtDenom, sdk.NewInt(1000)),
	)
	acc2InitCoins := sdk.NewCoins(
		sdk.NewCoin(dnConfig.MainDenom, sdk.NewInt(1)),
	)

	transferCoins := sdk.NewCoins(
		sdk.NewInt64Coin(dnConfig.MainDenom, 500),
		sdk.NewInt64Coin(dnConfig.EthDenom, 450),
		sdk.NewInt64Coin(dnConfig.BtcDenom, 400),
		sdk.NewInt64Coin(dnConfig.UsdtDenom, 350),
	)

	acc1ExpCoins := acc1InitCoins.
		Sub(transferCoins).
		Sub(sdk.NewCoins(sdk.NewInt64Coin(dnConfig.MainDenom, 1)))
	acc2ExpCoins := transferCoins

	acc1, acc1PrivKey := s.app.AddAccount(ctx, acc1InitCoins...)
	acc2, acc2PrivKey := s.app.AddAccount(ctx, acc2InitCoins...)

	// Account 1 -> Account 2: deposit and send via VM
	{
		s.CheckContractExecuted(
			s.ExecuteScript(acc1.GetAddress(), acc1PrivKey,
				"scripts/transfer_send.move", nil,
				s.BuildScriptArg(acc2.GetAddress().String(), client.NewAddressScriptArg),
				s.BuildScriptArg(transferCoins.AmountOf(dnConfig.MainDenom).String(), client.NewU128ScriptArg),
				s.BuildScriptArg(transferCoins.AmountOf(dnConfig.EthDenom).String(), client.NewU128ScriptArg),
				s.BuildScriptArg(transferCoins.AmountOf(dnConfig.BtcDenom).String(), client.NewU128ScriptArg),
				s.BuildScriptArg(transferCoins.AmountOf(dnConfig.UsdtDenom).String(), client.NewU128ScriptArg),
			),
		)

		// Check bank balance
		acc1CurCoins := bankKeeper.GetAllBalances(ctx, acc1.GetAddress())
		s.Require().True(acc1ExpCoins.IsEqual(acc1CurCoins))

		acc2CurCoins := bankKeeper.GetAllBalances(ctx, acc2.GetAddress())
		s.Require().True(acc2InitCoins.IsEqual(acc2CurCoins))

		// Check DelPool supply
		curDecPoolCoins := getDecPoolSupply()
		s.Require().True(
			prevDecPoolCoins.
				Add(transferCoins...).
				IsEqual(curDecPoolCoins),
		)
		prevDecPoolCoins = curDecPoolCoins
	}

	// Account 2: withdraw
	{
		s.CheckContractExecuted(
			s.ExecuteScript(acc2.GetAddress(), acc2PrivKey,
				"scripts/transfer_receive.move", nil,
				s.BuildScriptArg(transferCoins.AmountOf(dnConfig.MainDenom).String(), client.NewU128ScriptArg),
				s.BuildScriptArg(transferCoins.AmountOf(dnConfig.EthDenom).String(), client.NewU128ScriptArg),
				s.BuildScriptArg(transferCoins.AmountOf(dnConfig.BtcDenom).String(), client.NewU128ScriptArg),
				s.BuildScriptArg(transferCoins.AmountOf(dnConfig.UsdtDenom).String(), client.NewU128ScriptArg),
			),
		)

		// Check bank balance
		acc1CurCoins := bankKeeper.GetAllBalances(ctx, acc1.GetAddress())
		s.Require().True(acc1ExpCoins.IsEqual(acc1CurCoins))

		acc2CurCoins := bankKeeper.GetAllBalances(ctx, acc2.GetAddress())
		s.Require().True(acc2ExpCoins.IsEqual(acc2CurCoins))

		// Check DelPool supply
		curDecPoolCoins := getDecPoolSupply()
		s.Require().True(
			prevDecPoolCoins.
				Sub(transferCoins).
				IsEqual(curDecPoolCoins),
		)
		prevDecPoolCoins = curDecPoolCoins
	}
}

func (s *ModuleDVVTestSuite) TestDeployExecuteEdgeCases() {
	ctx := s.ctx

	// Case 1: deploy the same module twice, execute script
	// - Second deploy should fail
	// - Script should execute (first module is deployed)
	{
		acc, accPrivKey := s.app.AddAccount(ctx, sdk.NewCoin(dnConfig.MainDenom, sdk.NewInt(1000)))

		s.CheckContractExecuted(
			s.DeployModule(acc.GetAddress(), accPrivKey,
				"math/module.move", nil,
			),
		)

		s.CheckContractFailed(
			s.DeployModule(acc.GetAddress(), accPrivKey,
				"math/module.move", nil,
			),
		)

		s.CheckContractExecuted(
			s.ExecuteScript(acc.GetAddress(), accPrivKey,
				"math/script.template.move", []string{
					s.GetLibraAccAddressString(acc.GetAddress()),
				},
				s.BuildScriptArg("1", client.NewU64ScriptArg),
				s.BuildScriptArg("2", client.NewU64ScriptArg),
			),
		)
	}

	// Case 2: deploy module with address prefix (address 0x... { ... })
	// - Module with address prefix is OK (should deploy)
	// - Second deploy should fail
	// - Script should execute (first module is deployed)
	{
		acc, accPrivKey := s.app.AddAccount(ctx, sdk.NewCoin(dnConfig.MainDenom, sdk.NewInt(1000)))

		s.CheckContractExecuted(
			s.DeployModule(acc.GetAddress(), accPrivKey,
				"math/module_addr_wrapped.template.move", []string{
					s.GetLibraAccAddressString(acc.GetAddress()),
				},
			),
		)

		s.CheckContractFailed(
			s.DeployModule(acc.GetAddress(), accPrivKey,
				"math/module_addr_wrapped.template.move", []string{
					s.GetLibraAccAddressString(acc.GetAddress()),
				},
			),
		)

		s.CheckContractExecuted(
			s.ExecuteScript(acc.GetAddress(), accPrivKey,
				"math/script.template.move", []string{
					s.GetLibraAccAddressString(acc.GetAddress()),
				},
				s.BuildScriptArg("1", client.NewU64ScriptArg),
				s.BuildScriptArg("2", client.NewU64ScriptArg),
			),
		)
	}

	// Case 3: deploy multiple modules sourced from one file, execute scripts from one file
	// - Multi-module file should compile and deploy
	// - Multi-script file should compile (two CompiledUnits)
	// - Script execution only allowed for one script, but multiple Msgs might be included into one Tx
	{
		acc, accPrivKey := s.app.AddAccount(ctx, sdk.NewCoin(dnConfig.MainDenom, sdk.NewInt(1000)))

		// Deploy 4 modules
		s.CheckContractExecuted(
			s.DeployModule(acc.GetAddress(), accPrivKey,
				"math/module_quad.move", nil,
			),
		)

		// Compile 2 scripts
		scriptsCode := s.CompileMoveFile(acc.GetAddress(),
			"math/script_double.template.move",
			s.GetLibraAccAddressString(acc.GetAddress()),
			s.GetLibraAccAddressString(acc.GetAddress()),
			s.GetLibraAccAddressString(acc.GetAddress()),
			s.GetLibraAccAddressString(acc.GetAddress()),
		)
		s.Require().Len(scriptsCode, 2)

		// Execute scripts one by one (in one Tx)
		script1Msg := types.NewMsgExecuteScript(acc.GetAddress(), scriptsCode[0],
			s.BuildScriptArg("5", client.NewU64ScriptArg),
			s.BuildScriptArg("5", client.NewU64ScriptArg),
			s.BuildScriptArg("2", client.NewU64ScriptArg),
		)

		script2Msg := types.NewMsgExecuteScript(acc.GetAddress(), scriptsCode[1],
			s.BuildScriptArg("5", client.NewU64ScriptArg),
			s.BuildScriptArg("5", client.NewU64ScriptArg),
			s.BuildScriptArg("2", client.NewU64ScriptArg),
		)

		_, scriptsEvents := s.CheckContractExecuted(
			s.app.DeliverTx(
				s.ctx,
				acc.GetAddress(),
				accPrivKey,
				[]sdk.Msg{&script1Msg, &script2Msg},
			),
		)

		s.CheckABCIEventsContain(scriptsEvents, []sdk.Event{
			sdk.NewEvent(
				types.EventTypeMoveEvent,
				sdk.NewAttribute(types.AttributeVmEventSender, acc.GetAddress().String()),
				sdk.NewAttribute(types.AttributeVmEventSource, types.AttributeValueSourceScript),
				sdk.NewAttribute(types.AttributeVmEventType, "u64"),
				sdk.NewAttribute(types.AttributeVmEventData, "0800000000000000"),
			),
			sdk.NewEvent(
				types.EventTypeMoveEvent,
				sdk.NewAttribute(types.AttributeVmEventSender, acc.GetAddress().String()),
				sdk.NewAttribute(types.AttributeVmEventSource, types.AttributeValueSourceScript),
				sdk.NewAttribute(types.AttributeVmEventType, "u64"),
				sdk.NewAttribute(types.AttributeVmEventData, "0200000000000000"),
			),
		})
	}

	// Case 4: mixed module/script file
	// - Mixed file should compile
	// - Mixed module/script byteCode shouldn't be deployed
	{
		acc, accPrivKey := s.app.AddAccount(ctx, sdk.NewCoin(dnConfig.MainDenom, sdk.NewInt(1000)))

		s.CheckContractFailed(
			s.DeployModule(acc.GetAddress(), accPrivKey,
				"module_script.move", nil,
			),
		)
	}
}
