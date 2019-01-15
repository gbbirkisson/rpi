package rpi

import rpio "github.com/gbbirkisson/go-rpio"

type Pin = rpio.Pin
type PinEdge = rpio.Edge
type PinState = rpio.State

const (
	Input  = rpio.Input
	Output = rpio.Output
	Low    = rpio.Low
	High   = rpio.High
)
