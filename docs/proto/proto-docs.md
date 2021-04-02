<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [dfinance/dvm/common-types.proto](#dfinance/dvm/common-types.proto)
    - [u128](#dfinance.dvm.u128)
  
    - [VMTypeTag](#dfinance.dvm.VMTypeTag)
  
- [dfinance/dvm/compiler.proto](#dfinance/dvm/compiler.proto)
    - [CompilationResult](#dfinance.dvm.CompilationResult)
    - [CompilationUnit](#dfinance.dvm.CompilationUnit)
    - [CompiledUnit](#dfinance.dvm.CompiledUnit)
    - [SourceFiles](#dfinance.dvm.SourceFiles)
  
    - [DvmCompiler](#dfinance.dvm.DvmCompiler)
  
- [dfinance/dvm/data-source.proto](#dfinance/dvm/data-source.proto)
    - [CurrencyInfo](#dfinance.dvm.CurrencyInfo)
    - [CurrencyInfoRequest](#dfinance.dvm.CurrencyInfoRequest)
    - [CurrencyInfoResponse](#dfinance.dvm.CurrencyInfoResponse)
    - [DSAccessPath](#dfinance.dvm.DSAccessPath)
    - [DSAccessPaths](#dfinance.dvm.DSAccessPaths)
    - [DSRawResponse](#dfinance.dvm.DSRawResponse)
    - [DSRawResponses](#dfinance.dvm.DSRawResponses)
    - [NativeBalanceRequest](#dfinance.dvm.NativeBalanceRequest)
    - [NativeBalanceResponse](#dfinance.dvm.NativeBalanceResponse)
    - [OraclePriceRequest](#dfinance.dvm.OraclePriceRequest)
    - [OraclePriceResponse](#dfinance.dvm.OraclePriceResponse)
  
    - [ErrorCode](#dfinance.dvm.ErrorCode)
  
    - [DSService](#dfinance.dvm.DSService)
  
- [dfinance/dvm/metadata.proto](#dfinance/dvm/metadata.proto)
    - [Bytecode](#dfinance.dvm.Bytecode)
    - [Field](#dfinance.dvm.Field)
    - [Function](#dfinance.dvm.Function)
    - [Metadata](#dfinance.dvm.Metadata)
    - [ModuleMeta](#dfinance.dvm.ModuleMeta)
    - [ScriptMeta](#dfinance.dvm.ScriptMeta)
    - [Struct](#dfinance.dvm.Struct)
  
    - [DVMBytecodeMetadata](#dfinance.dvm.DVMBytecodeMetadata)
  
- [dfinance/dvm/vm.proto](#dfinance/dvm/vm.proto)
    - [Abort](#dfinance.dvm.Abort)
    - [AbortLocation](#dfinance.dvm.AbortLocation)
    - [Failure](#dfinance.dvm.Failure)
    - [FunctionLoc](#dfinance.dvm.FunctionLoc)
    - [LcsTag](#dfinance.dvm.LcsTag)
    - [Message](#dfinance.dvm.Message)
    - [ModuleIdent](#dfinance.dvm.ModuleIdent)
    - [MoveError](#dfinance.dvm.MoveError)
    - [MultipleCompilationResult](#dfinance.dvm.MultipleCompilationResult)
    - [StructIdent](#dfinance.dvm.StructIdent)
    - [VMAccessPath](#dfinance.dvm.VMAccessPath)
    - [VMArgs](#dfinance.dvm.VMArgs)
    - [VMBalanceChange](#dfinance.dvm.VMBalanceChange)
    - [VMBalanceChangeSet](#dfinance.dvm.VMBalanceChangeSet)
    - [VMEvent](#dfinance.dvm.VMEvent)
    - [VMExecuteResponse](#dfinance.dvm.VMExecuteResponse)
    - [VMExecuteScript](#dfinance.dvm.VMExecuteScript)
    - [VMPublishModule](#dfinance.dvm.VMPublishModule)
    - [VMStatus](#dfinance.dvm.VMStatus)
    - [VMValue](#dfinance.dvm.VMValue)
  
    - [LcsType](#dfinance.dvm.LcsType)
    - [VmWriteOp](#dfinance.dvm.VmWriteOp)
  
    - [VMModulePublisher](#dfinance.dvm.VMModulePublisher)
    - [VMScriptExecutor](#dfinance.dvm.VMScriptExecutor)
  
- [dfinance/vm/genesis.proto](#dfinance/vm/genesis.proto)
    - [GenesisState](#dfinance.vm.v1beta1.GenesisState)
    - [GenesisState.WriteOp](#dfinance.vm.v1beta1.GenesisState.WriteOp)
  
- [dfinance/vm/vm.proto](#dfinance/vm/vm.proto)
    - [CompiledItem](#dfinance.vm.v1beta1.CompiledItem)
    - [MsgDeployModule](#dfinance.vm.v1beta1.MsgDeployModule)
    - [MsgExecuteScript](#dfinance.vm.v1beta1.MsgExecuteScript)
    - [MsgExecuteScript.ScriptArg](#dfinance.vm.v1beta1.MsgExecuteScript.ScriptArg)
    - [TxVmStatus](#dfinance.vm.v1beta1.TxVmStatus)
    - [VmStatus](#dfinance.vm.v1beta1.VmStatus)
  
    - [CompiledItem.CodeType](#dfinance.vm.v1beta1.CompiledItem.CodeType)
  
- [dfinance/vm/query.proto](#dfinance/vm/query.proto)
    - [QueryCompileRequest](#dfinance.vm.v1beta1.QueryCompileRequest)
    - [QueryCompileResponse](#dfinance.vm.v1beta1.QueryCompileResponse)
    - [QueryDataRequest](#dfinance.vm.v1beta1.QueryDataRequest)
    - [QueryDataResponse](#dfinance.vm.v1beta1.QueryDataResponse)
    - [QueryMetadataRequest](#dfinance.vm.v1beta1.QueryMetadataRequest)
    - [QueryMetadataResponse](#dfinance.vm.v1beta1.QueryMetadataResponse)
    - [QueryTxVmStatusRequest](#dfinance.vm.v1beta1.QueryTxVmStatusRequest)
    - [QueryTxVmStatusResponse](#dfinance.vm.v1beta1.QueryTxVmStatusResponse)
  
    - [Query](#dfinance.vm.v1beta1.Query)
  
- [dfinance/vm/tx.proto](#dfinance/vm/tx.proto)
    - [MsgDeployModuleResponse](#dfinance.vm.v1beta1.MsgDeployModuleResponse)
    - [MsgExecuteScriptResponse](#dfinance.vm.v1beta1.MsgExecuteScriptResponse)
  
    - [Msg](#dfinance.vm.v1beta1.Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="dfinance/dvm/common-types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dfinance/dvm/common-types.proto



<a name="dfinance.dvm.u128"></a>

### u128
u128 type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `buf` | [bytes](#bytes) |  | Little-endian unsigned 128. |





 <!-- end messages -->


<a name="dfinance.dvm.VMTypeTag"></a>

### VMTypeTag
Type of contract argument.

| Name | Number | Description |
| ---- | ------ | ----------- |
| Bool | 0 | Bool 0x0 - false, 0x1 - true. |
| U64 | 1 | Uint64. Little-endian unsigned 64 bits integer. |
| Vector | 2 | Vector of bytes. |
| Address | 3 | Address, in bech32 form. 20 bytes. |
| U8 | 4 | U8 |
| U128 | 5 | U128 Little-endian unsigned 128 bits integer. |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="dfinance/dvm/compiler.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dfinance/dvm/compiler.proto



<a name="dfinance.dvm.CompilationResult"></a>

### CompilationResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `units` | [CompiledUnit](#dfinance.dvm.CompiledUnit) | repeated |  |
| `errors` | [string](#string) | repeated | list of error messages, empty if successful |






<a name="dfinance.dvm.CompilationUnit"></a>

### CompilationUnit
Compilation unit.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `text` | [string](#string) |  | utf8 encoded source code with libra/bech32 addresses |
| `name` | [string](#string) |  | name of the unit. |






<a name="dfinance.dvm.CompiledUnit"></a>

### CompiledUnit
Compiled source.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  | name of the module/script. |
| `bytecode` | [bytes](#bytes) |  | bytecode of the compiled module/script |






<a name="dfinance.dvm.SourceFiles"></a>

### SourceFiles
Compiler API


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `units` | [CompilationUnit](#dfinance.dvm.CompilationUnit) | repeated | Compilation units. |
| `address` | [bytes](#bytes) |  | address of the sender, in bech32 form |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="dfinance.dvm.DvmCompiler"></a>

### DvmCompiler


| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Compile` | [SourceFiles](#dfinance.dvm.SourceFiles) | [CompilationResult](#dfinance.dvm.CompilationResult) |  | |

 <!-- end services -->



<a name="dfinance/dvm/data-source.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dfinance/dvm/data-source.proto



<a name="dfinance.dvm.CurrencyInfo"></a>

### CurrencyInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [bytes](#bytes) |  |  |
| `decimals` | [uint32](#uint32) |  |  |
| `is_token` | [bool](#bool) |  |  |
| `address` | [bytes](#bytes) |  |  |
| `total_supply` | [u128](#dfinance.dvm.u128) |  |  |






<a name="dfinance.dvm.CurrencyInfoRequest"></a>

### CurrencyInfoRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `ticker` | [string](#string) |  |  |






<a name="dfinance.dvm.CurrencyInfoResponse"></a>

### CurrencyInfoResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `info` | [CurrencyInfo](#dfinance.dvm.CurrencyInfo) |  |  |
| `error_code` | [ErrorCode](#dfinance.dvm.ErrorCode) |  |  |
| `error_message` | [string](#string) |  | error message from libra, empty if ErrorCode::None |






<a name="dfinance.dvm.DSAccessPath"></a>

### DSAccessPath



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [bytes](#bytes) |  | AccountAddress |
| `path` | [bytes](#bytes) |  |  |






<a name="dfinance.dvm.DSAccessPaths"></a>

### DSAccessPaths



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `paths` | [DSAccessPath](#dfinance.dvm.DSAccessPath) | repeated |  |






<a name="dfinance.dvm.DSRawResponse"></a>

### DSRawResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `blob` | [bytes](#bytes) |  |  |
| `error_code` | [ErrorCode](#dfinance.dvm.ErrorCode) |  |  |
| `error_message` | [string](#string) |  | error message from libra, empty if ErrorCode::None |






<a name="dfinance.dvm.DSRawResponses"></a>

### DSRawResponses



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `blobs` | [bytes](#bytes) | repeated |  |






<a name="dfinance.dvm.NativeBalanceRequest"></a>

### NativeBalanceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [bytes](#bytes) |  |  |
| `ticker` | [string](#string) |  |  |






<a name="dfinance.dvm.NativeBalanceResponse"></a>

### NativeBalanceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `balance` | [u128](#dfinance.dvm.u128) |  |  |
| `error_code` | [ErrorCode](#dfinance.dvm.ErrorCode) |  |  |
| `error_message` | [string](#string) |  | error message from libra, empty if ErrorCode::None |






<a name="dfinance.dvm.OraclePriceRequest"></a>

### OraclePriceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `currency_1` | [string](#string) |  |  |
| `currency_2` | [string](#string) |  |  |






<a name="dfinance.dvm.OraclePriceResponse"></a>

### OraclePriceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `price` | [u128](#dfinance.dvm.u128) |  |  |
| `error_code` | [ErrorCode](#dfinance.dvm.ErrorCode) |  |  |
| `error_message` | [string](#string) |  | error message from libra, empty if ErrorCode::None |





 <!-- end messages -->


<a name="dfinance.dvm.ErrorCode"></a>

### ErrorCode


| Name | Number | Description |
| ---- | ------ | ----------- |
| NONE | 0 | no error |
| BAD_REQUEST | 1 | crash of compilation, logs will show stacktrace |
| NO_DATA | 2 | no such module |


 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="dfinance.dvm.DSService"></a>

### DSService
GRPC service

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `GetRaw` | [DSAccessPath](#dfinance.dvm.DSAccessPath) | [DSRawResponse](#dfinance.dvm.DSRawResponse) |  | |
| `MultiGetRaw` | [DSAccessPaths](#dfinance.dvm.DSAccessPaths) | [DSRawResponses](#dfinance.dvm.DSRawResponses) |  | |
| `GetOraclePrice` | [OraclePriceRequest](#dfinance.dvm.OraclePriceRequest) | [OraclePriceResponse](#dfinance.dvm.OraclePriceResponse) |  | |
| `GetNativeBalance` | [NativeBalanceRequest](#dfinance.dvm.NativeBalanceRequest) | [NativeBalanceResponse](#dfinance.dvm.NativeBalanceResponse) |  | |
| `GetCurrencyInfo` | [CurrencyInfoRequest](#dfinance.dvm.CurrencyInfoRequest) | [CurrencyInfoResponse](#dfinance.dvm.CurrencyInfoResponse) |  | |

 <!-- end services -->



<a name="dfinance/dvm/metadata.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dfinance/dvm/metadata.proto



<a name="dfinance.dvm.Bytecode"></a>

### Bytecode
Bytecode.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `code` | [bytes](#bytes) |  | bytecode of script |






<a name="dfinance.dvm.Field"></a>

### Field
Struct field.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  |  |
| `type` | [string](#string) |  |  |






<a name="dfinance.dvm.Function"></a>

### Function
Function representation.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  |  |
| `isPublic` | [bool](#bool) |  |  |
| `isNative` | [bool](#bool) |  |  |
| `type_parameters` | [string](#string) | repeated |  |
| `arguments` | [string](#string) | repeated |  |
| `returns` | [string](#string) | repeated |  |






<a name="dfinance.dvm.Metadata"></a>

### Metadata
Bytecode metadata.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `script` | [ScriptMeta](#dfinance.dvm.ScriptMeta) |  | In case the provided bytecode is a script. |
| `module` | [ModuleMeta](#dfinance.dvm.ModuleMeta) |  | In case the provided bytecode is a module. |






<a name="dfinance.dvm.ModuleMeta"></a>

### ModuleMeta
Module metadata.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  | module name. |
| `types` | [Struct](#dfinance.dvm.Struct) | repeated | Types defined in a module. |
| `functions` | [Function](#dfinance.dvm.Function) | repeated | Functions defined in a module. |






<a name="dfinance.dvm.ScriptMeta"></a>

### ScriptMeta
Script metadata.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `signers_count` | [uint32](#uint32) |  |  |
| `type_parameters` | [string](#string) | repeated |  |
| `arguments` | [VMTypeTag](#dfinance.dvm.VMTypeTag) | repeated |  |






<a name="dfinance.dvm.Struct"></a>

### Struct
Struct representation.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `name` | [string](#string) |  |  |
| `isResource` | [bool](#bool) |  |  |
| `type_parameters` | [string](#string) | repeated |  |
| `field` | [Field](#dfinance.dvm.Field) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="dfinance.dvm.DVMBytecodeMetadata"></a>

### DVMBytecodeMetadata
Returns bytecode metadata.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `GetMetadata` | [Bytecode](#dfinance.dvm.Bytecode) | [Metadata](#dfinance.dvm.Metadata) |  | |

 <!-- end services -->



<a name="dfinance/dvm/vm.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dfinance/dvm/vm.proto



<a name="dfinance.dvm.Abort"></a>

### Abort
VmStatus `MoveAbort` case.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `abort_location` | [AbortLocation](#dfinance.dvm.AbortLocation) |  | Abort location. (optional). Null if abort occurred in the script. |
| `abort_code` | [uint64](#uint64) |  | Abort code. |






<a name="dfinance.dvm.AbortLocation"></a>

### AbortLocation
An `AbortLocation` specifies where a Move program `abort` occurred, either in a function in
a module, or in a script.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [bytes](#bytes) |  | Indicates `abort` occurred in the specified module. |
| `module` | [string](#string) |  | Indicates the `abort` occurred in a script. |






<a name="dfinance.dvm.Failure"></a>

### Failure
VmStatus `ExecutionFailure` case.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `status_code` | [uint64](#uint64) |  | Status code. |
| `abort_location` | [AbortLocation](#dfinance.dvm.AbortLocation) |  | Abort location. (optional). Null if abort occurred in the script. |
| `function_loc` | [FunctionLoc](#dfinance.dvm.FunctionLoc) |  | Function location. |






<a name="dfinance.dvm.FunctionLoc"></a>

### FunctionLoc
Function location.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `function` | [uint64](#uint64) |  | Function index. |
| `code_offset` | [uint64](#uint64) |  | Code offset. |






<a name="dfinance.dvm.LcsTag"></a>

### LcsTag



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `type_tag` | [LcsType](#dfinance.dvm.LcsType) |  | type tag. |
| `vector_type` | [LcsTag](#dfinance.dvm.LcsTag) |  | vector type. Has a non-null value if the type_tag is equal to a LcsVector. |
| `struct_ident` | [StructIdent](#dfinance.dvm.StructIdent) |  | struct identifier. Has a non-null value if the type_tag is equal to a LcsStruct. |






<a name="dfinance.dvm.Message"></a>

### Message
Message.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `text` | [string](#string) |  | Message with error details if needed. |






<a name="dfinance.dvm.ModuleIdent"></a>

### ModuleIdent
Module identifier.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [bytes](#bytes) |  | module address. |
| `name` | [string](#string) |  | module name. |






<a name="dfinance.dvm.MoveError"></a>

### MoveError
VmStatus `Error` case.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `status_code` | [uint64](#uint64) |  | Status code. |






<a name="dfinance.dvm.MultipleCompilationResult"></a>

### MultipleCompilationResult



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `units` | [CompiledUnit](#dfinance.dvm.CompiledUnit) | repeated |  |
| `errors` | [string](#string) | repeated | list of error messages, empty if successful |






<a name="dfinance.dvm.StructIdent"></a>

### StructIdent
Full name of the structure.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [bytes](#bytes) |  | address of module owner |
| `module` | [string](#string) |  | module name. |
| `name` | [string](#string) |  | name of structure. |
| `type_params` | [LcsTag](#dfinance.dvm.LcsTag) | repeated | Structure type parameters. |






<a name="dfinance.dvm.VMAccessPath"></a>

### VMAccessPath
Storage path


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [bytes](#bytes) |  | account address. |
| `path` | [bytes](#bytes) |  | storage path. |






<a name="dfinance.dvm.VMArgs"></a>

### VMArgs
Contract arguments.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `type` | [VMTypeTag](#dfinance.dvm.VMTypeTag) |  | Argument type. |
| `value` | [bytes](#bytes) |  | Argument value. |






<a name="dfinance.dvm.VMBalanceChange"></a>

### VMBalanceChange



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [bytes](#bytes) |  |  |
| `ticker` | [string](#string) |  |  |
| `deposit` | [u128](#dfinance.dvm.u128) |  |  |
| `withdraw` | [u128](#dfinance.dvm.u128) |  |  |






<a name="dfinance.dvm.VMBalanceChangeSet"></a>

### VMBalanceChangeSet



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `change_set` | [VMBalanceChange](#dfinance.dvm.VMBalanceChange) | repeated |  |






<a name="dfinance.dvm.VMEvent"></a>

### VMEvent
VM event returns after contract execution.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender_address` | [bytes](#bytes) |  | Event sender address. |
| `sender_module` | [ModuleIdent](#dfinance.dvm.ModuleIdent) |  | sender module. |
| `event_type` | [LcsTag](#dfinance.dvm.LcsTag) |  | Type of value inside event. |
| `event_data` | [bytes](#bytes) |  | Event data in bytes to parse. |






<a name="dfinance.dvm.VMExecuteResponse"></a>

### VMExecuteResponse
Response from VM contains write_set, events, gas used and status for specific contract.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `write_set` | [VMValue](#dfinance.dvm.VMValue) | repeated | using string instead of bytes for now, as map support only ints and strings as keys |
| `events` | [VMEvent](#dfinance.dvm.VMEvent) | repeated | list of events executed during contract execution |
| `balance_change_set` | [VMBalanceChange](#dfinance.dvm.VMBalanceChange) | repeated | list of native balance updates. |
| `gas_used` | [uint64](#uint64) |  | Gas used during execution. |
| `status` | [VMStatus](#dfinance.dvm.VMStatus) |  | Main status of execution, might contain an error. |






<a name="dfinance.dvm.VMExecuteScript"></a>

### VMExecuteScript
VM contract object to process.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `senders` | [bytes](#bytes) | repeated | owners of contract. |
| `max_gas_amount` | [uint64](#uint64) |  | maximal total gas specified by wallet to spend for this transaction. |
| `gas_unit_price` | [uint64](#uint64) |  | maximal price can be paid per gas. |
| `block` | [uint64](#uint64) |  | block. |
| `timestamp` | [uint64](#uint64) |  | timestamp. |
| `code` | [bytes](#bytes) |  | compiled contract code. |
| `type_params` | [StructIdent](#dfinance.dvm.StructIdent) | repeated | type parameters. |
| `args` | [VMArgs](#dfinance.dvm.VMArgs) | repeated | Contract arguments. |






<a name="dfinance.dvm.VMPublishModule"></a>

### VMPublishModule
Publish module.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [bytes](#bytes) |  | owner of contract. |
| `max_gas_amount` | [uint64](#uint64) |  | maximal total gas specified by wallet to spend for this transaction. |
| `gas_unit_price` | [uint64](#uint64) |  | maximal price can be paid per gas. |
| `code` | [bytes](#bytes) |  | compiled contract code. |






<a name="dfinance.dvm.VMStatus"></a>

### VMStatus
A `VMStatus` is represented as either
- `Null` indicating successful execution.
- `Error` indicating an error from the VM itself.
- `MoveAbort` indicating an `abort` ocurred inside of a Move program
- `ExecutionFailure` indicating an runtime error.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `move_error` | [MoveError](#dfinance.dvm.MoveError) |  | Indicates an error from the VM, e.g. OUT_OF_GAS, INVALID_AUTH_KEY, RET_TYPE_MISMATCH_ERROR etc. The code will neither EXECUTED nor ABORTED |
| `abort` | [Abort](#dfinance.dvm.Abort) |  | Indicates an error from the VM, e.g. OUT_OF_GAS, INVALID_AUTH_KEY, RET_TYPE_MISMATCH_ERROR etc. The code will neither EXECUTED nor ABORTED |
| `execution_failure` | [Failure](#dfinance.dvm.Failure) |  | Indicates an failure from inside Move code, where the VM could not continue exection, e.g. dividing by zero or a missing resource |
| `message` | [Message](#dfinance.dvm.Message) |  | Message with error details if needed (optional). |






<a name="dfinance.dvm.VMValue"></a>

### VMValue
VM value should be passed before execution and return after execution (with opcodes), write_set in nutshell.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `type` | [VmWriteOp](#dfinance.dvm.VmWriteOp) |  | Type of operation |
| `value` | [bytes](#bytes) |  | Value returns from vm. |
| `path` | [VMAccessPath](#dfinance.dvm.VMAccessPath) |  | Access path. |





 <!-- end messages -->


<a name="dfinance.dvm.LcsType"></a>

### LcsType


| Name | Number | Description |
| ---- | ------ | ----------- |
| LcsBool | 0 | Bool |
| LcsU64 | 1 | Uint64 |
| LcsVector | 2 | Vector of bytes. |
| LcsAddress | 3 | Address, in bech32 form |
| LcsU8 | 4 | U8 |
| LcsU128 | 5 | U128 |
| LcsSigner | 6 | Signer. |
| LcsStruct | 7 | Struct. |



<a name="dfinance.dvm.VmWriteOp"></a>

### VmWriteOp
Write set operation type.

| Name | Number | Description |
| ---- | ------ | ----------- |
| Value | 0 | Insert or update value |
| Deletion | 1 | Delete. |


 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="dfinance.dvm.VMModulePublisher"></a>

### VMModulePublisher
GRPC service

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `PublishModule` | [VMPublishModule](#dfinance.dvm.VMPublishModule) | [VMExecuteResponse](#dfinance.dvm.VMExecuteResponse) |  | |


<a name="dfinance.dvm.VMScriptExecutor"></a>

### VMScriptExecutor


| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `ExecuteScript` | [VMExecuteScript](#dfinance.dvm.VMExecuteScript) | [VMExecuteResponse](#dfinance.dvm.VMExecuteResponse) |  | |

 <!-- end services -->



<a name="dfinance/vm/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dfinance/vm/genesis.proto



<a name="dfinance.vm.v1beta1.GenesisState"></a>

### GenesisState



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `write_set` | [GenesisState.WriteOp](#dfinance.vm.v1beta1.GenesisState.WriteOp) | repeated |  |






<a name="dfinance.vm.v1beta1.GenesisState.WriteOp"></a>

### GenesisState.WriteOp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | Move address (HEX string) |
| `path` | [string](#string) |  | Move module path (HEX string) |
| `value` | [string](#string) |  | Module code (HEX string) |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="dfinance/vm/vm.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dfinance/vm/vm.proto



<a name="dfinance.vm.v1beta1.CompiledItem"></a>

### CompiledItem
CompiledItem contains VM compilation result.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `byte_code` | [bytes](#bytes) |  |  |
| `name` | [string](#string) |  |  |
| `methods` | [dfinance.dvm.Function](#dfinance.dvm.Function) | repeated |  |
| `types` | [dfinance.dvm.Struct](#dfinance.dvm.Struct) | repeated |  |
| `code_type` | [CompiledItem.CodeType](#dfinance.vm.v1beta1.CompiledItem.CodeType) |  |  |






<a name="dfinance.vm.v1beta1.MsgDeployModule"></a>

### MsgDeployModule
MsgDeployModule defines a SDK message to deploy a module (contract) to VM.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `signer` | [string](#string) |  | Script sender address |
| `modules` | [bytes](#bytes) | repeated | Module code |






<a name="dfinance.vm.v1beta1.MsgExecuteScript"></a>

### MsgExecuteScript
MsgExecuteScript defines a SDK message to execute a script with args to VM.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `signer` | [string](#string) |  | Script sender address |
| `script` | [bytes](#bytes) |  | Script code |
| `args` | [MsgExecuteScript.ScriptArg](#dfinance.vm.v1beta1.MsgExecuteScript.ScriptArg) | repeated |  |






<a name="dfinance.vm.v1beta1.MsgExecuteScript.ScriptArg"></a>

### MsgExecuteScript.ScriptArg



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `type` | [dfinance.dvm.VMTypeTag](#dfinance.dvm.VMTypeTag) |  |  |
| `value` | [bytes](#bytes) |  |  |






<a name="dfinance.vm.v1beta1.TxVmStatus"></a>

### TxVmStatus
TxVmStatus keeps VM statuses and errors for Tx.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `hash` | [string](#string) |  | Tx hash [HEX string] |
| `vm_statuses` | [VmStatus](#dfinance.vm.v1beta1.VmStatus) | repeated | VM statuses for the Tx |






<a name="dfinance.vm.v1beta1.VmStatus"></a>

### VmStatus
VmStatus is a VM error response.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `status` | [string](#string) |  | Error Status: error / discard |
| `major_code` | [string](#string) |  | Major code |
| `sub_code` | [string](#string) |  | Sub code |
| `str_code` | [string](#string) |  | Detailed explanation of major code |
| `message` | [string](#string) |  | Error message |





 <!-- end messages -->


<a name="dfinance.vm.v1beta1.CompiledItem.CodeType"></a>

### CompiledItem.CodeType


| Name | Number | Description |
| ---- | ------ | ----------- |
| MODULE | 0 |  |
| SCRIPT | 1 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="dfinance/vm/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dfinance/vm/query.proto



<a name="dfinance.vm.v1beta1.QueryCompileRequest"></a>

### QueryCompileRequest
QueryCompileRequest is request type for Query/Compile RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [bytes](#bytes) |  | VM address (Libra address) |
| `code` | [string](#string) |  | Move code [Plain text] |






<a name="dfinance.vm.v1beta1.QueryCompileResponse"></a>

### QueryCompileResponse
QueryCompileResponse is response type for Query/Compile RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `compiled_items` | [CompiledItem](#dfinance.vm.v1beta1.CompiledItem) | repeated | Compiled items |






<a name="dfinance.vm.v1beta1.QueryDataRequest"></a>

### QueryDataRequest
QueryDataRequest is request type for Query/Data RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [bytes](#bytes) |  | VM address (Libra address) |
| `path` | [bytes](#bytes) |  | VM path |






<a name="dfinance.vm.v1beta1.QueryDataResponse"></a>

### QueryDataResponse
QueryDataResponse is response type for Query/Data RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `value` | [bytes](#bytes) |  | VMStorage value for address:path pair |






<a name="dfinance.vm.v1beta1.QueryMetadataRequest"></a>

### QueryMetadataRequest
QueryMetadataRequest is request type for Query/Metadata RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `code` | [bytes](#bytes) |  |  |






<a name="dfinance.vm.v1beta1.QueryMetadataResponse"></a>

### QueryMetadataResponse
QueryMetadataResponse is response type for Query/Metadata RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `metadata` | [dfinance.dvm.Metadata](#dfinance.dvm.Metadata) |  |  |






<a name="dfinance.vm.v1beta1.QueryTxVmStatusRequest"></a>

### QueryTxVmStatusRequest
QueryTxVmStatusRequest is request type for Query/TxVmStatus RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_meta` | [cosmos.base.abci.v1beta1.TxResponse](#cosmos.base.abci.v1beta1.TxResponse) |  | Tx meta received from /cosmos/tx/v1beta1/txs/{hash} |






<a name="dfinance.vm.v1beta1.QueryTxVmStatusResponse"></a>

### QueryTxVmStatusResponse
QueryTxVmStatusResponse is response type for Query/TxVmStatus RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `vm_status` | [TxVmStatus](#dfinance.vm.v1beta1.TxVmStatus) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="dfinance.vm.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Data` | [QueryDataRequest](#dfinance.vm.v1beta1.QueryDataRequest) | [QueryDataResponse](#dfinance.vm.v1beta1.QueryDataResponse) | Data queries VMStorage value | GET|/dfinance/vm/v1beta1/data|
| `TxVmStatus` | [QueryTxVmStatusRequest](#dfinance.vm.v1beta1.QueryTxVmStatusRequest) | [QueryTxVmStatusResponse](#dfinance.vm.v1beta1.QueryTxVmStatusResponse) | TxVmStatus queries VM status for Tx | GET|/dfinance/vm/v1beta1/tx_vm_status|
| `Compile` | [QueryCompileRequest](#dfinance.vm.v1beta1.QueryCompileRequest) | [QueryCompileResponse](#dfinance.vm.v1beta1.QueryCompileResponse) | Compile compiles provided Move code and returns byte code. | GET|/dfinance/vm/v1beta1/compile|
| `Metadata` | [QueryMetadataRequest](#dfinance.vm.v1beta1.QueryMetadataRequest) | [QueryMetadataResponse](#dfinance.vm.v1beta1.QueryMetadataResponse) | Metadata queries VM for byteCode metadata (metadata.proto/GetMetadata RPC wrapper). | |

 <!-- end services -->



<a name="dfinance/vm/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dfinance/vm/tx.proto



<a name="dfinance.vm.v1beta1.MsgDeployModuleResponse"></a>

### MsgDeployModuleResponse







<a name="dfinance.vm.v1beta1.MsgExecuteScriptResponse"></a>

### MsgExecuteScriptResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="dfinance.vm.v1beta1.Msg"></a>

### Msg
Msg defines the VM module Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `DeployModule` | [MsgDeployModule](#dfinance.vm.v1beta1.MsgDeployModule) | [MsgDeployModuleResponse](#dfinance.vm.v1beta1.MsgDeployModuleResponse) | DeployModule deploys Move module/modules to VMStorage. | |
| `ExecuteScript` | [MsgExecuteScript](#dfinance.vm.v1beta1.MsgExecuteScript) | [MsgExecuteScriptResponse](#dfinance.vm.v1beta1.MsgExecuteScriptResponse) | ExecuteScript executes provided Move script. | |

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

