package rpi

import (
	"context"

	proto "github.com/gbbirkisson/rpi/pkg/proto"
)

// NewGpioServer creates a new gpio server that uses the gpio interface provided
func NewGpioServer(gpio Gpio) proto.GpioServer {
	return &gpioServer{gpio: gpio}
}

type gpioServer struct {
	gpio Gpio
}

func (s *gpioServer) Open(ctx context.Context, void *proto.Void) (*proto.Void, error) {
	return &proto.Void{}, s.gpio.Open(ctx)
}

func (s *gpioServer) Close(ctx context.Context, void *proto.Void) (*proto.Void, error) {
	return &proto.Void{}, s.gpio.Close(ctx)
}

func (s *gpioServer) Input(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	return &proto.Void{}, s.gpio.Input(ctx, Pin(pin.Pin))
}

func (s *gpioServer) Output(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	return &proto.Void{}, s.gpio.Output(ctx, Pin(pin.Pin))
}

func (s *gpioServer) Clock(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	return &proto.Void{}, s.gpio.Clock(ctx, Pin(pin.Pin))
}

func (s *gpioServer) Pwm(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	return &proto.Void{}, s.gpio.Pwm(ctx, Pin(pin.Pin))
}

func (s *gpioServer) PullUp(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	return &proto.Void{}, s.gpio.PullUp(ctx, Pin(pin.Pin))
}

func (s *gpioServer) PullDown(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	return &proto.Void{}, s.gpio.PullDown(ctx, Pin(pin.Pin))
}

func (s *gpioServer) PullOff(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	return &proto.Void{}, s.gpio.PullOff(ctx, Pin(pin.Pin))
}

func (s *gpioServer) High(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	return &proto.Void{}, s.gpio.High(ctx, Pin(pin.Pin))
}

func (s *gpioServer) Low(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	return &proto.Void{}, s.gpio.Low(ctx, Pin(pin.Pin))
}

func (s *gpioServer) Toggle(ctx context.Context, pin *proto.GpioPin) (*proto.Void, error) {
	return &proto.Void{}, s.gpio.Toggle(ctx, Pin(pin.Pin))
}

func (s *gpioServer) Write(ctx context.Context, req *proto.RequestWrite) (*proto.Void, error) {
	return &proto.Void{}, s.gpio.Write(ctx, Pin(req.Pin), PinState(req.State))
}

func (s *gpioServer) Read(ctx context.Context, pin *proto.GpioPin) (*proto.ResponseRead, error) {
	res, err := s.gpio.Read(ctx, Pin(pin.Pin))
	if err != nil {
		return nil, err
	}
	return &proto.ResponseRead{State: int32(res)}, nil
}

func (s *gpioServer) Freq(ctx context.Context, req *proto.RequestFreq) (*proto.Void, error) {
	return &proto.Void{}, s.gpio.Freq(ctx, Pin(req.Pin), req.Freq)
}

func (s *gpioServer) DutyCycle(ctx context.Context, req *proto.RequestDutyCycle) (*proto.Void, error) {
	return &proto.Void{}, s.gpio.DutyCycle(ctx, Pin(req.Pin), req.DutyLen, req.CycleLen)
}

func (s *gpioServer) Detect(ctx context.Context, req *proto.RequestEdgeDetect) (*proto.Void, error) {
	return &proto.Void{}, s.gpio.Detect(ctx, Pin(req.Pin), PinEdge(req.Edge))
}

func (s *gpioServer) EdgeDetected(ctx context.Context, pin *proto.GpioPin) (*proto.ResponseEdgeDetected, error) {
	res, err := s.gpio.EdgeDetected(ctx, Pin(pin.Pin))
	if err != nil {
		return nil, err
	}
	return &proto.ResponseEdgeDetected{Detected: res}, nil
}
