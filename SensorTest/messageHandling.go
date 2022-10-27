package main

import (
	"encoding/hex"
)

var (
// earlyStop bool = false
)

// Initialize the AT module
/* Non unitTest compatible
func InitAT() {
	println("Initializing AT...")
	time.Sleep(time.Millisecond * 50)
	(machine.UARTConfig{BaudRate: 9600, TX: machine.D1, RX: machine.D0})
	_, err := machine.UART0.Write([]byte("AT+JOIN=DR3\r\n"))
	if err != nil {
		println("Error: " + err.Error())
	}
	ReadMessage(1)
}*/

// Send a message to the serial port of the lora module with the given payload
func SendMessage(payload string) string {
	//println("Sending message...")
	a := string([]byte(`AT+MSG= "` + payload + `"` + "\r\n"))
	return a
}

// Read the serial port of the lora module and return the message
/* Non unitTest compatible
func ReadMessage(wT int8, UART0Buff int, UART0Read byte) string {
	var msg string
	timer := 0
	msg1 := ""
	for {
		if timer >= int(wT)*60*1000 || earlyStop {
			return ""
		}
		if UART0Buff > 0 {
			rb := UART0Read
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
*/

// handling received message
func msgTreating(msg string) (bool, bool, int8, int8) {
	//println("treating...")
	//println(msg)
	if len(msg) < 4 {
		return false, false, 0, 0
	}
	str := Hex2Bin(msg[0])
	var st bool
	//msg is 2 bytes long, take first byte from str and take the last bit if it's 0, turn off the led, if it's 1 turn on the led
	if str[7] == '0' {
		//println("turn off")
		st = true
	} else {
		//println("turn on")
		st = false
	}

	var est bool
	//take second bit and if true earlyStop = true
	if str[6] == '1' {
		//println("earlyStop = false")
		est = false
	} else {
		//println("earlyStop = true")
		est = true
	}
	bytearray := msg[2:]
	maxB, minB := bitsManager(hex.EncodeToString([]byte(bytearray)))

	return st, est, minB, maxB
}

func bitsManager(num string) (int8, int8) {
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
	return int8(a * 10), int8(b * 10)
}

func Hex2Bin(in byte) string {
	var out []byte
	for i := 7; i >= 0; i-- {
		b := (in >> uint(i))
		out = append(out, (b%2)+48)
	}
	return string(out)
}
