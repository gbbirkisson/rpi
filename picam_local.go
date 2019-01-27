package rpi

import (
	"context"
	"fmt"

	picamera "github.com/gbbirkisson/piCamera"
)

// PiCamArgs is a struct of arguments when initializing the PiCamera
type PiCamArgs = picamera.RaspividArgs

// NewPiCamArgs creates the default arguments for the PiCam
func NewPiCamArgs() *PiCamArgs {
	return picamera.NewArgs()
}

// NewPiCamLocal creates a new local PiCam
func NewPiCamLocal(args *PiCamArgs) (PiCam, error) {
	if args.Brightness == 0 || args.Mode == 0 {
		return nil, fmt.Errorf("invalid camera arguments, use NewPiCamArgs() to create PiCam arguments")
	}
	cam, err := picamera.New(nil, args)
	if err != nil {
		return nil, err
	}
	return &piCamLocal{camera: cam}, nil
}

type piCamLocal struct {
	camera *picamera.PiCamera
}

func (c *piCamLocal) Open(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}

	err := c.camera.Start()
	if err != nil {
		return fmt.Errorf("unable to start camera: %v", err)
	}
	return nil
}

func (c *piCamLocal) Close(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	c.camera.Stop()
	return nil
}

func (c *piCamLocal) GetFrame(ctx context.Context) ([]byte, error) {
	return getFrame(ctx, c)
}

func (c *piCamLocal) GetFrames(ctx context.Context, byteCh chan<- []byte, errCh chan<- error) (<-chan struct{}, error) {
	if ctx.Err() != nil {
		return nil, ctx.Err()
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
