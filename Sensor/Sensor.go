package main

import (
	"machine"
	"time"
)

func main() {

	machine.InitPWM()
	led := machine.PWM{machine.D11}
	led.Configure()
	machine.InitADC()
	ldr := machine.ADC{machine.ADC0}
	hV := machine.ADC{machine.ADC2}
	ldr.Configure(machine.ADCConfig{})
	for {
		A := lightSensor(ldr)
		changeLight(A, led)
		highVoltageSensor(hV)
		time.Sleep(time.Second)
	}
}

func lightSensor(ldr machine.ADC) float32 {
	fl := float32(ldr.Get())
	println((fl / 65535.0) * 5.0)
	return (fl / 65535.0) * 5.0
}

func highVoltageSensor(hV machine.ADC) float32 {
	fl := float32(hV.Get())
	println((fl / 65535.0) * 5.0)
	return (fl / 65535.0) * 5.0
}

func changeLight(inLight float32, led machine.PWM) {
	if inLight > 2.7 {
		led.Set(uint16(0))
	} else if inLight > 2.9 {
		led.Set(uint16(12000))
	} else if inLight > 3.2 && inLight < 4 {
		led.Set(uint16(42000))
	} else {
		led.Set(uint(3000))
	}
}
