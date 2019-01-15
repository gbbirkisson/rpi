package rpi

import (
	gpio "github.com/gbbirkisson/rpi/pkg/gpio"
	proto "github.com/gbbirkisson/rpi/pkg/proto"
	"google.golang.org/grpc"
)

func GetGpio(client proto.GpioServiceClient) (*gpio.Gpio, error) {
	g := gpio.Gpio{Client: client}
	err := g.Validate()
	if err != nil {
		return nil, err
	}
	return &g, nil
}

func GetGpioClient(conn *grpc.ClientConn) proto.GpioServiceClient {
	return proto.NewGpioServiceClient(conn)
}

func GetGpioServer(g *gpio.Gpio) (*gpio.GpioServer, error) {
	err := g.Validate()
	if err != nil {
		return nil, err
	}
	return &gpio.GpioServer{Gpio: g}, nil
}
