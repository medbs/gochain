package rest

import (
	"fmt"
	"github.com/spf13/cobra"
	"gochain/core"
	"log"
	"os"
)

var CommandServe *cobra.Command
var GlobalChain *core.Chain

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

	//log.Println("running http server")
	//go func() error {
	//	//c := <-passedChain
	//	err := ledger.GlobalChain.Run(":8090")
	//	if err != nil {
	//		return err
	//	}
	//	return nil
	//}()
	//
	//return nil

	log.Println("running http server")
	//c := <-passedChain
	err := GlobalChain.Run(":8090")
	if err != nil {
		return err
	}
	return nil

	return nil

}
