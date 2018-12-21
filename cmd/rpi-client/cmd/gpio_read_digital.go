package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// digitalCmd represents the digital command
var gpioReadDigitalCmd = &cobra.Command{
	Use:   "digital",
	Short: "Read a digital value from a pin",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("digital called")
	},
}

func init() {
	gpioReadCmd.AddCommand(gpioReadDigitalCmd)
}
