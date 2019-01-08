package cmd

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/gbbirkisson/rpi"
	proto "github.com/gbbirkisson/rpi/proto"
	"github.com/spf13/cobra"
)

func getGpio() (*rpi.GPIO, error) {
	conn, err := getGrpcClient()
	if err != nil {
		return nil, err
	}
	return &rpi.GPIO{Client: proto.NewGpioClient(conn)}, nil
}

var gpioCmd = &cobra.Command{
	Use:   "gpio",
	Short: "Control the GPIO pins on the device",
}

var gpioLayoutCmd = &cobra.Command{
	Use:   "layout",
	Short: "Print pin layout",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(layout)
	},
}

const layout = `The server uses raw BCM2835 pinouts, not the ports as they are mapped on the output pins for the
raspberry pi, and not the wiringPi convention.

            Rev 2 and 3 Raspberry Pi                        Rev 1 Raspberry Pi (legacy)
  +-----+---------+----------+---------+-----+      +-----+--------+----------+--------+-----+
  | BCM |   Name  | Physical | Name    | BCM |      | BCM | Name   | Physical | Name   | BCM |
  +-----+---------+----++----+---------+-----+      +-----+--------+----++----+--------+-----+
  |     |    3.3v |  1 || 2  | 5v      |     |      |     | 3.3v   |  1 ||  2 | 5v     |     |
  |   2 |   SDA 1 |  3 || 4  | 5v      |     |      |   0 | SDA    |  3 ||  4 | 5v     |     |
  |   3 |   SCL 1 |  5 || 6  | 0v      |     |      |   1 | SCL    |  5 ||  6 | 0v     |     |
  |   4 | GPIO  7 |  7 || 8  | TxD     | 14  |      |   4 | GPIO 7 |  7 ||  8 | TxD    |  14 |
  |     |      0v |  9 || 10 | RxD     | 15  |      |     | 0v     |  9 || 10 | RxD    |  15 |
  |  17 | GPIO  0 | 11 || 12 | GPIO  1 | 18  |      |  17 | GPIO 0 | 11 || 12 | GPIO 1 |  18 |
  |  27 | GPIO  2 | 13 || 14 | 0v      |     |      |  21 | GPIO 2 | 13 || 14 | 0v     |     |
  |  22 | GPIO  3 | 15 || 16 | GPIO  4 | 23  |      |  22 | GPIO 3 | 15 || 16 | GPIO 4 |  23 |
  |     |    3.3v | 17 || 18 | GPIO  5 | 24  |      |     | 3.3v   | 17 || 18 | GPIO 5 |  24 |
  |  10 |    MOSI | 19 || 20 | 0v      |     |      |  10 | MOSI   | 19 || 20 | 0v     |     |
  |   9 |    MISO | 21 || 22 | GPIO  6 | 25  |      |   9 | MISO   | 21 || 22 | GPIO 6 |  25 |
  |  11 |    SCLK | 23 || 24 | CE0     | 8   |      |  11 | SCLK   | 23 || 24 | CE0    |   8 |
  |     |      0v | 25 || 26 | CE1     | 7   |      |     | 0v     | 25 || 26 | CE1    |   7 |
  |   0 |   SDA 0 | 27 || 28 | SCL 0   | 1   |      +-----+--------+----++----+--------+-----+
  |   5 | GPIO 21 | 29 || 30 | 0v      |     |
  |   6 | GPIO 22 | 31 || 32 | GPIO 26 | 12  |
  |  13 | GPIO 23 | 33 || 34 | 0v      |     |
  |  19 | GPIO 24 | 35 || 36 | GPIO 27 | 16  |
  |  26 | GPIO 25 | 37 || 38 | GPIO 28 | 20  |
  |     |      0v | 39 || 40 | GPIO 29 | 21  |
  +-----+---------+----++----+---------+-----+`

var gpioOpenCmd = &cobra.Command{
	Use:   "open",
	Short: "Open GPIO interface",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := getContext()
		defer cancel()
		gpio, err := getGpio()
		rpi.ExitOnError("could not create client", err)
		rpi.ExitOnError("error repsonse from server", gpio.Open(ctx))
	},
}

var gpioCloseCmd = &cobra.Command{
	Use:   "close",
	Short: "Close GPIO interface",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := getContext()
		defer cancel()
		gpio, err := getGpio()
		rpi.ExitOnError("could not create client", err)
		rpi.ExitOnError("error repsonse from server", gpio.Close(ctx))
	},
}

var gpioToggleCmd = &cobra.Command{
	Use:   "toggle [pin]",
	Short: "Toggle pin on/off",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("command requires [pin] arguments")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		pin, err := strconv.Atoi(args[0])
		rpi.ExitOnError("invalid pin", err)
		ctx, cancel := getContext()
		defer cancel()
		gpio, err := getGpio()
		rpi.ExitOnError("could not create client", err)
		rpi.ExitOnError("error repsonse from server", gpio.Toggle(ctx, rpi.Pin(pin)))
	},
}

var gpioOutputCmd = &cobra.Command{
	Use:   "output [pin]",
	Short: "Set pin as output pin",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("command requires [pin] arguments")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		pin, err := strconv.Atoi(args[0])
		rpi.ExitOnError("invalid pin", err)

		ctx, cancel := getContext()
		defer cancel()
		gpio, err := getGpio()
		rpi.ExitOnError("could not create client", err)
		rpi.ExitOnError("error repsonse from server", gpio.Output(ctx, rpi.Pin(pin)))
	},
}

func init() {
	rootCmd.AddCommand(gpioCmd)
	gpioCmd.AddCommand(gpioLayoutCmd)
	gpioCmd.AddCommand(gpioOpenCmd)
	gpioCmd.AddCommand(gpioCloseCmd)
	gpioCmd.AddCommand(gpioToggleCmd)
	gpioCmd.AddCommand(gpioOutputCmd)
}
