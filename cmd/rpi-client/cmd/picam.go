package cmd

import (
	"bytes"
	"image"
	"image/jpeg"
	"os"

	"github.com/spf13/viper"

	"github.com/gbbirkisson/rpi"
	helper "github.com/gbbirkisson/rpi/cmd"
	"github.com/spf13/cobra"
)

var picamCmd = &cobra.Command{
	Use:   "picam",
	Short: "Get frame from the PiCam",
	Run: func(_ *cobra.Command, args []string) {
		ctx, cancel := getContext()
		defer cancel()

		conn, err := rpi.GrpcClientConnectionInsecure(viper.GetString("host"), viper.GetString("port"))
		helper.ExitOnError("could not create client", err)

		cam := &rpi.PiCam{Connection: conn}
		raw, err := cam.GetFrame(ctx)
		helper.ExitOnError("could not get image", err)

		r := bytes.NewReader(raw)
		imageData, _, err := image.Decode(r)

		opts := jpeg.Options{}
		opts.Quality = 80
		jpeg.Encode(os.Stdout, imageData, &opts)
	},
}

func init() {
	rootCmd.AddCommand(picamCmd)

	picamCmd.Flags().BoolP("stream", "s", false, "Get a stream of frames")
	viper.BindPFlag("stream", picamCmd.Flags().Lookup("stream"))
}
