package rpi_test

import (
	"context"
	"log"

	"github.com/gbbirkisson/rpi"
)

func ExamplePiCam_Open() {
	ctx := context.Background()

	conn, err := rpi.GrpcClientConnectionInsecure("localhost", "8000")
	if err != nil {
		log.Fatalf("unable to create grpc client connection: %v\n", err)
	}
	defer conn.Close()

	// When using GRPC
	picam := rpi.PiCam{
		Connection: conn,
	}

	err = picam.Open(ctx)
	if err != nil {
		log.Fatalf("Unable to open picam: %v", err)
	}
}

func ExamplePiCam_Open_pi() {
	ctx := context.Background()

	picam := rpi.PiCam{
		Width:  648,
		Height: 486,
	}

	err := picam.Open(ctx)
	if err != nil {
		log.Fatalf("Unable to open picam: %v", err)
	}
}
