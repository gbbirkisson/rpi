package rpi_test

import (
	"context"
	"fmt"
	"log"

	"github.com/gbbirkisson/rpi"
)

func ExamplePiCam() {
	conn, err := rpi.GrpcClientConnectionInsecure("localhost", "8000")
	if err != nil {
		log.Fatalf("unable to create grpc client connection: %v\n", err)
	}
	defer conn.Close()

	picam := rpi.PiCam{
		Connection: conn,
		Width:      648,
		Height:     486,
		Rotation:   180,
	}

	fmt.Printf("%+v", picam.Connection)
}

func ExamplePiCam_pi() {
	picam := rpi.PiCam{
		Width:    648,
		Height:   486,
		Rotation: 180,
	}

	fmt.Printf("%d %d %d", picam.Width, picam.Height, picam.Rotation)
}

func ExamplePiCam_Open() {
	ctx := context.Background()

	conn, err := rpi.GrpcClientConnectionInsecure("localhost", "8000")
	if err != nil {
		log.Fatalf("unable to create grpc client connection: %v\n", err)
	}
	defer conn.Close()

	picam := rpi.PiCam{
		Connection: conn,
		Width:      648,
		Height:     486,
		Rotation:   180,
	}

	err = picam.Open(ctx)
	if err != nil {
		log.Fatalf("Unable to open picam: %v", err)
	}
}

func ExamplePiCam_Open_pi() {
	ctx := context.Background()

	picam := rpi.PiCam{
		Width:    648,
		Height:   486,
		Rotation: 180,
	}

	err := picam.Open(ctx)
	if err != nil {
		log.Fatalf("Unable to open picam: %v", err)
	}
}
