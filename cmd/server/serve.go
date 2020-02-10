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
			if err := serve(); err != nil {
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

func serve() error {

	c:= ledger.NewBlockChain(&ledger.P2pConfig{
			ListenF: port,
			Target: target,
			Secio: secio,
			Seed: seed,
		})

	err := ledger.Launch(c)
	if err != nil {
		return err
	}

	return nil
}
