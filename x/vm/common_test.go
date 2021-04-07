package vm_test

import (
	"context"
	"encoding/hex"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	cryptoTypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/dfinance/dstation/pkg"
	"github.com/dfinance/dstation/pkg/tests"
	"github.com/dfinance/dstation/x/vm/config"
	"github.com/dfinance/dstation/x/vm/keeper"
	"github.com/dfinance/dstation/x/vm/types"
)

type ModuleDVVTestSuite struct {
	suite.Suite

	app         *tests.DSimApp
	ctx         sdk.Context
	keeper      keeper.Keeper
	queryClient types.QueryClient
	//
	dvmStop func() error
}

func (s *ModuleDVVTestSuite) SetupSuite() {
	// Init DVM connections and config
	_, dsPort, err := server.FreeTCPAddr()
	if err != nil {
		panic(fmt.Errorf("free TCP port request for DS server: %w", err))
	}

	_, vmPort, err := server.FreeTCPAddr()
	if err != nil {
		panic(fmt.Errorf("free TCP port request for VM connection: %w", err))
	}

	vmConfig := config.VMConfig{
		Address:        fmt.Sprintf("tcp://127.0.0.1:%s", vmPort),
		DataListen:     fmt.Sprintf("tcp://127.0.0.1:%s", dsPort),
		MaxAttempts:    0,
		ReqTimeoutInMs: 0,
	}

	// Start the DVM
	// Debug helpers:
	// - printLogs = true
	// - extra args: "--log=debug" / "-vvv"
	dvmStop, err := tests.LaunchDVMWithNetTransport(vmPort, dsPort, false)
	if err != nil {
		panic(fmt.Errorf("launch DVM: %w", err))
	}
	s.dvmStop = dvmStop

	// Init the SimApp
	s.app = tests.SetupDSimApp(tests.WithDVMConfig(vmConfig))
	s.ctx = s.app.GetContext(false)
	s.keeper = s.app.DnApp.VmKeeper

	// Init the querier client
	querier := keeper.Querier{Keeper: s.keeper}
	queryHelper := baseapp.NewQueryServerTestHelper(s.ctx, s.app.DnApp.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, querier)
	s.queryClient = types.NewQueryClient(queryHelper)
}

func (s *ModuleDVVTestSuite) TearDownSuite() {
	if err := s.dvmStop(); err != nil {
		s.T().Logf("TearDownSuite: DVM: %v", err)
	}

	s.app.TearDown()
}

func (s *ModuleDVVTestSuite) SetupTest() {
	s.ctx = s.ctx.WithEventManager(sdk.NewEventManager())
}

// GetProjectPath returns project root dir ("dfinance" base path).
func (s *ModuleDVVTestSuite) GetProjectPath() string {
	workingDir, err := os.Getwd()
	s.Require().NoError(err)

	projectDir := workingDir
	for {
		projectDir = filepath.Dir(projectDir)
		if projectDir == "." {
			s.Require().True(false, "dstation path not found within the current working dir: %s", workingDir)
		}

		if filepath.Base(projectDir) == "dstation" {
			break
		}
	}

	return projectDir
}

// GetMoveFileContent reads Move file content withing "move" sub-directory.
func (s *ModuleDVVTestSuite) GetMoveFileContent(fileName string, templateValues ...string) []byte {
	filePath := path.Join(s.GetProjectPath(), "x/vm/move", fileName)
	fileContent, err := pkg.ParseFilePath("fileName", filePath, pkg.ParamTypeCliArg)
	s.Require().NoError(err, "VM contract file: read failed")

	if len(templateValues) > 0 {
		fmtArgs := make([]interface{}, 0 , len(templateValues))
		for _, templateValue := range templateValues {
			fmtArgs = append(fmtArgs, templateValue)
		}

		fileContentStr := fmt.Sprintf(string(fileContent), fmtArgs...)
		fileContent = []byte(fileContentStr)
	}

	return fileContent
}

// CompileMoveFile reads Move file and compiles it.
func (s *ModuleDVVTestSuite) CompileMoveFile(accAddr sdk.AccAddress, fileName string, templateValues ...string) [][]byte {
	scriptSrc := s.GetMoveFileContent(fileName, templateValues...)

	resp, err := s.queryClient.Compile(context.Background(), &types.QueryCompileRequest{
		Address: types.Bech32ToLibra(accAddr),
		Code:    string(scriptSrc),
	})
	s.Require().NoError(err, "VM contract compilation: failed")
	s.Require().NotEmpty(resp.CompiledItems, "VM contract compilation: no compiled item found")

	byteCodes := make([][]byte, 0, len(resp.CompiledItems))
	for i, item := range resp.CompiledItems {
		s.Require().NotEmpty(item, "VM contract compilation: compiled item [%d] is empty", i)
		byteCodes = append(byteCodes, item.ByteCode)
	}

	return byteCodes
}

// DeployModule deploys VM module via Tx.
func (s *ModuleDVVTestSuite) DeployModule(accAddr sdk.AccAddress, accPrivKey cryptoTypes.PrivKey, fileName string, templateValues []string) (sdk.GasInfo, *sdk.Result, error) {
	byteCodes := s.CompileMoveFile(accAddr, fileName, templateValues...)

	msg := types.NewMsgDeployModule(accAddr, byteCodes...)

	return s.app.DeliverTx(s.ctx, accAddr, accPrivKey, []sdk.Msg{&msg})
}

