// +build !pi

package rpi

import (
	"context"
	"fmt"
	"image"

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

func (c *PiCam) GetFrames(ctx context.Context, imageChan chan<- *image.Image, errorChan chan<- error) {
	defer close(imageChan)
	defer close(errorChan)

	for {
		select {
		case <-ctx.Done():
			errorChan <- ctx.Err()
			return
		default:
			if ctx.Err() != nil {
				return
			}

			frame, err := c.camera.GetFrame()
			if err != nil {
				errorChan <- err
			} else {
				imageChan <- c.createImage(frame, c.Width, c.Rotation)
			}
		}
	}
}
