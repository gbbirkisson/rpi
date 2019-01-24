package rpi

import (
	"context"

	rpio "github.com/gbbirkisson/go-rpio"
)

// Get a new local PiCam
func NewGpioLocal() Gpio {
	return &gpioLocal{}
}

type gpioLocal struct{}

func (g *gpioLocal) Open(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	return rpio.Open()
}

func (g *gpioLocal) Close(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	return rpio.Close()
}

func (g *gpioLocal) Input(ctx context.Context, pin Pin) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).Input()
	})
}

func (g *gpioLocal) Output(ctx context.Context, pin Pin) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).Output()
	})
}

func (g *gpioLocal) Clock(ctx context.Context, pin Pin) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).Clock()
	})
}

func (g *gpioLocal) Pwm(ctx context.Context, pin Pin) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).Pwm()
	})
}

func (g *gpioLocal) PullUp(ctx context.Context, pin Pin) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).PullUp()
	})
}

func (g *gpioLocal) PullDown(ctx context.Context, pin Pin) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).PullDown()
	})
}

func (g *gpioLocal) PullOff(ctx context.Context, pin Pin) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).PullOff()
	})
}

func (g *gpioLocal) High(ctx context.Context, pin Pin) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).High()
	})
}

func (g *gpioLocal) Low(ctx context.Context, pin Pin) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).Low()
	})
}

func (g *gpioLocal) Toggle(ctx context.Context, pin Pin) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).Toggle()
	})
}

func (g *gpioLocal) Write(ctx context.Context, pin Pin, state PinState) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).Write(state)
	})
}

func (g *gpioLocal) Read(ctx context.Context, pin Pin) (PinState, error) {
	if ctx.Err() != nil {
		return 255, ctx.Err()
	}
	return rpio.Pin(pin).Read(), nil
}

func (g *gpioLocal) Freq(ctx context.Context, pin Pin, freq int32) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).Freq(int(freq))
	})
}

func (g *gpioLocal) DutyCycle(ctx context.Context, pin Pin, dutyLen, cycleLen int32) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).DutyCycle(uint32(dutyLen), uint32(cycleLen))
	})
}

func (g *gpioLocal) Detect(ctx context.Context, pin Pin, edge PinEdge) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).Detect(edge)
	})
}

func (g *gpioLocal) EdgeDetected(ctx context.Context, pin Pin) (bool, error) {
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
