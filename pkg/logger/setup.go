package logger

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/server"
	"github.com/spf13/cobra"
)

// OverrideCmdCtxLogger adds Sentry hook for server.Context ZeroLog logger.
func OverrideCmdCtxLogger(cmd *cobra.Command) error {
	if cmd == nil {
		return fmt.Errorf("cmd: nil")
	}

	serverCtx := server.GetServerContextFromCmd(cmd)
	if serverCtx == nil {
		return fmt.Errorf("server.GetServerContextFromCmd: nil")
	}

	zlWrapper, ok := serverCtx.Logger.(server.ZeroLogWrapper)
	if !ok {
		return fmt.Errorf("serverCtx.Logger type assert failed: %T", serverCtx.Logger)
	}
	zlWrapper.Logger = zlWrapper.Hook(NewZeroLogSentryHook())

	serverCtx.Logger = zlWrapper
	if err := server.SetCmdServerContext(cmd, serverCtx); err != nil {
		return fmt.Errorf("server.SetCmdServerContext: %w", err)
	}

	return nil
}
