package rpi

import (
	"context"

	rpio "github.com/gbbirkisson/go-rpio"
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

type Common interface {
	GetVersion(ctx context.Context) (string, string, error)
	Modprobe(ctx context.Context, mod string) error
}

type PiCam interface {
	Open(ctx context.Context) error
	Close(ctx context.Context) error
	GetFrame(ctx context.Context) ([]byte, error)
	GetFrames(ctx context.Context, byteCh chan<- []byte, errCh chan<- error) (<-chan struct{}, error)
}

type Gpio interface {
	Open(ctx context.Context) error
	Close(ctx context.Context) error

	Input(ctx context.Context, pin Pin) error
	Output(ctx context.Context, pin Pin) error
	Clock(ctx context.Context, pin Pin) error
	Pwm(ctx context.Context, pin Pin) error
	PullUp(ctx context.Context, pin Pin) error
	PullDown(ctx context.Context, pin Pin) error
	PullOff(ctx context.Context, pin Pin) error

	High(ctx context.Context, pin Pin) error
	Low(ctx context.Context, pin Pin) error
	Toggle(ctx context.Context, pin Pin) error
	Write(ctx context.Context, pin Pin, state PinState) error
	Read(ctx context.Context, pin Pin) (PinState, error)

	Freq(ctx context.Context, pin Pin, freq int32) error
	DutyCycle(ctx context.Context, pin Pin, dutyLen, cycleLen int32) error
	Detect(ctx context.Context, pin Pin, edge PinEdge) error
	EdgeDetected(ctx context.Context, pin Pin) (bool, error)
}
