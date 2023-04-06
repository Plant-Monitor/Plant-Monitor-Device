package pch

import "pcs/models"

type metricConversionStrategy func(collection rawReadingsCollection) models.HealthProperty

const (
	temperatureSensor = "AHT20"
	humiditySensor
)

func getTemperature(collection rawReadingsCollection) models.HealthProperty {
	data := collection[temperatureSensor]

	temp := float32(((uint32(data[3]) & 0xF) << 16) | (uint32(data[4]) << 8) | (uint32(data[5])))
	temp = (temp * 200.0 / 0x100000) - 50

	return models.HealthProperty{
		Level: temp,
		Unit:  "deg C",
	}
}

func getHumidity(collection rawReadingsCollection) models.HealthProperty {
	data := collection[humiditySensor]

	humidity := float32((uint32(data[1]) << 12) | (uint32(data[2]) << 4) | (uint32(data[3]) >> 4))
	humidity = (humidity * 100) / 0x100000

	return models.HealthProperty{
		Level: humidity,
		Unit:  "%",
	}
}
