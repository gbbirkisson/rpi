package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var digitalCmd = &cobra.Command{
	Use:   "digital",
	Short: "Write a digital value to a pin",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("digital called")
	},
}

func init() {
	gpioWriteCmd.AddCommand(digitalCmd)
}
