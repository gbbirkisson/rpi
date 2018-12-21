package cmd

import (
	"fmt"
	"os"
	"strings"

	proto "github.com/gbbirkisson/rpi/proto"
	"github.com/spf13/cobra"
)

var gpioPinsCmd = &cobra.Command{
	Use:   "pins",
	Short: "List the available pins on the device",
	Run: func(cmd *cobra.Command, args []string) {
		client, ctx := getGpioClientAndContext(cmd)
		pinRes, err := client.Pins(ctx, &proto.Void{})
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to get pins: %v", err)
			os.Exit(1)
		}
		fmt.Println("ID\tAliases\t\t\t\tCaps\tDigitalLogical\tAnalogLogical")
		for _, pin := range pinRes.Info {
			fmt.Printf("%-7s %-31s %d\t%d\t\t%d\n", pin.Id, strings.Join(pin.Alias, ", "), pin.Caps, pin.DigitalLogical, pin.AnalogLogical)
		}
	},
}

func init() {
	gpioCmd.AddCommand(gpioPinsCmd)
}
