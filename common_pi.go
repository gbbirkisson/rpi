// +build pi

package rpi

import (
	"context"
	"fmt"
	"os/exec"
)

func (c *Common) Modprobe(ctx context.Context, mod string) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	output, err := exec.Command("modprobe", mod).CombinedOutput()
	if err != nil {
		return fmt.Errorf("%v: %s", err, output)
	}
	return nil
}

func (c *Common) GetVersion(ctx context.Context) (string, string, error) {
	if ctx.Err() != nil {
		return "", "", ctx.Err()
	}
	ver, res := GetLocalVersion()
	return ver, res, nil
}
