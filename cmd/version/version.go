package version

import (
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

func CommandVersion() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version and exit",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf(`go-chain: %s
API Version: %s
Go Version: %s
Go OS/ARCH: %s %s
`, "1.0.0", "1.0.0", runtime.Version(), runtime.GOOS, runtime.GOARCH)
		},
	}
}