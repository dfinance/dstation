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
  
- [dfinance/oracle/oracle.proto](#dfinance/oracle/oracle.proto)
    - [Asset](#dfinance.oracle.v1beta1.Asset)
    - [CurrentPrice](#dfinance.oracle.v1beta1.CurrentPrice)
    - [Oracle](#dfinance.oracle.v1beta1.Oracle)
    - [RawPrice](#dfinance.oracle.v1beta1.RawPrice)
  
- [dfinance/oracle/genesis.proto](#dfinance/oracle/genesis.proto)
    - [GenesisState](#dfinance.oracle.v1beta1.GenesisState)
    - [Params](#dfinance.oracle.v1beta1.Params)
    - [Params.PostPriceParams](#dfinance.oracle.v1beta1.Params.PostPriceParams)
  
- [dfinance/oracle/query.proto](#dfinance/oracle/query.proto)
    - [QueryAssetsRequest](#dfinance.oracle.v1beta1.QueryAssetsRequest)
    - [QueryAssetsResponse](#dfinance.oracle.v1beta1.QueryAssetsResponse)
    - [QueryCurrentPriceRequest](#dfinance.oracle.v1beta1.QueryCurrentPriceRequest)
    - [QueryCurrentPriceResponse](#dfinance.oracle.v1beta1.QueryCurrentPriceResponse)
    - [QueryCurrentPricesRequest](#dfinance.oracle.v1beta1.QueryCurrentPricesRequest)
    - [QueryCurrentPricesResponse](#dfinance.oracle.v1beta1.QueryCurrentPricesResponse)
    - [QueryOraclesRequest](#dfinance.oracle.v1beta1.QueryOraclesRequest)
    - [QueryOraclesResponse](#dfinance.oracle.v1beta1.QueryOraclesResponse)
  
    - [Query](#dfinance.oracle.v1beta1.Query)
  
- [dfinance/oracle/tx.proto](#dfinance/oracle/tx.proto)
    - [MsgPostPrice](#dfinance.oracle.v1beta1.MsgPostPrice)
    - [MsgPostPriceResponse](#dfinance.oracle.v1beta1.MsgPostPriceResponse)
    - [MsgSetAsset](#dfinance.oracle.v1beta1.MsgSetAsset)
    - [MsgSetAssetResponse](#dfinance.oracle.v1beta1.MsgSetAssetResponse)
    - [MsgSetOracle](#dfinance.oracle.v1beta1.MsgSetOracle)
    - [MsgSetOracleResponse](#dfinance.oracle.v1beta1.MsgSetOracleResponse)
  
    - [Msg](#dfinance.oracle.v1beta1.Msg)
  
- [dfinance/vm/genesis.proto](#dfinance/vm/genesis.proto)
    - [GenesisState](#dfinance.vm.v1beta1.GenesisState)
    - [GenesisState.WriteOp](#dfinance.vm.v1beta1.GenesisState.WriteOp)
  
- [dfinance/vm/gov.proto](#dfinance/vm/gov.proto)
    - [PlannedProposal](#dfinance.vm.v1beta1.PlannedProposal)
    - [StdLibUpdateProposal](#dfinance.vm.v1beta1.StdLibUpdateProposal)
  
- [dfinance/vm/vm.proto](#dfinance/vm/vm.proto)
    - [CompiledItem](#dfinance.vm.v1beta1.CompiledItem)
    - [TxVmStatus](#dfinance.vm.v1beta1.TxVmStatus)
    - [VmStatus](#dfinance.vm.v1beta1.VmStatus)
  
    - [CompiledItem.CodeType](#dfinance.vm.v1beta1.CompiledItem.CodeType)
  
- [dfinance/vm/query.proto](#dfinance/vm/query.proto)
    - [QueryCompileRequest](#dfinance.vm.v1beta1.QueryCompileRequest)
    - [QueryCompileResponse](#dfinance.vm.v1beta1.QueryCompileResponse)
    - [QueryDataRequest](#dfinance.vm.v1beta1.QueryDataRequest)
    - [QueryDataResponse](#dfinance.vm.v1beta1.QueryDataResponse)
    - [QueryDelegatedPoolSupplyRequest](#dfinance.vm.v1beta1.QueryDelegatedPoolSupplyRequest)
    - [QueryDelegatedPoolSupplyResponse](#dfinance.vm.v1beta1.QueryDelegatedPoolSupplyResponse)
    - [QueryMetadataRequest](#dfinance.vm.v1beta1.QueryMetadataRequest)
    - [QueryMetadataResponse](#dfinance.vm.v1beta1.QueryMetadataResponse)
    - [QueryTxVmStatusRequest](#dfinance.vm.v1beta1.QueryTxVmStatusRequest)
    - [QueryTxVmStatusResponse](#dfinance.vm.v1beta1.QueryTxVmStatusResponse)
  
    - [Query](#dfinance.vm.v1beta1.Query)
  
- [dfinance/vm/tx.proto](#dfinance/vm/tx.proto)
    - [MsgDeployModule](#dfinance.vm.v1beta1.MsgDeployModule)
    - [MsgDeployModuleResponse](#dfinance.vm.v1beta1.MsgDeployModuleResponse)
    - [MsgExecuteScript](#dfinance.vm.v1beta1.MsgExecuteScript)
    - [MsgExecuteScript.ScriptArg](#dfinance.vm.v1beta1.MsgExecuteScript.ScriptArg)
    - [MsgExecuteScriptResponse](#dfinance.vm.v1beta1.MsgExecuteScriptResponse)
  
    - [Msg](#dfinance.vm.v1beta1.Msg)
  
- [gravity/v1/attestation.proto](#gravity/v1/attestation.proto)
    - [Attestation](#gravity.v1.Attestation)
    - [ERC20Token](#gravity.v1.ERC20Token)
  
    - [ClaimType](#gravity.v1.ClaimType)
  
- [gravity/v1/batch.proto](#gravity/v1/batch.proto)
    - [OutgoingLogicCall](#gravity.v1.OutgoingLogicCall)
    - [OutgoingTransferTx](#gravity.v1.OutgoingTransferTx)
    - [OutgoingTxBatch](#gravity.v1.OutgoingTxBatch)
  
- [gravity/v1/ethereum_signer.proto](#gravity/v1/ethereum_signer.proto)
    - [SignType](#gravity.v1.SignType)
  
- [gravity/v1/types.proto](#gravity/v1/types.proto)
    - [BridgeValidator](#gravity.v1.BridgeValidator)
    - [ERC20ToDenom](#gravity.v1.ERC20ToDenom)
    - [LastObservedEthereumBlockHeight](#gravity.v1.LastObservedEthereumBlockHeight)
    - [Valset](#gravity.v1.Valset)
  
- [gravity/v1/msgs.proto](#gravity/v1/msgs.proto)
    - [MsgCancelSendToEth](#gravity.v1.MsgCancelSendToEth)
    - [MsgCancelSendToEthResponse](#gravity.v1.MsgCancelSendToEthResponse)
    - [MsgConfirmBatch](#gravity.v1.MsgConfirmBatch)
    - [MsgConfirmBatchResponse](#gravity.v1.MsgConfirmBatchResponse)
    - [MsgConfirmLogicCall](#gravity.v1.MsgConfirmLogicCall)
    - [MsgConfirmLogicCallResponse](#gravity.v1.MsgConfirmLogicCallResponse)
    - [MsgDepositClaim](#gravity.v1.MsgDepositClaim)
    - [MsgDepositClaimResponse](#gravity.v1.MsgDepositClaimResponse)
    - [MsgERC20DeployedClaim](#gravity.v1.MsgERC20DeployedClaim)
    - [MsgERC20DeployedClaimResponse](#gravity.v1.MsgERC20DeployedClaimResponse)
    - [MsgLogicCallExecutedClaim](#gravity.v1.MsgLogicCallExecutedClaim)
    - [MsgLogicCallExecutedClaimResponse](#gravity.v1.MsgLogicCallExecutedClaimResponse)
    - [MsgRequestBatch](#gravity.v1.MsgRequestBatch)
    - [MsgRequestBatchResponse](#gravity.v1.MsgRequestBatchResponse)
    - [MsgSendToEth](#gravity.v1.MsgSendToEth)
    - [MsgSendToEthResponse](#gravity.v1.MsgSendToEthResponse)
    - [MsgSetOrchestratorAddress](#gravity.v1.MsgSetOrchestratorAddress)
    - [MsgSetOrchestratorAddressResponse](#gravity.v1.MsgSetOrchestratorAddressResponse)
    - [MsgValsetConfirm](#gravity.v1.MsgValsetConfirm)
    - [MsgValsetConfirmResponse](#gravity.v1.MsgValsetConfirmResponse)
    - [MsgWithdrawClaim](#gravity.v1.MsgWithdrawClaim)
    - [MsgWithdrawClaimResponse](#gravity.v1.MsgWithdrawClaimResponse)
  
    - [Msg](#gravity.v1.Msg)
  
- [gravity/v1/genesis.proto](#gravity/v1/genesis.proto)
    - [GenesisState](#gravity.v1.GenesisState)
    - [Params](#gravity.v1.Params)
  
- [gravity/v1/pool.proto](#gravity/v1/pool.proto)
    - [BatchFees](#gravity.v1.BatchFees)
    - [IDSet](#gravity.v1.IDSet)
  
- [gravity/v1/query.proto](#gravity/v1/query.proto)
    - [QueryBatchConfirmsRequest](#gravity.v1.QueryBatchConfirmsRequest)
    - [QueryBatchConfirmsResponse](#gravity.v1.QueryBatchConfirmsResponse)
    - [QueryBatchFeeRequest](#gravity.v1.QueryBatchFeeRequest)
    - [QueryBatchFeeResponse](#gravity.v1.QueryBatchFeeResponse)
    - [QueryBatchRequestByNonceRequest](#gravity.v1.QueryBatchRequestByNonceRequest)
    - [QueryBatchRequestByNonceResponse](#gravity.v1.QueryBatchRequestByNonceResponse)
    - [QueryCurrentValsetRequest](#gravity.v1.QueryCurrentValsetRequest)
    - [QueryCurrentValsetResponse](#gravity.v1.QueryCurrentValsetResponse)
    - [QueryDelegateKeysByEthAddress](#gravity.v1.QueryDelegateKeysByEthAddress)
    - [QueryDelegateKeysByEthAddressResponse](#gravity.v1.QueryDelegateKeysByEthAddressResponse)
    - [QueryDelegateKeysByOrchestratorAddress](#gravity.v1.QueryDelegateKeysByOrchestratorAddress)
    - [QueryDelegateKeysByOrchestratorAddressResponse](#gravity.v1.QueryDelegateKeysByOrchestratorAddressResponse)
    - [QueryDelegateKeysByValidatorAddress](#gravity.v1.QueryDelegateKeysByValidatorAddress)
    - [QueryDelegateKeysByValidatorAddressResponse](#gravity.v1.QueryDelegateKeysByValidatorAddressResponse)
    - [QueryDenomToERC20Request](#gravity.v1.QueryDenomToERC20Request)
    - [QueryDenomToERC20Response](#gravity.v1.QueryDenomToERC20Response)
    - [QueryERC20ToDenomRequest](#gravity.v1.QueryERC20ToDenomRequest)
    - [QueryERC20ToDenomResponse](#gravity.v1.QueryERC20ToDenomResponse)
    - [QueryLastEventNonceByAddrRequest](#gravity.v1.QueryLastEventNonceByAddrRequest)
    - [QueryLastEventNonceByAddrResponse](#gravity.v1.QueryLastEventNonceByAddrResponse)
    - [QueryLastPendingBatchRequestByAddrRequest](#gravity.v1.QueryLastPendingBatchRequestByAddrRequest)
    - [QueryLastPendingBatchRequestByAddrResponse](#gravity.v1.QueryLastPendingBatchRequestByAddrResponse)
    - [QueryLastPendingLogicCallByAddrRequest](#gravity.v1.QueryLastPendingLogicCallByAddrRequest)
    - [QueryLastPendingLogicCallByAddrResponse](#gravity.v1.QueryLastPendingLogicCallByAddrResponse)
    - [QueryLastPendingValsetRequestByAddrRequest](#gravity.v1.QueryLastPendingValsetRequestByAddrRequest)
    - [QueryLastPendingValsetRequestByAddrResponse](#gravity.v1.QueryLastPendingValsetRequestByAddrResponse)
    - [QueryLastValsetRequestsRequest](#gravity.v1.QueryLastValsetRequestsRequest)
    - [QueryLastValsetRequestsResponse](#gravity.v1.QueryLastValsetRequestsResponse)
    - [QueryLogicConfirmsRequest](#gravity.v1.QueryLogicConfirmsRequest)
    - [QueryLogicConfirmsResponse](#gravity.v1.QueryLogicConfirmsResponse)
    - [QueryOutgoingLogicCallsRequest](#gravity.v1.QueryOutgoingLogicCallsRequest)
    - [QueryOutgoingLogicCallsResponse](#gravity.v1.QueryOutgoingLogicCallsResponse)
    - [QueryOutgoingTxBatchesRequest](#gravity.v1.QueryOutgoingTxBatchesRequest)
    - [QueryOutgoingTxBatchesResponse](#gravity.v1.QueryOutgoingTxBatchesResponse)
    - [QueryParamsRequest](#gravity.v1.QueryParamsRequest)
    - [QueryParamsResponse](#gravity.v1.QueryParamsResponse)
    - [QueryPendingSendToEth](#gravity.v1.QueryPendingSendToEth)
    - [QueryPendingSendToEthResponse](#gravity.v1.QueryPendingSendToEthResponse)
    - [QueryValsetConfirmRequest](#gravity.v1.QueryValsetConfirmRequest)
    - [QueryValsetConfirmResponse](#gravity.v1.QueryValsetConfirmResponse)
    - [QueryValsetConfirmsByNonceRequest](#gravity.v1.QueryValsetConfirmsByNonceRequest)
    - [QueryValsetConfirmsByNonceResponse](#gravity.v1.QueryValsetConfirmsByNonceResponse)
    - [QueryValsetRequestRequest](#gravity.v1.QueryValsetRequestRequest)
    - [QueryValsetRequestResponse](#gravity.v1.QueryValsetRequestResponse)
  
    - [Query](#gravity.v1.Query)
  
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



<a name="dfinance/oracle/oracle.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dfinance/oracle/oracle.proto



<a name="dfinance.oracle.v1beta1.Asset"></a>

### Asset
Asset represents an Oracle asset.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `asset_code` | [string](#string) |  | Asset code (for ex.: btc_usdt) |
| `oracles` | [string](#string) | repeated | List of registered RawPrice sources (Oracle addresses) If none - asset is essentially disabled |
| `decimals` | [uint32](#uint32) |  | Number of decimals for Asset's CurrentPrice values |






<a name="dfinance.oracle.v1beta1.CurrentPrice"></a>

### CurrentPrice
CurrentPrice contains meta of the current price for a particular asset (aggregated from multiple sources).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `asset_code` | [string](#string) |  | Asset code (for ex.: btc_usdt) |
| `ask_price` | [string](#string) |  | The latest lowest seller price |
| `bid_price` | [string](#string) |  | The latest highest buyer price |
| `received_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | The latest price update timestamp |
| `is_reversed` | [bool](#bool) |  | CurrentPrice is reversed flag: price is not received from Oracle sources, exchange rates were reversed programmatically |






<a name="dfinance.oracle.v1beta1.Oracle"></a>

### Oracle
Oracle contains Oracle source info.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `acc_address` | [string](#string) |  | Oracle account address |
| `description` | [string](#string) |  | Optional Oracle description |
| `price_max_bytes` | [uint32](#uint32) |  | Maximum number of bytes for PostPrice values |
| `price_decimals` | [uint32](#uint32) |  | Number of decimals for PostPrice values |






<a name="dfinance.oracle.v1beta1.RawPrice"></a>

### RawPrice
RawPrice is used to store normalized asset prices per Oracle.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `ask_price` | [string](#string) |  |  |
| `bid_price` | [string](#string) |  |  |
| `received_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="dfinance/oracle/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dfinance/oracle/genesis.proto



<a name="dfinance.oracle.v1beta1.GenesisState"></a>

### GenesisState



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#dfinance.oracle.v1beta1.Params) |  |  |
| `oracles` | [Oracle](#dfinance.oracle.v1beta1.Oracle) | repeated |  |
| `assets` | [Asset](#dfinance.oracle.v1beta1.Asset) | repeated |  |
| `current_prices` | [CurrentPrice](#dfinance.oracle.v1beta1.CurrentPrice) | repeated |  |






<a name="dfinance.oracle.v1beta1.Params"></a>

### Params
Params keeps keeper parameters (which might be changed via Gov).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nominees` | [string](#string) | repeated | Admin account addresses |
| `post_price` | [Params.PostPriceParams](#dfinance.oracle.v1beta1.Params.PostPriceParams) |  |  |






<a name="dfinance.oracle.v1beta1.Params.PostPriceParams"></a>

### Params.PostPriceParams
PostPriceParams keeps price posting parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `received_at_diff_in_s` | [uint32](#uint32) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="dfinance/oracle/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dfinance/oracle/query.proto



<a name="dfinance.oracle.v1beta1.QueryAssetsRequest"></a>

### QueryAssetsRequest
QueryAssetsRequest is request type for Query/Assets RPC method.






<a name="dfinance.oracle.v1beta1.QueryAssetsResponse"></a>

### QueryAssetsResponse
QueryAssetsResponse is response type for Query/Assets RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `assets` | [Asset](#dfinance.oracle.v1beta1.Asset) | repeated |  |






<a name="dfinance.oracle.v1beta1.QueryCurrentPriceRequest"></a>

### QueryCurrentPriceRequest
QueryAssetsRequest is request type for Query/CurrentPrice RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `left_denom` | [string](#string) |  |  |
| `right_denom` | [string](#string) |  |  |






<a name="dfinance.oracle.v1beta1.QueryCurrentPriceResponse"></a>

### QueryCurrentPriceResponse
QueryAssetsResponse is response type for Query/CurrentPrice RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `price` | [CurrentPrice](#dfinance.oracle.v1beta1.CurrentPrice) |  |  |






<a name="dfinance.oracle.v1beta1.QueryCurrentPricesRequest"></a>

### QueryCurrentPricesRequest
QueryCurrentPricesRequest is request type for Query/CurrentPrices RPC method.






<a name="dfinance.oracle.v1beta1.QueryCurrentPricesResponse"></a>

### QueryCurrentPricesResponse
QueryCurrentPricesResponse is response type for Query/CurrentPrices RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `prices` | [CurrentPrice](#dfinance.oracle.v1beta1.CurrentPrice) | repeated |  |






<a name="dfinance.oracle.v1beta1.QueryOraclesRequest"></a>

### QueryOraclesRequest
QueryOraclesRequest is request type for Query/Oracles RPC method.






<a name="dfinance.oracle.v1beta1.QueryOraclesResponse"></a>

### QueryOraclesResponse
QueryOraclesResponse is response type for Query/Oracles RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `oracles` | [Oracle](#dfinance.oracle.v1beta1.Oracle) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="dfinance.oracle.v1beta1.Query"></a>

### Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Oracles` | [QueryOraclesRequest](#dfinance.oracle.v1beta1.QueryOraclesRequest) | [QueryOraclesResponse](#dfinance.oracle.v1beta1.QueryOraclesResponse) | Oracles queries registered Oracle source list | GET|/dfinance/oracle/v1beta1/oracles|
| `Assets` | [QueryAssetsRequest](#dfinance.oracle.v1beta1.QueryAssetsRequest) | [QueryAssetsResponse](#dfinance.oracle.v1beta1.QueryAssetsResponse) | Assets queries registered Asset list | GET|/dfinance/oracle/v1beta1/assets|
| `CurrentPrice` | [QueryCurrentPriceRequest](#dfinance.oracle.v1beta1.QueryCurrentPriceRequest) | [QueryCurrentPriceResponse](#dfinance.oracle.v1beta1.QueryCurrentPriceResponse) | CurrentPrice queries current price for an Asset | GET|/dfinance/oracle/v1beta1/current_price|
| `CurrentPrices` | [QueryCurrentPricesRequest](#dfinance.oracle.v1beta1.QueryCurrentPricesRequest) | [QueryCurrentPricesResponse](#dfinance.oracle.v1beta1.QueryCurrentPricesResponse) | CurrentPrices queries current prices for all registered Assets | GET|/dfinance/oracle/v1beta1/current_prices|

 <!-- end services -->



<a name="dfinance/oracle/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dfinance/oracle/tx.proto



<a name="dfinance.oracle.v1beta1.MsgPostPrice"></a>

### MsgPostPrice
MsgPostPrice defines a SDK message to post a raw price from source (Oracle).


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `asset_code` | [string](#string) |  | Asset code (for ex.: btc_usdt) |
| `oracle_address` | [string](#string) |  | Price source (Oracle address) |
| `ask_price` | [string](#string) |  | The lowest seller price |
| `bid_price` | [string](#string) |  | The highest buyer price |
| `received_at` | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | Price timestamp |






<a name="dfinance.oracle.v1beta1.MsgPostPriceResponse"></a>

### MsgPostPriceResponse







<a name="dfinance.oracle.v1beta1.MsgSetAsset"></a>

### MsgSetAsset
MsgSetAsset defines a SDK message to create/update an Asset.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nominee` | [string](#string) |  | Nominee account address |
| `asset` | [Asset](#dfinance.oracle.v1beta1.Asset) |  | Target Asset to create/update |






<a name="dfinance.oracle.v1beta1.MsgSetAssetResponse"></a>

### MsgSetAssetResponse







<a name="dfinance.oracle.v1beta1.MsgSetOracle"></a>

### MsgSetOracle
MsgSetOracle defines a SDK message to create/update an Oracle.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nominee` | [string](#string) |  | Nominee account address |
| `oracle` | [Oracle](#dfinance.oracle.v1beta1.Oracle) |  | Target Oracle to create/update |






<a name="dfinance.oracle.v1beta1.MsgSetOracleResponse"></a>

### MsgSetOracleResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="dfinance.oracle.v1beta1.Msg"></a>

### Msg
Msg defines the Oracle module Msg service.

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `SetOracle` | [MsgSetOracle](#dfinance.oracle.v1beta1.MsgSetOracle) | [MsgSetOracleResponse](#dfinance.oracle.v1beta1.MsgSetOracleResponse) | SetOracle creates/updates an Oracle source (nominee authorized). | |
| `SetAsset` | [MsgSetAsset](#dfinance.oracle.v1beta1.MsgSetAsset) | [MsgSetAssetResponse](#dfinance.oracle.v1beta1.MsgSetAssetResponse) | SetAsset creates/updates an Asset (nominee authorized). | |
| `PostPrice` | [MsgPostPrice](#dfinance.oracle.v1beta1.MsgPostPrice) | [MsgPostPriceResponse](#dfinance.oracle.v1beta1.MsgPostPriceResponse) | PostPrice posts a raw price from a source (Oracle) | |

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



<a name="dfinance/vm/gov.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dfinance/vm/gov.proto



<a name="dfinance.vm.v1beta1.PlannedProposal"></a>

### PlannedProposal
PlannedProposal defines VM Gov proposal with apply schedule and wrapped proposal content.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `height` | [int64](#int64) |  | Height is a block height proposal should be applied at |
| `content` | [google.protobuf.Any](#google.protobuf.Any) |  | Content is a Gov proposal content |






<a name="dfinance.vm.v1beta1.StdLibUpdateProposal"></a>

### StdLibUpdateProposal



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `url` | [string](#string) |  | Url contains Stdlib update source code |
| `update_description` | [string](#string) |  | UpdateDescription contains some update description |
| `code` | [bytes](#bytes) | repeated | Code is a DVM byteCode of updated modules |





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






<a name="dfinance.vm.v1beta1.QueryDelegatedPoolSupplyRequest"></a>

### QueryDelegatedPoolSupplyRequest
QueryDelegatedPoolSupplyRequest is request type for Query/DelegatedPoolSupply RPC method.






<a name="dfinance.vm.v1beta1.QueryDelegatedPoolSupplyResponse"></a>

### QueryDelegatedPoolSupplyResponse
QueryDelegatedPoolSupplyResponse is response type for Query/DelegatedPoolSupply RPC method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `coins` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) | repeated |  |






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
| `DelegatedPoolSupply` | [QueryDelegatedPoolSupplyRequest](#dfinance.vm.v1beta1.QueryDelegatedPoolSupplyRequest) | [QueryDelegatedPoolSupplyResponse](#dfinance.vm.v1beta1.QueryDelegatedPoolSupplyResponse) | DelegatedPoolSupply queries Delegated pool module balance. | |

 <!-- end services -->



<a name="dfinance/vm/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## dfinance/vm/tx.proto



<a name="dfinance.vm.v1beta1.MsgDeployModule"></a>

### MsgDeployModule
MsgDeployModule defines a SDK message to deploy a module (contract) to VM.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `signer` | [string](#string) |  | Script sender address |
| `modules` | [bytes](#bytes) | repeated | Module code |






<a name="dfinance.vm.v1beta1.MsgDeployModuleResponse"></a>

### MsgDeployModuleResponse







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



<a name="gravity/v1/attestation.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gravity/v1/attestation.proto



<a name="gravity.v1.Attestation"></a>

### Attestation
Attestation is an aggregate of `claims` that eventually becomes `observed` by
all orchestrators
EVENT_NONCE:
EventNonce a nonce provided by the gravity contract that is unique per event fired
These event nonces must be relayed in order. This is a correctness issue,
if relaying out of order transaction replay attacks become possible
OBSERVED:
Observed indicates that >67% of validators have attested to the event,
and that the event should be executed by the gravity state machine

The actual content of the claims is passed in with the transaction making the claim
and then passed through the call stack alongside the attestation while it is processed
the key in which the attestation is stored is keyed on the exact details of the claim
but there is no reason to store those exact details becuause the next message sender
will kindly provide you with them.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `observed` | [bool](#bool) |  |  |
| `votes` | [string](#string) | repeated |  |
| `height` | [uint64](#uint64) |  |  |
| `claim` | [google.protobuf.Any](#google.protobuf.Any) |  |  |






<a name="gravity.v1.ERC20Token"></a>

### ERC20Token
ERC20Token unique identifier for an Ethereum ERC20 token.
CONTRACT:
The contract address on ETH of the token, this could be a Cosmos
originated token, if so it will be the ERC20 address of the representation
(note: developers should look up the token symbol using the address on ETH to display for UI)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `contract` | [string](#string) |  |  |
| `amount` | [string](#string) |  |  |





 <!-- end messages -->


<a name="gravity.v1.ClaimType"></a>

### ClaimType
ClaimType is the cosmos type of an event from the counterpart chain that can
be handled

| Name | Number | Description |
| ---- | ------ | ----------- |
| CLAIM_TYPE_UNSPECIFIED | 0 |  |
| CLAIM_TYPE_DEPOSIT | 1 |  |
| CLAIM_TYPE_WITHDRAW | 2 |  |
| CLAIM_TYPE_ERC20_DEPLOYED | 3 |  |
| CLAIM_TYPE_LOGIC_CALL_EXECUTED | 4 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gravity/v1/batch.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gravity/v1/batch.proto



<a name="gravity.v1.OutgoingLogicCall"></a>

### OutgoingLogicCall
OutgoingLogicCall represents an individual logic call from gravity to ETH


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `transfers` | [ERC20Token](#gravity.v1.ERC20Token) | repeated |  |
| `fees` | [ERC20Token](#gravity.v1.ERC20Token) | repeated |  |
| `logic_contract_address` | [string](#string) |  |  |
| `payload` | [bytes](#bytes) |  |  |
| `timeout` | [uint64](#uint64) |  |  |
| `invalidation_id` | [bytes](#bytes) |  |  |
| `invalidation_nonce` | [uint64](#uint64) |  |  |






<a name="gravity.v1.OutgoingTransferTx"></a>

### OutgoingTransferTx
OutgoingTransferTx represents an individual send from gravity to ETH


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `id` | [uint64](#uint64) |  |  |
| `sender` | [string](#string) |  |  |
| `dest_address` | [string](#string) |  |  |
| `erc20_token` | [ERC20Token](#gravity.v1.ERC20Token) |  |  |
| `erc20_fee` | [ERC20Token](#gravity.v1.ERC20Token) |  |  |






<a name="gravity.v1.OutgoingTxBatch"></a>

### OutgoingTxBatch
OutgoingTxBatch represents a batch of transactions going from gravity to ETH


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `batch_nonce` | [uint64](#uint64) |  |  |
| `batch_timeout` | [uint64](#uint64) |  |  |
| `transactions` | [OutgoingTransferTx](#gravity.v1.OutgoingTransferTx) | repeated |  |
| `token_contract` | [string](#string) |  |  |
| `block` | [uint64](#uint64) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gravity/v1/ethereum_signer.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gravity/v1/ethereum_signer.proto


 <!-- end messages -->


<a name="gravity.v1.SignType"></a>

### SignType
SignType defines messages that have been signed by an orchestrator

| Name | Number | Description |
| ---- | ------ | ----------- |
| SIGN_TYPE_UNSPECIFIED | 0 |  |
| SIGN_TYPE_ORCHESTRATOR_SIGNED_MULTI_SIG_UPDATE | 1 |  |
| SIGN_TYPE_ORCHESTRATOR_SIGNED_WITHDRAW_BATCH | 2 |  |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gravity/v1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gravity/v1/types.proto



<a name="gravity.v1.BridgeValidator"></a>

### BridgeValidator
BridgeValidator represents a validator's ETH address and its power


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `power` | [uint64](#uint64) |  |  |
| `ethereum_address` | [string](#string) |  |  |






<a name="gravity.v1.ERC20ToDenom"></a>

### ERC20ToDenom
This records the relationship between an ERC20 token and the denom
of the corresponding Cosmos originated asset


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `erc20` | [string](#string) |  |  |
| `denom` | [string](#string) |  |  |






<a name="gravity.v1.LastObservedEthereumBlockHeight"></a>

### LastObservedEthereumBlockHeight
LastObservedEthereumBlockHeight stores the last observed
Ethereum block height along with the Cosmos block height that
it was observed at. These two numbers can be used to project
outward and always produce batches with timeouts in the future
even if no Ethereum block height has been relayed for a long time


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `cosmos_block_height` | [uint64](#uint64) |  |  |
| `ethereum_block_height` | [uint64](#uint64) |  |  |






<a name="gravity.v1.Valset"></a>

### Valset
Valset is the Ethereum Bridge Multsig Set, each gravity validator also
maintains an ETH key to sign messages, these are used to check signatures on
ETH because of the significant gas savings


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nonce` | [uint64](#uint64) |  |  |
| `members` | [BridgeValidator](#gravity.v1.BridgeValidator) | repeated |  |
| `height` | [uint64](#uint64) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gravity/v1/msgs.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gravity/v1/msgs.proto



<a name="gravity.v1.MsgCancelSendToEth"></a>

### MsgCancelSendToEth
This call allows the sender (and only the sender)
to cancel a given MsgSendToEth and recieve a refund
of the tokens


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `transaction_id` | [uint64](#uint64) |  |  |
| `sender` | [string](#string) |  |  |






<a name="gravity.v1.MsgCancelSendToEthResponse"></a>

### MsgCancelSendToEthResponse







<a name="gravity.v1.MsgConfirmBatch"></a>

### MsgConfirmBatch
MsgConfirmBatch
When validators observe a MsgRequestBatch they form a batch by ordering
transactions currently in the txqueue in order of highest to lowest fee,
cutting off when the batch either reaches a hardcoded maximum size (to be
decided, probably around 100) or when transactions stop being profitable
(TODO determine this without nondeterminism) This message includes the batch
as well as an Ethereum signature over this batch by the validator
-------------


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nonce` | [uint64](#uint64) |  |  |
| `token_contract` | [string](#string) |  |  |
| `eth_signer` | [string](#string) |  |  |
| `orchestrator` | [string](#string) |  |  |
| `signature` | [string](#string) |  |  |






<a name="gravity.v1.MsgConfirmBatchResponse"></a>

### MsgConfirmBatchResponse







<a name="gravity.v1.MsgConfirmLogicCall"></a>

### MsgConfirmLogicCall
MsgConfirmLogicCall
When validators observe a MsgRequestBatch they form a batch by ordering
transactions currently in the txqueue in order of highest to lowest fee,
cutting off when the batch either reaches a hardcoded maximum size (to be
decided, probably around 100) or when transactions stop being profitable
(TODO determine this without nondeterminism) This message includes the batch
as well as an Ethereum signature over this batch by the validator
-------------


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `invalidation_id` | [string](#string) |  |  |
| `invalidation_nonce` | [uint64](#uint64) |  |  |
| `eth_signer` | [string](#string) |  |  |
| `orchestrator` | [string](#string) |  |  |
| `signature` | [string](#string) |  |  |






<a name="gravity.v1.MsgConfirmLogicCallResponse"></a>

### MsgConfirmLogicCallResponse







<a name="gravity.v1.MsgDepositClaim"></a>

### MsgDepositClaim
EthereumBridgeDepositClaim
When more than 66% of the active validator set has
claimed to have seen the deposit enter the ethereum blockchain coins are
issued to the Cosmos address in question
-------------


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `event_nonce` | [uint64](#uint64) |  |  |
| `block_height` | [uint64](#uint64) |  |  |
| `token_contract` | [string](#string) |  |  |
| `amount` | [string](#string) |  |  |
| `ethereum_sender` | [string](#string) |  |  |
| `cosmos_receiver` | [string](#string) |  |  |
| `orchestrator` | [string](#string) |  |  |






<a name="gravity.v1.MsgDepositClaimResponse"></a>

### MsgDepositClaimResponse







<a name="gravity.v1.MsgERC20DeployedClaim"></a>

### MsgERC20DeployedClaim
ERC20DeployedClaim allows the Cosmos module
to learn about an ERC20 that someone deployed
to represent a Cosmos asset


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `event_nonce` | [uint64](#uint64) |  |  |
| `block_height` | [uint64](#uint64) |  |  |
| `cosmos_denom` | [string](#string) |  |  |
| `token_contract` | [string](#string) |  |  |
| `name` | [string](#string) |  |  |
| `symbol` | [string](#string) |  |  |
| `decimals` | [uint64](#uint64) |  |  |
| `orchestrator` | [string](#string) |  |  |






<a name="gravity.v1.MsgERC20DeployedClaimResponse"></a>

### MsgERC20DeployedClaimResponse







<a name="gravity.v1.MsgLogicCallExecutedClaim"></a>

### MsgLogicCallExecutedClaim
This informs the Cosmos module that a logic
call has been executed


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `event_nonce` | [uint64](#uint64) |  |  |
| `block_height` | [uint64](#uint64) |  |  |
| `invalidation_id` | [bytes](#bytes) |  |  |
| `invalidation_nonce` | [uint64](#uint64) |  |  |
| `orchestrator` | [string](#string) |  |  |






<a name="gravity.v1.MsgLogicCallExecutedClaimResponse"></a>

### MsgLogicCallExecutedClaimResponse







<a name="gravity.v1.MsgRequestBatch"></a>

### MsgRequestBatch
MsgRequestBatch
this is a message anyone can send that requests a batch of transactions to
send across the bridge be created for whatever block height this message is
included in. This acts as a coordination point, the handler for this message
looks at the AddToOutgoingPool tx's in the store and generates a batch, also
available in the store tied to this message. The validators then grab this
batch, sign it, submit the signatures with a MsgConfirmBatch before a relayer
can finally submit the batch
-------------


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `denom` | [string](#string) |  |  |






<a name="gravity.v1.MsgRequestBatchResponse"></a>

### MsgRequestBatchResponse







<a name="gravity.v1.MsgSendToEth"></a>

### MsgSendToEth
MsgSendToEth
This is the message that a user calls when they want to bridge an asset
it will later be removed when it is included in a batch and successfully
submitted tokens are removed from the users balance immediately
-------------
AMOUNT:
the coin to send across the bridge, note the restriction that this is a
single coin not a set of coins that is normal in other Cosmos messages
FEE:
the fee paid for the bridge, distinct from the fee paid to the chain to
actually send this message in the first place. So a successful send has
two layers of fees for the user


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender` | [string](#string) |  |  |
| `eth_dest` | [string](#string) |  |  |
| `amount` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |
| `bridge_fee` | [cosmos.base.v1beta1.Coin](#cosmos.base.v1beta1.Coin) |  |  |






<a name="gravity.v1.MsgSendToEthResponse"></a>

### MsgSendToEthResponse







<a name="gravity.v1.MsgSetOrchestratorAddress"></a>

### MsgSetOrchestratorAddress
MsgSetOrchestratorAddress
this message allows validators to delegate their voting responsibilities
to a given key. This key is then used as an optional authentication method
for sigining oracle claims
VALIDATOR
The validator field is a cosmosvaloper1... string (i.e. sdk.ValAddress)
that references a validator in the active set
ORCHESTRATOR
The orchestrator field is a cosmos1... string  (i.e. sdk.AccAddress) that
references the key that is being delegated to
ETH_ADDRESS
This is a hex encoded 0x Ethereum public key that will be used by this validator
on Ethereum


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator` | [string](#string) |  |  |
| `orchestrator` | [string](#string) |  |  |
| `eth_address` | [string](#string) |  |  |






<a name="gravity.v1.MsgSetOrchestratorAddressResponse"></a>

### MsgSetOrchestratorAddressResponse







<a name="gravity.v1.MsgValsetConfirm"></a>

### MsgValsetConfirm
MsgValsetConfirm
this is the message sent by the validators when they wish to submit their
signatures over the validator set at a given block height. A validator must
first call MsgSetEthAddress to set their Ethereum address to be used for
signing. Then someone (anyone) must make a ValsetRequest, the request is
essentially a messaging mechanism to determine which block all validators
should submit signatures over. Finally validators sign the validator set,
powers, and Ethereum addresses of the entire validator set at the height of a
ValsetRequest and submit that signature with this message.

If a sufficient number of validators (66% of voting power) (A) have set
Ethereum addresses and (B) submit ValsetConfirm messages with their
signatures it is then possible for anyone to view these signatures in the
chain store and submit them to Ethereum to update the validator set
-------------


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nonce` | [uint64](#uint64) |  |  |
| `orchestrator` | [string](#string) |  |  |
| `eth_address` | [string](#string) |  |  |
| `signature` | [string](#string) |  |  |






<a name="gravity.v1.MsgValsetConfirmResponse"></a>

### MsgValsetConfirmResponse







<a name="gravity.v1.MsgWithdrawClaim"></a>

### MsgWithdrawClaim
WithdrawClaim claims that a batch of withdrawal
operations on the bridge contract was executed.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `event_nonce` | [uint64](#uint64) |  |  |
| `block_height` | [uint64](#uint64) |  |  |
| `batch_nonce` | [uint64](#uint64) |  |  |
| `token_contract` | [string](#string) |  |  |
| `orchestrator` | [string](#string) |  |  |






<a name="gravity.v1.MsgWithdrawClaimResponse"></a>

### MsgWithdrawClaimResponse






 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gravity.v1.Msg"></a>

### Msg
Msg defines the state transitions possible within gravity

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `ValsetConfirm` | [MsgValsetConfirm](#gravity.v1.MsgValsetConfirm) | [MsgValsetConfirmResponse](#gravity.v1.MsgValsetConfirmResponse) |  | POST|/gravity/v1/valset_confirm|
| `SendToEth` | [MsgSendToEth](#gravity.v1.MsgSendToEth) | [MsgSendToEthResponse](#gravity.v1.MsgSendToEthResponse) |  | POST|/gravity/v1/send_to_eth|
| `RequestBatch` | [MsgRequestBatch](#gravity.v1.MsgRequestBatch) | [MsgRequestBatchResponse](#gravity.v1.MsgRequestBatchResponse) |  | POST|/gravity/v1/request_batch|
| `ConfirmBatch` | [MsgConfirmBatch](#gravity.v1.MsgConfirmBatch) | [MsgConfirmBatchResponse](#gravity.v1.MsgConfirmBatchResponse) |  | POST|/gravity/v1/confirm_batch|
| `ConfirmLogicCall` | [MsgConfirmLogicCall](#gravity.v1.MsgConfirmLogicCall) | [MsgConfirmLogicCallResponse](#gravity.v1.MsgConfirmLogicCallResponse) |  | POST|/gravity/v1/confim_logic|
| `DepositClaim` | [MsgDepositClaim](#gravity.v1.MsgDepositClaim) | [MsgDepositClaimResponse](#gravity.v1.MsgDepositClaimResponse) |  | POST|/gravity/v1/deposit_claim|
| `WithdrawClaim` | [MsgWithdrawClaim](#gravity.v1.MsgWithdrawClaim) | [MsgWithdrawClaimResponse](#gravity.v1.MsgWithdrawClaimResponse) |  | POST|/gravity/v1/withdraw_claim|
| `ERC20DeployedClaim` | [MsgERC20DeployedClaim](#gravity.v1.MsgERC20DeployedClaim) | [MsgERC20DeployedClaimResponse](#gravity.v1.MsgERC20DeployedClaimResponse) |  | POST|/gravity/v1/erc20_deployed_claim|
| `LogicCallExecutedClaim` | [MsgLogicCallExecutedClaim](#gravity.v1.MsgLogicCallExecutedClaim) | [MsgLogicCallExecutedClaimResponse](#gravity.v1.MsgLogicCallExecutedClaimResponse) |  | POST|/gravity/v1/logic_call_executed_claim|
| `SetOrchestratorAddress` | [MsgSetOrchestratorAddress](#gravity.v1.MsgSetOrchestratorAddress) | [MsgSetOrchestratorAddressResponse](#gravity.v1.MsgSetOrchestratorAddressResponse) |  | POST|/gravity/v1/set_orchestrator_address|
| `CancelSendToEth` | [MsgCancelSendToEth](#gravity.v1.MsgCancelSendToEth) | [MsgCancelSendToEthResponse](#gravity.v1.MsgCancelSendToEthResponse) |  | POST|/gravity/v1/cancel_send_to_eth|

 <!-- end services -->



<a name="gravity/v1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gravity/v1/genesis.proto



<a name="gravity.v1.GenesisState"></a>

### GenesisState
GenesisState struct


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#gravity.v1.Params) |  |  |
| `last_observed_nonce` | [uint64](#uint64) |  |  |
| `valsets` | [Valset](#gravity.v1.Valset) | repeated |  |
| `valset_confirms` | [MsgValsetConfirm](#gravity.v1.MsgValsetConfirm) | repeated |  |
| `batches` | [OutgoingTxBatch](#gravity.v1.OutgoingTxBatch) | repeated |  |
| `batch_confirms` | [MsgConfirmBatch](#gravity.v1.MsgConfirmBatch) | repeated |  |
| `logic_calls` | [OutgoingLogicCall](#gravity.v1.OutgoingLogicCall) | repeated |  |
| `logic_call_confirms` | [MsgConfirmLogicCall](#gravity.v1.MsgConfirmLogicCall) | repeated |  |
| `attestations` | [Attestation](#gravity.v1.Attestation) | repeated |  |
| `delegate_keys` | [MsgSetOrchestratorAddress](#gravity.v1.MsgSetOrchestratorAddress) | repeated |  |
| `erc20_to_denoms` | [ERC20ToDenom](#gravity.v1.ERC20ToDenom) | repeated |  |
| `unbatched_transfers` | [OutgoingTransferTx](#gravity.v1.OutgoingTransferTx) | repeated |  |






<a name="gravity.v1.Params"></a>

### Params
Params represent the Gravity genesis and store parameters
gravity_id:
a random 32 byte value to prevent signature reuse, for example if the
cosmos validators decided to use the same Ethereum keys for another chain
also running Gravity we would not want it to be possible to play a deposit
from chain A back on chain B's Gravity. This value IS USED ON ETHEREUM so
it must be set in your genesis.json before launch and not changed after
deploying Gravity

contract_hash:
the code hash of a known good version of the Gravity contract
solidity code. This can be used to verify the correct version
of the contract has been deployed. This is a reference value for
goernance action only it is never read by any Gravity code

bridge_ethereum_address:
is address of the bridge contract on the Ethereum side, this is a
reference value for governance only and is not actually used by any
Gravity code

bridge_chain_id:
the unique identifier of the Ethereum chain, this is a reference value
only and is not actually used by any Gravity code

These reference values may be used by future Gravity client implemetnations
to allow for saftey features or convenience features like the Gravity address
in your relayer. A relayer would require a configured Gravity address if
governance had not set the address on the chain it was relaying for.

signed_valsets_window
signed_batches_window
signed_claims_window

These values represent the time in blocks that a validator has to submit
a signature for a batch or valset, or to submit a claim for a particular
attestation nonce. In the case of attestations this clock starts when the
attestation is created, but only allows for slashing once the event has passed

target_batch_timeout:

This is the 'target' value for when batches time out, this is a target becuase
Ethereum is a probabalistic chain and you can't say for sure what the block
frequency is ahead of time.

average_block_time
average_ethereum_block_time

These values are the average Cosmos block time and Ethereum block time repsectively
and they are used to copute what the target batch timeout is. It is important that
governance updates these in case of any major, prolonged change in the time it takes
to produce a block

slash_fraction_valset
slash_fraction_batch
slash_fraction_claim
slash_fraction_conflicting_claim

The slashing fractions for the various gravity related slashing conditions. The first three
refer to not submitting a particular message, the third for submitting a different claim
for the same Ethereum event


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `gravity_id` | [string](#string) |  |  |
| `contract_source_hash` | [string](#string) |  |  |
| `bridge_ethereum_address` | [string](#string) |  |  |
| `bridge_chain_id` | [uint64](#uint64) |  |  |
| `signed_valsets_window` | [uint64](#uint64) |  |  |
| `signed_batches_window` | [uint64](#uint64) |  |  |
| `signed_claims_window` | [uint64](#uint64) |  |  |
| `target_batch_timeout` | [uint64](#uint64) |  |  |
| `average_block_time` | [uint64](#uint64) |  |  |
| `average_ethereum_block_time` | [uint64](#uint64) |  |  |
| `slash_fraction_valset` | [bytes](#bytes) |  |  |
| `slash_fraction_batch` | [bytes](#bytes) |  |  |
| `slash_fraction_claim` | [bytes](#bytes) |  |  |
| `slash_fraction_conflicting_claim` | [bytes](#bytes) |  |  |
| `unbond_slashing_valsets_window` | [uint64](#uint64) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gravity/v1/pool.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gravity/v1/pool.proto



<a name="gravity.v1.BatchFees"></a>

### BatchFees



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `token` | [string](#string) |  |  |
| `total_fees` | [string](#string) |  |  |






<a name="gravity.v1.IDSet"></a>

### IDSet
IDSet represents a set of IDs


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `ids` | [uint64](#uint64) | repeated |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="gravity/v1/query.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## gravity/v1/query.proto



<a name="gravity.v1.QueryBatchConfirmsRequest"></a>

### QueryBatchConfirmsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nonce` | [uint64](#uint64) |  |  |
| `contract_address` | [string](#string) |  |  |






<a name="gravity.v1.QueryBatchConfirmsResponse"></a>

### QueryBatchConfirmsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `confirms` | [MsgConfirmBatch](#gravity.v1.MsgConfirmBatch) | repeated |  |






<a name="gravity.v1.QueryBatchFeeRequest"></a>

### QueryBatchFeeRequest







<a name="gravity.v1.QueryBatchFeeResponse"></a>

### QueryBatchFeeResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `batch_fees` | [BatchFees](#gravity.v1.BatchFees) | repeated |  |






<a name="gravity.v1.QueryBatchRequestByNonceRequest"></a>

### QueryBatchRequestByNonceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nonce` | [uint64](#uint64) |  |  |
| `contract_address` | [string](#string) |  |  |






<a name="gravity.v1.QueryBatchRequestByNonceResponse"></a>

### QueryBatchRequestByNonceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `batch` | [OutgoingTxBatch](#gravity.v1.OutgoingTxBatch) |  |  |






<a name="gravity.v1.QueryCurrentValsetRequest"></a>

### QueryCurrentValsetRequest







<a name="gravity.v1.QueryCurrentValsetResponse"></a>

### QueryCurrentValsetResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `valset` | [Valset](#gravity.v1.Valset) |  |  |






<a name="gravity.v1.QueryDelegateKeysByEthAddress"></a>

### QueryDelegateKeysByEthAddress



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `eth_address` | [string](#string) |  |  |






<a name="gravity.v1.QueryDelegateKeysByEthAddressResponse"></a>

### QueryDelegateKeysByEthAddressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator_address` | [string](#string) |  |  |
| `orchestrator_address` | [string](#string) |  |  |






<a name="gravity.v1.QueryDelegateKeysByOrchestratorAddress"></a>

### QueryDelegateKeysByOrchestratorAddress



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `orchestrator_address` | [string](#string) |  |  |






<a name="gravity.v1.QueryDelegateKeysByOrchestratorAddressResponse"></a>

### QueryDelegateKeysByOrchestratorAddressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator_address` | [string](#string) |  |  |
| `eth_address` | [string](#string) |  |  |






<a name="gravity.v1.QueryDelegateKeysByValidatorAddress"></a>

### QueryDelegateKeysByValidatorAddress



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `validator_address` | [string](#string) |  |  |






<a name="gravity.v1.QueryDelegateKeysByValidatorAddressResponse"></a>

### QueryDelegateKeysByValidatorAddressResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `eth_address` | [string](#string) |  |  |
| `orchestrator_address` | [string](#string) |  |  |






<a name="gravity.v1.QueryDenomToERC20Request"></a>

### QueryDenomToERC20Request



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |






<a name="gravity.v1.QueryDenomToERC20Response"></a>

### QueryDenomToERC20Response



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `erc20` | [string](#string) |  |  |
| `cosmos_originated` | [bool](#bool) |  |  |






<a name="gravity.v1.QueryERC20ToDenomRequest"></a>

### QueryERC20ToDenomRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `erc20` | [string](#string) |  |  |






<a name="gravity.v1.QueryERC20ToDenomResponse"></a>

### QueryERC20ToDenomResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `denom` | [string](#string) |  |  |
| `cosmos_originated` | [bool](#bool) |  |  |






<a name="gravity.v1.QueryLastEventNonceByAddrRequest"></a>

### QueryLastEventNonceByAddrRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |






<a name="gravity.v1.QueryLastEventNonceByAddrResponse"></a>

### QueryLastEventNonceByAddrResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `event_nonce` | [uint64](#uint64) |  |  |






<a name="gravity.v1.QueryLastPendingBatchRequestByAddrRequest"></a>

### QueryLastPendingBatchRequestByAddrRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |






<a name="gravity.v1.QueryLastPendingBatchRequestByAddrResponse"></a>

### QueryLastPendingBatchRequestByAddrResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `batch` | [OutgoingTxBatch](#gravity.v1.OutgoingTxBatch) |  |  |






<a name="gravity.v1.QueryLastPendingLogicCallByAddrRequest"></a>

### QueryLastPendingLogicCallByAddrRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |






<a name="gravity.v1.QueryLastPendingLogicCallByAddrResponse"></a>

### QueryLastPendingLogicCallByAddrResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `call` | [OutgoingLogicCall](#gravity.v1.OutgoingLogicCall) |  |  |






<a name="gravity.v1.QueryLastPendingValsetRequestByAddrRequest"></a>

### QueryLastPendingValsetRequestByAddrRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `address` | [string](#string) |  |  |






<a name="gravity.v1.QueryLastPendingValsetRequestByAddrResponse"></a>

### QueryLastPendingValsetRequestByAddrResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `valsets` | [Valset](#gravity.v1.Valset) | repeated |  |






<a name="gravity.v1.QueryLastValsetRequestsRequest"></a>

### QueryLastValsetRequestsRequest







<a name="gravity.v1.QueryLastValsetRequestsResponse"></a>

### QueryLastValsetRequestsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `valsets` | [Valset](#gravity.v1.Valset) | repeated |  |






<a name="gravity.v1.QueryLogicConfirmsRequest"></a>

### QueryLogicConfirmsRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `invalidation_id` | [bytes](#bytes) |  |  |
| `invalidation_nonce` | [uint64](#uint64) |  |  |






<a name="gravity.v1.QueryLogicConfirmsResponse"></a>

### QueryLogicConfirmsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `confirms` | [MsgConfirmLogicCall](#gravity.v1.MsgConfirmLogicCall) | repeated |  |






<a name="gravity.v1.QueryOutgoingLogicCallsRequest"></a>

### QueryOutgoingLogicCallsRequest







<a name="gravity.v1.QueryOutgoingLogicCallsResponse"></a>

### QueryOutgoingLogicCallsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `calls` | [OutgoingLogicCall](#gravity.v1.OutgoingLogicCall) | repeated |  |






<a name="gravity.v1.QueryOutgoingTxBatchesRequest"></a>

### QueryOutgoingTxBatchesRequest







<a name="gravity.v1.QueryOutgoingTxBatchesResponse"></a>

### QueryOutgoingTxBatchesResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `batches` | [OutgoingTxBatch](#gravity.v1.OutgoingTxBatch) | repeated |  |






<a name="gravity.v1.QueryParamsRequest"></a>

### QueryParamsRequest







<a name="gravity.v1.QueryParamsResponse"></a>

### QueryParamsResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `params` | [Params](#gravity.v1.Params) |  |  |






<a name="gravity.v1.QueryPendingSendToEth"></a>

### QueryPendingSendToEth



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `sender_address` | [string](#string) |  |  |






<a name="gravity.v1.QueryPendingSendToEthResponse"></a>

### QueryPendingSendToEthResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `transfers_in_batches` | [OutgoingTransferTx](#gravity.v1.OutgoingTransferTx) | repeated |  |
| `unbatched_transfers` | [OutgoingTransferTx](#gravity.v1.OutgoingTransferTx) | repeated |  |






<a name="gravity.v1.QueryValsetConfirmRequest"></a>

### QueryValsetConfirmRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nonce` | [uint64](#uint64) |  |  |
| `address` | [string](#string) |  |  |






<a name="gravity.v1.QueryValsetConfirmResponse"></a>

### QueryValsetConfirmResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `confirm` | [MsgValsetConfirm](#gravity.v1.MsgValsetConfirm) |  |  |






<a name="gravity.v1.QueryValsetConfirmsByNonceRequest"></a>

### QueryValsetConfirmsByNonceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nonce` | [uint64](#uint64) |  |  |






<a name="gravity.v1.QueryValsetConfirmsByNonceResponse"></a>

### QueryValsetConfirmsByNonceResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `confirms` | [MsgValsetConfirm](#gravity.v1.MsgValsetConfirm) | repeated |  |






<a name="gravity.v1.QueryValsetRequestRequest"></a>

### QueryValsetRequestRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `nonce` | [uint64](#uint64) |  |  |






<a name="gravity.v1.QueryValsetRequestResponse"></a>

### QueryValsetRequestResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| `valset` | [Valset](#gravity.v1.Valset) |  |  |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="gravity.v1.Query"></a>

### Query
Query defines the gRPC querier service

| Method Name | Request Type | Response Type | Description | HTTP Verb | Endpoint |
| ----------- | ------------ | ------------- | ------------| ------- | -------- |
| `Params` | [QueryParamsRequest](#gravity.v1.QueryParamsRequest) | [QueryParamsResponse](#gravity.v1.QueryParamsResponse) | Deployments queries deployments | GET|/gravity/v1beta/params|
| `CurrentValset` | [QueryCurrentValsetRequest](#gravity.v1.QueryCurrentValsetRequest) | [QueryCurrentValsetResponse](#gravity.v1.QueryCurrentValsetResponse) |  | GET|/gravity/v1beta/valset/current|
| `ValsetRequest` | [QueryValsetRequestRequest](#gravity.v1.QueryValsetRequestRequest) | [QueryValsetRequestResponse](#gravity.v1.QueryValsetRequestResponse) |  | GET|/gravity/v1beta/valset|
| `ValsetConfirm` | [QueryValsetConfirmRequest](#gravity.v1.QueryValsetConfirmRequest) | [QueryValsetConfirmResponse](#gravity.v1.QueryValsetConfirmResponse) |  | GET|/gravity/v1beta/valset/confirm|
| `ValsetConfirmsByNonce` | [QueryValsetConfirmsByNonceRequest](#gravity.v1.QueryValsetConfirmsByNonceRequest) | [QueryValsetConfirmsByNonceResponse](#gravity.v1.QueryValsetConfirmsByNonceResponse) |  | GET|/gravity/v1beta/confirms/{nonce}|
| `LastValsetRequests` | [QueryLastValsetRequestsRequest](#gravity.v1.QueryLastValsetRequestsRequest) | [QueryLastValsetRequestsResponse](#gravity.v1.QueryLastValsetRequestsResponse) |  | GET|/gravity/v1beta/valset/requests|
| `LastPendingValsetRequestByAddr` | [QueryLastPendingValsetRequestByAddrRequest](#gravity.v1.QueryLastPendingValsetRequestByAddrRequest) | [QueryLastPendingValsetRequestByAddrResponse](#gravity.v1.QueryLastPendingValsetRequestByAddrResponse) |  | GET|/gravity/v1beta/valset/last|
| `LastPendingBatchRequestByAddr` | [QueryLastPendingBatchRequestByAddrRequest](#gravity.v1.QueryLastPendingBatchRequestByAddrRequest) | [QueryLastPendingBatchRequestByAddrResponse](#gravity.v1.QueryLastPendingBatchRequestByAddrResponse) |  | GET|/gravity/v1beta/batch/{address}|
| `LastPendingLogicCallByAddr` | [QueryLastPendingLogicCallByAddrRequest](#gravity.v1.QueryLastPendingLogicCallByAddrRequest) | [QueryLastPendingLogicCallByAddrResponse](#gravity.v1.QueryLastPendingLogicCallByAddrResponse) |  | GET|/gravity/v1beta/logic/{address}|
| `LastEventNonceByAddr` | [QueryLastEventNonceByAddrRequest](#gravity.v1.QueryLastEventNonceByAddrRequest) | [QueryLastEventNonceByAddrResponse](#gravity.v1.QueryLastEventNonceByAddrResponse) |  | GET|/gravity/v1beta/oracle/eventnonce/{address}|
| `BatchFees` | [QueryBatchFeeRequest](#gravity.v1.QueryBatchFeeRequest) | [QueryBatchFeeResponse](#gravity.v1.QueryBatchFeeResponse) |  | GET|/gravity/v1beta/batchfees|
| `OutgoingTxBatches` | [QueryOutgoingTxBatchesRequest](#gravity.v1.QueryOutgoingTxBatchesRequest) | [QueryOutgoingTxBatchesResponse](#gravity.v1.QueryOutgoingTxBatchesResponse) |  | GET|/gravity/v1beta/batch/outgoingtx|
| `OutgoingLogicCalls` | [QueryOutgoingLogicCallsRequest](#gravity.v1.QueryOutgoingLogicCallsRequest) | [QueryOutgoingLogicCallsResponse](#gravity.v1.QueryOutgoingLogicCallsResponse) |  | GET|/gravity/v1beta/batch/outgoinglogic|
| `BatchRequestByNonce` | [QueryBatchRequestByNonceRequest](#gravity.v1.QueryBatchRequestByNonceRequest) | [QueryBatchRequestByNonceResponse](#gravity.v1.QueryBatchRequestByNonceResponse) |  | GET|/gravity/v1beta/batch/{nonce}|
| `BatchConfirms` | [QueryBatchConfirmsRequest](#gravity.v1.QueryBatchConfirmsRequest) | [QueryBatchConfirmsResponse](#gravity.v1.QueryBatchConfirmsResponse) |  | GET|/gravity/v1beta/batch/confirms|
| `LogicConfirms` | [QueryLogicConfirmsRequest](#gravity.v1.QueryLogicConfirmsRequest) | [QueryLogicConfirmsResponse](#gravity.v1.QueryLogicConfirmsResponse) |  | GET|/gravity/v1beta/logic/confirms|
| `ERC20ToDenom` | [QueryERC20ToDenomRequest](#gravity.v1.QueryERC20ToDenomRequest) | [QueryERC20ToDenomResponse](#gravity.v1.QueryERC20ToDenomResponse) |  | GET|/gravity/v1beta/cosmos_originated/erc20_to_denom|
| `DenomToERC20` | [QueryDenomToERC20Request](#gravity.v1.QueryDenomToERC20Request) | [QueryDenomToERC20Response](#gravity.v1.QueryDenomToERC20Response) |  | GET|/gravity/v1beta/cosmos_originated/denom_to_erc20|
| `GetDelegateKeyByValidator` | [QueryDelegateKeysByValidatorAddress](#gravity.v1.QueryDelegateKeysByValidatorAddress) | [QueryDelegateKeysByValidatorAddressResponse](#gravity.v1.QueryDelegateKeysByValidatorAddressResponse) |  | GET|/gravity/v1beta/query_delegate_keys_by_validator|
| `GetDelegateKeyByEth` | [QueryDelegateKeysByEthAddress](#gravity.v1.QueryDelegateKeysByEthAddress) | [QueryDelegateKeysByEthAddressResponse](#gravity.v1.QueryDelegateKeysByEthAddressResponse) |  | GET|/gravity/v1beta/query_delegate_keys_by_eth|
| `GetDelegateKeyByOrchestrator` | [QueryDelegateKeysByOrchestratorAddress](#gravity.v1.QueryDelegateKeysByOrchestratorAddress) | [QueryDelegateKeysByOrchestratorAddressResponse](#gravity.v1.QueryDelegateKeysByOrchestratorAddressResponse) |  | GET|/gravity/v1beta/query_delegate_keys_by_orchestrator|
| `GetPendingSendToEth` | [QueryPendingSendToEth](#gravity.v1.QueryPendingSendToEth) | [QueryPendingSendToEthResponse](#gravity.v1.QueryPendingSendToEthResponse) |  | GET|/gravity/v1beta/query_pending_send_to_eth|

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

