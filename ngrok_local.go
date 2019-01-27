package rpi

import (
	"context"
	"fmt"
	"os"
	"os/exec"
)

// NewNgrokLocal creates a new local Ngrok interface
func NewNgrokLocal(tunnelType, port, token, region string) (Ngrok, error) {
	if tunnelType == "" {
		return nil, fmt.Errorf("tunnelType cannot be empty, use 'http' or 'tcp'")
	}
	if port == "" {
		return nil, fmt.Errorf("port cannot be empty")
	}
	if tunnelType == "tcp" && token == "" {
		return nil, fmt.Errorf("token cannot be empty if tunnel type is tcp")
	}
	if region == "" {
		return nil, fmt.Errorf("region cannot be empty, use 'us' or 'eu'")
	}
	return &nGrokLocal{tunnelType: tunnelType, port: port, token: token, region: region}, nil
}

type nGrokLocal struct {
	tunnelType, port, token, region string
	command                         *exec.Cmd
}

func (n *nGrokLocal) Open(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	if n.command != nil {
		return fmt.Errorf("ngrok already running")
	}
	n.command = exec.Command(
		"ngrok",
		n.tunnelType,
		n.port,
		"--authtoken",
		n.token,
		"--log=stdout",
		"--region",
		n.region,
	)

	n.command.Stdout = os.Stdout
	n.command.Stderr = os.Stderr

	return n.command.Start()
}

func (n *nGrokLocal) Close(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	if n.command == nil {
		return fmt.Errorf("ngrok not running")
	}
	err := n.command.Process.Signal(os.Kill)
	if err != nil {
		return fmt.Errorf("unable to send kill signal to ngrok process: %v", err)
	}
	n.command.Wait()
	n.command = nil
	return nil
}
