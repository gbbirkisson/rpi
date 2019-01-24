package rpi

import (
	"context"
	"log"

	proto "github.com/gbbirkisson/rpi/pkg/proto"
)

// Get new common server
func NewCommonServer(common Common) proto.CommonServer {
	return &commonServer{common: common}
}

type commonServer struct {
	common Common
}

func (s *commonServer) GetVersion(ctx context.Context, _ *proto.Void) (*proto.VersionRes, error) {
	log.Printf("Common.GetVersion called\n")
	defer log.Printf("Common.GetVersion finished\n")

	version, revision, err := s.common.GetVersion(ctx)
	if err != nil {
		return nil, err
	}

	return &proto.VersionRes{Version: version, Revision: revision}, nil
}

func (s *commonServer) Modprobe(ctx context.Context, req *proto.ModprobeRequest) (*proto.Void, error) {
	return &proto.Void{}, s.common.Modprobe(ctx, req.Module)
}
