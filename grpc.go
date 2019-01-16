package rpi

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

// Returns an insecure grpc client connection that can used by grpc clients
func GrpcClientConnectionInsecure(host, port string) (*grpc.ClientConn, error) {
	address := host + ":" + port
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("could not connect to backend: %v\n", err)
	}
	return conn, nil
}

// Returns an insecure grpc server that can serve grpc services
func GrpcServerInsecure(host, port string) (*grpc.Server, net.Listener, error) {
	address := host + ":" + port
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return nil, nil, err
	}
	return grpc.NewServer(), lis, nil
}
