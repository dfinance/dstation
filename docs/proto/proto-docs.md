<!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [dfinance/vm/genesis.proto](#dfinance/vm/genesis.proto)
    - [GenesisState](#dfinance.vm.v1beta1.GenesisState)
    - [GenesisState.WriteOp](#dfinance.vm.v1beta1.GenesisState.WriteOp)
  
- [dfinance/vm/vm.proto](#dfinance/vm/vm.proto)
    - [MsgDeployModule](#dfinance.vm.v1beta1.MsgDeployModule)
    - [MsgExecuteScript](#dfinance.vm.v1beta1.MsgExecuteScript)
    - [MsgExecuteScript.ScriptArg](#dfinance.vm.v1beta1.MsgExecuteScript.ScriptArg)
    - [TxVmStatus](#dfinance.vm.v1beta1.TxVmStatus)
    - [VmStatus](#dfinance.vm.v1beta1.VmStatus)
  
- [dfinance/vm/query.proto](#dfinance/vm/query.proto)
    - [QueryDataRequest](#dfinance.vm.v1beta1.QueryDataRequest)
    - [QueryDataResponse](#dfinance.vm.v1beta1.QueryDataResponse)
    - [QueryTxVmStatusRequest](#dfinance.vm.v1beta1.QueryTxVmStatusRequest)
    - [QueryTxVmStatusResponse](#dfinance.vm.v1beta1.QueryTxVmStatusResponse)
  
    - [Query](#dfinance.vm.v1beta1.Query)
  
- [dfinance/vm/tx.proto](#dfinance/vm/tx.proto)
    - [MsgDeployModuleResponse](#dfinance.vm.v1beta1.MsgDeployModuleResponse)
    - [MsgExecuteScriptResponse](#dfinance.vm.v1beta1.MsgExecuteScriptResponse)
  
    - [Msg](#dfinance.vm.v1beta1.Msg)
  
- [Scalar Value Types](#scalar-value-types)



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
| `type` | [types.VMTypeTag](#types.VMTypeTag) |  |  |
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

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="dfinance/vm/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dfinance/vm/query.proto



<a name="dfinance.vm.v1beta1.QueryDataRequest"></a>

### QueryDataRequest
QueryDataRequest is request type for Query/Data RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  | VM address [HEX string] |
| `path` | [string](#string) |  | VM path [HEX string] |






<a name="dfinance.vm.v1beta1.QueryDataResponse"></a>

### QueryDataResponse
QueryDataResponse is response type for Query/Data RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `value` | [string](#string) |  | VMStorage value for address:path pair [HEX string] |






<a name="dfinance.vm.v1beta1.QueryTxVmStatusRequest"></a>

### QueryTxVmStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `tx_meta` | [cosmos.base.abci.v1beta1.TxResponse](#cosmos.base.abci.v1beta1.TxResponse) |  | Tx meta received from /cosmos/tx/v1beta1/txs/{hash} |






<a name="dfinance.vm.v1beta1.QueryTxVmStatusResponse"></a>

### QueryTxVmStatusResponse



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

