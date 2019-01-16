package rpi_test

import (
	"log"

	"github.com/gbbirkisson/rpi"
)

func ExampleGrpcClientConnectionInsecure() {
	conn, err := rpi.GrpcClientConnectionInsecure("localhost", "8000")
	if err != nil {
		log.Fatalf("unable to create grpc client connection: %v\n", err)
	}
	defer conn.Close()

	// Use client
}

func ExampleGrpcServerInsecure() {
	srv, lis, err := rpi.GrpcServerInsecure("0.0.0.0", "8000")
	if err != nil {
		log.Fatalf("unable to create grpc server: %v\n", err)
	}

	// Register services with srv.RegisterService(...)

	log.Fatal(srv.Serve(lis))
}
