package rpi

import (
	"fmt"

	common "github.com/gbbirkisson/rpi/pkg/common"
	"google.golang.org/grpc"
)

var Revision string = "development"
var Version string = "development"

func GetCommonServer() *common.CommonServer {
	return &common.CommonServer{Version: Version, Revision: Revision}
}

func GetGrpcClient(host, port string) (*grpc.ClientConn, error) {
	address := host + ":" + port
	c, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("could not connect to backend: %v\n", err)
	}
	return c, nil
}
