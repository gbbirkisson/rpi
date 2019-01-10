package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/gbbirkisson/rpi"
	proto "github.com/gbbirkisson/rpi/proto"
	"github.com/spf13/cobra"
)

var picamCmd = &cobra.Command{
	Use:   "picam",
	Short: "Get a photo from the PiCam",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := getGrpcClient()
		rpi.ExitOnError("could not create client", err)
		client := proto.NewPiCamClient(conn)
		ctx, cancel := getContext()
		defer cancel()

		width, err := cmd.Flags().GetInt32("width")
		rpi.ExitOnError("invalid width argument", err)

		height, err := cmd.Flags().GetInt32("height")
		rpi.ExitOnError("invalid height argument", err)

		res, err := client.GetPhoto(ctx, &proto.RequestImage{Width: width, Height: height})
		rpi.ExitOnError("error repsonse from server", err)

		if len(args) > 0 {
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
	},
}

var picamStreamCmd = &cobra.Command{
	Use:   "stream",
	Short: "Stream frames from the raspberry pi",
	Run: func(cmd *cobra.Command, args []string) {
		conn, err := getGrpcClient()
		rpi.ExitOnError("could not create client", err)
		client := proto.NewPiCamClient(conn)
		ctx, cancel := getContext()
		defer cancel()

		width, err := cmd.Flags().GetInt32("width")
		rpi.ExitOnError("invalid width argument", err)

		height, err := cmd.Flags().GetInt32("height")
		rpi.ExitOnError("invalid height argument", err)

		stream, err := client.GetVideo(ctx, &proto.RequestImage{Width: width, Height: height})
		rpi.ExitOnError("error repsonse from server", err)

		for i := 0; i < 10; i++ {
			res, err := stream.Recv()
			fmt.Println(i)
			if err != nil {
				fmt.Printf("error: %v\n", err)
			} else {
				fmt.Printf("bytes: %d\n", len(res.ImageBytes))
			}
		}

		stream.CloseSend()

		// if len(args) > 0 {
		// 	// Create file
		// 	f, err := os.Create(args[0])
		// 	rpi.ExitOnError("could not create file", err)
		// 	defer f.Close()
		// 	f.Write(res.ImageBytes)
		// } else {
		// 	r := bytes.NewReader(res.ImageBytes)
		// 	_, err := io.Copy(os.Stdout, r)
		// 	rpi.ExitOnError("failed to write to std out", err)
		// }
	},
}

func init() {
	rootCmd.AddCommand(picamCmd)
	picamCmd.AddCommand(picamStreamCmd)
	picamCmd.PersistentFlags().Int32P("width", "x", 648, "Width of the image")
	picamCmd.PersistentFlags().Int32P("height", "y", 486, "Height of the image")
}
