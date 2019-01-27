package rpi

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

// NewGrpcClientConnectionInsecure creates a new insecure grpc client connection that can used by grpc clients
func NewGrpcClientConnectionInsecure(host, port string) (*grpc.ClientConn, error) {
	address := host + ":" + port
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("could not connect to backend: %v", err)
	}
	return conn, nil
}

// NewGrpcServerInsecure creates a new insecure grpc server that can serve grpc services
func NewGrpcServerInsecure(host, port string) (*grpc.Server, net.Listener, error) {
	address := host + ":" + port
	log.Printf("listening to %s\n", address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		return nil, nil, err
	}
	return grpc.NewServer(grpc.UnaryInterceptor(loggingInterceptorUnary), grpc.StreamInterceptor(loggingInterceptorStream)), lis, nil
}

func loggingInterceptorUnary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()

	log.Printf("%s(%+v) called\n", info.FullMethod, req)
	h, err := handler(ctx, req)
	log.Printf("%s(%+v) finished in %s, error: %v\n", info.FullMethod, req, time.Since(start), err)

	return h, err
}

func loggingInterceptorStream(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	start := time.Now()
	log.Printf("%s(...) stream started\n", info.FullMethod)
	err := handler(srv, stream)
	log.Printf("%s(...) finished in %s, error: %v\n", info.FullMethod, time.Since(start), err)
	return err
}
