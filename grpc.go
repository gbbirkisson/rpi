package rpi

import (
	"fmt"

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
