package gpio

import (
	"pcs/models"
	"pcs/utils"
	"sync"
)

type GpioClient struct {
	config *gpioConfig
}

var gpioClientInstance *GpioClient
var gpioClientLock *sync.Mutex = &sync.Mutex{}

func GetGpioClientInstance() *GpioClient {
	return utils.GetSingletonInstance(
		gpioClientInstance,
		gpioClientLock,
		newGpioClient,
		nil,
	)
}

func newGpioClient(initParams ...any) *GpioClient {
	return &GpioClient{GetGpioConfigInstance()}
}

func (client *GpioClient) readDigitalValue(pin models.Pin) models.DigitalReading {
	/*
		TODO:
		Implement digital signal reading as a function of the target metric.
		Logic will depend on chosen communication protocol.
	*/
	return models.DigitalReading(0)
}

func (client *GpioClient) Read() models.DigitalReadingsCollection {
	readsColl := make(models.DigitalReadingsCollection)

	for metric, pin := range *client.config {
		readsColl[metric] = client.readDigitalValue(pin)
	}
	return readsColl
}
