package rpi

import (
	context "context"
	"log"

	proto "github.com/gbbirkisson/rpi/pkg/proto"
)

type CommonServer struct {
	Version, Revision string
}

func (s *CommonServer) GetVersion(context.Context, *proto.Void) (*proto.VersionRes, error) {
	log.Printf("Common.Version()\n")
	return &proto.VersionRes{Version: s.Version, Revision: s.Revision}, nil
}
