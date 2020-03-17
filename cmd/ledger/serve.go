package ledger

import (
	"fmt"
	"github.com/spf13/cobra"
	core "gochain/core"
	"log"
	"os"
)

var CommandServe *cobra.Command

var p2pPort int
var target string
var secio bool
var seed int64
var httpPort string

func init() {
	CommandServe = &cobra.Command{
		Use:   "bc",
		Short: "Start the BlockChain",
		Long:  ``,
		Run: func(CommandServe *cobra.Command, args []string) {

			if err := serve(); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		},
	}

	CommandServe.Flags().IntVar(&p2pPort, "l", 8199, "port: wait for incoming connections")
	CommandServe.Flags().StringVar(&target, "d", "", "target peer to dial")
	CommandServe.Flags().BoolVar(&secio, "secio", true, "enable secio")
	CommandServe.Flags().Int64Var(&seed, "seed", 0, "set random seed for id generation")
	CommandServe.Flags().StringVar(&httpPort, "p", ":8090", "port of the http ledger")
}

func serve() error {

	c := core.NewBlockChain(&core.P2pConfig{
		ListenF: p2pPort,
		Target:  target,
		Secio:   secio,
		Seed:    seed,
	}, &core.HttpConfig{
		HttpPort: httpPort})

	log.Println("running p2p server")
	nc, err := core.Launch(c)
	if err != nil {
		return err
	}

	log.Println("running http server")
	err = nc.Run(httpPort)
	if err != nil {
		return err
	}
	return nil

}
