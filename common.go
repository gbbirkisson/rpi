package rpi

import (
	proto "github.com/gbbirkisson/rpi/pkg/proto"
	"google.golang.org/grpc"
)

var version string = ""
var revision string = "development"

type Common struct {
	Connection *grpc.ClientConn
	client     proto.CommonServiceClient
}

func GetLocalVersion() (string, string) {
	return version, revision
}
