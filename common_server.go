package rpi

import (
	"context"

	proto "github.com/gbbirkisson/rpi/pkg/proto"
)

// NewCommonServer creates a new common server using the common interface provided
func NewCommonServer(common Common) proto.CommonServer {
	return &commonServer{common: common}
}

type commonServer struct {
	common Common
}

func (s *commonServer) GetVersion(ctx context.Context, _ *proto.Void) (*proto.VersionRes, error) {
	version, revision, err := s.common.GetVersion(ctx)
	if err != nil {
		return nil, err
	}

	return &proto.VersionRes{Version: version, Revision: revision}, nil
}

func (s *commonServer) Modprobe(ctx context.Context, req *proto.ModprobeRequest) (*proto.Void, error) {
	return &proto.Void{}, s.common.Modprobe(ctx, req.Module)
}
