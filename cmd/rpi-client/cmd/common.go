package cmd

import (
	"errors"

	"github.com/gbbirkisson/rpi"
	helper "github.com/gbbirkisson/rpi/cmd"
	"github.com/spf13/cobra"
)

func getCommon() rpi.Common {
	return rpi.NewCommonRemote(getConnection())
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := getContext()
		defer cancel()

		commonLocal := rpi.NewCommonLocal()
		localVer, localRev, _ := commonLocal.GetVersion(ctx)
		helper.PrintVersion("rpi-client", localVer, localRev)

		commonRemote := getCommon()

		remoteVer, remoteRev, err := commonRemote.GetVersion(ctx)
		helper.ExitOnError("could not get server version", err)

		helper.PrintVersion("rpi-server", remoteVer, remoteRev)
	},
}

var modprobeCmd = &cobra.Command{
	Use:   "modprobe [module]",
	Short: "Run modbprobe command on server",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("command requires [module] argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := getContext()
		defer cancel()

		commonRemote := getCommon()
		err := commonRemote.Modprobe(ctx, args[0])
		helper.ExitOnError("unable to modprobe", err)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd, modprobeCmd)
}
