package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "dev"

// Private helpers ------------------------------------------------------------

func command() *cobra.Command {
	return &cobra.Command{
		Use:     "premise-cobra-cli",
		Version: version,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(cmd.OutOrStdout(), "Hello from premise-cobra-cli")
		},
	}
}

func main() {
	if err := command().Execute(); err != nil {
		os.Exit(1)
	}
}
