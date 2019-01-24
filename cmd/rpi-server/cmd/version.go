package cmd

import (
	"context"
	"fmt"

	"github.com/gbbirkisson/rpi"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Run: func(cmd *cobra.Command, args []string) {
		common := rpi.NewCommonLocal()
		ver, rev, _ := common.GetVersion(context.Background())
		fmt.Printf("rpi-server version %s %s\n", ver, rev)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
