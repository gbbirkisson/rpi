package rpi

import (
	context "context"
	"fmt"
	"log"
	"os"

	proto "github.com/gbbirkisson/rpi/proto"
)

var Version string = "development"

func ExitOnError(msg string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, msg+": %v\n", err)
		os.Exit(1)
	}
}

type CommonServerImpl struct{}

func (s *CommonServerImpl) Version(context.Context, *proto.Void) (*proto.VersionRes, error) {
	log.Printf("Common.Version()\n")
	return &proto.VersionRes{Version: Version}, nil
}
