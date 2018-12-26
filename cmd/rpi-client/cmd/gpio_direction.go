package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/gbbirkisson/rpi"
	proto "github.com/gbbirkisson/rpi/proto"
	"github.com/spf13/cobra"
)

var gpioDirectionCmd = &cobra.Command{
	Use:   "direction [pin] [direction]",
	Short: "Set direction of GPIO pins",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("command requires [pin] and [direction] arguments")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

		var dir int
		_, err := fmt.Sscanf(args[1], "%d", &dir)

		if err != nil {
			fmt.Fprintf(os.Stderr, "'%s' is not a valid int: %v\n", args[1], err)
			os.Exit(1)
		}

		client, ctx := getGpioClientAndContext(cmd)
		_, err = client.SetDirection(ctx, &proto.SetDirectionReq{Pin: args[0], Direction: int32(dir)})
		rpi.ExitOnError(err, "unable to set direction")
	},
}

func init() {
	gpioCmd.AddCommand(gpioDirectionCmd)
}
