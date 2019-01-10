package cmd

import (
	"fmt"

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
		fmt.Printf("%c", res.ImageBytes)
	},
}

func init() {
	rootCmd.AddCommand(picamCmd)
}
