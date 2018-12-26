package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/gbbirkisson/rpi"
	proto "github.com/gbbirkisson/rpi/proto"
	"github.com/spf13/cobra"
)

var digitalCmd = &cobra.Command{
	Use:   "digital [pin] [value]",
	Short: "Write a digital value to a pin",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("command requires [pin] and [value] arguments")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var value int
		_, err := fmt.Sscanf(args[1], "%d", &value)

		if err != nil {
			fmt.Fprintf(os.Stderr, "'%s' is not a valid int: %v\n", args[1], err)
			os.Exit(1)
		}

		client, ctx := getGpioClientAndContext(cmd)
		_, err = client.DigitalWrite(ctx, &proto.DigitalWriteReq{Pin: args[0], Value: int32(value)})
		rpi.ExitOnError(err, "unable to write value gpio")
	},
}

func init() {
	gpioWriteCmd.AddCommand(digitalCmd)
}
