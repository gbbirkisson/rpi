// +build !pi

package rpi

import (
	"context"
	"fmt"

	proto "github.com/gbbirkisson/rpi/pkg/proto"
)

func (c *PiCam) Open(ctx context.Context) error {
	cli, err := c.getClient()
	if err != nil {
		return fmt.Errorf("unable to get grpc client: %v", err)
	}
	_, err = cli.Open(ctx, &proto.RequestOpen{Width: c.Width, Height: c.Height, Rotation: c.Rotation})
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

func (c *PiCam) GetFrames(ctx context.Context, byteCh chan<- []byte, errCh chan<- error) {
	defer close(byteCh)
	defer close(errCh)

	cli, err := c.getClient()
	if err != nil {
		errCh <- fmt.Errorf("unable to get grpc client: %v", err)
		return
	}

	stream, err := cli.GetFrames(ctx, &proto.Void{})
	defer stream.CloseSend()

	if err != nil {
		errCh <- fmt.Errorf("unable to get frames: %v", err)
		return
	}

	for {
		select {
		case <-ctx.Done():
			return
		default:
			if ctx.Err() != nil {
				errCh <- ctx.Err()
				return
			}

			res, err := stream.Recv()

			if err != nil {
				errCh <- err
			} else {
				byteCh <- res.ImageBytes
			}
		}
	}
}

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
