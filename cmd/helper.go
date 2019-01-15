package cmd

import (
	"fmt"
	"os"
)

func ExitOnError(msg string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, msg+": %v\n", err)
		os.Exit(1)
	}
}
