package gpio

import (
	"pcs/config"
	"pcs/utils"
)

type GpioClient struct {
	config *config.GpioConfig
}

var gpioClientInstance *GpioClient

func GetGpioClientInstance() *GpioClient {
	return utils.GetSingletonInstance(
		gpioClientInstance,
		newGpioClient,
		nil,
	)
}

func newGpioClient(initParams interface{}) *GpioClient {
	return &GpioClient{config.GetGpioConfigInstance()}
}

func (client *GpioClient) readPin(metric string) {

}

func (client *GpioClient) Read() map[string]float32 {
	return make(map[string]float32)
}
