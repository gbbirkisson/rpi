package cmd

import (
	"context"

	proto "github.com/gbbirkisson/rpi/proto"
	"github.com/spf13/cobra"
)

func getGpioClientAndContext(cmd *cobra.Command) (proto.GPIOClient, context.Context) {
	client, ctx := getGrpcClientAndContext(cmd)
	return proto.NewGPIOClient(client), ctx
}

var gpioCmd = &cobra.Command{
	Use:   "gpio",
	Short: "Control the GPIO pins on the device",
}

func init() {
	rootCmd.AddCommand(gpioCmd)
}
