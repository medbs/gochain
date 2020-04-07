package rest

import (
	"fmt"
	"github.com/spf13/cobra"
	"gochain/cmd/ledger"
	"log"
	"os"
)

var CommandServe *cobra.Command

func init() {
	CommandServe = &cobra.Command{
		Use:   "hs",
		Short: "Start the http server",
		Long:  ``,
		Run: func(CommandServe *cobra.Command, args []string) {
			if err := serve(); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		},
	}

}

func serve() error {

	log.Println("running http server")
	err := ledger.GlobalChain.Run(":8090")
	if err != nil {
		return err
	}
	return nil

	return nil

}
