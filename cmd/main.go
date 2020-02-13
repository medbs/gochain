package main

import (
	"fmt"
	"github.com/spf13/cobra"
	blockChainCmd "gochain/cmd/ledger"
	restCmd "gochain/cmd/rest"
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
	rootCmd.AddCommand(blockChainCmd.CommandServe)
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
