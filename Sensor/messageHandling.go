package main

import (
	"machine"
	"time"
)

var (
	failSend    int
	failReceive int
)

func InitAT() {
	println("Initializing AT...")
	time.Sleep(time.Millisecond * 200)
	machine.UART0.Configure(machine.UARTConfig{BaudRate: 9600, TX: machine.D1, RX: machine.D0})
	_, err := machine.UART0.Write([]byte("AT+JOIN\r\n"))
	time.Sleep(time.Second * 15)
	if err != nil {
		println("Error: " + err.Error())
	}
}

// Send a message to the serial port of the lora module with the given payload
func SendMessage(payload string) {
	println("Sending message...")
	_, err := machine.UART0.Write([]byte(`AT+MSG= "` + payload + `"` + "\r\n"))
	if err != nil {
		println("Error: " + err.Error())
		failSend++
	}
}

// Read the serial port of the lora module and return the message
func ReadMessage() string {
	println("reading...")
	var msg string
	for {
		if machine.UART0.Buffered() > 0 {
			rb, err := machine.UART0.ReadByte()
			if err != nil {
				println("Error: " + err.Error())
				failReceive++
				return ""
			}
			msg += string(rb)
			if msg[len(msg)-1] == '\n' {
				return msg
			}
		} else {
			time.Sleep(time.Millisecond * 500)
		}
	}
}
