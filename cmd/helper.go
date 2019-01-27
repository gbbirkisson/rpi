package cmd

import (
	"fmt"
	"os"
)

// Prints the message, error and exits the program with status code 1
func ExitOnError(msg string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, msg+": %v\n", err)
		os.Exit(1)
	}
}
