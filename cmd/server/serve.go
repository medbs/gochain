package server

import (
	"fmt"
	"github.com/spf13/cobra"
	ledger "gochain/core"
	"os"
)

var CommandServe *cobra.Command

var port int

func init() {
	CommandServe = &cobra.Command{
		Use:   "serve",
		Short: "Start the app",
		Long:  ``,
		Run: func(CommandServe *cobra.Command, args []string) {
			if err := serve(CommandServe, args); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(2)
			}
		},
	}

	CommandServe.Flags().IntVar(&port, "port", 8099, "port port of the app")
}

func serve(cmd *cobra.Command, args []string) error {

	//var chain ledger.Chain
	//ledger.Launch(&chain)
	c:= ledger.NewChain(&ledger.P2pConfig{
			ListenF: 8987,
			Target: "",
			Secio: true,
			Seed: 0,
		})

	ledger.Launch(c)
	return nil
}
