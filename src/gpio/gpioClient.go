package gpio

import (
	"fmt"
	"pcs/config"
	"sync"
)

type GpioClient struct {
	config *config.GpioConfig
}

var gpioClientLock = &sync.Mutex{}
var gpioClientInstance *GpioClient

func GetGpioClientInstance() *GpioClient {
	if gpioClientInstance == nil {
		gpioClientLock.Lock()
		defer gpioClientLock.Unlock()
		if gpioClientInstance == nil {
			fmt.Println("Creating GpioClient instance now.")
			gpioClientInstance = newGpioClient()
		} else {
			fmt.Println("GpioClient instance already created.")
		}
	} else {
		fmt.Println("GpioClient instance already created.")
	}

	return gpioClientInstance
}

func newGpioClient() *GpioClient {
	return &GpioClient{config.GetGpioConfigInstance()}
}

func (client *GpioClient) readPin(metric string) {

}

func (client *GpioClient) Read() map[string]float32 {
	return make(map[string]float32)
}
