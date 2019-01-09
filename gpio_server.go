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
	log.Printf("GPIO.Input(%d)\n", pin.Pin)
	return &proto.Void{}, s.gpio.Input(ctx, Pin(pin.Pin))
}

func (s *GpioServerImpl) Output(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	log.Printf("GPIO.Output(%d)\n", pin.Pin)
	return &proto.Void{}, s.gpio.Output(ctx, Pin(pin.Pin))
}

func (s *GpioServerImpl) Clock(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	log.Printf("GPIO.Clock(%d)\n", pin.Pin)
	return &proto.Void{}, s.gpio.Clock(ctx, Pin(pin.Pin))
}

func (s *GpioServerImpl) Pwm(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	log.Printf("GPIO.Pwm(%d)\n", pin.Pin)
	return &proto.Void{}, s.gpio.Pwm(ctx, Pin(pin.Pin))
}

func (s *GpioServerImpl) PullUp(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	log.Printf("GPIO.PullUp(%d)\n", pin.Pin)
	return &proto.Void{}, s.gpio.PullUp(ctx, Pin(pin.Pin))
}

func (s *GpioServerImpl) PullDown(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	log.Printf("GPIO.PullDown(%d)\n", pin.Pin)
	return &proto.Void{}, s.gpio.PullDown(ctx, Pin(pin.Pin))
}

func (s *GpioServerImpl) PullOff(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	log.Printf("GPIO.PullOff(%d)\n", pin.Pin)
	return &proto.Void{}, s.gpio.PullOff(ctx, Pin(pin.Pin))
}

func (s *GpioServerImpl) High(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	log.Printf("GPIO.High(%d)\n", pin.Pin)
	return &proto.Void{}, s.gpio.High(ctx, Pin(pin.Pin))
}

func (s *GpioServerImpl) Low(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	log.Printf("GPIO.Low(%d)\n", pin.Pin)
	return &proto.Void{}, s.gpio.Low(ctx, Pin(pin.Pin))
}

func (s *GpioServerImpl) Toggle(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	log.Printf("GPIO.Toggle(%d)\n", pin.Pin)
	return &proto.Void{}, s.gpio.Toggle(ctx, Pin(pin.Pin))
}

func (s *GpioServerImpl) Write(ctx context.Context, req *proto.RequestWrite) (*proto.Void, error) {
	log.Printf("GPIO.Write(%d, %d)\n", req.Pin, req.State)
	return &proto.Void{}, s.gpio.Write(ctx, Pin(req.Pin), PinState(req.State))
}

func (s *GpioServerImpl) Read(ctx context.Context, pin *proto.GpioPin) (*proto.ResponseRead, error) {
	log.Printf("GPIO.Read(%d)\n", pin.Pin)
	res, err := s.gpio.Read(ctx, Pin(pin.Pin))
	if err != nil {
		return nil, err
	}
	return &proto.ResponseRead{State: int32(res)}, nil
}

func (s *GpioServerImpl) Freq(ctx context.Context, req *proto.RequestFreq) (*proto.Void, error) {
	log.Printf("GPIO.Freq(%d, %d)\n", req.Pin, req.Freq)
	return &proto.Void{}, s.gpio.Freq(ctx, Pin(req.Pin), req.Freq)
}

func (s *GpioServerImpl) DutyCycle(ctx context.Context, req *proto.RequestDutyCycle) (*proto.Void, error) {
	log.Printf("GPIO.DutyCycle(%d, %d, %d)\n", req.Pin, req.DutyLen, req.CycleLen)
	return &proto.Void{}, s.gpio.DutyCycle(ctx, Pin(req.Pin), req.DutyLen, req.CycleLen)
}

func (s *GpioServerImpl) Detect(ctx context.Context, req *proto.RequestEdgeDetect) (*proto.Void, error) {
	log.Printf("GPIO.Detect(%d, %d)\n", req.Pin, req.Edge)
	return &proto.Void{}, s.gpio.Detect(ctx, Pin(req.Pin), PinEdge(req.Edge))
}

func (s *GpioServerImpl) EdgeDetected(ctx context.Context, pin *proto.GpioPin) (*proto.ResponseEdgeDetected, error) {
	log.Printf("GPIO.EdgeDetected(%d)\n", pin.Pin)
	res, err := s.gpio.EdgeDetected(ctx, Pin(pin.Pin))
	if err != nil {
		return nil, err
	}
	return &proto.ResponseEdgeDetected{Detected: res}, nil
}
