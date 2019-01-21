package rpi

import (
	context "context"
	"log"

	"github.com/gbbirkisson/rpi"
	proto "github.com/gbbirkisson/rpi/pkg/proto"
)

// A GRPC server for basic operations
type CommonServer struct{}

func (s *CommonServer) GetVersion(context.Context, *proto.Void) (*proto.VersionRes, error) {
	log.Printf("Common.GetVersion called\n")
	defer log.Printf("Common.GetVersion finished\n")

	version, revision := rpi.GetVersion()

	return &proto.VersionRes{Version: version, Revision: revision}, nil
}
