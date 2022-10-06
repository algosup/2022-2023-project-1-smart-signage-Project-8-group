package main

import (
	"machine"
	"time"
)

const led = machine.PA5 //test LED
const pa4 = machine.PA4 //LED band
const pb6 = machine.PB6 //High voltage sensor
const pb7 = machine.PB7 //Low Voltage sensor
const a3 = machine.A3

func main() {
	// test if flash is succesfull
	blinkLED()
}

//turn on/of LED band
func powerLED() {
	if pa4.Get() {
		pa4.Low()
	} else {
		pa4.High()
	}
}

func checkPowerStatusH() bool {
	return pb6.Get()
}

func checkPowerStatusL() bool {
	return pb7.Get()
}

func checkHeat() float32 {
	return -1.2
	//figure out how to get anything from the temp sensor
}

func checkLight() bool {
	return true
}

func blinkLED() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.Low()
		time.Sleep(time.Millisecond * 1000)

		led.High()
		time.Sleep(time.Millisecond * 1000)
	}
}
