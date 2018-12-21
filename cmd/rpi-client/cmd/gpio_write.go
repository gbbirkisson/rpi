package cmd

import (
	"github.com/spf13/cobra"
)

var gpioWriteCmd = &cobra.Command{
	Use:   "write",
	Short: "Write a value to a pin",
}

func init() {
	gpioCmd.AddCommand(gpioWriteCmd)
}
