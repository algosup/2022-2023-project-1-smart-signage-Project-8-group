package main

import (
	"encoding/hex"
	"machine"
	"strconv"
	"strings"
	"time"
)

var (
	earlyStop   bool    = false
	switchFunc  bool    = true
	lSV         float32 // The value of the light sensor
	highVoltage float32 // The value of the high voltage sensor
	lowVoltage  float32 // The value of the high voltage sensor
	// The value of the low voltage sensor
	waitTime      int8        = 15  // The time to wait before sending the next message
	maxBrightness uint8       = 100 // The maximum brightness of the LED
	minBrightness uint8       = 0   // The minimum brightness of the LED
	stop          bool        = false
	timeCounter   uint8       = 0
	timeMachine   Device      = New(machine.I2C0)
	led           machine.PWM = machine.PWM{Pin: machine.D3}   // D3 is the pin for PWM ~3
	lS            machine.ADC = machine.ADC{Pin: machine.ADC0} // A2 is the pin for the light sensor A2
)

func main() {
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
			ReadMessage(waitTime+1, 0)
			switchFunc = true
		}
	}
}

// Handle the PWM LED and set the brightness based on the light sensor value
func changeLight(light uint16, led machine.PWM) uint8 {
	LEDsBrightness := uint8(float32(light) / 65535 * 100) // Get the percentage of the light sensor value

	if LEDsBrightness > maxBrightness {
		LEDsBrightness = maxBrightness
	}

	if LEDsBrightness < minBrightness {
		LEDsBrightness = minBrightness
	}

	led.Set(uint16(float32(LEDsBrightness) / 100.0 * 65535.0))
	return LEDsBrightness
}

func mainProg(hV machine.ADC, lV machine.ADC) {
	lightSensorValue := changeLight(lS.Get(), led) // Get the value of the light sensor

	highVoltage = float32(hV.Get()) / 65535 * 100 // Read the high voltage sensor
	lowVoltage = float32(lV.Get()) / 65535 * 100  // Read the low voltage sensor
	println("lowVoltage: ", lowVoltage)
	println("highVoltage: ", highVoltage)
	println("lightSensorValue: ", lightSensorValue)

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

	data := make([]byte, 8)
	timeMachine.bus.ReadRegister(timeMachine.Address, uint8(0x00), data)
	seconds := bcdToDec(data[0] & 0x7F)
	minute := bcdToDec(data[1])
	hour := hoursBCDToInt(data[2])
	day := bcdToDec(data[3])
	println("TIME;")
	println("Day : ", day, "Hour : ", hour, " Minute : ", minute, " Seconds : ", seconds)
	InitAT() //make sure the AT module is ready

	println("str: ", str)
	SendMessage(str) //send the message to the gateway
}

// Initialize the AT module
func InitAT() {
	println("Initializing AT...")
	time.Sleep(time.Millisecond * 50)
	machine.UART0.Configure(machine.UARTConfig{BaudRate: 9600, TX: machine.D1, RX: machine.D0})
	_, err := machine.UART0.Write([]byte("AT+JOIN=DR3\r\n"))
	if err != nil {
		println("Error: " + err.Error())
	}
	ReadMessage(0, 15)
}

// Send a message to the serial port of the lora module with the given payload
func SendMessage(payload string) {
	println("Sending message...")
	_, err := machine.UART0.Write([]byte(`AT+MSG= "` + payload + `"` + "\r\n"))
	if err != nil {
		println("Error: " + err.Error())
	}
}

// Read the serial port of the lora module and return the message
func ReadMessage(wT int8, wTS uint16) string {

	led.Configure()
	var msg string
	timer := 0
	msg1 := ""
	var timeCheck uint16
	if wT == 0 && wTS != 0 {
		timeCheck = wTS
	} else {
		timeCheck = uint16(wT) * 60
	}

	for {
		changeLight(lS.Get(), led) // Get the value of the light sensor                                        // Change the LED brightness based on the light sensor value
		if timer >= int(timeCheck)*1000 || earlyStop {
			return ""
		}
		if machine.UART0.Buffered() > 0 {
			rb, err := machine.UART0.ReadByte()
			if err != nil {
				println("Error: " + err.Error())
				continue
				// return ""
			}
			msg1 += string(rb)
			if msg1[len(msg1)-1] == '\n' {
				msg = msg1
				continue
				// return msg
			}
		} else {
			if msg != "" {
				if msg != "+AT: ERROR(-11)\r\n" && msg != "+AT: ERROR(-24)\r\n" {
					//if msg contains "G: PORT:"
					println(msg)
					if strings.Contains(msg, "RX: ") {
						//remove everything out of "" after RX:
						msg = msg[strings.Index(msg, "RX: ")+4:]
						msg = msg[:strings.Index(msg, "\"")]
						msgTreating(msg)
					}
				}
				msg = ""
				msg1 = ""
			}
			timer += 20
			time.Sleep(time.Millisecond * 20)
		}
	}
}

// handling received message
func msgTreating(msg string) {
	println("treating...")
	println(msg)
	if len(msg) < 4 {
		return
	}
	str := Hex2Bin(msg[0])
	//msg is 2 bytes long, take first byte from str and take the last bit if it's 0, turn off the led, if it's 1 turn on the led
	if str[7] == '0' {
		println("turn off")
		stop = true
	} else {
		println("turn on")
		stop = false
	}

	//take second bit and if true earlyStop = true
	if str[6] == '1' {
		println("earlyStop = false")
		earlyStop = false
	} else {
		println("earlyStop = true")
		earlyStop = true
	}
	bytearray := msg[2:]
	maxBrightness, minBrightness = bitsManager(hex.EncodeToString([]byte(bytearray)))

}

func bitsManager(num string) (uint8, uint8) {
	//take two first bytes num[0:1] and convert them to int8 without parseint
	a := int8(num[0])
	a2 := int8(num[1])
	a = a + a2*10 - 30
	b := int8(num[2])
	b2 := int8(num[3])
	b = b + b2*10 - 30

	if a < b {
		b = a
	}
	if b > 10 {
		b = 10
	}
	if a > 10 {
		a = 10
	}
	return uint8(a * 10), uint8(b * 10)
}

func Hex2Bin(in byte) string {
	var out []byte
	for i := 7; i >= 0; i-- {
		b := (in >> uint(i))
		out = append(out, (b%2)+48)
	}
	return string(out)
}
