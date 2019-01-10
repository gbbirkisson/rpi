package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"

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
			if err != nil {
				fmt.Fprintf(os.Stderr, "error getting frame: %v\n", err)
				time.Sleep(100 * time.Millisecond)
			} else {
				if len(args) > 0 && !shouldStream {
					// Create file
					f, err := os.Create(args[0])
					rpi.ExitOnError("could not create file", err)
					defer f.Close()
					f.Write(res.ImageBytes)
				} else {
					r := bytes.NewReader(res.ImageBytes)
					_, err := io.Copy(os.Stdout, r)
					rpi.ExitOnError("failed to write to std out", err)
				}
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
