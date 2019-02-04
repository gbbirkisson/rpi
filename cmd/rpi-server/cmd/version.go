package cmd

import (
	"context"

	"github.com/gbbirkisson/rpi"
	helper "github.com/gbbirkisson/rpi/cmd"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Run: func(cmd *cobra.Command, args []string) {
		common := rpi.NewCommonLocal()
		ver, rev, _ := common.GetVersion(context.Background())
		helper.PrintVersion("rpi-server", ver, rev)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
