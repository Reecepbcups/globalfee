package main

import (
	"fmt"
	"os"

	"github.com/reecepbcups/globalfee/app"
	"github.com/reecepbcups/globalfee/cmd/globald/cmd"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
)

func main() {
	rootCmd := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, "globalfee", app.DefaultNodeHome); err != nil {
		fmt.Fprintln(rootCmd.OutOrStderr(), err)
		os.Exit(1)
	}
}
