package cmd

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// func writeConfigYaml(path string, filename string) error {
// 	file := filepath.Join(path, filename+".yaml")

// 	err := os.MkdirAll(path, os.ModePerm)
// 	if err != nil {
// 		return fmt.Errorf("Failed creating configuration folder: %v", err)
// 	}
// 	viper.WriteConfig()
// 	err = viper.WriteConfigAs(file)
// 	if err != nil {
// 		return fmt.Errorf("Failed creating configuration file: %v", err)
// 	}
// 	return nil
// }

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configuration operations",
}

var configPrintCmd = &cobra.Command{
	Use:   "print",
	Short: "Print effective config",
	Run: func(cmd *cobra.Command, args []string) {
		PrintConfig()
	},
}

var configWriteCmd = &cobra.Command{
	Use:   "write",
	Short: "Write config to configuration file",
	Run: func(cmd *cobra.Command, args []string) {
		err := viper.WriteConfig()
		ExitOnError("unable to write config", err)
	},
}

var configSetCmd = &cobra.Command{
	Use:   "set [key] [value]",
	Short: "Set config variable and save configuration file",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("command requires [key] and [value] arguments")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		viper.WriteConfig()

		viper.Set(args[0], args[1])
		err := viper.WriteConfig()
		ExitOnError("unable to write config", err)
	},
}

// AddConfigCommand adds configuration commands that are mutual between server and client
func AddConfigCommand(root *cobra.Command) {
	configWriteCmd.Flags().BoolP("force", "f", false, "Force writing configuration file event though it exists")

	configCmd.AddCommand(configPrintCmd)
	configCmd.AddCommand(configWriteCmd)
	configCmd.AddCommand(configSetCmd)
	root.AddCommand(configCmd)
}
