package main

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/ds1307"
	//rtclib
)

var (
	pr          machine.ADC   = machine.ADC{Pin: machine.ADC0}
	low         machine.ADC   = machine.ADC{Pin: machine.ADC3}
	high        machine.ADC   = machine.ADC{Pin: machine.ADC2}
	led         machine.PWM   = machine.PWM{Pin: machine.D3}
	timeMachine ds1307.Device = ds1307.New(machine.I2C0)

	LEDsUp         bool   = true
	LEDsUsage      uint8  = 0
	LEDsStatus     uint8  = 98
	LEDsBrightness uint8  = 80
	minBrightness  uint8  = 0
	maxBrightness  uint8  = 100
	up1            uint8  = 30
	up2            uint8  = 40
	down1          uint8  = 60
	down2          uint8  = 70
	upDays         uint8  = 20
	prValue        uint16 = 0
	lowValue       uint16 = 0
	highValue      uint16 = 0
	errors         uint8  = 0
	needsBackup    bool   = true
	dateTime       time.Time
)

func main() {
	machine.I2C0.Configure(machine.I2CConfig{})
	//set time  of timeMachine to current time
	timeMachine.SetTime(time.Now())
	machine.InitPWM()
	machine.InitADC()

	led.Configure()
	pr.Configure(machine.ADCConfig{Reference: 5, Samples: 8})
	low.Configure(machine.ADCConfig{Reference: 5, Samples: 8})
	high.Configure(machine.ADCConfig{Reference: 12, Samples: 8})

	for {
		loop()
	}
}

func loop() {
	addToLedUsage()
	// ReadDownlink()
	getSensorsValues()
	setLEDs()
	// SendUplink()

	//call reading function
	time.Sleep(time.Millisecond * 5000)
}

func getSensorsValues() {
	prValue = pr.Get()
	lowValue = low.Get()
	highValue = high.Get()

	errors = 0

	//use lib ds1307
	dateTime, _ = timeMachine.ReadTime()
}

func setLEDs() {
	LEDsBrightness = uint8(float32(prValue) / 65535.0 * 100.0)
	if LEDsBrightness > maxBrightness {
		LEDsBrightness = maxBrightness
	}

	if LEDsBrightness < minBrightness {
		LEDsBrightness = minBrightness
	}

	led.Set(uint16(float32(LEDsBrightness) / 100.0 * 65535.0))
}

func addToLedUsage() {
	if LEDsUp && errors == 0 {
		LEDsUsage = LEDsUsage*2 + 1
	} else {
		LEDsUsage = LEDsUsage * 2
	}
}

// print all values func
func printAllValues() {
	println("prValue: ", prValue)
	println("lowValue: ", lowValue)
	println("highValue: ", highValue)
	println("LEDsBrightness: ", LEDsBrightness)
	println("LEDsUsage: ", LEDsUsage)
	println("LEDsStatus: ", LEDsStatus)
	println("dateTime: ", dateTime.String())
}
