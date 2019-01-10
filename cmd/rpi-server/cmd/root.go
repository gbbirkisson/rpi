package cmd

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gbbirkisson/rpi"
	proto "github.com/gbbirkisson/rpi/proto"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "rpi-server",
	Short: "Raspberry PI IO server",
	Long:  `A gRPC server that allows you to do IO operations on the Raspberry PI`,
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Printf("rpi server started")
		ip := viper.GetString("ip")
		port := viper.GetString("port")
		address := ip + ":" + port

		srv := grpc.NewServer()

		proto.RegisterCommonServer(srv, &rpi.CommonServerImpl{})

		if cmd.Flag("gpio").Value.String() == "true" {
			log.Printf("adding gpio service")
			proto.RegisterGpioServer(srv, &rpi.GpioServerImpl{})
		}

		if cmd.Flag("picam").Value.String() == "true" {
			log.Printf("adding picam service\n")
			err := modprobe()
			if err != nil {
				log.Printf("unable to modprobe: %v\n", err)
				log.Printf("not startin picam service\n")
			} else {
				log.Printf("successfully ran modprobe\n")
				proto.RegisterPiCamServer(srv, &rpi.PiCamServerImpl{})
			}
		}

		lis, err := net.Listen("tcp", address)
		if err != nil {
			log.Printf("failed starting server: %v\n", err)
			os.Exit(1)
		}

		log.Printf("listening to %s\n", address)
		log.Fatal(srv.Serve(lis))
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().IntP("port", "p", 8000, "server port")
	rootCmd.PersistentFlags().StringP("ip", "i", "0.0.0.0", "server ip")
	rootCmd.PersistentFlags().BoolP("gpio", "g", false, "gpio service enabled")
	rootCmd.PersistentFlags().BoolP("picam", "c", false, "picam service enabled")

	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("ip", rootCmd.PersistentFlags().Lookup("ip"))
	viper.BindPFlag("gpio", rootCmd.PersistentFlags().Lookup("gpio"))
	viper.BindPFlag("picam", rootCmd.PersistentFlags().Lookup("picam"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	configPath := "/etc/rpi-server"
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		file := filepath.Join(configPath, "config.yaml")

		fmt.Fprintf(os.Stderr, "Config file not found, creating it: %s\n", file)

		err = os.MkdirAll(configPath, os.ModePerm)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed creating configuration folder: %v\n", err)
			return
		}

		err = viper.WriteConfigAs(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed creating configuration file: %v\n", err)
			return
		}
	}
}

func modprobe() error {
	_, err := exec.Command("modprobe", os.Getenv("MODPROBE")).CombinedOutput()
	return err
}
