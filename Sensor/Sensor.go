package main

import (
	"machine"
	"time"
)

var (
	lSV         float32 // The value of the light sensor
	highVoltage float32 // The value of the high voltage sensor
	lowVoltage  float32 // The value of the low voltage sensor
)

func main() {
	machine.InitPWM()
	machine.InitADC()
	lS := machine.ADC{Pin: machine.ADC0} // A2 is the pin for the light sensor A2
	hV := machine.ADC{Pin: machine.ADC4} // A4 is the pin for the high voltage sensor A4
	lV := machine.ADC{Pin: machine.ADC5} // A5 is the pin for the low voltage sensor A5
	led := machine.PWM{Pin: machine.D3}  // D3 is the pin for PWM ~3

	led.Configure()                   // Configure the PWM LED
	lS.Configure(machine.ADCConfig{}) // Configure the ADC light sensor
	hV.Configure(machine.ADCConfig{}) // Configure the ADC high voltage sensor
	lV.Configure(machine.ADCConfig{}) // Configure the ADC low voltage sensor

	//main loop
	for {
		highVoltage = ADCSensor(hV) // Read the high voltage sensor
		lowVoltage = ADCSensor(lV)  // Read the low voltage sensor

		// get time of day, if night time between 01:00 and 06:00, set lights to 0%
		// if time.Now().Hour() >= 1 && time.Now().Hour() <= 6 {
		// 	led.Set(0) // Set the brightness of the LED to 0%
		// } else {
		lightSensorValue := changeLight(ADCSensor(lS))
		led.Set(lightSensorValue) // Change the LED brightness based on the light sensor value
		// }

		//create percentage of light sensor value and return it in rune
		str := string(int(float32(lightSensorValue) * 100 / 65535))

		InitAT() //make sure the AT module is ready

		SendMessage(str) //send the message to the gateway

		time.Sleep(time.Second * 1) // Sleep for 15 minutes
	}
}

// Handle the ADC sensors and return the value in volts (float32)
func ADCSensor(adc machine.ADC) float32 {
	ui := adc.Get()
	return (float32(ui) / 65535.0) * 5.0
}

// Handle the PWM LED and set the brightness based on the light sensor value
func changeLight(inLight float32) uint16 {
	//return percentage of inLight out of 65535
	return uint16((100 - (inLight * 20)) * 65535 * 0.01)
}
