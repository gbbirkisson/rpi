// Package rpi helps you to develop software for the Raspberry Pi that does
// IO operations on the Raspberry PI. It enables you to develop your code
// locally on any type of architecture, by using gRPC to control a Raspberry PI
// remotely. This is very convenient in conjunction with services like balena.io.
//
// This makes developing applications for the RaspberryPi extremely easy.
// Once your software is ready, you have the option of continuing to use
// gRPC calls, or switch over to a local version of the interfaces to compile
// a binary that runs directly on the RaspberryPi.
//
// If you have any suggestion or comments, please feel free to open an issue on
// the projects GitHub page.
package rpi

import (
	"context"
)

// Pin is the raw BCM2835 pinout of a GPIO pin
type Pin uint8

// PinEdge is edge events detection modes
type PinEdge uint8

// PinState is either high or low
type PinState uint8

// PinMode is the mode of the pin, Input, Output, Clock, Pwm or Spi
type PinMode uint8

const (
	// Input is the constant used to set a pin to input mode
	Input PinMode = iota
	// Output is the constant used to set a pin to output mode
	Output
)

const (
	// Low is the constant used to set a pin to low (0v)
	Low PinState = iota
	// High is the constant used to set a pin to high (+5v)
	High
)

// Common interface are the basic operations sometimes needed to use other the other interfaces
type Common interface {
	GetVersion(ctx context.Context) (string, string, error)
	Modprobe(ctx context.Context, mod string) error
}

// PiCam interface provides a way to fetch frames from a PiCam connected to a RaspberryPi
type PiCam interface {
	Open(ctx context.Context) error
	Close(ctx context.Context) error
	GetFrame(ctx context.Context) ([]byte, error)
	GetFrames(ctx context.Context, byteCh chan<- []byte, errCh chan<- error) (<-chan struct{}, error)
}

// Gpio interface provides a way to control and read from the GPIO pins on a RaspberryPi
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

// Ngrok interface provides a way to create an Ngrok tunnel to expose services to the internet
type Ngrok interface {
	Open(ctx context.Context) error
	Close(ctx context.Context) error
}
