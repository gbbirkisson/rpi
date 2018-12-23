package cmd

import (
	"fmt"
	"os"

	proto "github.com/gbbirkisson/rpi/proto"
	"github.com/spf13/cobra"
)

var gpioInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get info about GPIO constants",
	Run: func(cmd *cobra.Command, args []string) {
		client, ctx := getGpioClientAndContext(cmd)
		info, err := client.Info(ctx, &proto.Void{})
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to get gpio info: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Name\tConstant\tDescription\n")
		fmt.Printf("In\t%d\t\tFor setting pin direction to in\n", info.In)
		fmt.Printf("In\t%d\t\tFor setting pin direction to out\n", info.Out)
		fmt.Printf("In\t%d\t\tFor setting pin direction to high\n", info.High)
		fmt.Printf("In\t%d\t\tFor setting pin direction to low\n", info.Low)
	},
}

func init() {
	gpioCmd.AddCommand(gpioInfoCmd)
}
