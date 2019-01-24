package cmd

import (
	"fmt"

	"github.com/gbbirkisson/rpi"
	helper "github.com/gbbirkisson/rpi/cmd"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := getContext()
		defer cancel()

		commonLocal := rpi.NewCommonLocal()
		localVer, localRev, _ := commonLocal.GetVersion(ctx)
		fmt.Printf("rpi-client version %s %s\n", localVer, localRev)

		commonRemote, err := rpi.NewCommonRemote(getConnection())
		helper.ExitOnError("could not create client", err)

		remoteVer, remoteRev, err := commonRemote.GetVersion(ctx)
		helper.ExitOnError("could not get server version", err)

		fmt.Printf("rpi-server version %s %s\n", remoteVer, remoteRev)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
