package pch

import (
	"log"
	"pcs/models"
	"pcs/utils"
	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/host/v3"
	"sync"
)

var (
	i2cport i2c.BusCloser
)

type PCHClient struct {
	sensorConfig sensorConfig
	metricConfig metricConfig
}

var pchClientInstance *PCHClient
var pchClientLock *sync.Mutex = &sync.Mutex{}

func GetPCHClientInstance() *PCHClient {
	return utils.GetSingletonInstance(
		pchClientInstance,
		pchClientLock,
		newPchClient,
		nil,
	)
}

func newPchClient(initParams ...any) *PCHClient {
	setupPCH()
	return &PCHClient{
		loadSensorConfig(),
		loadMetricConfig(),
	}
}

func setupPCH() {
	_, err := host.Init()
	if err != nil {
		log.Fatal(err)
	}

	i2cport, err = i2creg.Open("")
	if err != nil {
		log.Fatalf("failed to open I2C port: %v", err)
	}

	// todo: setup spi port
	//defer port.Close()
}

func (client *PCHClient) GetReadings() models.ConvertedReadingsCollection {
	rawReadsColl := client.getRawReadingsCollection()

	coll := make(models.ConvertedReadingsCollection)
	for metric, conversionStrat := range client.metricConfig {
		coll[metric] = conversionStrat(rawReadsColl)
	}

	return coll
}

func (client *PCHClient) getRawReadingsCollection() rawReadingsCollection {
	coll := make(rawReadingsCollection)
	for sensor, driver := range client.sensorConfig {
		coll[sensor] = driver()
	}

	return coll
}

type sensorConfig map[sensor]sensorDriver
type rawReadingsCollection map[sensor][]byte
type metricConfig map[models.Metric]metricConversionStrategy
type sensor string

func loadSensorConfig() sensorConfig {
	return sensorConfig{
		"AHT20": getRawRead_AHTxx,
	}
}
func loadMetricConfig() metricConfig {
	return metricConfig{
		"temperature": getTemperature,
		"humidity":    getHumidity,
	}
}
