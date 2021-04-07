package types

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/types"
)

const (
	// Error codes in JSON format
	jsonErrorCodes = `
{
  "14": "MAX_GAS_UNITS_BELOW_MIN_TRANSACTION_GAS_UNITS",
  "3010": "DUPLICATE_TABLE",
  "4001": "EXECUTED",
  "4016": "ABORTED",
  "1": "INVALID_SIGNATURE",
  "3": "SEQUENCE_NUMBER_TOO_OLD",
  "1072": "DUPLICATE_ACQUIRES_RESOURCE_ANNOTATION_ERROR",
  "2017": "UNEXPECTED_VERIFIER_ERROR",
  "1046": "CALL_BORROWED_MUTABLE_REFERENCE_ERROR",
  "1006": "INVALID_RESOURCE_FIELD",
  "2005": "PC_OVERFLOW",
  "6": "TRANSACTION_EXPIRED",
  "1058": "EQUALITY_OP_TYPE_MISMATCH_ERROR",
  "3022": "VALUE_SERIALIZATION_ERROR",
  "1023": "POP_RESOURCE_ERROR",
  "1056": "INTEGER_OP_TYPE_MISMATCH_ERROR",
  "15": "GAS_UNIT_PRICE_BELOW_MIN_BOUND",
  "1017": "LOOKUP_FAILED",
  "4008": "MISSING_DATA",
  "2008": "STORAGE_ERROR",
  "1063": "MOVEFROM_NO_RESOURCE_ERROR",
  "2006": "VERIFICATION_ERROR",
  "1089": "TOO_MANY_LOCALS",
  "2009": "INTERNAL_TYPE_ERROR",
  "1007": "INVALID_FALL_THROUGH",
  "10": "EXCEEDED_MAX_TRANSACTION_SIZE",
  "4020": "EXECUTION_STACK_OVERFLOW",
  "5": "INSUFFICIENT_BALANCE_FOR_TRANSACTION_FEE",
  "1043": "BORROWLOC_UNAVAILABLE_ERROR",
  "3020": "BAD_U128",
  "7": "SENDING_ACCOUNT_DOES_NOT_EXIST",
  "1086": "INVALID_LOOP_BREAK",
  "1065": "MOVETO_NO_RESOURCE_ERROR",
  "1081": "LINKER_ERROR",
  "1000": "UNKNOWN_VERIFICATION_ERROR",
  "3002": "BAD_MAGIC",
  "3007": "UNKNOWN_OPCODE",
  "1045": "CALL_TYPE_MISMATCH_ERROR",
  "3008": "BAD_HEADER_TABLE",
  "1011": "INVALID_MAIN_FUNCTION_SIGNATURE",
  "1048": "UNPACK_TYPE_MISMATCH_ERROR",
  "1053": "WRITEREF_RESOURCE_ERROR",
  "1030": "RET_TYPE_MISMATCH_ERROR",
  "1040": "MOVELOC_UNAVAILABLE_ERROR",
  "1005": "RECURSIVE_STRUCT_DEFINITION",
  "2021": "TYPE_RESOLUTION_FAILURE",
  "4017": "ARITHMETIC_ERROR",
  "17": "INVALID_GAS_SPECIFIER",
  "18446744073709551615": "UNKNOWN_STATUS",
  "1041": "MOVELOC_EXISTS_BORROW_ERROR",
  "1020": "TYPE_MISMATCH",
  "12": "UNKNOWN_MODULE",
  "16": "GAS_UNIT_PRICE_ABOVE_MAX_BOUND",
  "1014": "UNIMPLEMENTED_HANDLE",
  "11": "UNKNOWN_SCRIPT",
  "1003": "INVALID_SIGNATURE_TOKEN",
  "1035": "BORROWFIELD_BAD_FIELD_ERROR",
  "1084": "EMPTY_CODE_UNIT",
  "2018": "UNEXPECTED_DESERIALIZATION_ERROR",
  "3000": "UNKNOWN_BINARY_ERROR",
  "2019": "FAILED_TO_SERIALIZE_WRITE_SET_CHANGES",
  "1026": "ABORT_TYPE_MISMATCH_ERROR",
  "22": "NO_ACCOUNT_ROLE",
  "1037": "COPYLOC_UNAVAILABLE_ERROR",
  "2020": "FAILED_TO_DESERIALIZE_RESOURCE",
  "3003": "UNKNOWN_VERSION",
  "3014": "UNKNOWN_NATIVE_STRUCT_FLAG",
  "4004": "RESOURCE_ALREADY_EXISTS",
  "1044": "BORROWLOC_EXISTS_BORROW_ERROR",
  "2012": "VM_STARTUP_FAILURE",
  "1055": "WRITEREF_NO_MUTABLE_REFERENCE_ERROR",
  "13": "MAX_GAS_UNITS_EXCEEDS_MAX_GAS_UNITS_BOUND",
  "4002": "OUT_OF_GAS",
  "4025": "VM_MAX_VALUE_DEPTH_REACHED",
  "2010": "EVENT_KEY_MISMATCH",
  "1047": "PACK_TYPE_MISMATCH_ERROR",
  "1067": "MODULE_ADDRESS_DOES_NOT_MATCH_SENDER",
  "1049": "READREF_TYPE_MISMATCH_ERROR",
  "2015": "UNEXPECTED_ERROR_FROM_KNOWN_MOVE_FUNCTION",
  "1074": "GLOBAL_REFERENCE_ERROR",
  "1057": "BOOLEAN_OP_TYPE_MISMATCH_ERROR",
  "19": "UNABLE_TO_DESERIALIZE_ACCOUNT",
  "1069": "POSITIVE_STACK_SIZE_AT_BLOCK_END",
  "1085": "INVALID_LOOP_SPLIT",
  "1062": "MOVEFROM_TYPE_MISMATCH_ERROR",
  "1075": "CONSTRAINT_KIND_MISMATCH",
  "3019": "BAD_U64",
  "1025": "BR_TYPE_MISMATCH_ERROR",
  "1001": "INDEX_OUT_OF_BOUNDS",
  "1091": "FUNCTION_RESOLUTION_FAILURE",
  "3004": "UNKNOWN_TABLE_TYPE",
  "1031": "RET_BORROWED_MUTABLE_REFERENCE_ERROR",
  "1051": "READREF_EXISTS_MUTABLE_BORROW_ERROR",
  "1064": "MOVETO_TYPE_MISMATCH_ERROR",
  "1082": "INVALID_CONSTANT_TYPE",
  "4003": "RESOURCE_DOES_NOT_EXIST",
  "4024": "VM_MAX_TYPE_DEPTH_REACHED",
  "1073": "INVALID_ACQUIRES_RESOURCE_ANNOTATION_ERROR",
  "1032": "FREEZEREF_TYPE_MISMATCH_ERROR",
  "1012": "DUPLICATE_ELEMENT",
  "1068": "NO_MODULE_HANDLES",
  "4009": "DATA_FORMAT_ERROR",
  "1009": "NEGATIVE_STACK_SIZE_WITHIN_BLOCK",
  "9": "INVALID_WRITE_SET",
  "18": "SENDING_ACCOUNT_FROZEN",
  "20": "CURRENCY_INFO_DOES_NOT_EXIST",
  "1027": "STLOC_TYPE_MISMATCH_ERROR",
  "3006": "UNKNOWN_SERIALIZED_TYPE",
  "1021": "MISSING_DEPENDENCY",
  "1038": "COPYLOC_RESOURCE_ERROR",
  "23": "BAD_CHAIN_ID",
  "1076": "NUMBER_OF_TYPE_ARGUMENTS_MISMATCH",
  "1087": "INVALID_LOOP_CONTINUE",
  "1095": "DUPLICATE_MODULE_NAME",
  "1083": "MALFORMED_CONSTANT_DATA",
  "3024": "CODE_DESERIALIZATION_ERROR",
  "2003": "EMPTY_VALUE_STACK",
  "2": "INVALID_AUTH_KEY",
  "1039": "COPYLOC_EXISTS_BORROW_ERROR",
  "3001": "MALFORMED",
  "3023": "VALUE_DESERIALIZATION_ERROR",
  "2011": "UNREACHABLE",
  "0": "UNKNOWN_VALIDATION_STATUS",
  "4": "SEQUENCE_NUMBER_TOO_NEW",
  "1052": "WRITEREF_TYPE_MISMATCH_ERROR",
  "1060": "BORROWGLOBAL_TYPE_MISMATCH_ERROR",
  "2016": "VERIFIER_INVARIANT_VIOLATION",
  "4021": "CALL_STACK_OVERFLOW",
  "1090": "GENERIC_MEMBER_OPCODE_MISMATCH",
  "2000": "UNKNOWN_INVARIANT_VIOLATION_ERROR",
  "1042": "BORROWLOC_REFERENCE_ERROR",
  "1028": "STLOC_UNSAFE_TO_DESTROY_ERROR",
  "1071": "EXTRANEOUS_ACQUIRES_RESOURCE_ANNOTATION_ERROR",
  "1094": "INVALID_OPERATION_IN_SCRIPT",
  "3009": "UNEXPECTED_SIGNATURE_TYPE",
  "1070": "MISSING_ACQUIRES_RESOURCE_ANNOTATION_ERROR",
  "1013": "INVALID_MODULE_HANDLE",
  "1061": "BORROWGLOBAL_NO_RESOURCE_ERROR",
  "1033": "FREEZEREF_EXISTS_MUTABLE_BORROW_ERROR",
  "1036": "BORROWFIELD_EXISTS_MUTABLE_BORROW_ERROR",
  "1054": "WRITEREF_EXISTS_BORROW_ERROR",
  "1059": "EXISTS_RESOURCE_TYPE_MISMATCH_ERROR",
  "21": "INVALID_MODULE_PUBLISHER",
  "1077": "LOOP_IN_INSTANTIATION_GRAPH",
  "8": "REJECTED_WRITE_SET",
  "1034": "BORROWFIELD_TYPE_MISMATCH_ERROR",
  "1088": "UNSAFE_RET_UNUSED_RESOURCES",
  "1080": "ZERO_SIZED_STRUCT",
  "3005": "UNKNOWN_SIGNATURE_TYPE",
  "3012": "UNKNOWN_NOMINAL_RESOURCE",
  "3013": "UNKNOWN_KIND",
  "1050": "READREF_RESOURCE_ERROR",
  "4000": "UNKNOWN_RUNTIME_STATUS",
  "1029": "UNSAFE_RET_LOCAL_OR_RESOURCE_STILL_BORROWED"
}
`

	VMErrUnknown   = "unknown"
	VMExecutedCode = 4001
	VMAbortedCode  = 4016
)

