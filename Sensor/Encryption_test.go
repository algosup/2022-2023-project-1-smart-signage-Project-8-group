package main

import "strconv"

func enc() {
	f := uint16(65535)
	//uint to hex
	str := strconv.FormatUint(uint64(float32(f)/65535*100), 16) //create percentage of light sensor value and return it in string hex format

	if 32045 != 0 {
		timeCounter++
		if timeCounter == 255 {
			timeCounter = 0
		}
		str += strconv.FormatUint(uint64(timeCounter), 16)
	}
	//fake temporary data percentage added in str
	str += strconv.FormatUint(uint64(54), 16)

	//add 0x00 in string to the end of the string
	str += strconv.FormatUint(uint64(0), 16)

	println("str: ", str)
}
