package rpi

import (
	"context"
	"fmt"

	"github.com/gbbirkisson/rpi"
	proto "github.com/gbbirkisson/rpi/pkg/proto"
)

// A PiCam GRPC server that needs a PiCam to operate
type PiCamServer struct {
	Camera                  *rpi.PiCam
	Width, Height, Rotation int32
}

func (s *PiCamServer) Open(ctx context.Context, req *proto.Void) (*proto.Void, error) {

	if s.Camera != nil {
		return &proto.Void{}, s.Camera.Open(ctx)
	}

	cam := rpi.PiCam{
		Width:    s.Width,
		Height:   s.Height,
		Rotation: s.Rotation,
	}

	err := cam.Open(ctx)
	if err != nil {
		return nil, err
	}

	s.Camera = &cam

	return &proto.Void{}, nil
}

func (s *PiCamServer) Close(ctx context.Context, req *proto.Void) (*proto.Void, error) {
	if s.Camera == nil {
		return nil, fmt.Errorf("server camera already closed")
	}

	err := s.Camera.Close(ctx)
	if s.Camera == nil {
		return nil, err
	}

	s.Camera = nil

	return &proto.Void{}, nil
}

func (s *PiCamServer) GetFrames(_ *proto.Void, stream proto.PiCamService_GetFramesServer) error {
	if s.Camera == nil {
		return fmt.Errorf("server camera closed")
	}

	imgChan := make(chan []byte)
	errChan := make(chan error)

	done, err := s.Camera.GetFrames(stream.Context(), imgChan, errChan)

	if err == nil {
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