var (
	// VM execution status majorCode to string error matching.
	errorCodes map[string]string
)

// StringifyVMStatusMajorCode returns dvm.VMStatus majorCode string representation.
func StringifyVMStatusMajorCode(majorCode string) string {
	if v, ok := errorCodes[majorCode]; ok {
		return v
	}

	return VMErrUnknown
}

func (m VmStatus) String() string {
	return fmt.Sprintf("VM status:\n"+
		"  Status: %s\n"+
		"  Major code: %s\n"+
		"  String code: %s\n"+
		"  Sub code: %s\n"+
		"  Message:  %s",
		m.Status, m.MajorCode, m.StrCode, m.SubCode, m.Message,
	)
}

// StringifyVmStatuses build string representation of VmStatus list.
func StringifyVmStatuses(list []VmStatus) string {
	strBuilder := strings.Builder{}
	strBuilder.WriteString("VMStatuses:\n")
	for i, status := range list {
		strBuilder.WriteString(status.String())
		if i < len(list)-1 {
			strBuilder.WriteString("\n")
		}
	}

	return strBuilder.String()
}

// NewVmStatus creates a new VMStatus error.
func NewVmStatus(status, majorCode, subCode, message string) VmStatus {
	strCode := ""
	if status != AttributeValueStatusKeep {
		strCode = StringifyVMStatusMajorCode(majorCode)
	}

	return VmStatus{
		Status:    status,
		MajorCode: majorCode,
		SubCode:   subCode,
		Message:   message,
		StrCode:   strCode,
	}
}

