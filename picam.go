package rpi

import (
	"fmt"

	picamera "github.com/gbbirkisson/piCamera"
	picam "github.com/gbbirkisson/rpi/pkg/picam"
)

func GetPiCam(args *picamera.RaspividArgs) (*picamera.PiCamera, error) {
	return picamera.New(nil, args)
}

func GetPiCamServer(camera *picamera.PiCamera) (*picam.PiCamServer, error) {
	if camera == nil {
		return nil, fmt.Errorf("picamera cannot be nil")
	}
	return &picam.PiCamServer{Camera: camera}, nil
}
