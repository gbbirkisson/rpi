package cmd

import (
	"errors"

	"github.com/gbbirkisson/rpi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configSetCmd = &cobra.Command{
	Use:   "set [key] [value]",
	Short: "Set config variable",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("command requires [key] and [value] arguments")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		viper.Set(args[0], args[1])
		err := viper.WriteConfig()
		rpi.ExitOnError(err, "could not write config file")
	},
}

func init() {
	configCmd.AddCommand(configSetCmd)
}
