package rpi

import (
	context "context"
	"log"

	common "github.com/gbbirkisson/rpi"
	rpi "github.com/gbbirkisson/rpi/proto"
)

type CommonServerImpl struct{}

func (s *CommonServerImpl) Version(context.Context, *rpi.Void) (*rpi.VersionRes, error) {
	log.Printf("Common.Version()\n")
	return &rpi.VersionRes{Version: common.Version}, nil
}
