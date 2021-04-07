package keeper

import (
	"context"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	dvmTypes "github.com/dfinance/dstation/pkg/types/dvm"
)

// VMExecRetryReq contains VM "execution" request meta (request details and retry settings).
type VMExecRetryReq struct {
	// Request to retry (module publish).
	RawModule *dvmTypes.VMPublishModule
	// Request to retry (script execution)
	RawScript *dvmTypes.VMExecuteScript
	// Max number of request attempts (0 - infinite)
	MaxAttempts uint
	// Request timeout per attempt (0 - infinite) [ms]
	ReqTimeoutInMs uint
}

// sendExecuteReq sends request with retry mechanism.
func (k Keeper) sendExecuteReq(ctx sdk.Context, moduleReq *dvmTypes.VMPublishModule, scriptReq *dvmTypes.VMExecuteScript) (*dvmTypes.VMExecuteResponse, error) {
	if moduleReq == nil && scriptReq == nil {
		return nil, fmt.Errorf("request (module / script) not specified")
	}
	if moduleReq != nil && scriptReq != nil {
		return nil, fmt.Errorf("only single request (module / script) is supported")
	}

	retryReq := VMExecRetryReq{
		RawModule:      moduleReq,
		RawScript:      scriptReq,
		MaxAttempts:    k.config.MaxAttempts,
		ReqTimeoutInMs: k.config.ReqTimeoutInMs,
	}

	return k.retryExecReq(ctx, retryReq)
}

// retryExecReq sends request with retry mechanism and waits for connection and execution.
// Contract: either RawModule or RawScript must be specified for RetryExecReq.
func (k Keeper) retryExecReq(ctx sdk.Context, req VMExecRetryReq) (retResp *dvmTypes.VMExecuteResponse, retErr error) {
	const failedRetryLogPeriod = 100

	doneCh := make(chan bool)
	curAttempt := uint(0)
	reqTimeout := time.Duration(req.ReqTimeoutInMs) * time.Millisecond
	reqStartedAt := time.Now()

	go func() {
		defer func() {
			close(doneCh)
		}()

		for {
			var connCtx context.Context
			var connCancel context.CancelFunc
			var resp *dvmTypes.VMExecuteResponse
			var err error

			curAttempt++

			connCtx = context.Background()
			if reqTimeout > 0 {
				connCtx, connCancel = context.WithTimeout(context.Background(), reqTimeout)
			}

			curReqStartedAt := time.Now()
			if req.RawModule != nil {
				resp, err = k.vmClient.VMModulePublisherClient.PublishModule(connCtx, req.RawModule)
			} else if req.RawScript != nil {
				resp, err = k.vmClient.VMScriptExecutorClient.ExecuteScript(connCtx, req.RawScript)
			}
			if connCancel != nil {
				connCancel()
			}
			curReqDur := time.Since(curReqStartedAt)

			if err == nil {
				retResp, retErr = resp, nil
				return
			}

			if req.MaxAttempts != 0 && curAttempt == req.MaxAttempts {
				retResp, retErr = nil, err
				return
			}

			if curReqDur < reqTimeout {
				time.Sleep(reqTimeout - curReqDur)
			}

			if curAttempt%failedRetryLogPeriod == 0 {
				msg := fmt.Sprintf("Failing VM request: attempt %d / %d with %v timeout: %v", curAttempt, req.MaxAttempts, reqTimeout, time.Since(reqStartedAt))
				k.Logger(ctx).Info(msg)
			}
		}
	}()
	<-doneCh

	reqDur := time.Since(reqStartedAt)
	msg := fmt.Sprintf("in %d attempt(s) with %v timeout (%v)", curAttempt, reqTimeout, reqDur)
	if retErr == nil {
		k.Logger(ctx).Info(fmt.Sprintf("Successfull VM request (%s)", msg))
	} else {
		k.Logger(ctx).Error(fmt.Sprintf("Failed VM request (%s): %v", msg, retErr))
		retErr = fmt.Errorf("%s: %w", msg, retErr)
	}

	return
}
