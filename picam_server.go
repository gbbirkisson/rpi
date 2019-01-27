package rpi

import (
	"context"
	"fmt"
	"log"

	proto "github.com/gbbirkisson/rpi/pkg/proto"
)

// NewPicamServer creates a picam server that uses the picam interface provided
func NewPicamServer(cam PiCam) proto.PiCamServer {
	return &piCamServer{camera: cam}
}

type piCamServer struct {
	camera PiCam
}

func (s *piCamServer) Open(ctx context.Context, req *proto.Void) (*proto.Void, error) {
	log.Printf("PiCam.Open called\n")
	defer log.Printf("PiCam.Open finished\n")

	return &proto.Void{}, s.camera.Open(ctx)
}

func (s *piCamServer) Close(ctx context.Context, req *proto.Void) (*proto.Void, error) {
	log.Printf("PiCam.Close called\n")
	defer log.Printf("PiCam.Close finished\n")

	return &proto.Void{}, s.camera.Close(ctx)
}

func (s *piCamServer) GetFrames(_ *proto.Void, stream proto.PiCam_GetFramesServer) error {
	log.Printf("PiCam.GetFrames called\n")
	defer log.Printf("PiCam.GetFrames finished\n")

	imgChan := make(chan []byte)
	errChan := make(chan error)

	done, err := s.camera.GetFrames(stream.Context(), imgChan, errChan)

	if err != nil {
		return fmt.Errorf("unable to start getting frames")
	}

	for {
		select {
		case <-done:
			return nil
		case err := <-errChan:
			return err
		case img := <-imgChan:
			stream.Send(&proto.ResponseImage{ImageBytes: img})
		}
	}
}
