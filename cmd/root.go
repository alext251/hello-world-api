package cmd

import (
	"fmt"
	"os"

	"github.com/alext251/hello-world-api/web"
	"github.com/spf13/cobra"
)

var (
	debug bool

	rootCmd = &cobra.Command{
		Use:   "hello-world",
		Short: "hello-world creates a simple Hello World API with a POST and GET request",
		Run: func(cmd *cobra.Command, args []string) {
			web.StartAPI(debug)
		},
	}
)

// Execute handles root command execution and
// error handling.
func Execute() {
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "toggle to enable debug logging")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
