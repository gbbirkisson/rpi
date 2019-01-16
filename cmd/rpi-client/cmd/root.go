package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/gbbirkisson/rpi"
	proto "github.com/gbbirkisson/rpi/pkg/proto"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

func getCommonClient() (proto.CommonServiceClient, error) {
	conn, err := rpi.GrpcClientConnectionInsecure(viper.GetString("host"), viper.GetString("port"))
	if err != nil {
		return nil, err
	}
	return proto.NewCommonServiceClient(conn), nil
}

func getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(viper.GetInt64("timeout"))*time.Millisecond)
}

var rootCmd = &cobra.Command{
	Use:   "rpi-client",
	Short: "A client to run commands on the rpi-server",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().String("host", "127.0.0.1", "server ip")
	rootCmd.PersistentFlags().Int("port", 8000, "server port")
	rootCmd.PersistentFlags().Int("timeout", 3000, "server timeout in milliseconds")

	viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("timeout", rootCmd.PersistentFlags().Lookup("timeout"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Search config in home directory with name ".rpi-client" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigName(".rpi-client")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		file := filepath.Join(home, ".rpi-client.yaml")
		fmt.Fprintf(os.Stderr, "Config file not found, creating it: %s\n", file)
		viper.WriteConfigAs(file)
	}
}
