package cmd

import (
	"fmt"
	"os"
)

// ExitOnError prints the provided message and error, and then exits the program with status code 1
func ExitOnError(msg string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, msg+": %v\n", err)
		os.Exit(1)
	}
}
