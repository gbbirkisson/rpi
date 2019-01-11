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
	picamera "github.com/technomancers/piCamera"
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
			m, err := cmd.Flags().GetBool("cmod")
			if err != nil {
				rpi.ExitOnError("cmod flag invalid", err)
			}
			if m {
				err := modprobe()
				rpi.ExitOnError("unable to modprobe", err)
			}

			width, err := cmd.Flags().GetInt("cwidth")
			if err != nil {
				rpi.ExitOnError("cwidth flag invalid", err)
			}

			height, err := cmd.Flags().GetInt("cheight")
			if err != nil {
				rpi.ExitOnError("cheight flag invalid", err)
			}

			rot, err := cmd.Flags().GetInt("crotation")
			if err != nil {
				rpi.ExitOnError("crotation flag invalid", err)
			}

			camargs := picamera.NewArgs()
			camargs.Width = width
			camargs.Height = height
			camargs.Rotation = rot
			cam, err := picamera.New(nil, camargs)
			rpi.ExitOnError("unable to create camera", err)

			err = cam.Start()
			rpi.ExitOnError("unable to start camera", err)

			defer cam.Stop()

			proto.RegisterPiCamServer(srv, &rpi.PiCamServerImpl{Camera: cam})
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
	rootCmd.PersistentFlags().Bool("cmod", false, "modprobe on start (for pi camera)")
	rootCmd.PersistentFlags().Int("cwidth", 648, "Width of the image from picam")
	rootCmd.PersistentFlags().Int("cheight", 486, "Height of the image from picam")
	rootCmd.PersistentFlags().Int("crotation", 0, "Rotation of camera image")

	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("ip", rootCmd.PersistentFlags().Lookup("ip"))
	viper.BindPFlag("gpio", rootCmd.PersistentFlags().Lookup("gpio"))
	viper.BindPFlag("picam", rootCmd.PersistentFlags().Lookup("picam"))
	viper.BindPFlag("cmod", rootCmd.PersistentFlags().Lookup("cmod"))
	viper.BindPFlag("cwidth", rootCmd.PersistentFlags().Lookup("cwidth"))
	viper.BindPFlag("cheight", rootCmd.PersistentFlags().Lookup("cheight"))
	viper.BindPFlag("crotation", rootCmd.PersistentFlags().Lookup("crotation"))
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
	output, err := exec.Command("modprobe", os.Getenv("MODPROBE")).CombinedOutput()
	if err != nil {
		return fmt.Errorf("%v: %s", err, output)
	}
	return nil
}
