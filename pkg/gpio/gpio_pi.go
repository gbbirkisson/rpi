// +build pi

package rpi

import (
	"context"

	rpio "github.com/gbbirkisson/go-rpio"
	proto "github.com/gbbirkisson/rpi/proto"
)

type Gpio struct {
	Client proto.GpioClient
}

func (g *Gpio) Validate() error {
	return nil
}

func (g *Gpio) Open(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	return rpio.Open()
}

func (g *Gpio) Close(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	return rpio.Close()
}

func (g *Gpio) Input(ctx context.Context, pin Pin) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).Input()
	})
}

func (g *Gpio) Output(ctx context.Context, pin Pin) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).Output()
	})
}

func (g *Gpio) Clock(ctx context.Context, pin Pin) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).Clock()
	})
}

func (g *Gpio) Pwm(ctx context.Context, pin Pin) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).Pwm()
	})
}

func (g *Gpio) PullUp(ctx context.Context, pin Pin) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).PullUp()
	})
}

func (g *Gpio) PullDown(ctx context.Context, pin Pin) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).PullDown()
	})
}

func (g *Gpio) PullOff(ctx context.Context, pin Pin) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).PullOff()
	})
}

func (g *Gpio) High(ctx context.Context, pin Pin) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).High()
	})
}

func (g *Gpio) Low(ctx context.Context, pin Pin) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).Low()
	})
}

func (g *Gpio) Toggle(ctx context.Context, pin Pin) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).Toggle()
	})
}

func (g *Gpio) Write(ctx context.Context, pin Pin, state PinState) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).Write(state)
	})
}

func (g *Gpio) Read(ctx context.Context, pin Pin) (PinState, error) {
	if ctx.Err() != nil {
		return 255, ctx.Err()
	}
	return rpio.Pin(pin).Read(), nil
}

func (g *Gpio) Freq(ctx context.Context, pin Pin, freq int32) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).Freq(int(freq))
	})
}

func (g *Gpio) DutyCycle(ctx context.Context, pin Pin, dutyLen, cycleLen int32) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).DutyCycle(uint32(dutyLen), uint32(cycleLen))
	})
}

func (g *Gpio) Detect(ctx context.Context, pin Pin, edge PinEdge) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).Detect(edge)
	})
}

func (g *Gpio) EdgeDetected(ctx context.Context, pin Pin) (bool, error) {
	if ctx.Err() != nil {
		return false, ctx.Err()
	}
	return rpio.Pin(pin).EdgeDetected(), nil
}

func checkContextAndExec(ctx context.Context, f func()) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	f()
	return nil
}
