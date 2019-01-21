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
