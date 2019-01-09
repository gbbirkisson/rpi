package rpi

import (
	proto "github.com/gbbirkisson/rpi/proto"
	rpio "github.com/stianeikeland/go-rpio"
)

type Pin = rpio.Pin
type PinEdge = rpio.Edge
type PinState = rpio.State

const (
	Input  = rpio.Input
	Output = rpio.Output
	Low    = rpio.Low
	High   = rpio.High
)

type GPIO struct {
	Client proto.GpioClient
}
