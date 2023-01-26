package gpio

import (
	"os"
	"pcs/models"
	"pcs/utils"
	"strings"
	"sync"
)

type gpioConfig map[models.Metric]models.Pin

var gpioConfigInstance *gpioConfig
var gpioConfigLock = &sync.Mutex{}

func GetGpioConfigInstance() *gpioConfig {
	return utils.GetSingletonInstance(
		gpioConfigInstance,
		gpioConfigLock,
		newGpioConfig,
		nil,
	)
}

func newGpioConfig(initParams ...any) *gpioConfig {
	pinMapString := os.Getenv("PIN_MAP_STRING")
	return parsePinMapString(pinMapString)
}

func parsePinMapString(pinMapString string) *gpioConfig {
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
	instance := gpioConfig(rslt)

	return &instance
}
