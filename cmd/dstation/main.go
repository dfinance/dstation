package main

import (
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrCmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/dfinance/dstation/cmd/dstation/cmd"
	"github.com/dfinance/dstation/cmd/dstation/config"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd()

	if err := svrCmd.Execute(rootCmd, config.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}
