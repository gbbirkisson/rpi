package main

import (
	"context"
	"fmt"
	"os"
	"time"

	proto "github.com/gbbirkisson/rpi/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("10.116.32.168:8000", grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not connect to backend: %v\n", err)
		os.Exit(1)
	}
	rpiGPIO := proto.NewGPIOClient(conn)

	ctx := context.Background()

	_, err = rpiGPIO.Init(ctx, &proto.Void{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "could init gpio: %v\n", err)
		os.Exit(1)
	}
	defer func() {
		_, err := rpiGPIO.Close(ctx, &proto.Void{})
		if err != nil {
			fmt.Fprintf(os.Stderr, "could close gpio: %v\n", err)

		}
	}()

	_, err = rpiGPIO.SetDirection(ctx, &proto.SetDirectionReq{Pin: "GPIO_18", Direction: 1})
	if err != nil {
		fmt.Fprintf(os.Stderr, "could set pin direction: %v\n", err)
		return
	}

	_, err = rpiGPIO.DigitalWrite(ctx, &proto.DigitalWriteReq{Pin: "GPIO_18", Value: 1})
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not write to pin: %v\n", err)
		return
	}

	time.Sleep(5 * time.Second)
}
