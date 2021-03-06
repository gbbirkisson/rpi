package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

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

		log.Println("starting rpi server")
		srv, lis, err := rpi.NewGrpcServerInsecure(viper.GetString("server.host"), viper.GetString("server.port"))
		if err != nil {
			helper.ExitOnError("unable to create server", err)
		}

		common := rpi.NewCommonLocal()
		proto.RegisterCommonServer(srv, rpi.NewCommonServer(common))

		if viper.GetBool("gpio.enabled") {
			log.Println("adding gpio service")
			gpio := rpi.NewGpioLocal()
			if viper.GetBool("gpio.open") {
				log.Println("opening gpio interface")
				err := gpio.Open(ctx)
				helper.ExitOnError("unable to open gpio", err)
			}
			defer gpio.Close(ctx)
			proto.RegisterGpioServer(srv, rpi.NewGpioServer(gpio))
		}

		modprobe := viper.GetStringSlice("modprobe")
		if len(modprobe) > 0 {
			log.Printf("running modprobe for %s\n", modprobe)
			for _, mod := range modprobe {
				err := common.Modprobe(ctx, mod)
				if err != nil {
					helper.ExitOnError(fmt.Sprintf("unable modprobe module '%s'", mod), err)
				}
			}
		}

		if viper.GetBool("picam.enabled") {
			log.Println("adding picam service")
			camArgs := rpi.NewPiCamArgs()
			camArgs.Width = viper.GetInt("picam.width")
			camArgs.Height = viper.GetInt("picam.height")
			camArgs.Rotation = viper.GetInt("picam.rotation")
			cam, err := rpi.NewPiCamLocal(camArgs)

			helper.ExitOnError("unable to create camera", err)

			if viper.GetBool("picam.open") {
				log.Println("opening picam interface")
				err = cam.Open(ctx)
				helper.ExitOnError("unable to open camera", err)
			}
			defer cam.Close(ctx)

			proto.RegisterPiCamServer(srv, rpi.NewPicamServer(cam))
		}

		if viper.GetBool("ngrok.enabled") {
			log.Println("adding ngrok service")
			ngrok, err := rpi.NewNgrokLocal("tcp", viper.GetString("server.port"), viper.GetString("ngrok.token"), viper.GetString("ngrok.region"))
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

	rootCmd.PersistentFlags().BoolP("gpio", "g", false, "gpio service enabled")
	rootCmd.PersistentFlags().Bool("gpio_open", false, "open gpio service on start")
	viper.BindPFlag("gpio.enabled", rootCmd.PersistentFlags().Lookup("gpio"))
	viper.BindPFlag("gpio.open", rootCmd.PersistentFlags().Lookup("gpio_open"))

	rootCmd.PersistentFlags().StringSlice("modprobe", []string{}, "modprobe on start")
	viper.BindPFlag("modprobe", rootCmd.PersistentFlags().Lookup("modprobe"))

	rootCmd.PersistentFlags().Bool("ngrok", false, "Start a ngrok tunnel")
	rootCmd.PersistentFlags().String("ngrok_token", "", "Ngrok authentication token")
	rootCmd.PersistentFlags().String("ngrok_region", "eu", "Ngrok region")
	viper.BindPFlag("ngrok.enabled", rootCmd.PersistentFlags().Lookup("ngrok"))
	viper.BindPFlag("ngrok.token", rootCmd.PersistentFlags().Lookup("ngrok_token"))
	viper.BindPFlag("ngrok.region", rootCmd.PersistentFlags().Lookup("ngrok_region"))

	rootCmd.PersistentFlags().BoolP("picam", "c", false, "picam service enabled")
	rootCmd.PersistentFlags().Bool("picam_open", false, "open picam service on start")
	rootCmd.PersistentFlags().Int("picam_width", 648, "Width of the image from pi camera")
	rootCmd.PersistentFlags().Int("picam_height", 486, "Height of the image from pi camera")
	rootCmd.PersistentFlags().Int("picam_rotation", 0, "Rotation of pi camera image")
	viper.BindPFlag("picam.enabled", rootCmd.PersistentFlags().Lookup("picam"))
	viper.BindPFlag("picam.open", rootCmd.PersistentFlags().Lookup("picam_open"))
	viper.BindPFlag("picam.width", rootCmd.PersistentFlags().Lookup("picam_width"))
	viper.BindPFlag("picam.height", rootCmd.PersistentFlags().Lookup("picam_height"))
	viper.BindPFlag("picam.rotation", rootCmd.PersistentFlags().Lookup("picam_rotation"))

	rootCmd.PersistentFlags().StringP("host", "s", "0.0.0.0", "server ip")
	rootCmd.PersistentFlags().IntP("port", "p", 8000, "server port")
	viper.BindPFlag("server.host", rootCmd.PersistentFlags().Lookup("host"))
	viper.BindPFlag("server.port", rootCmd.PersistentFlags().Lookup("port"))

	helper.AddConfigCommand(rootCmd)
}

var configFileName = "config"
var configPath = "/etc/rpi-server"

func initConfig() {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configFileName)

	readErr := viper.ReadInConfig()
	if readErr != nil {
		fmt.Fprintf(os.Stderr, "unable to read config file: %v\n", readErr)
	}
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix("rpi")
}
