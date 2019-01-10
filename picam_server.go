package rpi

import (
	"fmt"
	"log"
	"time"

	proto "github.com/gbbirkisson/rpi/proto"
	picamera "github.com/technomancers/piCamera"
)

type PiCamServerImpl struct {
	Camera *picamera.PiCamera
}

// func (s *PiCamServerImpl) GetPhoto(ctx context.Context, req *proto.RequestImage) (*proto.ResponseImage, error) {
// 	log.Println("PiCam.GetPhoto()")

// 	cam := raspicam.NewStill()
// 	cam.Width = int(req.Width)
// 	cam.Height = int(req.Height)

// 	if cam.Width == 0 {
// 		cam.Width = 648
// 	}

// 	if cam.Height == 0 {
// 		cam.Height = 486
// 	}
// 	log.Printf("%s", cam.String())

// 	errCh := make(chan error)
// 	go func() {
// 		for x := range errCh {
// 			fmt.Fprintf(os.Stderr, "%v\n", x)
// 		}
// 	}()

// 	var b bytes.Buffer
// 	logTime(func() {
// 		w := bufio.NewWriter(&b)
// 		raspicam.Capture(cam, w, errCh)
// 		w.Flush()
// 	})

// 	return &proto.ResponseImage{ImageBytes: b.Bytes()}, nil
// }

func (s *PiCamServerImpl) GetFrames(req *proto.RequestImage, stream proto.PiCam_GetFramesServer) error {
	log.Println("PiCam.GetFrames()")
	for {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		default:
			start := time.Now()
			frame, err := s.Camera.GetFrame()
			if err != nil {
				return fmt.Errorf("unable to get frame: %v", err)
			}
			stream.Send(&proto.ResponseImage{ImageBytes: frame})
			elapsed := time.Since(start)
			log.Printf("Rendering frame took %s", elapsed)
		}
	}
}

func logTime(f func()) {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	log.Printf("Rendering frame took %s", elapsed)
}
