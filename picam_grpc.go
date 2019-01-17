// +build pi

package rpi

import (
	"context"
	"fmt"
	"image"

	proto "github.com/gbbirkisson/rpi/pkg/proto"
)

func (c *PiCam) getClient() (proto.PiCamServiceClient, error) {
	if c.client != nil {
		return c.client, nil
	}
	if c.Connection == nil {
		return nil, fmt.Errorf("PiCam.Connection is nil")
	}
	c.client = proto.NewPiCamServiceClient(c.Connection)
	return c.client, nil
}

func (c *PiCam) Open(ctx context.Context, width, height, rotation int32) error {
	cli, err := c.getClient()
	if err != nil {
		return fmt.Errorf("unable to get grpc client: %v", err)
	}
	_, err = cli.Open(ctx, &proto.RequestOpen{Width: width, Height: height, Rotation: rotation})
	return err
}

func (c *PiCam) Close(ctx context.Context) error {
	cli, err := c.getClient()
	if err != nil {
		return fmt.Errorf("unable to get grpc client: %v", err)
	}
	_, err = cli.Close(ctx, &proto.Void{})
	return err
}

func (c *PiCam) GetFrames(ctx context.Context, imageChan chan<- image.Image, errorChan chan<- error) {
	defer close(imageChan)
	defer close(errorChan)

	cli, err := c.getClient()
	if err != nil {
		errorChan <- fmt.Errorf("unable to get grpc client: %v", err)
		return
	}

	stream, err := cli.GetFrames(ctx, &proto.Void{})
	defer stream.CloseSend()

	if err != nil {
		errorChan <- fmt.Errorf("unable to get frames: %v", err)
		return
	}

	for {
		select {
		case <-ctx.Done():
			errorChan <- ctx.Err()
			return
		default:
			if ctx.Err() != nil {
				return
			}

			res, err := stream.Recv()

			if err != nil {
				errorChan <- err
			} else {
				imageChan <- c.createImage(res.ImageBytes, res.Width, res.Height)
			}
		}
	}
}
