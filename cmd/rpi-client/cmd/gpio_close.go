package cmd

import (
	"github.com/gbbirkisson/rpi"
	proto "github.com/gbbirkisson/rpi/proto"
	"github.com/spf13/cobra"
)

var gpioCloseCmd = &cobra.Command{
	Use:   "close",
	Short: "Close GPIO interface",
	Run: func(cmd *cobra.Command, args []string) {
		client, ctx := getGpioClientAndContext(cmd)
		_, err := client.Close(ctx, &proto.Void{})
		rpi.ExitOnError(err, "unable to close gpio")
	},
}

func init() {
	gpioCmd.AddCommand(gpioCloseCmd)
}
