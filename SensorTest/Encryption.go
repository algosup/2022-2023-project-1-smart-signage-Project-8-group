package main

import (
	"encoding/base64"
	"strconv"
)

func main() {
	timeCounter := 0
	f := uint16(65535)
	//uint to hex
	str := strconv.FormatUint(uint64(float32(f)/65535*100), 16) //create percentage of light sensor value and return it in string hex format
	if len(str) == 1 {
		str = "0" + str
	}

	if 32045 != 0 {
		timeCounter++
		if timeCounter == 255 {
			timeCounter = 0
		}
		str += strconv.FormatUint(uint64(timeCounter), 16)
	}
	if len(str) == 3 {
		str = str[:2] + "0" + str[2:]
	}

	//fake temporary data percentage added in str
	str += strconv.FormatUint(uint64(54), 16)
	if len(str) == 5 {
		str = str[:4] + "0" + str[4:]
	}

	//add 0x00 in string to the end of the string
	str += strconv.FormatUint(uint64(0), 16)
	if len(str) == 7 {
		str = str[:6] + "0" + str[6:]
	}

	println("str: ", str)
	println("base64: ", toBase64(str))
}

func toBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}
