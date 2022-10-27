package main

import (
	"machine"
	"strconv"
)

var (
	earlyStop   bool    = false
	switchFunc  bool    = true
	lSV         float32 // The value of the light sensor
	highVoltage float32 // The value of the high voltage sensor
	lowVoltage  float32 // The value of the high voltage sensor
	// The value of the low voltage sensor
	waitTime      uint16      = 15  // The time to wait before sending the next message
	maxBrightness uint8       = 100 // The maximum brightness of the LED
	minBrightness uint8       = 0   // The minimum brightness of the LED
	stop          bool        = false
	timeCounter   uint8       = 0
	timeMachine   Device      = New(machine.I2C0)
	led           machine.PWM = machine.PWM{Pin: machine.D3}   // D3 is the pin for PWM ~3
	lS            machine.ADC = machine.ADC{Pin: machine.ADC0} // A2 is the pin for the light sensor A2
)

func main() {
	machine.I2C0.Configure(machine.I2CConfig{Frequency: machine.TWI_FREQ_100KHZ})
	machine.InitPWM()
	machine.InitADC()
	hV := machine.ADC{Pin: machine.ADC2} // A4 is the pin for the high voltage sensor A4
	lV := machine.ADC{Pin: machine.ADC3} // A5 is the pin for the low voltage sensor A5

	led.Configure() // Configure the PWM LED
	//input mode
	lS.Configure(machine.ADCConfig{Reference: 5, Samples: 8})  // Configure the ADC light sensor
	hV.Configure(machine.ADCConfig{Reference: 12, Samples: 8}) // Configure the ADC high voltage sensor
	lV.Configure(machine.ADCConfig{Reference: 5, Samples: 8})  // Configure the ADC low voltage sensor

	//main loop
	for {
		if switchFunc && !stop {
			println("begin")
			mainProg(hV, lV)
			switchFunc = false
		} else {
			if stop {
				led.Set(0)
			}
			println("begin 2")
			ReadMessage(waitTime * 60)
			switchFunc = true
		}
	}
}

func mainProg(hV machine.ADC, lV machine.ADC) {
	lightSensorValue := changeLight(lS.Get(), led) // Get the value of the light sensor

	highVoltage = float32(hV.Get()) / 65535 * 100 // Read the high voltage sensor
	lowVoltage = float32(lV.Get()) / 65535 * 100  // Read the low voltage sensor

	//uint to hex
	str := strconv.FormatUint(uint64(float32(lightSensorValue)), 16) //create percentage of light sensor value and return it in string hex format
	if len(str) == 1 {
		str = "0" + str
	}

	if lightSensorValue != 0 {
		timeCounter++
		if timeCounter == 255 {
			timeCounter = 0
		}
		str += strconv.FormatUint(uint64(timeCounter), 16)
	}
	if len(str) == 3 {
		str = str[:2] + "0" + str[2:]
	}

	// deadLeds := (2.5 - lowVoltage) / 0.12
	//fake temporary data percentage added in str
	// str += strconv.FormatUint(uint64(deadLeds), 16)
	str += strconv.FormatUint(uint64(54), 16)
	if len(str) == 5 {
		str = str[:4] + "0" + str[4:]
	}

	//add 0x00 in string to the end of the string
	str += strconv.FormatUint(uint64(0), 16)
	if len(str) == 7 {
		str = str[:6] + "0" + str[6:]
	}

	InitAT() //make sure the AT module is ready

	SendMessage(str) //send the message to the gateway
}

// Handle the PWM LED and set the brightness based on the light sensor value
func changeLight(light uint16, led machine.PWM) uint8 {
	println("light: ", light)
	LEDsBrightness := uint8(float32(65535-light) / 65535 * 100) // Get the percentage of the light sensor value

	if LEDsBrightness > maxBrightness {
		LEDsBrightness = maxBrightness
	}

	if LEDsBrightness < minBrightness {
		LEDsBrightness = minBrightness
	}
	println("done")
	led.Set(uint16(float32(LEDsBrightness) / 100.0 * 65535.0))
	println("LEDsBrightness: ", LEDsBrightness)
	return LEDsBrightness
}
