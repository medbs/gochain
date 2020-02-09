package main

import (
	"fmt"
	"github.com/spf13/cobra"
	servecmd "gochain/cmd/server"
	versioncmd "gochain/cmd/version"
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
	rootCmd.AddCommand(servecmd.CommandServe)
	rootCmd.AddCommand(versioncmd.CommandVersion())
	return rootCmd
}


func main() {

	if err := commandRoot().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}

}
