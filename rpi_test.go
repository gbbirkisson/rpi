package rpi

import (
	"testing"

	rpio "github.com/gbbirkisson/go-rpio"
)

func TestConstants(t *testing.T) {
	tables := []struct {
		rpiConstant  int32
		rpioConstant int32
	}{
		{int32(Input), int32(rpio.Input)},
		{int32(Output), int32(rpio.Output)},
		{int32(Low), int32(rpio.Low)},
		{int32(High), int32(rpio.High)},
	}

	for _, table := range tables {
		if table.rpiConstant != table.rpioConstant {
			t.Errorf("rpi and rpio constant mismatch")
		}
	}
}
