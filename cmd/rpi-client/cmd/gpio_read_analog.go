package cmd

import (
	"errors"
	"fmt"

	"github.com/gbbirkisson/rpi"
	proto "github.com/gbbirkisson/rpi/proto"
	"github.com/spf13/cobra"
)

var gpioReadAnalogCmd = &cobra.Command{
	Use:   "analog [pin]",
	Short: "Read a analog value from a pin",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("command requires [pin] argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		client, ctx := getGpioClientAndContext(cmd)
		res, err := client.AnalogRead(ctx, &proto.PinReq{Pin: args[0]})
		rpi.ExitOnError(err, "unable to read value from gpio")
		fmt.Println(res.Value)
	},
}

func init() {
	gpioReadCmd.AddCommand(gpioReadAnalogCmd)
}
