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
	machine.UART0.Configure(machine.UARTConfig{TX: machine.D1, RX: machine.D0})
	var err error
	for err != nil {
		_, err = machine.UART0.Write([]byte("AT+JOIN\r\n"))
		time.Sleep(time.Second * 5)
		if err != nil {
			println(err.Error())
		} else {
			println(ReadMessage())
		}
	}
}

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
