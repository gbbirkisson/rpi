// +build !pi

package rpi

import (
	"context"

	proto "github.com/gbbirkisson/rpi/proto"
)

func (g *GPIO) Open(ctx context.Context) error {
	_, err := g.Client.Open(ctx, &proto.Void{})
	return err
}

func (g *GPIO) Close(ctx context.Context) error {
	_, err := g.Client.Close(ctx, &proto.Void{})
	return err
}

func (g *GPIO) Input(ctx context.Context, pin Pin) error {
	_, err := g.Client.Input(ctx, &proto.GpioPin{Pin: int32(pin)})
	return err
}

func (g *GPIO) Output(ctx context.Context, pin Pin) error {
	_, err := g.Client.Output(ctx, &proto.GpioPin{Pin: int32(pin)})
	return err
}

func (g *GPIO) Clock(ctx context.Context, pin Pin) error {
	_, err := g.Client.Clock(ctx, &proto.GpioPin{Pin: int32(pin)})
	return err
}

func (g *GPIO) Pwm(ctx context.Context, pin Pin) error {
	_, err := g.Client.Pwm(ctx, &proto.GpioPin{Pin: int32(pin)})
	return err
}

func (g *GPIO) PullUp(ctx context.Context, pin Pin) error {
	_, err := g.Client.PullUp(ctx, &proto.GpioPin{Pin: int32(pin)})
	return err
}

func (g *GPIO) PullDown(ctx context.Context, pin Pin) error {
	_, err := g.Client.PullDown(ctx, &proto.GpioPin{Pin: int32(pin)})
	return err
}

func (g *GPIO) PullOff(ctx context.Context, pin Pin) error {
	_, err := g.Client.PullOff(ctx, &proto.GpioPin{Pin: int32(pin)})
	return err
}

func (g *GPIO) High(ctx context.Context, pin Pin) error {
	_, err := g.Client.High(ctx, &proto.GpioPin{Pin: int32(pin)})
	return err
}

func (g *GPIO) Low(ctx context.Context, pin Pin) error {
	_, err := g.Client.Low(ctx, &proto.GpioPin{Pin: int32(pin)})
	return err
}

func (g *GPIO) Toggle(ctx context.Context, pin Pin) error {
	_, err := g.Client.Toggle(ctx, &proto.GpioPin{Pin: int32(pin)})
	return err
}

func (g *GPIO) Write(ctx context.Context, pin Pin, state PinState) error {
	_, err := g.Client.Write(ctx, &proto.RequestWrite{Pin: int32(pin), State: int32(state)})
	return err
}

func (g *GPIO) Read(ctx context.Context, pin Pin) (PinState, error) {
	state, err := g.Client.Read(ctx, &proto.GpioPin{Pin: int32(pin)})
	if err != nil {
		return 255, err
	}
	return PinState(state.State), err
}

func (g *GPIO) Freq(ctx context.Context, pin Pin, freq int32) error {
	_, err := g.Client.Freq(ctx, &proto.RequestFreq{Pin: int32(pin), Freq: freq})
	return err
}

func (g *GPIO) DutyCycle(ctx context.Context, pin Pin, dutyLen, cycleLen int32) error {
	_, err := g.Client.DutyCycle(ctx, &proto.RequestDutyCycle{Pin: int32(pin), DutyLen: dutyLen, CycleLen: cycleLen})
	return err
}

func (g *GPIO) Detect(ctx context.Context, pin Pin, edge PinEdge) error {
	_, err := g.Client.Detect(ctx, &proto.RequestEdgeDetect{Pin: int32(pin), Edge: int32(edge)})
	return err
}

func (g *GPIO) EdgeDetected(ctx context.Context, pin Pin) (bool, error) {
	res, err := g.Client.EdgeDetected(ctx, &proto.GpioPin{Pin: int32(pin)})
	if err != nil {
		return false, err
	}
	return res.Detected, nil
}
