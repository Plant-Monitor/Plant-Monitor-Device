package config

import (
	"os"
	"pcs/models"
	"pcs/utils"
	"strings"
	"sync"
)

type GpioConfig map[models.Metric]models.Pin

var gpioConfigInstance *GpioConfig
var gpioConfigLock = &sync.Mutex{}

func GetGpioConfigInstance() *GpioConfig {
	return utils.GetSingletonInstance(
		gpioConfigInstance,
		gpioConfigLock,
		newGpioConfig,
		nil,
	)
}

func newGpioConfig(initParams ...any) *GpioConfig {
	pinMapString := os.Getenv("PIN_MAP_STRING")
	return parsePinMapString(pinMapString)
}

func parsePinMapString(pinMapString string) *GpioConfig {
	var rslt = make(map[models.Metric]models.Pin)

	pinMapSlice := strings.Fields(pinMapString)
	var metric string

	for index, elem := range pinMapSlice {
		if index%2 == 0 {
			metric = elem
		} else {
			rslt[models.Metric(metric)] = models.Pin(elem)
		}
	}
	instance := GpioConfig(rslt)

	return &instance
}
