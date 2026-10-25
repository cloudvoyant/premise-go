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
		RunE: func(cmd *cobra.Command, args []string) error {
			_, err := fmt.Fprintln(cmd.OutOrStdout(), "Hello from premise-cobra-cli")
			return err
		},
	}
}

func main() {
	if err := command().Execute(); err != nil {
		os.Exit(1)
	}
}
