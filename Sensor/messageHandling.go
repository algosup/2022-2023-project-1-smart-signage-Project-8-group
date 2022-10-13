package main

import (
	"machine"
)

var (
	failSend    int
	failReceive int
)

// Send a message to the serial port of the lora module with the given payload
func SendMessage(payload string) {
	_, err := machine.UART0.Write([]byte("AT+SEND=" + payload + "\r\n"))
	if err != nil {
		println(err.Error())
		failSend++
	}
}

// Read the serial port of the lora module and return the message
func ReadMessage() string {
	var msg string
	for {
		if machine.UART0.Buffered() > 0 {
			rb, err := machine.UART0.ReadByte()
			if err != nil {
				println(err.Error())
				failReceive++
				return ""
			}
			msg += string(rb)
			if msg[len(msg)-1] == '\n' {
				return msg
			}
		}
	}
}
