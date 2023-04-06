package pch

import "pcs/models"

const (
	temperatureSensor = "AHT20"
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
