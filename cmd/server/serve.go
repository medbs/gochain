package server

import (
	"fmt"
	"github.com/spf13/cobra"
	ledger "gochain/core"
	"os"
)

var CommandServe *cobra.Command

var port int
var target string
var secio bool
var seed int64

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

	CommandServe.Flags().IntVar(&port, "l", 8199, "port: wait for incoming connections")
	CommandServe.Flags().StringVar(&target, "d", "", "target peer to dial")
	CommandServe.Flags().BoolVar(&secio, "secio", true, "enable secio")
	CommandServe.Flags().Int64Var(&seed, "seed", 0, "set random seed for id generation")
}

func serve(cmd *cobra.Command, args []string) error {

	c:= ledger.NewChain(&ledger.P2pConfig{
			ListenF: port,
			Target: target,
			Secio: secio,
			Seed: seed,
		})

	ledger.Launch(c)
	return nil
}
