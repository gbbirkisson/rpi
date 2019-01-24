package rpi

import (
	"context"
	"fmt"
	"os/exec"

	proto "github.com/gbbirkisson/rpi/pkg/proto"
	"google.golang.org/grpc"
)

var version string = ""
var revision string = "development"

// Get new local common functions
func NewCommonLocal() Common {
	return &commonLocal{}
}

type commonLocal struct{}

func (c *commonLocal) GetVersion(ctx context.Context) (string, string, error) {
	return version, revision, nil
}

func (c *commonLocal) Modprobe(ctx context.Context, mod string) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	output, err := exec.Command("modprobe", mod).CombinedOutput()
	if err != nil {
		return fmt.Errorf("%v: %s", err, output)
	}
	return nil
}

// Get new remote common functions
func NewCommonRemote(connection *grpc.ClientConn) (Common, error) {
	return &commonRemote{client: proto.NewCommonClient(connection)}, nil
}

type commonRemote struct {
	client proto.CommonClient
}

func (c *commonRemote) GetVersion(ctx context.Context) (string, string, error) {
	res, err := c.client.GetVersion(ctx, &proto.Void{})
	if err != nil {
		return "", "", err
	}
	return res.Version, res.Revision, nil
}

func (c *commonRemote) Modprobe(ctx context.Context, module string) error {
	_, err := c.client.Modprobe(ctx, &proto.ModprobeRequest{Module: module})
	return err
}
