package main

import (
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrCmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/dfinance/dstation/pkg/logger"

	"github.com/dfinance/dstation/cmd/dstation/cmd"
	"github.com/dfinance/dstation/cmd/dstation/config"
)

func main() {
	// build the root cmd
	rootCmd, _ := cmd.NewRootCmd()

	// configure Sentry integration and crash logging
	if err := logger.SetupSentry(version.Name, version.Version, version.Commit); err != nil {
		fmt.Printf("logger.SetupSentry: %v\n", err)
		os.Exit(1)
	}
	defer logger.CrashDeferHandler()

	// execute the root cmd
	if err := svrCmd.Execute(rootCmd, config.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)
		default:
			os.Exit(1)
		}
	}
}
