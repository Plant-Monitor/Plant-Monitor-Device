package pch

import (
	"encoding/binary"
	"log"
	"periph.io/x/conn/v3/i2c"
	"time"
)

type sensorDriver func() []byte

const (
	AHT20_ADDRESS         = 0x38
	AHT20_MEASUREMENT_CMD = 0xAC3300
)

func getRawRead_AHTxx() []byte {
	// Connect to the AHT20 sensor using the I2C address
	dev := &i2c.Dev{Addr: AHT20_ADDRESS, Bus: i2cport}

	// Trigger a measurement by sending the command 0xAC
	data := make([]byte, 6)
	binary.BigEndian.PutUint32(data[0:3], AHT20_MEASUREMENT_CMD)

	if _, err := dev.Write(data[0:3]); err != nil {
		log.Fatalf("failed to send command: %v", err)
	}

	// Wait for the measurement to complete (16 ms for AHT20)
	time.Sleep(10 * time.Millisecond)

	// Read the temperature and humidity data by sending the command 0x33
	if err := dev.Tx(nil, data); err != nil {
		log.Fatalf("failed to read data: %v", err)
	}

	return data
}
