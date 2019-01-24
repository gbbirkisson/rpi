package rpi

import (
	"context"
	"fmt"
	"io"

	"github.com/LK4D4/joincontext"
	proto "github.com/gbbirkisson/rpi/pkg/proto"
	"google.golang.org/grpc"
)

// Get a new remote PiCam
func NewPiCamRemote(connection *grpc.ClientConn) (PiCam, error) {
	return &piCamGrpc{client: proto.NewPiCamClient(connection)}, nil
}

type piCamGrpc struct {
	client proto.PiCamClient
}

func (c *piCamGrpc) Open(ctx context.Context) error {
	_, err := c.client.Open(ctx, &proto.Void{})
	return err
}

func (c *piCamGrpc) Close(ctx context.Context) error {
	_, err := c.client.Close(ctx, &proto.Void{})
	return err
}

func (c *piCamGrpc) GetFrame(ctx context.Context) ([]byte, error) {
	return getFrame(c, ctx)
}

func (c *piCamGrpc) GetFrames(ctx context.Context, byteCh chan<- []byte, errCh chan<- error) (<-chan struct{}, error) {
	stream, err := c.client.GetFrames(ctx, &proto.Void{})
	if err != nil {
		return nil, fmt.Errorf("unable to get frame stream: %v", err)
	}

	ctx, _ = joincontext.Join(ctx, stream.Context())

	go func() {
		defer close(byteCh)
		defer close(errCh)

		for {
			if ctx.Err() != nil {
				// Context canceled
				return
			}

			res, err := stream.Recv()
			if err == io.EOF {
				// End of stream
				return
			}

			if err != nil {
				// If some other error
				errCh <- err
				continue
			}

			// Send image to byte channel
			byteCh <- res.ImageBytes
		}
	}()

	return ctx.Done(), nil
}
