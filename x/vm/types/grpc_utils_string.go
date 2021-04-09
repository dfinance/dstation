package types

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"

	dvmTypes "github.com/dfinance/dstation/pkg/types/dvm"
)

const (
	//
	VmGasPrice = 1                   // gas unit price for VM execution
	VmGasLimit = ^uint64(0)/1000 - 1 // gas limit for VM execution

	// VM Event to sdk.Event conversion params
	EventTypeProcessingGas = 10000 // initial gas for processing event type.
	EventTypeNoGasLevels   = 2     // defines number of nesting levels that do not charge gas
)

// StringifyVMAccessPath converts dvmTypes.VMAccessPath to HEX string.
func StringifyVMAccessPath(path *dvmTypes.VMAccessPath) string {
	if path == nil {
		return "nil"
	}

	return fmt.Sprintf("Access path:\n"+
		"  Address: %s\n"+
		"  Path:    %s\n"+
		"  Key:     %s",
		hex.EncodeToString(path.Address),
		hex.EncodeToString(path.Path),
		hex.EncodeToString(GetVMStorageKey(path)),
	)
}

// StringifyVMTypeTag convert dvmTypes.VMTypeTag to string representation.
func StringifyVMTypeTag(tag dvmTypes.VMTypeTag) (string, error) {
	if val, ok := dvmTypes.VMTypeTag_name[int32(tag)]; !ok {
		return "", fmt.Errorf("can't find string representation of VMTypeTag %d, check correctness of type value", tag)
	} else {
		return val, nil
	}
}

// StringifyVMTypeTagPanic wraps StringifyVMTypeTag and panics on error.
func StringifyVMTypeTagPanic(tag dvmTypes.VMTypeTag) string {
	val, err := StringifyVMTypeTag(tag)
	if err != nil {
		panic(err)
	}

	return val
}

// StringifyVMLCSTag converts dvmTypes.LcsTag to string representation (recursive).
// <indentCount> defines number of prefixed indent string for each line.
func StringifyVMLCSTag(tag *dvmTypes.LcsTag, indentCount ...int) (string, error) {
	const strIndent = "  "

	curIndentCount := 0
	if len(indentCount) > 1 {
		return "", fmt.Errorf("invalid indentCount length")
	}
	if len(indentCount) == 1 {
		curIndentCount = indentCount[0]
	}
	if curIndentCount < 0 {
		return "", fmt.Errorf("invalid indentCount")
	}

	strBuilder := strings.Builder{}

	// Helper funcs
	buildStrIndent := func() string {
		str := ""
		for i := 0; i < curIndentCount; i++ {
			str += strIndent
		}
		return str
	}

	buildErr := func(comment string, err error) error {
		return fmt.Errorf("indent %d: %s: %w", curIndentCount, comment, err)
	}

	buildLcsTypeStr := func(t dvmTypes.LcsType) (string, error) {
		val, ok := dvmTypes.LcsType_name[int32(t)]
		if !ok {
			return "", fmt.Errorf("can't find string representation of LcsTag %d, check correctness of type value", t)
		}
		return val, nil
	}

	// Print current tag with recursive func call for fields
	if tag == nil {
		strBuilder.WriteString("nil")
		return strBuilder.String(), nil
	}

	indentStr := buildStrIndent()
	strBuilder.WriteString("LcsTag:\n")

	// Field: TypeTag
	typeTagStr, err := buildLcsTypeStr(tag.TypeTag)
	if err != nil {
		return "", buildErr("TypeTag", err)
	}
	strBuilder.WriteString(fmt.Sprintf("%sTypeTag: %s\n", indentStr, typeTagStr))

	// Field: VectorType
	vectorTypeStr, err := StringifyVMLCSTag(tag.VectorType, curIndentCount+1)
	if err != nil {
		return "", buildErr("VectorType", err)
	}
	strBuilder.WriteString(fmt.Sprintf("%sVectorType: %s\n", indentStr, vectorTypeStr))

	// Field: StructIdent
	if tag.StructIdent != nil {
		strBuilder.WriteString(fmt.Sprintf("%sStructIdent.Address: %s\n", indentStr, hex.EncodeToString(tag.StructIdent.Address)))
		strBuilder.WriteString(fmt.Sprintf("%sStructIdent.Module: %s\n", indentStr, tag.StructIdent.Module))
		strBuilder.WriteString(fmt.Sprintf("%sStructIdent.Name: %s\n", indentStr, tag.StructIdent.Name))
		if len(tag.StructIdent.TypeParams) > 0 {
			for structParamIdx, structParamTag := range tag.StructIdent.TypeParams {
				structParamTagStr, err := StringifyVMLCSTag(structParamTag, curIndentCount+1)
				if err != nil {
					return "", buildErr(fmt.Sprintf("StructIdent.TypeParams[%d]", structParamIdx), err)
				}
				strBuilder.WriteString(fmt.Sprintf("%sStructIdent.TypeParams[%d]: %s", indentStr, structParamIdx, structParamTagStr))
				if structParamIdx < len(tag.StructIdent.TypeParams)-1 {
					strBuilder.WriteString("\n")
				}
			}
		} else {
			strBuilder.WriteString(fmt.Sprintf("%sStructIdent.TypeParams: empty", indentStr))
		}
	} else {
		strBuilder.WriteString(fmt.Sprintf("%sStructIdent: nil", indentStr))
	}

	return strBuilder.String(), nil
}

