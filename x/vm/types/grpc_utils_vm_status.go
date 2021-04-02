package types

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/types"
)

const (
	// Error codes in JSON format
	jsonErrorCodes = `{
		"3001": "MALFORMED",
		"4016": "ABORTED",
		"7": "SENDING_ACCOUNT_DOES_NOT_EXIST",
		"18": "SENDING_ACCOUNT_FROZEN",
		"1027": "STLOC_TYPE_MISMATCH_ERROR",
		"1077": "LOOP_IN_INSTANTIATION_GRAPH",
		"2011": "UNREACHABLE",
		"1025": "BR_TYPE_MISMATCH_ERROR",
		"1095": "DUPLICATE_MODULE_NAME",
		"1038": "COPYLOC_RESOURCE_ERROR",
		"2014": "INVALID_CODE_CACHE",
		"3004": "UNKNOWN_TABLE_TYPE",
		"3017": "BAD_U16",
		"1074": "GLOBAL_REFERENCE_ERROR",
		"1029": "UNSAFE_RET_LOCAL_OR_RESOURCE_STILL_BORROWED",
		"1049": "READREF_TYPE_MISMATCH_ERROR",
		"1022": "POP_REFERENCE_ERROR",
		"1052": "WRITEREF_TYPE_MISMATCH_ERROR",
		"1040": "MOVELOC_UNAVAILABLE_ERROR",
		"1005": "RECURSIVE_STRUCT_DEFINITION",
		"1050": "READREF_RESOURCE_ERROR",
		"1030": "RET_TYPE_MISMATCH_ERROR",
		"4025": "VM_MAX_VALUE_DEPTH_REACHED",
		"21": "INVALID_MODULE_PUBLISHER",
		"1070": "MISSING_ACQUIRES_RESOURCE_ANNOTATION_ERROR",
		"3006": "UNKNOWN_SERIALIZED_TYPE",
		"16": "GAS_UNIT_PRICE_ABOVE_MAX_BOUND",
		"1053": "WRITEREF_RESOURCE_ERROR",
		"1046": "CALL_BORROWED_MUTABLE_REFERENCE_ERROR",
		"1011": "INVALID_MAIN_FUNCTION_SIGNATURE",
		"1028": "STLOC_UNSAFE_TO_DESTROY_ERROR",
		"1065": "MOVETO_NO_RESOURCE_ERROR",
		"1": "INVALID_SIGNATURE",
		"1013": "INVALID_MODULE_HANDLE",
		"3007": "UNKNOWN_OPCODE",
		"3010": "DUPLICATE_TABLE",
		"3014": "UNKNOWN_NATIVE_STRUCT_FLAG",
		"1084": "EMPTY_CODE_UNIT",
		"3018": "BAD_U32",
		"1009": "NEGATIVE_STACK_SIZE_WITHIN_BLOCK",
		"1004": "INVALID_FIELD_DEF",
		"4": "SEQUENCE_NUMBER_TOO_NEW",
		"12": "UNKNOWN_MODULE",
		"0": "UNKNOWN_VALIDATION_STATUS",
		"1037": "COPYLOC_UNAVAILABLE_ERROR",
		"1047": "PACK_TYPE_MISMATCH_ERROR",
		"1055": "WRITEREF_NO_MUTABLE_REFERENCE_ERROR",
		"1091": "FUNCTION_RESOLUTION_FAILURE",
		"1094": "INVALID_OPERATION_IN_SCRIPT",
		"4001": "EXECUTED",
		"1026": "ABORT_TYPE_MISMATCH_ERROR",
		"1057": "BOOLEAN_OP_TYPE_MISMATCH_ERROR",
		"1069": "POSITIVE_STACK_SIZE_AT_BLOCK_END",
		"4007": "TYPE_ERROR",
		"4012": "CANNOT_WRITE_EXISTING_RESOURCE",
		"14": "MAX_GAS_UNITS_BELOW_MIN_TRANSACTION_GAS_UNITS",
		"1068": "NO_MODULE_HANDLES",
		"4009": "DATA_FORMAT_ERROR",
		"2000": "UNKNOWN_INVARIANT_VIOLATION_ERROR",
		"1024": "RELEASEREF_TYPE_MISMATCH_ERROR",
		"3005": "UNKNOWN_SIGNATURE_TYPE",
		"5": "INSUFFICIENT_BALANCE_FOR_TRANSACTION_FEE",
		"4003": "RESOURCE_DOES_NOT_EXIST",
		"1008": "JOIN_FAILURE",
		"1066": "CREATEACCOUNT_TYPE_MISMATCH_ERROR",
		"2006": "VERIFICATION_ERROR",
		"4021": "CALL_STACK_OVERFLOW",
		"1042": "BORROWLOC_REFERENCE_ERROR",
		"1061": "BORROWGLOBAL_NO_RESOURCE_ERROR",
		"1071": "EXTRANEOUS_ACQUIRES_RESOURCE_ANNOTATION_ERROR",
		"1002": "RANGE_OUT_OF_BOUNDS",
		"4004": "RESOURCE_ALREADY_EXISTS",
		"4018": "DYNAMIC_REFERENCE_ERROR",
		"3012": "UNKNOWN_NOMINAL_RESOURCE",
		"1056": "INTEGER_OP_TYPE_MISMATCH_ERROR",
		"3023": "VALUE_DESERIALIZATION_ERROR",
		"1072": "DUPLICATE_ACQUIRES_RESOURCE_ANNOTATION_ERROR",
		"1018": "VISIBILITY_MISMATCH",
		"3021": "BAD_ULEB_U8",
		"19": "UNABLE_TO_DESERIALIZE_ACCOUNT",
		"1064": "MOVETO_TYPE_MISMATCH_ERROR",
		"23": "BAD_CHAIN_ID",
		"1062": "MOVEFROM_TYPE_MISMATCH_ERROR",
		"3003": "UNKNOWN_VERSION",
		"2005": "PC_OVERFLOW",
		"3019": "BAD_U64",
		"1015": "INCONSISTENT_FIELDS",
		"1060": "BORROWGLOBAL_TYPE_MISMATCH_ERROR",
		"1054": "WRITEREF_EXISTS_BORROW_ERROR",
		"2": "INVALID_AUTH_KEY",
		"1059": "EXISTS_RESOURCE_TYPE_MISMATCH_ERROR",
		"1078": "UNUSED_LOCALS_SIGNATURE",
		"4000": "UNKNOWN_RUNTIME_STATUS",
		"10": "EXCEEDED_MAX_TRANSACTION_SIZE",
		"1031": "RET_BORROWED_MUTABLE_REFERENCE_ERROR",
		"2009": "INTERNAL_TYPE_ERROR",
		"2001": "OUT_OF_BOUNDS_INDEX",
		"3015": "BAD_ULEB_U16",
		"1020": "TYPE_MISMATCH",
		"1067": "MODULE_ADDRESS_DOES_NOT_MATCH_SENDER",
		"1016": "UNUSED_FIELD",
		"3009": "UNEXPECTED_SIGNATURE_TYPE",
		"1044": "BORROWLOC_EXISTS_BORROW_ERROR",
		"2002": "OUT_OF_BOUNDS_RANGE",
		"1012": "DUPLICATE_ELEMENT",
		"1080": "ZERO_SIZED_STRUCT",
		"4008": "MISSING_DATA",
		"1021": "MISSING_DEPENDENCY",
		"15": "GAS_UNIT_PRICE_BELOW_MIN_BOUND",
		"3000": "UNKNOWN_BINARY_ERROR",
		"4017": "ARITHMETIC_ERROR",
		"1083": "MALFORMED_CONSTANT_DATA",
		"3022": "VALUE_SERIALIZATION_ERROR",
		"1001": "INDEX_OUT_OF_BOUNDS",
		"2008": "STORAGE_ERROR",
		"1048": "UNPACK_TYPE_MISMATCH_ERROR",
		"2004": "EMPTY_CALL_STACK",
		"13": "MAX_GAS_UNITS_EXCEEDS_MAX_GAS_UNITS_BOUND",
		"1003": "INVALID_SIGNATURE_TOKEN",
		"1079": "UNUSED_TYPE_SIGNATURE",
		"1088": "UNSAFE_RET_UNUSED_RESOURCES",
		"1081": "LINKER_ERROR",
		"11": "UNKNOWN_SCRIPT",
		"1058": "EQUALITY_OP_TYPE_MISMATCH_ERROR",
		"20": "CURRENCY_INFO_DOES_NOT_EXIST",
		"1086": "INVALID_LOOP_BREAK",
		"1023": "POP_RESOURCE_ERROR",
		"1082": "INVALID_CONSTANT_TYPE",
		"3": "SEQUENCE_NUMBER_TOO_OLD",
		"1014": "UNIMPLEMENTED_HANDLE",
		"1045": "CALL_TYPE_MISMATCH_ERROR",
		"2012": "VM_STARTUP_FAILURE",
		"4011": "REMOTE_DATA_ERROR",
		"22": "NO_ACCOUNT_ROLE",
		"1089": "TOO_MANY_LOCALS",
		"1085": "INVALID_LOOP_SPLIT",
		"1033": "FREEZEREF_EXISTS_MUTABLE_BORROW_ERROR",
		"1010": "UNBALANCED_STACK",
		"1035": "BORROWFIELD_BAD_FIELD_ERROR",
		"1017": "LOOKUP_FAILED",
		"1075": "CONSTRAINT_KIND_MISMATCH",
		"1000": "UNKNOWN_VERIFICATION_ERROR",
		"1043": "BORROWLOC_UNAVAILABLE_ERROR",
		"1073": "INVALID_ACQUIRES_RESOURCE_ANNOTATION_ERROR",
		"4024": "VM_MAX_TYPE_DEPTH_REACHED",
		"9": "INVALID_WRITE_SET",
		"4005": "EVICTED_ACCOUNT_ACCESS",
		"3020": "BAD_U128",
		"2003": "EMPTY_VALUE_STACK",
		"4006": "ACCOUNT_ADDRESS_ALREADY_EXISTS",
		"1041": "MOVELOC_EXISTS_BORROW_ERROR",
		"1051": "READREF_EXISTS_MUTABLE_BORROW_ERROR",
		"4010": "INVALID_DATA",
		"3016": "BAD_ULEB_U32",
		"1007": "INVALID_FALL_THROUGH",
		"1087": "INVALID_LOOP_CONTINUE",
		"1090": "GENERIC_MEMBER_OPCODE_MISMATCH",
		"4023": "GAS_SCHEDULE_ERROR",
		"2013": "NATIVE_FUNCTION_INTERNAL_INCONSISTENCY",
		"1036": "BORROWFIELD_EXISTS_MUTABLE_BORROW_ERROR",
		"8": "REJECTED_WRITE_SET",
		"3011": "VERIFIER_INVARIANT_VIOLATION",
		"18446744073709551615": "UNKNOWN_STATUS",
		"2010": "EVENT_KEY_MISMATCH",
		"3008": "BAD_HEADER_TABLE",
		"17": "INVALID_GAS_SPECIFIER",
		"3024": "CODE_DESERIALIZATION_ERROR",
		"1019": "TYPE_RESOLUTION_FAILURE",
		"1063": "MOVEFROM_NO_RESOURCE_ERROR",
		"1006": "INVALID_RESOURCE_FIELD",
		"2007": "LOCAL_REFERENCE_ERROR",
		"1076": "NUMBER_OF_TYPE_ARGUMENTS_MISMATCH",
		"3002": "BAD_MAGIC",
		"4020": "EXECUTION_STACK_OVERFLOW",
		"4002": "OUT_OF_GAS",
		"6": "TRANSACTION_EXPIRED",
		"1032": "FREEZEREF_TYPE_MISMATCH_ERROR",
		"1034": "BORROWFIELD_TYPE_MISMATCH_ERROR",
		"1039": "COPYLOC_EXISTS_BORROW_ERROR",
		"3013": "UNKNOWN_KIND"
	}`

	VMErrUnknown   = "unknown"
	VMExecutedCode = 4001
	VMAbortedCode  = 4016
)

var (
	// VM execution status majorCode to string error matching.
	errorCodes map[string]string
)

// StringifyVMStatusMajorCode returns vm_grpc.VMStatus majorCode string representation.
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
