package main

import (
	"machine"
	"strconv"
	"time"
)

var (
	lightSensorValue float32 // The value of the light sensor
	highVoltage      float32 // The value of the high voltage sensor
	lowVoltage       float32 // The value of the low voltage sensor
	// lowLight         uint8   = 10  // The value of the low light in percentage
	// mediumLight      uint8   = 50  // The value of the medium light in percentage
	// highLight        uint8   = 100 // The value of the high light in percentage
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
		lightSensorValue = ADCSensor(lS) // Read the light sensor
		highVoltage = ADCSensor(hV)      // Read the high voltage sensor
		lowVoltage = ADCSensor(lV)       // Read the low voltage sensor

		// get time of day, if night time between 01:00 and 06:00, set lights to 0%
		// if time.Now().Hour() >= 1 && time.Now().Hour() <= 6 {
		// 	led.Set(0) // Set the brightness of the LED to 0%
		// } else {
		led.Set(changeLight(lightSensorValue)) // Change the LED brightness based on the light sensor value
		// }

		temp := strconv.Itoa(int(lightSensorValue * 100))
		//put a dot before the last 2 digits
		temp = temp[:len(temp)-2] + "." + temp[len(temp)-2:]
		str := temp
		InitAT()
		SendMessage(str) //send the message to the gateway

		println(lowVoltage)         // Print the values to the serial monitor
		time.Sleep(time.Minute * 8) // Sleep for 15 minutes
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
	return uint16((5.0 / inLight) * 65535.0)
}
