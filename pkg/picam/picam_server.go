package rpi

import (
	"fmt"
	"log"
	"time"

	picamera "github.com/gbbirkisson/piCamera"
	proto "github.com/gbbirkisson/rpi/pkg/proto"
)

type PiCamServer struct {
	Camera *picamera.PiCamera
}

func (s *PiCamServer) GetFrames(stream proto.PiCamService_GetFramesServer) error {
	log.Println("PiCam.GetFrames()")
	for {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		default:
			start := time.Now()
			frame, err := s.Camera.GetFrame()
			if err != nil {
				return fmt.Errorf("unable to get frame: %v", err)
			}
			stream.Send(&proto.ResponseImage{ImageBytes: frame})
			elapsed := time.Since(start)
			log.Printf("Rendering frame took %s", elapsed)
		}
	}
}

func logTime(f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	log.Printf("Rendering frame took %s", elapsed)
}
