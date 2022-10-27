package main

import (
	"tinygo.org/x/drivers"
)

// Device wraps an I2C connection to a DS1307 device.
type Device struct {
	bus         drivers.I2C
	Address     uint8
	AddressSRAM uint8
}

const (
	I2CAddress = 0x68
	TimeDate   = 0x00
	Control    = 0x7
	//CH is oscillator halt bit
	CH              = 0x7
	SRAMBeginAddres = 0x8
	SRAMEndAddress  = 0x3F
)

// New creates a new DS1307 connection. I2C bus must be already configured.
func New(bus drivers.I2C) Device {
	return Device{bus: bus,
		Address:     uint8(I2CAddress),
		AddressSRAM: SRAMBeginAddres,
	}
}

// decToBcd converts int to BCD
func decToBcd(dec int) uint8 {
	return uint8(dec + 6*(dec/10))
}

// bcdToDec converts BCD to int
func bcdToDec(bcd uint8) int {
	return int(bcd - 6*(bcd>>4))
}

// hoursBCDToInt converts the BCD hours to int
func hoursBCDToInt(value uint8) (hour int) {
	if value&0x40 != 0x00 {
		hour = bcdToDec(value & 0x1F)
		if (value & 0x20) != 0x00 {
			hour += 12
		}
	} else {
		hour = bcdToDec(value)
	}
	return
}

// IsOscillatorRunning returns if the oscillator is running
func (d *Device) IsOscillatorRunning() bool {
	data := []byte{0}
	err := d.bus.ReadRegister(d.Address, uint8(TimeDate), data)
	if err != nil {
		return false
	}
	return (data[0] & (1 << CH)) == 0
}

// SetOscillatorRunning starts/stops internal oscillator by toggling halt bit
func (d *Device) SetOscillatorRunning(running bool) error {
	data := make([]byte, 3)
	err := d.bus.ReadRegister(d.Address, uint8(TimeDate), data)
	if err != nil {
		return err
	}
	if running {
		data[0] &^= (1 << CH)
	} else {
		data[0] |= (1 << CH)
	}
	data[1], data[0] = data[0], uint8(TimeDate)
	err = d.bus.Tx(uint16(d.Address), data[:2], nil)
	return err
}