// StringifyVMLCSTagPanic wraps StringifyVMLCSTag and panics on error.
func StringifyVMLCSTagPanic(tag *dvmTypes.LcsTag, indentCount ...int) string {
	val, err := StringifyVMLCSTag(tag, indentCount...)
	if err != nil {
		panic(err)
	}

	return val
}

// StringifyVMWriteOp converts dvmTypes.VmWriteOp to string representation.
func StringifyVMWriteOp(wOp dvmTypes.VmWriteOp) string {
	switch wOp {
	case dvmTypes.VmWriteOp_Value:
		return "write"
	case dvmTypes.VmWriteOp_Deletion:
		return "del"
	default:
		return "unknown"
	}
}

// StringifyVMValue converts dvmTypes.VMValue (writeSet) to string representation.
func StringifyVMValue(value *dvmTypes.VMValue) string {
	if value == nil {
		return "nil"
	}

	return fmt.Sprintf("\nWriteSet %q:\n"+
		"  Address: %s\n"+
		"  Path: %s\n"+
		"  Value: %s",
		StringifyVMWriteOp(value.Type),
		hex.EncodeToString(value.Path.Address),
		hex.EncodeToString(value.Path.Path),
		hex.EncodeToString(value.Value),
	)
}

// StringifyVMStatus converts dvmTypes.VMStatus (execution result) to string representation.
func StringifyVMStatus(status *dvmTypes.VMStatus) string {
	strBuilder := strings.Builder{}
	strBuilder.WriteString(fmt.Sprintf("Exec %q status:\n", status.String()))

	if status != nil {
		majorStatus, subStatus, _ := GetStatusCodesFromVMStatus(status)

		strBuilder.WriteString(fmt.Sprintf("  Major code: %d\n", majorStatus))
		strBuilder.WriteString(fmt.Sprintf("  Major status: %s\n", StringifyVMStatusMajorCode(strconv.FormatUint(majorStatus, 10))))
		strBuilder.WriteString(fmt.Sprintf("  Sub code: %d\n", subStatus))
		strBuilder.WriteString(fmt.Sprintf("  Message: %s", status.GetMessage()))
	} else {
		strBuilder.WriteString("  VMStatus: nil")
	}

	return strBuilder.String()
}

// StringifyVMEvent converts dvmTypes.VMEvent to string representation.
func StringifyVMEvent(event *dvmTypes.VMEvent) string {
	strBuilder := strings.Builder{}

	if event == nil {
		strBuilder.WriteString("nil")
		return strBuilder.String()
	}
	strBuilder.WriteString("\n")

	strBuilder.WriteString("Event:\n")
	strBuilder.WriteString(fmt.Sprintf("  SenderAddress: %s\n", hex.EncodeToString(event.SenderAddress)))
	if event.SenderModule != nil {
		strBuilder.WriteString(fmt.Sprintf("  SenderModule.Address: %s\n", hex.EncodeToString(event.SenderModule.Address)))
		strBuilder.WriteString(fmt.Sprintf("  SenderModule.Name: %s\n", event.SenderModule.Name))
	} else {
		strBuilder.WriteString("  SenderModule: nil\n")
	}
	strBuilder.WriteString(fmt.Sprintf("  EventType: %s\n", StringifyVMLCSTagPanic(event.EventType, 2)))
	strBuilder.WriteString(fmt.Sprintf("  EventData: %s", hex.EncodeToString(event.EventData)))

	return strBuilder.String()
}

