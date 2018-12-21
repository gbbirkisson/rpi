package cmd

import (
	"github.com/spf13/cobra"
)

var gpioReadCmd = &cobra.Command{
	Use:   "read",
	Short: "Read a value from a pin",
}

func init() {
	gpioCmd.AddCommand(gpioReadCmd)
}
