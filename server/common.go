package rpi

import (
	context "context"

	common "github.com/gbbirkisson/rpi"
	rpi "github.com/gbbirkisson/rpi/proto"
)

type CommonServerImpl struct{}

func (s *CommonServerImpl) Version(context.Context, *rpi.Void) (*rpi.VersionRes, error) {
	return &rpi.VersionRes{Version: common.Version}, nil
}
