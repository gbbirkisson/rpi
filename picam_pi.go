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

func (c *PiCam) GetFrames(ctx context.Context, byteCh chan<- []byte, errCh chan<- error) (<-chan struct{}, error) {

	if c.camera == nil {
		return nil, fmt.Errorf("picam is nil")
	}

	go func() {
		defer close(byteCh)
		defer close(errCh)

		for {
			if ctx.Err() != nil {
				return
			}

			frame, err := c.camera.GetFrame()

			if err != nil {
				errCh <- err
			} else {
				byteCh <- frame
			}
		}
	}()

	return ctx.Done(), nil
}
