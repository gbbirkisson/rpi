package rpi_test

import (
	"context"
	"log"

	"github.com/gbbirkisson/rpi"
)

var picam rpi.PiCam
var ctx context.Context

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

	log.Printf("%+v", picam.Connection)
}

func ExamplePiCam_pi() {
	picam := rpi.PiCam{
		Width:    648,
		Height:   486,
		Rotation: 180,
	}

	log.Printf("%d %d %d", picam.Width, picam.Height, picam.Rotation)
}

func ExamplePiCam_Open() {
	err := picam.Open(ctx)
	if err != nil {
		log.Fatalf("unable to open picam: %v", err)
	}
}

func ExamplePiCam_Close() {
	err := picam.Close(ctx)
	if err != nil {
		log.Fatalf("unable to close picam: %v", err)
	}
}

func ExamplePiCam_GetFrame() {
	imgBytes, err := picam.GetFrame(ctx)
	if err != nil {
		log.Fatalf("unable to get frame: %v", err)
	}
	log.Printf("%d", len(imgBytes))
}

func ExamplePiCam_GetFrames() {
	imgChan := make(chan []byte)
	errChan := make(chan error)

	done, err := picam.GetFrames(ctx, imgChan, errChan)
	if err != nil {
		log.Fatalf("unable to start getting frames: %v", err)
	}

	select {
	case imgBytes := <-imgChan:
		log.Printf("%d", len(imgBytes))
	case err := <-errChan:
		log.Fatalf("error getting frame from camera: %v", err)
	case <-done:
		return
	}
}
