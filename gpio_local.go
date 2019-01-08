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
