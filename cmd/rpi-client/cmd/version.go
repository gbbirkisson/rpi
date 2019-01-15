package cmd

import (
	"fmt"

	"github.com/gbbirkisson/rpi"
	proto "github.com/gbbirkisson/rpi/proto"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("rpi-client version %s %s\n", rpi.Version, rpi.Revision)
		client, err := getCommonClient()
		if err != nil {
			rpi.ExitOnError("could not get server version", err)
		}
		ctx, cancel := getContext()
		defer cancel()

		res, err := client.Version(ctx, &proto.Void{})
		if err != nil {
			rpi.ExitOnError("could not get server version", err)
		}

		fmt.Printf("Server version: %s\n", res.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
