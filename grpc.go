package rpi

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

// Get a new insecure grpc client connection that can used by grpc clients
func NewGrpcClientConnectionInsecure(host, port string) (*grpc.ClientConn, error) {
	address := host + ":" + port
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("could not connect to backend: %v\n", err)
	}
	return conn, nil
}

// Get a new insecure grpc server that can serve grpc services
func NewGrpcServerInsecure(host, port string) (*grpc.Server, net.Listener, error) {
	address := host + ":" + port
	log.Printf("listening to %s\n", address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return nil, nil, err
	}
	return grpc.NewServer(), lis, nil
}
