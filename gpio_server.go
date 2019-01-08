package rpi

import (
	"context"
	"log"

	proto "github.com/gbbirkisson/rpi/proto"
)

type GpioServerImpl struct {
	gpio GPIO
}

func (s *GpioServerImpl) Open(ctx context.Context, void *proto.Void) (*proto.Void, error) {
	log.Println("GPIO.Open()")
	return &proto.Void{}, s.gpio.Open(ctx)
}

func (s *GpioServerImpl) Close(ctx context.Context, void *proto.Void) (*proto.Void, error) {
	log.Println("GPIO.Close()")
	return &proto.Void{}, s.gpio.Close(ctx)
}

func (s *GpioServerImpl) Input(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	log.Println("GPIO.Input()")
	return &proto.Void{}, s.gpio.Input(ctx, Pin(pin.Pin))
}

func (s *GpioServerImpl) Output(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	log.Println("GPIO.Output()")
	return &proto.Void{}, s.gpio.Output(ctx, Pin(pin.Pin))
}

func (s *GpioServerImpl) Clock(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	log.Println("GPIO.Clock()")
	return &proto.Void{}, s.gpio.Clock(ctx, Pin(pin.Pin))
}

func (s *GpioServerImpl) Pwm(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	log.Println("GPIO.Pwm()")
	return &proto.Void{}, s.gpio.Pwm(ctx, Pin(pin.Pin))
}

func (s *GpioServerImpl) High(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	log.Println("GPIO.High()")
	return &proto.Void{}, s.gpio.High(ctx, Pin(pin.Pin))
}

func (s *GpioServerImpl) Low(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	log.Println("GPIO.Low()")
	return &proto.Void{}, s.gpio.Low(ctx, Pin(pin.Pin))
}

func (s *GpioServerImpl) Toggle(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	log.Println("GPIO.Toggle()")
	return &proto.Void{}, s.gpio.Toggle(ctx, Pin(pin.Pin))
}
