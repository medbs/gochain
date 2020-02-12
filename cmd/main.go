package main

import (
	"fmt"
	"github.com/spf13/cobra"
	restCmd "gochain/cmd/rest"
	serveCmd "gochain/cmd/server"
	versionCmd "gochain/cmd/version"
	"os"
)

var verbose bool

func commandRoot() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "go-chain",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
			os.Exit(2)
		},
	}

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.AddCommand(serveCmd.CommandServe)
	rootCmd.AddCommand(versionCmd.CommandVersion())
	rootCmd.AddCommand(restCmd.CommandServe)
	return rootCmd
}

func main() {

	if err := commandRoot().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}

}
