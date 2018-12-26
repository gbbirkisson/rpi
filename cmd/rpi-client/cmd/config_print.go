package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configPrintCmd = &cobra.Command{
	Use:   "print",
	Short: "Print configuration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Configuration file: %s\n", viper.ConfigFileUsed())
		fmt.Printf("Effective config:\n")
		for _, s := range [...]string{"ip", "port"} {
			fmt.Printf("  %-10s %v\n", s, viper.Get(s))
		}
	},
}

func init() {
	configCmd.AddCommand(configPrintCmd)
}
