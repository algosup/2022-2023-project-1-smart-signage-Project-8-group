package main

import (
	"machine"
	"time"
)

func main() {
	pin11 := machine.D11
	pin11.Configure(machine.PinConfig{Mode: machine.PinOutput})
	highTime := 1
	lowTime := 1
	//main loop
	for {
		pin11.High()
		time.Sleep(time.Millisecond / time.Duration(highTime))
		pin11.Low()
		time.Sleep(time.Millisecond * time.Duration(lowTime))
	}
}
