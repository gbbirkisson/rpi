// +build pi

package rpi

import (
	"context"
	"fmt"

	picamera "github.com/gbbirkisson/piCamera"
)

func (c *PiCam) Open(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	args := picamera.NewArgs()
	args.Width = int(c.Width)
	args.Height = int(c.Height)
	args.Rotation = int(c.Rotation)

	cam, err := picamera.New(nil, args)
	if err != nil {
		return fmt.Errorf("unable to create camera: %v", err)
	}

	err = cam.Start()
	if err != nil {
		return fmt.Errorf("unable to start camera: %v", err)
	}

	c.camera = cam
	return nil
}

func (c *PiCam) Close(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	c.camera.Stop()
	return nil
}

func (c *PiCam) GetFrames(ctx context.Context, byteCh chan<- []byte, errCh chan<- error) {
	defer close(byteCh)
	defer close(errCh)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			if ctx.Err() != nil {
				errCh <- ctx.Err()
				return
			}

			frame, err := c.camera.GetFrame()
			if err != nil {
				errCh <- err
			} else {
				byteCh <- frame
			}
		}
	}
}
