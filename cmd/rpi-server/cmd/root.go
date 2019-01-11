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
	RunE: func(_ *cobra.Command, args []string) error {
		log.Printf("rpi server started")
		ip := viper.GetString("ip")
		port := viper.GetString("port")
		address := ip + ":" + port

		srv := grpc.NewServer()

		proto.RegisterCommonServer(srv, &rpi.CommonServerImpl{})

		if viper.GetBool("gpio") {
			log.Printf("adding gpio service")
			proto.RegisterGpioServer(srv, &rpi.GpioServerImpl{})
		}

		if viper.GetBool("camera") {
			log.Printf("adding picam service\n")

			err := modprobe()
			rpi.ExitOnError("unable to modprobe", err)

			camargs := picamera.NewArgs()
			camargs.Width = viper.GetInt("camera_width")
			camargs.Height = viper.GetInt("camera_height")
			camargs.Rotation = viper.GetInt("camera_rotation")
			log.Printf("camera arguments: %+v\n", camargs)
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
	rootCmd.PersistentFlags().BoolP("camera", "c", false, "picam service enabled")
	rootCmd.PersistentFlags().String("modprobe", "", "modprobe on start (for pi camera)")
	rootCmd.PersistentFlags().Int("camera_width", 648, "Width of the image from pi camera")
	rootCmd.PersistentFlags().Int("camera_height", 486, "Height of the image from pi camera")
	rootCmd.PersistentFlags().Int("camera_rotation", 0, "Rotation of pi camera image")

	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("ip", rootCmd.PersistentFlags().Lookup("ip"))
	viper.BindPFlag("gpio", rootCmd.PersistentFlags().Lookup("gpio"))
	viper.BindPFlag("camera", rootCmd.PersistentFlags().Lookup("camera"))
	viper.BindPFlag("modprobe", rootCmd.PersistentFlags().Lookup("modprobe"))
	viper.BindPFlag("camera_width", rootCmd.PersistentFlags().Lookup("camera_width"))
	viper.BindPFlag("camera_height", rootCmd.PersistentFlags().Lookup("camera_height"))
	viper.BindPFlag("camera_rotation", rootCmd.PersistentFlags().Lookup("camera_rotation"))

	viper.SetEnvPrefix("rpi")
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
	param := viper.GetString("modprobe")
	if param != "" {
		log.Printf("running 'modprobe %s'\n", param)

		output, err := exec.Command("modprobe", param).CombinedOutput()
		if err != nil {
			return fmt.Errorf("%v: %s", err, output)
		}
	}
	return nil
}
