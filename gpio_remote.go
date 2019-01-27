package rpi

import (
	"context"

	proto "github.com/gbbirkisson/rpi/pkg/proto"
	"google.golang.org/grpc"
)

// NewGpioRemote creates a new Gpio interface that uses remote gpio pins
func NewGpioRemote(connection *grpc.ClientConn) Gpio {
	return &gpioRemote{client: proto.NewGpioClient(connection)}
}

type gpioRemote struct {
	client proto.GpioClient
}

func (g *gpioRemote) Open(ctx context.Context) error {
	_, err := g.client.Open(ctx, &proto.Void{})
	return err
}

func (g *gpioRemote) Close(ctx context.Context) error {
	_, err := g.client.Close(ctx, &proto.Void{})
	return err
}

func (g *gpioRemote) Input(ctx context.Context, pin Pin) error {
	_, err := g.client.Input(ctx, &proto.GpioPin{Pin: int32(pin)})
	return err
}

func (g *gpioRemote) Output(ctx context.Context, pin Pin) error {
	_, err := g.client.Output(ctx, &proto.GpioPin{Pin: int32(pin)})
	return err
}

func (g *gpioRemote) Clock(ctx context.Context, pin Pin) error {
	_, err := g.client.Clock(ctx, &proto.GpioPin{Pin: int32(pin)})
	return err
}

func (g *gpioRemote) Pwm(ctx context.Context, pin Pin) error {
	_, err := g.client.Pwm(ctx, &proto.GpioPin{Pin: int32(pin)})
	return err
}

func (g *gpioRemote) PullUp(ctx context.Context, pin Pin) error {
	_, err := g.client.PullUp(ctx, &proto.GpioPin{Pin: int32(pin)})
	return err
}

func (g *gpioRemote) PullDown(ctx context.Context, pin Pin) error {
	_, err := g.client.PullDown(ctx, &proto.GpioPin{Pin: int32(pin)})
	return err
}

func (g *gpioRemote) PullOff(ctx context.Context, pin Pin) error {
	_, err := g.client.PullOff(ctx, &proto.GpioPin{Pin: int32(pin)})
	return err
}

func (g *gpioRemote) High(ctx context.Context, pin Pin) error {
	_, err := g.client.High(ctx, &proto.GpioPin{Pin: int32(pin)})
	return err
}

func (g *gpioRemote) Low(ctx context.Context, pin Pin) error {
	_, err := g.client.Low(ctx, &proto.GpioPin{Pin: int32(pin)})
	return err
}

func (g *gpioRemote) Toggle(ctx context.Context, pin Pin) error {
	_, err := g.client.Toggle(ctx, &proto.GpioPin{Pin: int32(pin)})
	return err
}

func (g *gpioRemote) Write(ctx context.Context, pin Pin, state PinState) error {
	_, err := g.client.Write(ctx, &proto.RequestWrite{Pin: int32(pin), State: int32(state)})
	return err
}

func (g *gpioRemote) Read(ctx context.Context, pin Pin) (PinState, error) {
	state, err := g.client.Read(ctx, &proto.GpioPin{Pin: int32(pin)})
	if err != nil {
		return 255, err
	}
	return PinState(state.State), err
}

func (g *gpioRemote) Freq(ctx context.Context, pin Pin, freq int32) error {
	_, err := g.client.Freq(ctx, &proto.RequestFreq{Pin: int32(pin), Freq: freq})
	return err
}

func (g *gpioRemote) DutyCycle(ctx context.Context, pin Pin, dutyLen, cycleLen int32) error {
	_, err := g.client.DutyCycle(ctx, &proto.RequestDutyCycle{Pin: int32(pin), DutyLen: dutyLen, CycleLen: cycleLen})
	return err
}

func (g *gpioRemote) Detect(ctx context.Context, pin Pin, edge PinEdge) error {
	_, err := g.client.Detect(ctx, &proto.RequestEdgeDetect{Pin: int32(pin), Edge: int32(edge)})
	return err
}

func (g *gpioRemote) EdgeDetected(ctx context.Context, pin Pin) (bool, error) {
	res, err := g.client.EdgeDetected(ctx, &proto.GpioPin{Pin: int32(pin)})
	if err != nil {
		return false, err
	}
	return res.Detected, nil
}
