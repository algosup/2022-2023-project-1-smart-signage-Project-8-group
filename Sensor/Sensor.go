package main

import (
	"machine"
	// "strconv"
	"time"
)

var (
	lightSensorValue float32
	highVoltage      float32
	lowVoltage       uint16
)

func main() {
	machine.InitPWM()
	machine.InitADC()
	lS := machine.ADC{machine.ADC0} // A2 is the pin for the light sensor A2
	hV := machine.ADC{machine.ADC4} // A4 is the pin for the high voltage sensor A4
	lV := machine.ADC{machine.ADC5} // A5 is the pin for the low voltage sensor A5
	led := machine.PWM{machine.D3}  // D3 is the pin for PWM ~3
	led.Configure()
	lS.Configure(machine.ADCConfig{})
	hV.Configure(machine.ADCConfig{})
	lV.Configure(machine.ADCConfig{})

	// //goroutine that calls SendMessage every 15 minutes to send the data to the gateway with payload = "lightSensorValue,highVoltage,lowVoltage"
	// go func() {
	//  InitAT()
	// 	for range time.Tick(15 * time.Minute) {
	// 		//convert the values lightSensorValue, highVoltage and lowVoltage to string
	// 		str := (strconv.FormatFloat(float64(lightSensorValue), 'f', 2, 32) + "," + strconv.FormatFloat(float64(highVoltage), 'f', 2, 32) + "," + strconv.FormatFloat(float64(lowVoltage), 'f', 2, 32))
	// 		SendMessage(str) //send the message to the gateway
	// 	}
	// }()

	//main loop
	for {
		lightSensorValue = ADCSensorConv(ADCSensor(lS)) // Read the light sensor
		led.Set(changeLight(lightSensorValue))          // Change the LED brightness based on the light sensor value
		highVoltage = ADCSensorConv(ADCSensor(hV))      // Read the high voltage sensor
		lowVoltage = ADCSensor(lV)                      // Read the low voltage sensor
		println("Light sensor value: ", lightSensorValue)
		time.Sleep(time.Second) // Wait before acting again
	}
}

// Handle the ADC sensors and return the value in volts (float32)
func ADCSensorConv(fl uint16) float32 {
	return (float32(fl) / 65535.0) * 5.0
}

// Handle the ADC sensors and return the value in volts (float32)
func ADCSensor(adc machine.ADC) uint16 {
	return adc.Get()
}

// Handle the PWM LED and set the brightness based on the light sensor value
func changeLight(inLight float32) uint16 {
	if inLight > 2.7 {
		return (uint16(2000)) // Turn off the LED
	} else if inLight > 2.9 {
		return (uint16(12000)) // Turn on the LED at 18% brightness
	} else if inLight > 3.2 && inLight < 4 {
		return (uint16(42000)) // Turn on the LED at 64% brightness
	} else {
		return (uint16(3000)) // Turn on the LED with little brightness
	}
}
