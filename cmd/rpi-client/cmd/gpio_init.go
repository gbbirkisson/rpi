package cmd

import (
	"github.com/gbbirkisson/rpi"
	proto "github.com/gbbirkisson/rpi/proto"
	"github.com/spf13/cobra"
)

var gpioInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize GPIO interface",
	Run: func(cmd *cobra.Command, args []string) {
		client, ctx := getGpioClientAndContext(cmd)
		_, err := client.Init(ctx, &proto.Void{})
		rpi.ExitOnError(err, "unable to initialize gpio")
	},
}

func init() {
	gpioCmd.AddCommand(gpioInitCmd)
}
