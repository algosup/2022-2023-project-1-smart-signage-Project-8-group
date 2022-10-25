package main

import (
	"machine"
	"strings"
	"time"
)

var (
	earlyStop bool = false
)

// Initialize the AT module
func InitAT() {
	println("Initializing AT...")
	time.Sleep(time.Millisecond * 50)
	machine.UART0.Configure(machine.UARTConfig{BaudRate: 9600, TX: machine.D1, RX: machine.D0})
	_, err := machine.UART0.Write([]byte("AT+JOIN=DR3\r\n"))
	if err != nil {
		println("Error: " + err.Error())
	}
	ReadMessage(1)
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
func ReadMessage(wT int8) string {
	led := machine.PWM{Pin: machine.D3}  // D3 is the pin for PWM ~3
	lS := machine.ADC{Pin: machine.ADC0} // A2 is the pin for the light sensor A2

	led.Configure()
	var msg string
	timer := 0
	msg1 := ""
	for {
		lightSensorValue := changeLight(ADCSensor(lS)) // Get the value of the light sensor
		led.Set(lightSensorValue)                      // Change the LED brightness based on the light sensor value
		if timer >= int(wT)*60*1000 || earlyStop {
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

func msgTreating(msg string) {
	println("treating...")
	println(msg)
	//msg is 2 bytes long, take first byte and take the first bit if it's 0, turn off the led, if it's 1 turn on the led
	if (msg[0] & 0x80) == 0 {
		println("turn off")
		stop = true
	} else {
		println("turn on")
		stop = false
	}

	//take second bit and if true earlyStop = true
	if (msg[0] & 0x40) == 0 {
		println("earlyStop = false")
		earlyStop = false
	} else {
		println("earlyStop = true")
		earlyStop = true
	}

	maxBrightness, minBrightness = bitsManager(msg[1])

}

func bitsManager(num uint8) (int8, int8) {
	a, b := separateInt(num)
	if a < b {
		b = a
	}
	if b > 10 {
		b = 10
	}
	if a > 10 {
		a = 10
	}
	println("maxBrightness: ", a*10)
	println("minBrightness: ", b*10)
	return a * 10, b * 10
}

// func to separate an int into 2 array of 4 bits
func separateInt(num uint8) (int8, int8) {
	// separate the number into 2  using bitwise operators
	var arr uint8
	var arr1 uint8
	arr = num >> 4
	arr1 = num & 0x0F
	return int8(arr), int8(arr1)
}
