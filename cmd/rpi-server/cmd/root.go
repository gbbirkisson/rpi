package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gbbirkisson/rpi"
	helper "github.com/gbbirkisson/rpi/cmd"
	proto "github.com/gbbirkisson/rpi/pkg/proto"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "rpi-server",
	Short: "Raspberry PI IO server",
	Long:  `A gRPC server that allows you to do IO operations on the Raspberry PI`,
	Run: func(_ *cobra.Command, args []string) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		log.Printf("starting rpi server")
		srv, lis, err := rpi.NewGrpcServerInsecure(viper.GetString("host"), viper.GetString("port"))
		if err != nil {
			helper.ExitOnError("unable to create server", err)
		}

		common := rpi.NewCommonLocal()
		proto.RegisterCommonServer(srv, rpi.NewCommonServer(common))

		if viper.GetBool("gpio") {
			log.Printf("adding gpio service")
			gpio := rpi.NewGpioLocal()
			err := gpio.Open(ctx)
			helper.ExitOnError("unable to open gpio", err)
			defer gpio.Close(ctx)
			proto.RegisterGpioServer(srv, rpi.NewGpioServer(gpio))
		}

		modprobe := viper.GetStringSlice("modprobe")
		if len(modprobe) > 0 {
			log.Printf("running modprobe")
			for _, mod := range modprobe {
				err := common.Modprobe(ctx, mod)
				if err != nil {
					helper.ExitOnError(fmt.Sprintf("unable modprobe module '%s'", mod), err)
				}
			}
		}

		if viper.GetBool("picam") {
			log.Printf("adding picam service")
			camArgs := rpi.NewPiCamArgs()
			camArgs.Width = viper.GetInt("picam_width")
			camArgs.Height = viper.GetInt("picam_height")
			camArgs.Rotation = viper.GetInt("picam_rotation")
			cam, err := rpi.NewPiCamLocal(camArgs)

			helper.ExitOnError("unable to create camera", err)

			err = cam.Open(ctx)
			helper.ExitOnError("unable to open camera", err)
			defer cam.Close(ctx)

			proto.RegisterPiCamServer(srv, rpi.NewPicamServer(cam))
		}

		if viper.GetBool("ngrok") {
			log.Printf("adding ngrok service")
			ngrok, err := rpi.NewNgrokLocal("tcp", viper.GetString("port"), viper.GetString("ngrok_token"), viper.GetString("ngrok_region"))
			helper.ExitOnError("unable to setup ngrok", err)
			err = ngrok.Open(ctx)
			helper.ExitOnError("unable start ngrok", err)
			defer ngrok.Close(ctx)
		}

		log.Fatal(srv.Serve(lis))
	},
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

	rootCmd.PersistentFlags().String("host", "0.0.0.0", "server ip")
	rootCmd.PersistentFlags().Int("port", 8000, "server port")

	viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag("port", rootCmd.PersistentFlags().Lookup("port"))

	rootCmd.PersistentFlags().BoolP("gpio", "g", false, "gpio service enabled")
	viper.BindPFlag("gpio", rootCmd.PersistentFlags().Lookup("gpio"))

	rootCmd.PersistentFlags().StringSlice("modprobe", []string{}, "modprobe on start")
	viper.BindPFlag("modprobe", rootCmd.PersistentFlags().Lookup("modprobe"))

	rootCmd.PersistentFlags().BoolP("picam", "c", false, "picam service enabled")
	rootCmd.PersistentFlags().Int("picam_width", 648, "Width of the image from pi camera")
	rootCmd.PersistentFlags().Int("picam_height", 486, "Height of the image from pi camera")
	rootCmd.PersistentFlags().Int("picam_rotation", 0, "Rotation of pi camera image")
	viper.BindPFlag("picam", rootCmd.PersistentFlags().Lookup("picam"))
	viper.BindPFlag("picam_width", rootCmd.PersistentFlags().Lookup("picam_width"))
	viper.BindPFlag("picam_height", rootCmd.PersistentFlags().Lookup("picam_height"))
	viper.BindPFlag("picam_rotation", rootCmd.PersistentFlags().Lookup("picam_rotation"))

	rootCmd.PersistentFlags().Bool("ngrok", false, "Start a ngrok tunnel")
	rootCmd.PersistentFlags().String("ngrok_token", "", "Ngrok authentication token")
	rootCmd.PersistentFlags().String("ngrok_region", "eu", "Ngrok region")
	viper.BindPFlag("ngrok", rootCmd.PersistentFlags().Lookup("ngrok"))
	viper.BindPFlag("ngrok_token", rootCmd.PersistentFlags().Lookup("ngrok_token"))
	viper.BindPFlag("ngrok_region", rootCmd.PersistentFlags().Lookup("ngrok_region"))
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
