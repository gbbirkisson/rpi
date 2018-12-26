package cmd

import (
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Operations for configuring client",
}

func init() {
	rootCmd.AddCommand(configCmd)
}
