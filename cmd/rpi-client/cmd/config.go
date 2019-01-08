package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Operations for configuring client",
}

var configSetCmd = &cobra.Command{
	Use:   "set [key] [value]",
	Short: "Set config variable",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("command requires [key] and [value] arguments")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		viper.Set(args[0], args[1])
		return viper.WriteConfig()
	},
}

var configPrintCmd = &cobra.Command{
	Use:   "print",
	Short: "Print configuration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Configuration file: %s\n", viper.ConfigFileUsed())
		fmt.Printf("Effective config:\n")
		for _, s := range [...]string{"ip", "port", "timeout"} {
			fmt.Printf("  %-10s %v\n", s, viper.Get(s))
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configSetCmd)
	configCmd.AddCommand(configPrintCmd)
}
