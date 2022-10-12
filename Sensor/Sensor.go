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
		led.Set(uint16(3000))
		lightSensor(ldr)
		highVoltageSensor(hV)
		time.Sleep(time.Second)
	}
}

func lightSensor(ldr machine.ADC) {
	fl := float32(ldr.Get())
	println((fl / 65535.0) * 5.0)
}

func highVoltageSensor(hV machine.ADC) {
	fl := float32(hV.Get())
	println((fl / 65535.0) * 5.0)
}
