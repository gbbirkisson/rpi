package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var gpioWriteAnalogCmd = &cobra.Command{
	Use:   "analog",
	Short: "Write a analog value to a pin",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("analog called")
	},
}

func init() {
	gpioWriteCmd.AddCommand(gpioWriteAnalogCmd)
}
