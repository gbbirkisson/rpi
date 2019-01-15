package cmd

import (
	"fmt"

	"github.com/gbbirkisson/rpi"
	helper "github.com/gbbirkisson/rpi/cmd"
	proto "github.com/gbbirkisson/rpi/pkg/proto"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("rpi-client version %s %s\n", rpi.Version, rpi.Revision)
		client, err := getCommonClient()
		if err != nil {
			helper.ExitOnError("could not get server version", err)
		}
		ctx, cancel := getContext()
		defer cancel()

		res, err := client.GetVersion(ctx, &proto.Void{})
		if err != nil {
			helper.ExitOnError("could not get server version", err)
		}

		fmt.Printf("rpi-server version %s %s\n", res.Version, res.Revision)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
