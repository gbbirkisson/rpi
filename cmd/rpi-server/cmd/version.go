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
		ver, rev := rpi.GetVersion()
		fmt.Printf("rpi-server version %s %s\n", ver, rev)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
