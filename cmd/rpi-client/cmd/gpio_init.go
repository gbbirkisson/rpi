package cmd

import (
	"fmt"
	"os"

	proto "github.com/gbbirkisson/rpi/proto"
	"github.com/spf13/cobra"
)

var gpioInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize GPIO interface",
	Run: func(cmd *cobra.Command, args []string) {
		client, ctx := getGpioClientAndContext(cmd)
		_, err := client.Init(ctx, &proto.Void{})
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to initialize gpio: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	gpioCmd.AddCommand(gpioInitCmd)
}
