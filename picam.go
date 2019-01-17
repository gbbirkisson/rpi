package rpi

import (
	"context"

	picamera "github.com/gbbirkisson/piCamera"
	proto "github.com/gbbirkisson/rpi/pkg/proto"
	"google.golang.org/grpc"
)

// This is the PiCam client, when using it over GRPC you have to provide a *grpc.ClientConn. When the code is compiled with the tag "-tags pi" there is no GRPC involved so no client connection is required. Then however you can specify Width, Height and Rotation. See examples.
type PiCam struct {
	Connection              *grpc.ClientConn
	Width, Height, Rotation int32
	client                  proto.PiCamServiceClient
	camera                  *picamera.PiCamera
}

// Get a single frame from the PiCam
func (c *PiCam) GetFrame(ctx context.Context) ([]byte, error) {
	imgch := make(chan []byte)
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
