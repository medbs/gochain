package rest

import (
	"fmt"
	"github.com/spf13/cobra"
	rest "gochain/rest"
	"os"
)

var CommandServe *cobra.Command
var port string

func init() {
	CommandServe = &cobra.Command{
		Use:   "hs",
		Short: "Start the http ledger",
		Long:  ``,
		Run: func(CommandServe *cobra.Command, args []string) {
			if err := serve(); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(2)
			}
		},
	}

	CommandServe.Flags().StringVar(&port, "p", ":8199", "port of the http ledger")
}

func serve() error {
	//s := rest.NewHttpServer(&port)
	r := rest.NewRouter()
	err := r.Run(port)
	if err != nil {
		return err
	}
	return nil
}
