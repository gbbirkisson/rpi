package rpi_test

import (
	"log"

	"github.com/gbbirkisson/rpi"
)

func ExampleGetVersion() {
	version, revision := rpi.GetVersion()
	log.Printf("version: %s revision: %s", version, revision)
}
