package rpi

import (
	"bufio"
	"bytes"
	context "context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dhowden/raspicam"
	proto "github.com/gbbirkisson/rpi/proto"
)

type PiCamServerImpl struct {
}

func (s *PiCamServerImpl) GetPhoto(ctx context.Context, req *proto.RequestImage) (*proto.ResponseImage, error) {
	log.Println("PiCam.GetPhoto()")

	cam := raspicam.NewStill()
	cam.Width = int(req.Width)
	cam.Height = int(req.Height)

	if cam.Width == 0 {
		cam.Width = 648
	}

	if cam.Height == 0 {
		cam.Height = 486
	}
	log.Printf("%s", cam.String())

	errCh := make(chan error)
	go func() {
		for x := range errCh {
			fmt.Fprintf(os.Stderr, "%v\n", x)
		}
	}()

	var b bytes.Buffer
	logTime(func() {
		w := bufio.NewWriter(&b)
		raspicam.Capture(cam, w, errCh)
		w.Flush()
	})

	return &proto.ResponseImage{ImageBytes: b.Bytes()}, nil
}

func (s *PiCamServerImpl) GetVideo(req *proto.RequestImage, stream proto.PiCam_GetVideoServer) error {
	log.Println("PiCam.GetVideo()")
	cam := raspicam.NewStill()

	cam.Width = int(req.Width)
	cam.Height = int(req.Height)

	if cam.Width == 0 {
		cam.Width = 648
	}

	if cam.Height == 0 {
		cam.Height = 486
	}

	errCh := make(chan error)
	go func() {
		for x := range errCh {
			fmt.Fprintf(os.Stderr, "%v\n", x)
		}
	}()

	for {
		select {
		case err := <-errCh:
			return err
		default:
			var b bytes.Buffer
			logTime(func() {
				w := bufio.NewWriter(&b)
				raspicam.Capture(cam, w, errCh)
				w.Flush()
			})
			stream.Send(&proto.ResponseImage{ImageBytes: b.Bytes()})
		}
	}
}

func logTime(f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	log.Printf("Rendering frame took %s", elapsed)
}
