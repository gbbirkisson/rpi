package rpi

import (
	"context"
	"fmt"
	"log"

	"github.com/gbbirkisson/rpi"
	proto "github.com/gbbirkisson/rpi/pkg/proto"
)

// A PiCam GRPC service that needs a PiCam to operate
type PiCamServer struct {
	Camera *rpi.PiCam
}

func (s *PiCamServer) Open(ctx context.Context, req *proto.Void) (*proto.Void, error) {
	log.Printf("PiCam.Open called\n")
	defer log.Printf("PiCam.Open finished\n")

	if s.Camera == nil {
		return nil, fmt.Errorf("server picam is null")
	}

	return &proto.Void{}, s.Camera.Open(ctx)
}

func (s *PiCamServer) Close(ctx context.Context, req *proto.Void) (*proto.Void, error) {
	log.Printf("PiCam.Close called\n")
	defer log.Printf("PiCam.Close finished\n")

	if s.Camera == nil {
		return nil, fmt.Errorf("server picam is null")
	}

	return &proto.Void{}, s.Camera.Close(ctx)
}

func (s *PiCamServer) GetFrames(_ *proto.Void, stream proto.PiCamService_GetFramesServer) error {
	log.Printf("PiCam.GetFrames called\n")
	defer log.Printf("PiCam.GetFrames finished\n")

	if s.Camera == nil {
		return fmt.Errorf("server picam is null")
	}

	imgChan := make(chan []byte)
	errChan := make(chan error)

	done, err := s.Camera.GetFrames(stream.Context(), imgChan, errChan)

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
