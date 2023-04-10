package pch

import (
	"encoding/binary"
	"fmt"
	"log"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/conn/v3/spi"
	"time"
)

type sensorDriver func() []byte

const (
	AHT20_ADDRESS         = 0x38
	AHT20_MEASUREMENT_CMD = 0xAC3300
	TCS34725_ADDRESS      = 0x29
	mcp3008Channel        = 0
	Command               = 0x80 // Command register
	Enable                = 0x00 // Enable register
	EnablePowerOn         = 0x01 // Power on
	EnableRGBC            = 0x02 // Enable the RGBC sensor
	IntegrationTime       = 0x01 // Integration time register
	Gain                  = 0x0F // Gain register
	DataLow               = 0x14 // Low byte of clear channel data
	DataHigh              = 0x15 // High byte of clear channel data
)

func writeReg(dev *i2c.Dev, reg, value byte) error {
	cmd := []byte{Command | reg, value}
	_, err := dev.Write(cmd)
	return err
}

func readData(dev *i2c.Dev, reg byte) ([]byte, error) {
	cmd := []byte{Command | reg}
	data := make([]byte, 1)
	if err := dev.Tx(cmd, data); err != nil {
		return nil, err
	}
	return data, nil
}

func readDataOrDie(dev *i2c.Dev, reg byte) []byte {
	data, err := readData(dev, reg)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

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
	time.Sleep(16 * time.Millisecond)

	// Read the temperature and humidity data by sending the command 0x33
	if err := dev.Tx(nil, data); err != nil {
		log.Fatalf("failed to read data: %v", err)
	}

	return data
}

func getRawRead_MCP3008() []byte {
	// Connect to the MCP3008 ADC using the SPI parameters
	dev, err := spiport.Connect(1*physic.MegaHertz, spi.Mode0, 8)
	if err != nil {
		log.Fatalf("failed to connect to device: %v", err)
	}

	// Send the command to read the analog input from the specified channel
	tx := []byte{
		0x01,                      // Start bit
		(8 + mcp3008Channel) << 4, // Single-ended input mode
		0x00,                      // Placeholder byte for data
	}
	rx := make([]byte, 3) // Response buffer for 3 bytes of data
	if err := dev.Tx(tx, rx); err != nil {
		log.Fatalf("failed to send SPI command: %v", err)
	}

	return rx
}

func getRawRead_HCSR04() []byte {
	// Send a 10 microsecond pulse on the trigger pin
	err := trigPin.Out(gpio.High)
	if err != nil {
		return nil
	}
	time.Sleep(10 * time.Microsecond)
	err = trigPin.Out(gpio.Low)
	if err != nil {
		return nil
	}

	// Wait for the echo pin to go high
	startTime := time.Now()
	for echoPin.Read() == gpio.Low {
		if time.Since(startTime) > 500*time.Millisecond {
			fmt.Println("Timeout waiting for echo signal")
			break
		}
	}

	// Wait for the echo pin to go low and calculate the distance
	startTime = time.Now()
	for echoPin.Read() == gpio.High {
		if time.Since(startTime) > 500*time.Millisecond {
			fmt.Println("Timeout waiting for echo signal")
			break
		}
	}

	duration := time.Since(startTime)
	data := make([]byte, 1)
	//todo: convert duration to byte
	data[0] = byte(duration)
	return data
}

func getRawRead_TCS34725() []byte {
	// Connect to the AHT20 sensor using the I2C address
	dev := &i2c.Dev{Addr: TCS34725_ADDRESS, Bus: i2cport}

	// Initialize the TCS34725 sensor
	if err := writeReg(dev, Enable, EnablePowerOn|EnableRGBC); err != nil {
		log.Fatal(err)
	}
	if err := writeReg(dev, IntegrationTime, 0x00); err != nil { // 2.4ms integration time
		log.Fatal(err)
	}
	if err := writeReg(dev, Gain, 0x00); err != nil { // 1x gain
		log.Fatal(err)
	}

	// Read the light intensity from the sensor
	intensityLow, err := readData(dev, DataLow)
	if err != nil {
		log.Fatal(err)
	}
	intensityHigh := readDataOrDie(dev, DataHigh)
	data := []byte{intensityLow[0], intensityHigh[0]}

	return data
}
