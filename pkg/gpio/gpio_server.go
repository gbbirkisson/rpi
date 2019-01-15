package rpi

import (
	"context"
	"log"

	proto "github.com/gbbirkisson/rpi/pkg/proto"
)

type GpioServer struct {
	Gpio *Gpio
}

func (s *GpioServer) Open(ctx context.Context, void *proto.Void) (*proto.Void, error) {
	log.Println("GPIO.Open()")
	return &proto.Void{}, s.Gpio.Open(ctx)
}

func (s *GpioServer) Close(ctx context.Context, void *proto.Void) (*proto.Void, error) {
	log.Println("GPIO.Close()")
	return &proto.Void{}, s.Gpio.Close(ctx)
}

func (s *GpioServer) Input(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	log.Printf("GPIO.Input(%d)\n", pin.Pin)
	return &proto.Void{}, s.Gpio.Input(ctx, Pin(pin.Pin))
}

func (s *GpioServer) Output(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	log.Printf("GPIO.Output(%d)\n", pin.Pin)
	return &proto.Void{}, s.Gpio.Output(ctx, Pin(pin.Pin))
}

func (s *GpioServer) Clock(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	log.Printf("GPIO.Clock(%d)\n", pin.Pin)
	return &proto.Void{}, s.Gpio.Clock(ctx, Pin(pin.Pin))
}

func (s *GpioServer) Pwm(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	log.Printf("GPIO.Pwm(%d)\n", pin.Pin)
	return &proto.Void{}, s.Gpio.Pwm(ctx, Pin(pin.Pin))
}

func (s *GpioServer) PullUp(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	log.Printf("GPIO.PullUp(%d)\n", pin.Pin)
	return &proto.Void{}, s.Gpio.PullUp(ctx, Pin(pin.Pin))
}

func (s *GpioServer) PullDown(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	log.Printf("GPIO.PullDown(%d)\n", pin.Pin)
	return &proto.Void{}, s.Gpio.PullDown(ctx, Pin(pin.Pin))
}

func (s *GpioServer) PullOff(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	log.Printf("GPIO.PullOff(%d)\n", pin.Pin)
	return &proto.Void{}, s.Gpio.PullOff(ctx, Pin(pin.Pin))
}

func (s *GpioServer) High(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	log.Printf("GPIO.High(%d)\n", pin.Pin)
	return &proto.Void{}, s.Gpio.High(ctx, Pin(pin.Pin))
}

func (s *GpioServer) Low(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	log.Printf("GPIO.Low(%d)\n", pin.Pin)
	return &proto.Void{}, s.Gpio.Low(ctx, Pin(pin.Pin))
}

func (s *GpioServer) Toggle(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	log.Printf("GPIO.Toggle(%d)\n", pin.Pin)
	return &proto.Void{}, s.Gpio.Toggle(ctx, Pin(pin.Pin))
}

func (s *GpioServer) Write(ctx context.Context, req *proto.RequestWrite) (*proto.Void, error) {
	log.Printf("GPIO.Write(%d, %d)\n", req.Pin, req.State)
	return &proto.Void{}, s.Gpio.Write(ctx, Pin(req.Pin), PinState(req.State))
}

func (s *GpioServer) Read(ctx context.Context, pin *proto.GpioPin) (*proto.ResponseRead, error) {
	log.Printf("GPIO.Read(%d)\n", pin.Pin)
	res, err := s.Gpio.Read(ctx, Pin(pin.Pin))
	if err != nil {
		return nil, err
	}
	return &proto.ResponseRead{State: int32(res)}, nil
}

func (s *GpioServer) Freq(ctx context.Context, req *proto.RequestFreq) (*proto.Void, error) {
	log.Printf("GPIO.Freq(%d, %d)\n", req.Pin, req.Freq)
	return &proto.Void{}, s.Gpio.Freq(ctx, Pin(req.Pin), req.Freq)
}

func (s *GpioServer) DutyCycle(ctx context.Context, req *proto.RequestDutyCycle) (*proto.Void, error) {
	log.Printf("GPIO.DutyCycle(%d, %d, %d)\n", req.Pin, req.DutyLen, req.CycleLen)
	return &proto.Void{}, s.Gpio.DutyCycle(ctx, Pin(req.Pin), req.DutyLen, req.CycleLen)
}

func (s *GpioServer) Detect(ctx context.Context, req *proto.RequestEdgeDetect) (*proto.Void, error) {
	log.Printf("GPIO.Detect(%d, %d)\n", req.Pin, req.Edge)
	return &proto.Void{}, s.Gpio.Detect(ctx, Pin(req.Pin), PinEdge(req.Edge))
}

func (s *GpioServer) EdgeDetected(ctx context.Context, pin *proto.GpioPin) (*proto.ResponseEdgeDetected, error) {
	log.Printf("GPIO.EdgeDetected(%d)\n", pin.Pin)
	res, err := s.Gpio.EdgeDetected(ctx, Pin(pin.Pin))
	if err != nil {
		return nil, err
	}
	return &proto.ResponseEdgeDetected{Detected: res}, nil
}
