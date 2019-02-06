package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gbbirkisson/rpi"
	helper "github.com/gbbirkisson/rpi/cmd"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func getConnection() *grpc.ClientConn {
	conn, err := rpi.NewGrpcClientConnectionInsecure(viper.GetString("host"), viper.GetString("port"))
	helper.ExitOnError("could not create connection", err)
	return conn
}

func getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(viper.GetInt64("timeout"))*time.Millisecond)
}

var rootCmd = &cobra.Command{
	Use:   "rpi-client",
	Short: "A client to run commands on the rpi-server",
}

// Execute the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	viper.SetEnvPrefix("rpi")

	rootCmd.PersistentFlags().StringP("host", "s", "127.0.0.1", "server host address")
	rootCmd.PersistentFlags().IntP("port", "p", 8000, "server port")
	rootCmd.PersistentFlags().IntP("timeout", "t", 5000, "server timeout in milliseconds")

	viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("timeout", rootCmd.PersistentFlags().Lookup("timeout"))

	helper.AddConfigCommand(rootCmd)
}

var configFileName = ".rpi-client"

func initConfig() {
	home, err := homedir.Dir()
	helper.ExitOnError("unable to find home directory", err)

	viper.AddConfigPath(home)
	viper.SetConfigName(configFileName)

	viper.ReadInConfig()
	viper.AutomaticEnv()
}
