package config

import (
	"os"
	"pcs/models"
	"pcs/utils"
	"strings"
	"sync"
)

type GpioConfig map[models.Pin]models.Metric

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
	var rslt = make(map[models.Pin]models.Metric)

	pinMapSlice := strings.Fields(pinMapString)
	var pin string

	for index, elem := range pinMapSlice {
		if index%2 == 0 {
			pin = elem
		} else {
			rslt[models.Pin(pin)] = models.Metric(elem)
		}
	}
	instance := GpioConfig(rslt)

	return &instance
}
