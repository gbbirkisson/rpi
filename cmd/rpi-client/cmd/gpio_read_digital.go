package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var gpioReadDigitalCmd = &cobra.Command{
	Use:   "digital [pin]",
	Short: "Read a digital value from a pin",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("command requires [pin] argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("digital called")
	},
}

func init() {
	gpioReadCmd.AddCommand(gpioReadDigitalCmd)
}
