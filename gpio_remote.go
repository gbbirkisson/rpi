// +build !local

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
