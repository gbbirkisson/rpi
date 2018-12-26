package cmd

import (
	"fmt"
	"log"
	"net"
	"os"

	proto "github.com/gbbirkisson/rpi/proto"
	rpi "github.com/gbbirkisson/rpi/server"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	grpc "google.golang.org/grpc"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "rpi-server",
	Short: "Raspberry PI IO server",
	Long:  `A gRPC server that allows you to do IO operations on the Raspberry PI`,
	RunE:  run,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().IntP("port", "p", 8000, "server port")
	rootCmd.Flags().StringP("ip", "i", "0.0.0.0", "server ip")
	rootCmd.Flags().BoolP("gpio", "g", false, "gpio service enabled")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".rpi-server" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".rpi-server")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func run(cmd *cobra.Command, args []string) error {
	log.Printf("rpi server started")
	ip := cmd.Flag("ip").Value.String()
	port := cmd.Flag("port").Value.String()
	address := ip + ":" + port

	srv := grpc.NewServer()

	proto.RegisterCommonServer(srv, &rpi.CommonServerImpl{})

	if cmd.Flag("gpio").Value.String() == "true" {
		log.Printf("adding gpio service")
		proto.RegisterGPIOServer(srv, rpi.GetGPIOServer())
	}

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Printf("failed starting server: %v\n", err)
		os.Exit(1)
	}

	log.Printf("listening to %s\n", address)
	log.Fatal(srv.Serve(lis))
	return nil
}
