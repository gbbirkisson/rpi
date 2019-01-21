// +build !pi

package rpi

import (
	"context"
	"fmt"
	"io"

	"github.com/LK4D4/joincontext"
	proto "github.com/gbbirkisson/rpi/pkg/proto"
)

// Open the PiCamera
func (c *PiCam) Open(ctx context.Context) error {
	cli, err := c.getClient()
	if err != nil {
		return fmt.Errorf("unable to get grpc client: %v", err)
	}
	_, err = cli.Open(ctx, &proto.Void{})
	return err
}

// Close the PiCamera
func (c *PiCam) Close(ctx context.Context) error {
	cli, err := c.getClient()
	if err != nil {
		return fmt.Errorf("unable to get grpc client: %v", err)
	}
	_, err = cli.Close(ctx, &proto.Void{})
	return err
}

// Get a stream of frames from the camera. Each frame is sent to the byte channel. If an error occurs if is sent to the error channel without stopping to try to get new frames
func (c *PiCam) GetFrames(ctx context.Context, byteCh chan<- []byte, errCh chan<- error) (<-chan struct{}, error) {
	cli, err := c.getClient()
	if err != nil {
		return nil, fmt.Errorf("unable to get grpc client: %v", err)
	}

	stream, err := cli.GetFrames(ctx, &proto.Void{})
	if err != nil {
		return nil, fmt.Errorf("unable to get frame stream: %v", err)
	}

	ctx, _ = joincontext.Join(ctx, stream.Context())

	go func() {
		defer close(byteCh)
		defer close(errCh)

		for {
			if ctx.Err() != nil {
				// Context canceled
				return
			}

			res, err := stream.Recv()
			if err == io.EOF {
				// End of stream
				return
			}

			if err != nil {
				// If some other error
				errCh <- err
				continue
			}

			// Send image to byte channel
			byteCh <- res.ImageBytes
		}
	}()

	return ctx.Done(), nil
}

func (c *PiCam) getClient() (proto.PiCamServiceClient, error) {
	if c.client != nil {
		return c.client, nil
	}
	if c.Connection == nil {
		return nil, fmt.Errorf("piCam.Connection is nil")
	}
	c.client = proto.NewPiCamServiceClient(c.Connection)
	return c.client, nil
}
