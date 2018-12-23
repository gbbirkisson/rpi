package rpi

import (
	context "context"
	"log"

	rpi "github.com/gbbirkisson/rpi/proto"
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
)

type GpioServerImpl struct{}

func GetGPIOServer() *GpioServerImpl {
	driver, err := embd.DescribeHost()

	if err != nil {
		log.Fatalf("could not get pinmap: %v", err)
	}

	log.Printf("gpio driver")
	for _, pin := range driver.GPIODriver().PinMap() {
		log.Printf("- %+v\n", pin)
	}

	return &GpioServerImpl{}
}

func (s *GpioServerImpl) Init(ctx context.Context, void *rpi.Void) (*rpi.Void, error) {
	log.Printf("GPIO.Init()\n")
	return &rpi.Void{}, embd.InitGPIO()
}

func (s *GpioServerImpl) Close(ctx context.Context, void *rpi.Void) (*rpi.Void, error) {
	log.Printf("GPIO.Close()\n")
	return &rpi.Void{}, embd.CloseGPIO()
}

func (s *GpioServerImpl) Pins(ctx context.Context, void *rpi.Void) (*rpi.PinsRes, error) {
	driver, err := embd.DescribeHost()
	res := rpi.PinsRes{}

	if err != nil {
		return nil, err
	}

	for _, pin := range driver.GPIODriver().PinMap() {
		res.Info = append(res.Info, &rpi.PinInfo{Id: pin.ID, Alias: pin.Aliases, Caps: int32(pin.Caps), DigitalLogical: int32(pin.DigitalLogical), AnalogLogical: int32(pin.AnalogLogical)})
	}
	return &res, nil
}

func (s *GpioServerImpl) SetDirection(ctx context.Context, req *rpi.SetDirectionReq) (*rpi.Void, error) {
	pin := req.Pin
	dir := embd.Direction(int(req.Direction))
	log.Printf("GPIO.SetDirection(%s,%d)\n", pin, dir)
	return &rpi.Void{}, embd.SetDirection(req.Pin, dir)
}

func (s *GpioServerImpl) DigitalWrite(ctx context.Context, req *rpi.DigitalWriteReq) (*rpi.Void, error) {
	pin := req.Pin
	val := int(req.Value)
	log.Printf("GPIO.DigitalWrite(%s,%d)\n", pin, val)
	return &rpi.Void{}, embd.DigitalWrite(req.Pin, int(req.Value))
}

func (s *GpioServerImpl) DigitalRead(ctx context.Context, req *rpi.DigitalReadReq) (*rpi.DigitalReadRes, error) {
	pin := req.Pin
	log.Printf("GPIO.DigitalRead(%s)\n", pin)
	res, err := embd.DigitalRead(pin)
	if err != nil {
		return nil, err
	}
	return &rpi.DigitalReadRes{Value: int32(res)}, err
}

func (s *GpioServerImpl) Info(ctx context.Context, req *rpi.Void) (*rpi.InfoRes, error) {
	log.Printf("GPIO.Info()\n")
	return &rpi.InfoRes{In: int32(embd.In), Out: int32(embd.Out), High: int32(embd.High), Low: int32(embd.Low)}, nil
}
