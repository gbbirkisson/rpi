package rpi

import (
	"context"
	"fmt"
	"image"

	picamera "github.com/gbbirkisson/piCamera"
	picam "github.com/gbbirkisson/rpi/pkg/picam"
	proto "github.com/gbbirkisson/rpi/pkg/proto"
	"google.golang.org/grpc"
)

// This is the PiCam client, when using it over GRPC you have to provide a *grpc.ClientConn.
// When the code is compiled with the tag "-tags pi" there is no GRPC involved so no client
// connection is required.
type PiCam struct {
	Connection              *grpc.ClientConn
	Width, Height, Rotation int32
	client                  proto.PiCamServiceClient
	camera                  *picamera.PiCamera
}

// Get a single frame from the PiCam
func (c *PiCam) GetFrame(ctx context.Context) (*image.Image, error) {
	imgch := make(chan *image.Image)
	errCh := make(chan error)

	newCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	go c.GetFrames(newCtx, imgch, errCh)

	select {
	case img := <-imgch:
		return img, nil
	case err := <-errCh:
		return nil, err
	}
}

func (c *PiCam) createImage(bytes []byte, width, height int32) *image.Image {
	return nil // TODO: Create image
}

func GetPiCam(args *picamera.RaspividArgs) (*picamera.PiCamera, error) { // TODO: Remove
	return picamera.New(nil, args)
}

func GetPiCamServer(camera *picamera.PiCamera) (*picam.PiCamServer, error) { // TODO: Remove
	if camera == nil {
		return nil, fmt.Errorf("picamera cannot be nil")
	}
	return &picam.PiCamServer{Camera: camera}, nil
}
