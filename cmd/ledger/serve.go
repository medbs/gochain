package ledger

import (
	"fmt"
	"github.com/spf13/cobra"
	core "gochain/core"
	"log"
	"os"
	"strconv"
)

var CommandServe *cobra.Command

var address string
var p2pPort int
var target string
var secio bool
var seed int64
var httpPort int

func init() {
	CommandServe = &cobra.Command{
		Use:   "bc",
		Short: "Start the BlockChain",
		Long:  ``,
		Run: func(CommandServe *cobra.Command, args []string) {

			c := core.NewBlockChain(&core.P2pConfig{
				Address: address,
				Port: p2pPort,
				Target:  target,
				Secio:   secio,
				Seed:    seed,
			}, &core.HttpConfig{
				HttpPort: httpPort})

			log.Println("running http server")
			httpPort := strconv.Itoa(c.HttpConfig.HttpPort)
			go c.Run(":" + httpPort)

			_, err := serve(c)

			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}

		},
	}

	CommandServe.Flags().StringVar(&address, "address", "127.0.0.1", "address of the node")
	CommandServe.Flags().IntVar(&p2pPort, "ptpPort", 8199, "port: wait for incoming connections")
	CommandServe.Flags().StringVar(&target, "d", "", "target peer to dial")
	CommandServe.Flags().BoolVar(&secio, "secio", true, "enable secio")
	CommandServe.Flags().Int64Var(&seed, "seed", 0, "set random seed for id generation")
	CommandServe.Flags().IntVar(&httpPort, "httpPort", 8090, "port of the gin http server")

}

func serve(c *core.Chain) (*core.Chain, error) {

	log.Println("running p2p server")
	GlobalChain, err := c.Launch()

	if err != nil {
		return nil, err
	}

	return GlobalChain, nil

}
