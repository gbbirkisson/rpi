// +build !pi

package rpi

import (
	"context"
	"fmt"

	proto "github.com/gbbirkisson/rpi/pkg/proto"
)

func (c *Common) Modprobe(ctx context.Context, mod string) error {
	cli, err := c.getClient(ctx)
	if err != nil {
		return fmt.Errorf("unable to get grpc client: %v", err)
	}
	_, err = cli.Modprobe(ctx, &proto.ModprobeRequest{Params: mod})
	return err
}

func (c *Common) GetVersion(ctx context.Context) (string, string, error) {
	cli, err := c.getClient(ctx)
	if err != nil {
		return "", "", fmt.Errorf("unable to get grpc client: %v", err)
	}
	res, err := cli.GetVersion(ctx, &proto.Void{})
	if err != nil {
		return "", "", err
	}
	return res.Version, res.Revision, nil
}

func (c *Common) getClient(ctx context.Context) (proto.CommonServiceClient, error) {
	if c.client != nil {
		return c.client, nil
	}
	if c.Connection == nil {
		return nil, fmt.Errorf("Common.Connection is nil")
	}
	c.client = proto.NewCommonServiceClient(c.Connection)
	return c.client, nil
}
