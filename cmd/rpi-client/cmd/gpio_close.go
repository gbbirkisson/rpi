package cmd

import (
	"fmt"
	"os"

	proto "github.com/gbbirkisson/rpi/proto"
	"github.com/spf13/cobra"
)

var gpioCloseCmd = &cobra.Command{
	Use:   "close",
	Short: "Close GPIO interface",
	Run: func(cmd *cobra.Command, args []string) {
		client, ctx := getGpioClientAndContext(cmd)
		_, err := client.Close(ctx, &proto.Void{})
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to close gpio: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	gpioCmd.AddCommand(gpioCloseCmd)
}