// ExecuteScript executes VM script via Tx.
func (s *ModuleDVVTestSuite) ExecuteScript(accAddr sdk.AccAddress, accPrivKey cryptoTypes.PrivKey, fileName string, templateValues []string, args ...types.MsgExecuteScript_ScriptArg) (sdk.GasInfo, *sdk.Result, error) {
	byteCodes := s.CompileMoveFile(accAddr, fileName, templateValues...)
	s.Require().Len(byteCodes, 1, "VM script execution: compiledUnits len mismatch")

	msg := types.NewMsgExecuteScript(accAddr, byteCodes[0], args...)

	return s.app.DeliverTx(s.ctx, accAddr, accPrivKey, []sdk.Msg{&msg})
}

// BuildScriptArg wraps x/vm/client arg builders and can be used with ExecuteScript func.
func (s *ModuleDVVTestSuite) BuildScriptArg(value string, builder func(string) (types.MsgExecuteScript_ScriptArg, error)) types.MsgExecuteScript_ScriptArg {
	arg, err := builder(value)
	s.Require().NoError(err)

	return arg
}

// CheckContractExecuted checks that Tx result doesn't contain failed VM statues.
func (s *ModuleDVVTestSuite) CheckContractExecuted(gasInfo sdk.GasInfo, txRes *sdk.Result, txErr error) (sdk.GasInfo, []abci.Event) {
	discardEvents := s.getDiscardContractEvents(txRes, txErr)
	for _, event := range discardEvents {
		s.T().Logf("Failed ABCI event:\n%s", s.StringifyABCIEvent(event))
	}
	s.Require().Empty(discardEvents, "VM contract execution: failed")

	return gasInfo, txRes.Events
}

// CheckContractFailed checks that Tx result contains failed VM statues.
func (s *ModuleDVVTestSuite) CheckContractFailed(gasInfo sdk.GasInfo, txRes *sdk.Result, txErr error) (sdk.GasInfo, []abci.Event) {
	discardEvents := s.getDiscardContractEvents(txRes, txErr)
	s.Require().NotEmpty(discardEvents, "VM contract execution: succeeded")

	return gasInfo, txRes.Events
}

// CheckABCIEventsContain checks that eventsA contains all eventsB entries.
func (s *ModuleDVVTestSuite) CheckABCIEventsContain(eventsA []abci.Event, eventsB []sdk.Event) {
	s.Require().GreaterOrEqual(len(eventsA), len(eventsB), "comparing ABCI/SDK events: length mismatch: %d / %d", len(eventsA), len(eventsB))

	eventsBConv := make([]abci.Event, 0, len(eventsB))
	for _, event := range eventsB {
		attrsConv := make([]abci.EventAttribute, 0, len(event.Attributes))
		for _, attr := range event.Attributes {
			attrsConv = append(attrsConv, abci.EventAttribute{
				Key:   attr.Key,
				Value: attr.Value,
			})
		}

		eventsBConv = append(eventsBConv, abci.Event{
			Type:       event.Type,
			Attributes: attrsConv,
		})
	}

	for i, eventB := range eventsBConv {
		s.Require().Contains(eventsA, eventB, "comparing ABCI/SDK events: ABCIs have no SDK [%d]", i)
	}
}

// StringifySdkEvent builds abci.Event string representation.
func (s *ModuleDVVTestSuite) StringifyABCIEvent(event abci.Event) string {
	str := strings.Builder{}

	str.WriteString(fmt.Sprintf("- Event.Type: %s\n", event.Type))
	for _, attr := range event.Attributes {
		str.WriteString(fmt.Sprintf("  * %s = %s\n", attr.Key, attr.Value))
		if string(attr.Key) == types.AttributeErrMajorStatus {
			str.WriteString(fmt.Sprintf("    Description: %s\n", types.StringifyVMStatusMajorCode(string(attr.Value))))
		}
	}

	return str.String()
}

// PrintABCIEvents prints abci.Event list to test stdout.
func (s *ModuleDVVTestSuite) PrintABCIEvents(events []abci.Event) {
	for i, event := range events {
		s.T().Logf("Event [%d]:\n%s", i, s.StringifyABCIEvent(event))
	}
}

// GetLibraAccAddressString converts sdk.AccAddress to Libra address HEX string (used as a templateValue for VM scripts).
// acc.GetAddress().String() works as well with DVM (this func is optional).
func (s *ModuleDVVTestSuite) GetLibraAccAddressString(accAddr sdk.AccAddress) string {
	libraAddr := types.Bech32ToLibra(accAddr)
	moveAddr := hex.EncodeToString(libraAddr)

	return "0x" + moveAddr
}

// getDiscardContractEvents returns Discard abci.Event for Tx delivery result.
func (s *ModuleDVVTestSuite) getDiscardContractEvents(txRes *sdk.Result, txErr error) []abci.Event {
	s.Require().NoError(txErr, "VM contract Tx delivery: failed")
	s.Require().NotNil(txRes, "VM contract Tx delivery: result is nil")

	var discardEvents []abci.Event
	for _, event := range txRes.Events {
		if event.Type != types.EventTypeContractStatus {
			continue
		}

		for _, attr := range event.Attributes {
			if string(attr.Key) != types.AttributeStatus {
				continue
			}

			if string(attr.Value) == types.AttributeValueStatusDiscard {
				discardEvents = append(discardEvents, event)
				break
			}
		}
	}

	return discardEvents
}

func TestVMModule_DVM(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(ModuleDVVTestSuite))
}
