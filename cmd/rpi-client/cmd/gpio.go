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

var gpioInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Print GPIO info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(layout)
	},
}

const layout = `The server uses raw BCM2835 pinouts, not the ports as they are mapped on the output pins for the
raspberry pi, and not the wiringPi convention.

Modes:
* pwm is possible only for pins 12, 13, 18, 19.
* clock is possible only for pins 4, 5, 6, 20, 21.

Note that some pins share common pwm channel, so calling this function will set same duty cycle
for all pins belonging to channel:
* channel 1 (pwm0) for pins 12, 18, 40
* channel 2 (pwm1) for pins 13, 19, 41, 45.

Note that some pins share the same clock source, it means that changing frequency for one pin will 
change it also for all pins within a group:
* gp_clk0: pins 4, 20, 32, 34
* gp_clk1: pins 5, 21, 42, 44
* gp_clk2: pins 6 and 43
* pwm_clk: pins 12, 13, 18, 19, 40, 41, 45

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
	Short: "Open and memory map GPIO memory",
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
	Short: "Unmap GPIO memory",
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
	Short: "Toggle pin low/high",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("command requires [pin] arguments")
		}
		_, err := strconv.Atoi(args[0])
		if err != nil {
			return errors.New("argument [pin] is not a valid int")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		pin, _ := strconv.Atoi(args[0])

		ctx, cancel := getContext()
		defer cancel()

		gpio, err := getGpio()
		rpi.ExitOnError("could not create client", err)
		rpi.ExitOnError("error repsonse from server", gpio.Toggle(ctx, rpi.Pin(pin)))
	},
}

var gpioHighCmd = &cobra.Command{
	Use:   "high [pin]",
	Short: "Set pin to high",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("command requires [pin] arguments")
		}
		_, err := strconv.Atoi(args[0])
		if err != nil {
			return errors.New("argument [pin] is not a valid int")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		pin, _ := strconv.Atoi(args[0])

		ctx, cancel := getContext()
		defer cancel()

		gpio, err := getGpio()
		rpi.ExitOnError("could not create client", err)
		rpi.ExitOnError("error repsonse from server", gpio.High(ctx, rpi.Pin(pin)))
	},
}

var gpioLowCmd = &cobra.Command{
	Use:   "low [pin]",
	Short: "Set pin to low",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("command requires [pin] arguments")
		}
		_, err := strconv.Atoi(args[0])
		if err != nil {
			return errors.New("argument [pin] is not a valid int")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		pin, _ := strconv.Atoi(args[0])

		ctx, cancel := getContext()
		defer cancel()

		gpio, err := getGpio()
		rpi.ExitOnError("could not create client", err)
		rpi.ExitOnError("error repsonse from server", gpio.Low(ctx, rpi.Pin(pin)))
	},
}

var gpioModeCmd = &cobra.Command{
	Use:   "mode [pin] [mode]",
	Short: "Set pin mode",
	Long:  "Valid mode arguments are input, output, clock, pwm, pullup, pulldown and pulloff",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("command requires [pin] and [mode] arguments")
		}
		_, err := strconv.Atoi(args[0])
		if err != nil {
			return errors.New("argument [pin] is not a valid int")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		pin, _ := strconv.Atoi(args[0])

		var setMode func(gpio *rpi.GPIO) error

		ctx, cancel := getContext()
		defer cancel()

		switch args[1] {
		case "input":
			setMode = func(gpio *rpi.GPIO) error {
				return gpio.Input(ctx, rpi.Pin(pin))
			}
		case "output":
			setMode = func(gpio *rpi.GPIO) error {
				return gpio.Output(ctx, rpi.Pin(pin))
			}
		case "clock":
			setMode = func(gpio *rpi.GPIO) error {
				return gpio.Clock(ctx, rpi.Pin(pin))
			}
		case "pwm":
			setMode = func(gpio *rpi.GPIO) error {
				return gpio.Pwm(ctx, rpi.Pin(pin))
			}
		case "pullup":
			setMode = func(gpio *rpi.GPIO) error {
				return gpio.PullUp(ctx, rpi.Pin(pin))
			}
		case "pulldown":
			setMode = func(gpio *rpi.GPIO) error {
				return gpio.PullDown(ctx, rpi.Pin(pin))
			}
		case "pulloff":
			setMode = func(gpio *rpi.GPIO) error {
				return gpio.PullOff(ctx, rpi.Pin(pin))
			}
		default:
			return errors.New("argument [mode] is not a valid mode")
		}

		gpio, err := getGpio()
		rpi.ExitOnError("could not create client", err)
		rpi.ExitOnError("could not create client", setMode(gpio))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(gpioCmd)
	gpioCmd.AddCommand(gpioInfoCmd)

	// Init
	gpioCmd.AddCommand(gpioOpenCmd)
	gpioCmd.AddCommand(gpioCloseCmd)

	// Pin mode
	gpioCmd.AddCommand(gpioModeCmd)

	// Basic pin operations
	gpioCmd.AddCommand(gpioToggleCmd)
	gpioCmd.AddCommand(gpioHighCmd)
	gpioCmd.AddCommand(gpioLowCmd)

}
