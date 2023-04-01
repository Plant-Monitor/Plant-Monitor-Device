package gpio

import (
	"encoding/json"
	"fmt"
	"github.com/stianeikeland/go-rpio/v4"
	"log"
	"os"
	"pcs/models"
	"pcs/utils"
	"sync"
)

type GpioClient struct {
	config *gpioConfig
}

var gpioClientInstance *GpioClient
var gpioClientLock = &sync.Mutex{}

func GetGpioClientInstance() *GpioClient {
	return utils.GetSingletonInstance(
		gpioClientInstance,
		gpioClientLock,
		newGpioClient,
		nil,
	)
}

func newGpioClient(initParams ...any) *GpioClient {
	spiSetup()
	return &GpioClient{loadGpioConfig()}
}

func spiSetup() {
	if err := rpio.Open(); err != nil {
		panic(err)
	}

	if err := rpio.SpiBegin(rpio.Spi0); err != nil {
		panic(err)
	}
}

func (client *GpioClient) readDigitalValue(number models.PeripheralNumber) models.DigitalReading {
	rpio.SpiChipSelect(uint8(number))
	return models.DigitalReading(rpio.SpiReceive(1)[0])
}

func (client *GpioClient) Read() models.DigitalReadingsCollection {
	readsColl := make(models.DigitalReadingsCollection)

	for metric, periphNum := range *client.config {
		readsColl[metric] = client.readDigitalValue(periphNum)
	}
	return readsColl
}

type gpioConfig map[models.Metric]models.PeripheralNumber

func loadGpioConfig() *gpioConfig {
	content, err := os.ReadFile(
		fmt.Sprintf(
			"%s/src/config/metricPeripheralNumberMapping.json",
			os.Getenv("PATH_TO_PROJECT"),
		),
	)

	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Unmarshalling data
	var payload gpioConfig
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return &payload
}