func (m TxVmStatus) String() string {
	return fmt.Sprintf("Tx:\n"+
		"  Hash: %s\n"+
		"  Statuses: %s",
		m.Hash, StringifyVmStatuses(m.VmStatuses),
	)
}

// NewTxVmStatus creates a new TxVMStatus object.
func NewTxVmStatus(hash string, statuses []VmStatus) TxVmStatus {
	return TxVmStatus{
		Hash:       hash,
		VmStatuses: statuses,
	}
}

// NewVmStatusFromABCILogs converts SDK TxResponse log events to TxVMStatus.
func NewVmStatusFromABCILogs(tx types.TxResponse) TxVmStatus {
	statuses := make([]VmStatus, 0)

	for _, log := range tx.Logs {
		for _, event := range log.Events {
			isFound := false

			if event.Type == EventTypeContractStatus {
				status := ""
				majorCode := ""
				subCode := ""
				message := ""

				for _, attr := range event.Attributes {
					// find that it's event contains contract status.
					if attr.Key == AttributeStatus {
						status = attr.Value

						if status == AttributeValueStatusDiscard || status == AttributeValueStatusError {
							isFound = true
							break
						}
					}
				}

				// event found.
				if isFound {
					for _, attr := range event.Attributes {
						switch attr.Key {
						case AttributeErrMajorStatus:
							majorCode = attr.Value

						case AttributeErrSubStatus:
							subCode = attr.Value

						case AttributeErrMessage:
							message = attr.Value
						}
					}
				}

				statuses = append(statuses, NewVmStatus(status, majorCode, subCode, message))
			}
		}
	}

	return NewTxVmStatus(tx.TxHash, statuses)
}

func init() {
	errorCodes = make(map[string]string)
	if err := json.Unmarshal([]byte(jsonErrorCodes), &errorCodes); err != nil {
		panic(err)
	}
}
