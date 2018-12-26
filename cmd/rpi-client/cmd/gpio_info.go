package cmd

import (
	"fmt"
	"strings"

	"github.com/gbbirkisson/rpi"
	proto "github.com/gbbirkisson/rpi/proto"
	"github.com/spf13/cobra"
)

var gpioInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get info about GPIO pins",
	Run: func(cmd *cobra.Command, args []string) {
		client, ctx := getGpioClientAndContext(cmd)
		info, err := client.Info(ctx, &proto.Void{})
		rpi.ExitOnError(err, "unable to get gpio info")

		fmt.Printf("Name\tConstant\tDescription\n")
		fmt.Printf("In\t%d\t\tSet pin direction to in\n", info.In)
		fmt.Printf("Out\t%d\t\tSet pin direction to out\n", info.Out)
		fmt.Printf("Low\t%d\t\tSet pin output to low\n", info.Low)
		fmt.Printf("High\t%d\t\tSet pin output to high\n", info.High)

		fmt.Println()
		fmt.Println("ID\tAliases\t\t\t\tCaps\tDigitalLogical\tAnalogLogical")
		for _, pin := range info.Info {
			fmt.Printf("%-7s %-31s %d\t%d\t\t%d\n", pin.Id, strings.Join(pin.Alias, ", "), pin.Caps, pin.DigitalLogical, pin.AnalogLogical)
		}
	},
}

func init() {
	gpioCmd.AddCommand(gpioInfoCmd)
}
