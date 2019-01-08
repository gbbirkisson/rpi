package rpi

import (
	proto "github.com/gbbirkisson/rpi/proto"
	rpio "github.com/stianeikeland/go-rpio"
)

type Pin = rpio.Pin

const (
	Input  = rpio.Input
	Output = rpio.Output
)

type GPIO struct {
	Client proto.GpioClient
}
