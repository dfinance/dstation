package keeper_test

import (
	"context"
	"testing"
	"time"

	"github.com/dfinance/dvm-proto/go/ds_grpc"

	"github.com/dfinance/dstation/pkg/tests"
)

func DoDSClientRequest(t *testing.T, app *tests.DSimApp, handler func(ctx context.Context, client ds_grpc.DSServiceClient)) {
	client := ds_grpc.NewDSServiceClient(app.MockVMServer.GetDSClientConnection())
	ctx, ctxCancel := context.WithDeadline(context.Background(), time.Now().Add(100*time.Millisecond))
	defer ctxCancel()

	handler(ctx, client)
}
