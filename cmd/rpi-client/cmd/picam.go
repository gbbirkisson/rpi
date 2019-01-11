package cmd

import (
	"bufio"
	"bytes"
	"image"
	"image/png"
	"os"

	"github.com/gbbirkisson/rpi"
	proto "github.com/gbbirkisson/rpi/proto"
	"github.com/spf13/cobra"
)

var picamCmd = &cobra.Command{
	Use:   "picam",
	Short: "Get frame from the PiCam",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := getGrpcClient()
		rpi.ExitOnError("could not create client", err)
		client := proto.NewPiCamClient(conn)
		ctx, cancel := getContext()
		defer cancel()

		stream, err := client.GetFrames(ctx, &proto.RequestImage{})
		rpi.ExitOnError("error response from server", err)
		defer stream.CloseSend()

		shouldStream, err := cmd.Flags().GetBool("stream")
		rpi.ExitOnError("stream flag invalid", err)

		for {
			res, err := stream.Recv()
			rpi.ExitOnError("error getting frame", err)
			r := bytes.NewReader(res.ImageBytes)
			imageData, _, err := image.Decode(r)
			rpi.ExitOnError("unable to decode image", err)

			if len(args) > 0 && !shouldStream {
				// Create file
				f, err := os.Create(args[0])
				rpi.ExitOnError("could not create file", err)
				defer f.Close()
				w := bufio.NewWriter(f)
				err = png.Encode(w, imageData)
				rpi.ExitOnError("unable to encode image", err)
			} else {
				w := bufio.NewWriter(os.Stdout)
				err = png.Encode(w, imageData)
				rpi.ExitOnError("unable to encode image", err)
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
}
