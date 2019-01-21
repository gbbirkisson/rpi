package rpi

import (
	context "context"
	"log"

	"github.com/gbbirkisson/rpi"
	proto "github.com/gbbirkisson/rpi/pkg/proto"
)

// A GRPC server for basic operations
type CommonServer struct {
	Common *rpi.Common
}

func (s *CommonServer) GetVersion(ctx context.Context, _ *proto.Void) (*proto.VersionRes, error) {
	log.Printf("Common.GetVersion called\n")
	defer log.Printf("Common.GetVersion finished\n")

	version, revision, err := s.Common.GetVersion(ctx)
	if err != nil {
		return nil, err
	}

	return &proto.VersionRes{Version: version, Revision: revision}, nil
}

func (s *CommonServer) Modprobe(ctx context.Context, req *proto.ModprobeRequest) (*proto.Void, error) {
	return &proto.Void{}, s.Common.Modprobe(ctx, req.Params)
}
