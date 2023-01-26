package gpio

import (
	"pcs/config"
	"pcs/models"
	"pcs/utils"
	"sync"
)

type GpioClient struct {
	config *config.GpioConfig
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
	return &GpioClient{config.GetGpioConfigInstance()}
}

func (client *GpioClient) readDigitalValue(metric models.Metric) models.DigitalReading {
	/*
		TODO:
		Implement digital signal reading as a function of the target metric.
		Logic will depend on chosen communication protocol.
	*/
	return models.DigitalReading(0)
}

func (client *GpioClient) Read() models.DigitalReadingsCollection {
	readsColl := make(models.DigitalReadingsCollection)

	for metric, _ := range *client.config {
		readsColl[metric] = client.readDigitalValue(metric)
	}
	return readsColl
}
