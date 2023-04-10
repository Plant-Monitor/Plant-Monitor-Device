package pch

import (
	"pcs/models"
	"time"
)

type metricConversionStrategy func(collection rawReadingsCollection) *models.HealthProperty

const (
	temperatureSensor = "AHT20"
	humiditySensor    = "AHT20"
	moistureSensor    = "MCP3008"
	lightSensor       = "TCS34725"
	levelSensor       = "HCSR04"
)

func getTemperature(collection rawReadingsCollection) *models.HealthProperty {
	data := collection[temperatureSensor].([]byte)

	temp := float32(((uint32(data[3]) & 0xF) << 16) | (uint32(data[4]) << 8) | (uint32(data[5])))
	temp = (temp * 200.0 / 0x100000) - 50

	return &models.HealthProperty{
		Level: temp,
		Unit:  "deg C",
	}
}

func getHumidity(collection rawReadingsCollection) *models.HealthProperty {
	data := collection[humiditySensor].([]byte)

	humidity := float32((uint32(data[1]) << 12) | (uint32(data[2]) << 4) | (uint32(data[3]) >> 4))
	humidity = (humidity * 100) / 0x100000

	return &models.HealthProperty{
		Level: humidity,
		Unit:  "%",
	}
}

func getMoisture(collection rawReadingsCollection) *models.HealthProperty {
	data := collection[moistureSensor].([]byte)

	// Parse the raw data into a 10-bit analog voltage value
	rawValue := ((uint32(data[1]) & 0x03) << 8) | uint32(data[2])
	voltage := float32(rawValue) * 3.3 / 1023 // Assuming VCC=3.3V
	percentage := (voltage / 3.3) * 100

	return &models.HealthProperty{
		Level: percentage,
		Unit:  "%",
	}
}

func getWaterLevel(collection rawReadingsCollection) *models.HealthProperty {
	duration := (collection[levelSensor]).(time.Duration)

	distance := float32(duration.Seconds() * 340.0 / 2.0) // speed of sound is 340 m/s
	distance = distance * 100
	return &models.HealthProperty{
		Level: distance,
		Unit:  "cm",
	}
}

func getLightIntensity(collection rawReadingsCollection) *models.HealthProperty {
	data := collection[lightSensor].([]byte)
	intensity := uint16(data[0])
	addIntensity := uint16(data[1])
	intensity |= addIntensity << 8

	return &models.HealthProperty{
		Level: float32(intensity),
		Unit:  "lux",
	}
}
