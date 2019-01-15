// +build !pi

package rpi

import (
	"context"
	"fmt"

	proto "github.com/gbbirkisson/rpi/pkg/proto"
)

type Gpio struct {
	Client proto.GpioServiceClient
}

func (g *Gpio) Validate() error {
	if g.Client == nil {
		return fmt.Errorf("GpioClient cannot be nil")
	}
	return nil
}

func (g *Gpio) Open(ctx context.Context) error {
	_, err := g.Client.Open(ctx, &proto.Void{})
	return err
}

func (g *Gpio) Close(ctx context.Context) error {
	_, err := g.Client.Close(ctx, &proto.Void{})
	return err
}

func (g *Gpio) Input(ctx context.Context, pin Pin) error {
	_, err := g.Client.Input(ctx, &proto.GpioPin{Pin: int32(pin)})
	return err
}

func (g *Gpio) Output(ctx context.Context, pin Pin) error {
	_, err := g.Client.Output(ctx, &proto.GpioPin{Pin: int32(pin)})
	return err
}

func (g *Gpio) Clock(ctx context.Context, pin Pin) error {
	_, err := g.Client.Clock(ctx, &proto.GpioPin{Pin: int32(pin)})
	return err
}

func (g *Gpio) Pwm(ctx context.Context, pin Pin) error {
	_, err := g.Client.Pwm(ctx, &proto.GpioPin{Pin: int32(pin)})
	return err
}

func (g *Gpio) PullUp(ctx context.Context, pin Pin) error {
	_, err := g.Client.PullUp(ctx, &proto.GpioPin{Pin: int32(pin)})
	return err
}

func (g *Gpio) PullDown(ctx context.Context, pin Pin) error {
	_, err := g.Client.PullDown(ctx, &proto.GpioPin{Pin: int32(pin)})
	return err
}

func (g *Gpio) PullOff(ctx context.Context, pin Pin) error {
	_, err := g.Client.PullOff(ctx, &proto.GpioPin{Pin: int32(pin)})
	return err
}

func (g *Gpio) High(ctx context.Context, pin Pin) error {
	_, err := g.Client.High(ctx, &proto.GpioPin{Pin: int32(pin)})
	return err
}

func (g *Gpio) Low(ctx context.Context, pin Pin) error {
	_, err := g.Client.Low(ctx, &proto.GpioPin{Pin: int32(pin)})
	return err
}

func (g *Gpio) Toggle(ctx context.Context, pin Pin) error {
	_, err := g.Client.Toggle(ctx, &proto.GpioPin{Pin: int32(pin)})
	return err
}

func (g *Gpio) Write(ctx context.Context, pin Pin, state PinState) error {
	_, err := g.Client.Write(ctx, &proto.RequestWrite{Pin: int32(pin), State: int32(state)})
	return err
}

func (g *Gpio) Read(ctx context.Context, pin Pin) (PinState, error) {
	state, err := g.Client.Read(ctx, &proto.GpioPin{Pin: int32(pin)})
	if err != nil {
		return 255, err
	}
	return PinState(state.State), err
}

func (g *Gpio) Freq(ctx context.Context, pin Pin, freq int32) error {
	_, err := g.Client.Freq(ctx, &proto.RequestFreq{Pin: int32(pin), Freq: freq})
	return err
}

func (g *Gpio) DutyCycle(ctx context.Context, pin Pin, dutyLen, cycleLen int32) error {
	_, err := g.Client.DutyCycle(ctx, &proto.RequestDutyCycle{Pin: int32(pin), DutyLen: dutyLen, CycleLen: cycleLen})
	return err
}

func (g *Gpio) Detect(ctx context.Context, pin Pin, edge PinEdge) error {
	_, err := g.Client.Detect(ctx, &proto.RequestEdgeDetect{Pin: int32(pin), Edge: int32(edge)})
	return err
}

func (g *Gpio) EdgeDetected(ctx context.Context, pin Pin) (bool, error) {
	res, err := g.Client.EdgeDetected(ctx, &proto.GpioPin{Pin: int32(pin)})
	if err != nil {
		return false, err
	}
	return res.Detected, nil
}