// StringifyEventType returns dvmTypes.LcsTag Move serialization.
// Func is similar to StringifyVMLCSTag, but result is one lined Move representation.
func StringifyEventType(gasMeter sdk.GasMeter, tag *dvmTypes.LcsTag) (string, error) {
	// Start with initial gas for first event, and then go in progression based on depth.
	return processEventType(gasMeter, tag, EventTypeProcessingGas, 1)
}

// StringifyEventTypePanic wraps StringifyEventType and panic on error.
func StringifyEventTypePanic(gasMeter sdk.GasMeter, tag *dvmTypes.LcsTag) string {
	eventType, eventTypeErr := StringifyEventType(gasMeter, tag)
	if eventTypeErr != nil {
		debugMsg := ""
		if tagStr, tagErr := StringifyVMLCSTag(tag); tagErr != nil {
			debugMsg = fmt.Sprintf("StringifyVMLCSTag failed: %v", tagErr)
		} else {
			debugMsg = tagStr
		}

		panicErr := fmt.Sprintf("EventType serialization failed: %v\n%s", eventTypeErr, debugMsg)
		panic(panicErr)
	}

	return eventType
}

// StringifySenderAddress converts VM address to string (0x1 for stdlib and wallet1... otherwise).
func StringifySenderAddress(addr []byte) string {
	if bytes.Equal(addr, StdLibAddress) {
		return StdLibAddressShortStr
	} else {
		return sdk.AccAddress(addr).String()
	}
}

// PrintVMStackTrace prints VM stack trace if contract is not executed successfully.
func PrintVMStackTrace(txId []byte, log log.Logger, exec *dvmTypes.VMExecuteResponse) {
	strBuilder := strings.Builder{}

	strBuilder.WriteString(fmt.Sprintf("Stack trace %X:\n", txId))

	// Print common status
	if len(exec.Events) > 0 {
		for eventIdx, event := range exec.Events {
			strBuilder.WriteString(fmt.Sprintf("Events[%d]: %s\n", eventIdx, StringifyVMEvent(event)))
		}
	} else {
		strBuilder.WriteString("Events: empty\n")
	}

	// Print all writeSets
	if len(exec.WriteSet) > 0 {
		for wsIdx, ws := range exec.WriteSet {
			strBuilder.WriteString(fmt.Sprintf("WriteSet[%d]: %s", wsIdx, StringifyVMValue(ws)))
		}
	} else {
		strBuilder.WriteString("WriteSet: empty")
	}

	log.Debug(strBuilder.String())
}

// GetVMTypeTagByString converts {typeTag} gRPC enum string representation to dvmTypes.VMTypeTag.
func GetVMTypeTagByString(typeTag string) (dvmTypes.VMTypeTag, error) {
	if val, ok := dvmTypes.VMTypeTag_value[typeTag]; !ok {
		return -1, fmt.Errorf("can't find tag VMTypeTag %s, check correctness of type value", typeTag)
	} else {
		return dvmTypes.VMTypeTag(val), nil
	}
}

