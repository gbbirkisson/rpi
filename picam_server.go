package rpi

import (
	"bufio"
	"bytes"
	context "context"
	"fmt"
	"log"
	"os"

	"github.com/dhowden/raspicam"
	proto "github.com/gbbirkisson/rpi/proto"
)

type PiCamServerImpl struct {
}

func (s *PiCamServerImpl) GetPhoto(ctx context.Context, req *proto.Void) (*proto.ResponseImage, error) {
	log.Println("PiCam.GetPhoto()")

	cam := raspicam.NewStill()
	errCh := make(chan error)
	go func() {
		for x := range errCh {
			fmt.Fprintf(os.Stderr, "%v\n", x)
		}
	}()

	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	raspicam.Capture(cam, w, errCh)
	w.Flush()

	return &proto.ResponseImage{ImageBytes: b.Bytes()}, nil
}
