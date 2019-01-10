package cmd

import (
	"bytes"
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
		res, err := client.GetPhoto(ctx, &proto.Void{})
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

func init() {
	rootCmd.AddCommand(picamCmd)
}