// GetStatusCodesFromVMStatus extracts majorStatus, subStatus and abortLocation from dvmTypes.VMStatus
// panic if error exist but error object == nil
func GetStatusCodesFromVMStatus(status *dvmTypes.VMStatus) (majorStatus, subStatus uint64, location *dvmTypes.AbortLocation) {
	switch sErr := status.GetError().(type) {
	case *dvmTypes.VMStatus_Abort:
		majorStatus = VMAbortedCode
		if sErr.Abort == nil {
			panic(fmt.Errorf("getting status codes: VMStatus_Abort.Abort is nil"))
		}
		subStatus = sErr.Abort.GetAbortCode()
		if l := sErr.Abort.GetAbortLocation(); l != nil {
			location = l
		}
	case *dvmTypes.VMStatus_ExecutionFailure:
		if sErr.ExecutionFailure == nil {
			panic(fmt.Errorf("getting status codes: VMStatus_ExecutionFailure.ExecutionFailure is nil"))
		}
		majorStatus = sErr.ExecutionFailure.GetStatusCode()
		if l := sErr.ExecutionFailure.GetAbortLocation(); l != nil {
			location = l
		}
	case *dvmTypes.VMStatus_MoveError:
		if sErr.MoveError == nil {
			panic(fmt.Errorf("getting status codes: VMStatus_MoveError.MoveError is nil"))
		}
		majorStatus = sErr.MoveError.GetStatusCode()
	case nil:
		majorStatus = VMExecutedCode
	}

	return
}

// processEventType recursively processes event type and returns result event type as a string.
// If {depth} < 0 we do not charge gas as some nesting levels might be "free".
func processEventType(gasMeter sdk.GasMeter, tag *dvmTypes.LcsTag, gas, depth uint64) (string, error) {
	// We can't consume gas later (after recognizing the type), because it open doors for security holes.
	// Let's say dev will create type with a lot of generics, so transaction will take much more time to process.
	// In result it could be a situation when validator doesn't have enough time to process transaction.
	// Charging gas amount is geometry increased from depth to depth.

	if depth > EventTypeNoGasLevels {
		gas += EventTypeProcessingGas * (depth - EventTypeNoGasLevels - 1)
		gasMeter.ConsumeGas(gas, "event type processing")
	}

	if tag == nil {
		return "", nil
	}

	// Helper function: lcsTypeToString returns dvmTypes.LcsType Move representation
	lcsTypeToString := func(lcsType dvmTypes.LcsType) string {
		switch lcsType {
		case dvmTypes.LcsType_LcsBool:
			return "bool"
		case dvmTypes.LcsType_LcsU8:
			return "u8"
		case dvmTypes.LcsType_LcsU64:
			return "u64"
		case dvmTypes.LcsType_LcsU128:
			return "u128"
		case dvmTypes.LcsType_LcsSigner:
			return "signer"
		case dvmTypes.LcsType_LcsVector:
			return "vector"
		case dvmTypes.LcsType_LcsStruct:
			return "struct"
		default:
			return dvmTypes.LcsType_name[int32(lcsType)]
		}
	}

	// Check data consistency
	if tag.TypeTag == dvmTypes.LcsType_LcsVector && tag.VectorType == nil {
		return "", fmt.Errorf("TypeTag of type %q, but VectorType is nil", lcsTypeToString(tag.TypeTag))
	}
	if tag.TypeTag == dvmTypes.LcsType_LcsStruct && tag.StructIdent == nil {
		return "", fmt.Errorf("TypeTag of type %q, but StructIdent is nil", lcsTypeToString(tag.TypeTag))
	}

	// Vector tag
	if tag.VectorType != nil {
		vectorType, err := processEventType(gasMeter, tag.VectorType, gas, depth+1)
		if err != nil {
			return "", fmt.Errorf("VectorType serialization: %w", err)
		}
		return fmt.Sprintf("%s<%s>", lcsTypeToString(dvmTypes.LcsType_LcsVector), vectorType), nil
	}

	// Struct tag
	if tag.StructIdent != nil {
		structType := fmt.Sprintf("%s::%s::%s", StringifySenderAddress(tag.StructIdent.Address), tag.StructIdent.Module, tag.StructIdent.Name)
		if len(tag.StructIdent.TypeParams) == 0 {
			return structType, nil
		}

		structParams := make([]string, 0, len(tag.StructIdent.TypeParams))
		for paramIdx, paramTag := range tag.StructIdent.TypeParams {
			structParam, err := processEventType(gasMeter, paramTag, gas, depth+1)
			if err != nil {
				return "", fmt.Errorf("StructIdent serialization: TypeParam[%d]: %w", paramIdx, err)
			}
			structParams = append(structParams, structParam)
		}
		return fmt.Sprintf("%s<%s>", structType, strings.Join(structParams, ", ")), nil
	}

	// Single tag
	return lcsTypeToString(tag.TypeTag), nil
}
