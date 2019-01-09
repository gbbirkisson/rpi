// +build local

package rpi

import (
	"context"

	rpio "github.com/stianeikeland/go-rpio/v4"
)

func checkContextAndExec(ctx context.Context, f func()) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	f()
	return nil
}

func (g *GPIO) Open(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	return rpio.Open()
}

func (g *GPIO) Close(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	return rpio.Close()
}

func (g *GPIO) Input(ctx context.Context, pin Pin) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).Input()
	})
}

func (g *GPIO) Output(ctx context.Context, pin Pin) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).Output()
	})
}

func (g *GPIO) Clock(ctx context.Context, pin Pin) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).Clock()
	})
}

func (g *GPIO) Pwm(ctx context.Context, pin Pin) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).Pwm()
	})
}

func (g *GPIO) PullUp(ctx context.Context, pin Pin) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).PullUp()
	})
}

func (g *GPIO) PullDown(ctx context.Context, pin Pin) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).PullDown()
	})
}

func (g *GPIO) PullOff(ctx context.Context, pin Pin) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).PullOff()
	})
}

func (g *GPIO) High(ctx context.Context, pin Pin) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).High()
	})
}

func (g *GPIO) Low(ctx context.Context, pin Pin) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).Low()
	})
}

func (g *GPIO) Toggle(ctx context.Context, pin Pin) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).Toggle()
	})
}

func (g *GPIO) Write(ctx context.Context, pin Pin, state PinState) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).Write(state)
	})
}

func (g *GPIO) Read(ctx context.Context, pin Pin) (PinState, error) {
	if ctx.Err() != nil {
		return 255, ctx.Err()
	}
	return rpio.Pin(pin).Read(), nil
}

func (g *GPIO) Freq(ctx context.Context, pin Pin, freq int32) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).Freq(int(freq))
	})
}

func (g *GPIO) DutyCycle(ctx context.Context, pin Pin, dutyLen, cycleLen int32) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).DutyCycle(uint32(dutyLen), uint32(cycleLen))
	})
}

func (g *GPIO) Detect(ctx context.Context, pin Pin, edge PinEdge) error {
	return checkContextAndExec(ctx, func() {
		rpio.Pin(pin).Detect(edge)
	})
}

func (g *GPIO) EdgeDetected(ctx context.Context, pin Pin) (bool, error) {
	if ctx.Err() != nil {
		return false, ctx.Err()
	}
	return rpio.Pin(pin).EdgeDetected(), nil
}
