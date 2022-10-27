package main

import (
	"encoding/hex"
	"machine"
	"strings"
	"time"
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
	ReadMessage(16)
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
func ReadMessage(wTS uint16) string {
	data := make([]byte, 3)
	timeMachine.bus.ReadRegister(timeMachine.Address, uint8(0x00), data)
	println(timeMachine.IsOscillatorRunning())
	timer := uint32(hoursBCDToInt(data[2]))*3600 + uint32(bcdToDec(data[1]))*60 + uint32(bcdToDec(data[0]&0x7F)) + uint32(wTS)

	overflowHandler := false
	if timer >= 85319 {
		timer = 900
		overflowHandler = true
	}
	led.Configure()
	var msg string
	msg1 := ""

	for {
		changeLight(lS.Get(), led)                                                                                     // Get the value of the light sensor
		timeMachine.bus.ReadRegister(timeMachine.Address, uint8(0x00), data)                                           // Read the time from the RTC
		current := uint32(hoursBCDToInt(data[2]))*3600 + uint32(bcdToDec(data[1]))*60 + uint32(bcdToDec(data[0]&0x7F)) // Convert the time to seconds
		if overflowHandler {
			if current < 900 {
				current = 1080 + 1080 - current
			} else {
				current = 1080 - (86399 - current)
			}
		}

		if timer <= current || earlyStop {
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
