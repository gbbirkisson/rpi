package rpi

import "context"

func getFrame(ctx context.Context, cam PiCam) ([]byte, error) {
	imgch := make(chan []byte)
	errCh := make(chan error)

	newCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	_, err := cam.GetFrames(newCtx, imgch, errCh)
	if err != nil {
		return nil, err
	}

	select {
	case img := <-imgch:
		return img, nil
	case err := <-errCh:
		return nil, err
	}
}
