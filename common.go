package rpi

import (
	"fmt"
	"os"
)

var Version string = "development"

func ExitOnError(err error, msg string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, msg+": %v\n", err)
		os.Exit(1)
	}
}
