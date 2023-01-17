package gpio

import (
	"pcs/config"
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

func (client *GpioClient) readPin(metric string) {

}

func (client *GpioClient) Read() map[string]float32 {
	return make(map[string]float32)
}
