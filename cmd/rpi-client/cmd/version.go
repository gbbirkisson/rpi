package cmd

import (
	"fmt"

	"github.com/gbbirkisson/rpi"
	helper "github.com/gbbirkisson/rpi/cmd"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := getContext()
		defer cancel()

		ver, rev := rpi.GetLocalVersion()

		fmt.Printf("rpi-client version %s %s\n", ver, rev)

		conn, err := rpi.GrpcClientConnectionInsecure(viper.GetString("host"), viper.GetString("port"))
		helper.ExitOnError("could not create client", err)

		common := rpi.Common{Connection: conn}

		ver, rev, err = common.GetVersion(ctx)
		helper.ExitOnError("could not get server version", err)

		fmt.Printf("rpi-server version %s %s\n", ver, rev)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
