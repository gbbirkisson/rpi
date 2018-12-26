package cmd

import (
	"fmt"

	"github.com/gbbirkisson/rpi"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Server version: %s", rpi.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
