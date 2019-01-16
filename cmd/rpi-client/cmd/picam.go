package cmd

import (
	"bytes"
	"image"
	"image/jpeg"
	"io"
	"os"

	"github.com/spf13/viper"

	"github.com/gbbirkisson/rpi"
	helper "github.com/gbbirkisson/rpi/cmd"
	proto "github.com/gbbirkisson/rpi/pkg/proto"
	"github.com/spf13/cobra"
)

var picamCmd = &cobra.Command{
	Use:   "picam",
	Short: "Get frame from the PiCam",
	Run: func(_ *cobra.Command, args []string) {
		conn, err := rpi.GrpcClientConnectionInsecure(viper.GetString("host"), viper.GetString("port"))
		helper.ExitOnError("could not create client", err)
		client := proto.NewPiCamServiceClient(conn)
		ctx, cancel := getContext()
		defer cancel()

		stream, err := client.GetFrames(ctx, &proto.RequestImage{})
		helper.ExitOnError("error response from server", err)
		defer stream.CloseSend()

		shouldStream := viper.GetBool("stream")
		for {
			res, err := stream.Recv()
			helper.ExitOnError("error getting frame", err)
			r := bytes.NewReader(res.ImageBytes)
			imageData, _, err := image.Decode(r)
			helper.ExitOnError("unable to decode image", err)

			opts := jpeg.Options{}
			opts.Quality = 80

			if len(args) > 0 && !shouldStream {
				// Create file
				f, err := os.Create(args[0])
				helper.ExitOnError("could not create file", err)
				defer f.Close()
				io.Copy(f, bytes.NewReader(res.ImageBytes))
				// err = jpeg.Encode(f, imageData, &opts)
				// helper.ExitOnError("unable to encode image", err)
			} else {
				err = jpeg.Encode(os.Stdout, imageData, &opts)
				helper.ExitOnError("unable to encode image", err)
			}
			if !shouldStream {
				break
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(picamCmd)

	picamCmd.Flags().BoolP("stream", "s", false, "Get a stream of frames")
	viper.BindPFlag("stream", picamCmd.Flags().Lookup("stream"))
}
