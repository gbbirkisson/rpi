package rpi

import (
	"context"
	"fmt"

	proto "github.com/gbbirkisson/rpi/pkg/proto"
)

// A PiCam GRPC server that needs a PiCam to operate
type PiCamServer struct {
	Camera *PiCam
}

func (s *PiCamServer) Open(ctx context.Context, req *proto.RequestOpen) (*proto.Void, error) {

	if s.Camera != nil {
		return &proto.Void{}, s.Camera.Open(ctx)
	}

	cam := PiCam{
		Width:    req.Width,
		Height:   req.Height,
		Rotation: req.Rotation,
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

	imgch := make(chan []byte)
	errCh := make(chan error)
	go s.Camera.GetFrames(stream.Context(), imgch, errCh)
	for {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		case err := <-errCh:
			return err
		case img := <-imgch:
			stream.Send(&proto.ResponseImage{ImageBytes: img})
		}
	}
}
