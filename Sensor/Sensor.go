package main

import (
	"machine"
	"strconv"
)

var (
	switchFunc    bool    = true
	lSV           float32       // The value of the light sensor
	highVoltage   float32       // The value of the high voltage sensor
	lowVoltage    float32       // The value of the low voltage sensor
	waitTime      int8    = 15  // The time to wait before sending the next message
	maxBrightness int8    = 100 // The maximum brightness of the LED
	minBrightness int8    = 0   // The minimum brightness of the LED
	stop          bool    = false
	timeCounter   uint8   = 0
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
		if switchFunc && !stop {
			println("begin")
			mainProg(led, lS, hV, lV)
			switchFunc = false
		} else {
			if stop {
				led.Set(0)
			}
			println("begin 2")
			ReadMessage(waitTime + 1)
			switchFunc = true
		}
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
	val := uint16(((5 - inLight) * 20) * 65535 / 100)
	if val > uint16(float32(65535)/100*float32(maxBrightness)) {
		return uint16(float32(65535) / 100 * float32(maxBrightness))
	} else if val < uint16(float32(65535)/100*float32(minBrightness)) {
		return uint16(float32(65535) / 100 * float32(minBrightness))
	}
	return val
}

func mainProg(led machine.PWM, lS machine.ADC, hV machine.ADC, lV machine.ADC) {
	lightSensorValue := changeLight(ADCSensor(lS)) // Get the value of the light sensor
	led.Set(lightSensorValue)                      // Change the LED brightness based on the light sensor value

	highVoltage = ADCSensor(hV) // Read the high voltage sensor
	lowVoltage = ADCSensor(lV)  // Read the low voltage sensor

	InitAT() //make sure the AT module is ready

	//uint to hex
	str := strconv.FormatUint(uint64(float32(lightSensorValue)/65535*100), 16) //create percentage of light sensor value and return it in string hex format

	if lightSensorValue != 0 {
		timeCounter++
		if timeCounter == 255 {
			timeCounter = 0
		}
		str += strconv.FormatUint(uint64(timeCounter), 16)
	}

	//fake temporary data percentage added in str
	str += strconv.FormatUint(uint64(54), 16)

	//add 0x00 in string to the end of the string
	str += strconv.FormatUint(uint64(0), 16)
	println("str: ", str)
	SendMessage(str) //send the message to the gateway
}

// string to hex
func strToHex(str string) string {
	var hex string
	for _, char := range str {
		hex += string(rune(uint16(char)))
	}
	return hex
}
