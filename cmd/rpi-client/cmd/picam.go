package cmd

import (
	"bytes"
	"image"
	"image/jpeg"
	"io"
	"os"
	"os/exec"

	"github.com/gbbirkisson/rpi"
	helper "github.com/gbbirkisson/rpi/cmd"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

func getPicam() rpi.PiCam {
	return rpi.NewPiCamRemote(getConnection())
}

var picamCmd = &cobra.Command{
	Use:   "picam",
	Short: "Control the PiCam on server",
}

var picamOpenCmd = &cobra.Command{
	Use:   "open",
	Short: "Open the Picam",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := getContext()
		defer cancel()

		cam := getPicam()
		err := cam.Open(ctx)
		helper.ExitOnError("error opening the picam", err)
	},
}

var picamCloseCmd = &cobra.Command{
	Use:   "close",
	Short: "Close the Picam",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := getContext()
		defer cancel()

		cam := getPicam()
		err := cam.Close(ctx)
		helper.ExitOnError("error closing the picam", err)
	},
}

var picamFrameCmd = &cobra.Command{
	Use:   "frame",
	Short: "Get frame from the PiCam",
	Long: `To open up image in a viewer (feh by defult):

	rpi-client picam frame

To save image to file:

rpi-client picam frame > image.jpg
	`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := getContext()
		defer cancel()

		cam := getPicam()

		raw, err := cam.GetFrame(ctx)
		helper.ExitOnError("could not get image", err)

		r := bytes.NewReader(raw)
		imageData, _, err := image.Decode(r)
		helper.ExitOnError("unable to decode image", err)

		pr, pw := io.Pipe()
		defer pr.Close()

		opts := jpeg.Options{}
		opts.Quality = 80
		go func() {
			jpeg.Encode(pw, imageData, &opts)
			pw.Close()
		}()

		if terminal.IsTerminal(int(os.Stdout.Fd())) {
			err = viewer(cmd, pr)
			helper.ExitOnError("error running viewer command", err)
		} else {
			io.Copy(os.Stdout, pr)
		}
	},
}

func viewer(cmd *cobra.Command, reader io.Reader) error {
	viewerCmd, err := cmd.Flags().GetStringSlice("viewer")
	helper.ExitOnError("unable to get viewer command", err)

	vCmd := exec.Command(viewerCmd[0], viewerCmd[1:]...)
	vCmd.Stdin = reader
	vCmd.Stdout = os.Stdout

	return vCmd.Run()
}

func init() {
	picamFrameCmd.Flags().StringSliceP("viewer", "v", []string{"feh", "-x", "-"}, "Command that the image bytes will be piped to using stdin")

	picamCmd.AddCommand(
		picamFrameCmd,
		picamCloseCmd,
		picamOpenCmd,
	)

	rootCmd.AddCommand(picamCmd)
}
