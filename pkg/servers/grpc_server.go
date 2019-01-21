package rpi

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

// Returns an insecure grpc server that can serve grpc services
func GrpcServerInsecure(host, port string) (*grpc.Server, net.Listener, error) {
	address := host + ":" + port
	log.Printf("listening to %s\n", address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return nil, nil, err
	}
	return grpc.NewServer(), lis, nil
}
