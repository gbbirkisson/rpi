package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gbbirkisson/rpi"
	helper "github.com/gbbirkisson/rpi/cmd"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func getConnection() *grpc.ClientConn {
	conn, err := rpi.NewGrpcClientConnectionInsecure(viper.GetString("server.host"), viper.GetString("server.port"))
	helper.ExitOnError("could not create connection", err)
	return conn
}

func getContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(viper.GetInt64("server.timeout"))*time.Millisecond)
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

	rootCmd.PersistentFlags().StringP("host", "s", "127.0.0.1", "server host address")
	rootCmd.PersistentFlags().IntP("port", "p", 8000, "server port")
	rootCmd.PersistentFlags().IntP("timeout", "t", 5000, "server timeout in milliseconds")

	viper.BindPFlag("server.host", rootCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag("server.port", rootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("server.timeout", rootCmd.PersistentFlags().Lookup("timeout"))

	helper.AddConfigCommand(rootCmd)
}

var configFileName = ".rpi-client"

func initConfig() {
	home, err := homedir.Dir()
	if err == nil {
		viper.AddConfigPath(home)
		viper.SetConfigName(configFileName)
		readErr := viper.ReadInConfig()
		if readErr != nil {
			fmt.Fprintf(os.Stderr, "unable to read config: %v\n", readErr)
		}
	} else {
		fmt.Fprintf(os.Stderr, "unable to find home dir: %v\n", err)
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix("rpi")
}
