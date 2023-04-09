package pch

import (
	"log"
	"pcs/models"
	"pcs/utils"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/conn/v3/spi"
	"periph.io/x/conn/v3/spi/spireg"
	"periph.io/x/host/v3"
	"sync"
)

var (
	i2cport i2c.BusCloser
	spiport spi.PortCloser
	
	trigPin gpio.PinIO
	echPin gpio.PinIO
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

	// Open an SPI connection to the MCP3008 ADC
	spiport, err := spireg.Open("SPI0.0") // Use the SPI0.0 port on Raspberry Pi
	if err != nil {
		log.Fatalf("failed to open SPI port: %v", err)
	}
	//defer port.Close()
	
	//open and configure pins for ultrasonic sensor
	trigPin = gpioreg.ByName("GPIO22")
	if trigPin == nil{
		log.Fatal("Failed to find trigger pin")
	}
	echoPin = gpioreg.ByName("GPIO18")
	if echoPin == nil{
		log.Fatal("Failed to find echo pin")
	}
	// Configure the trigger pin as an output and the echo pin as an input
	if err := trigPin.Out(gpio.Low); err != nil {
		fmt.Println("Failed to configure trigger pin:", err)
	}
	if err := echoPin.In(gpio.PullDown, gpio.NoEdge); err != nil {
		fmt.Println("Failed to configure echo pin:", err)
	}
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
		"MCP3008": getRawRead_MCP3008,
		"HCSR04": getRawRead_HCSR04,
		"TCS34725": getRawRead_TCS34725,
	}
}
func loadMetricConfig() metricConfig {
	return metricConfig{
		"temperature": getTemperature,
		"humidity":    getHumidity,
		"moisture":    getMoisture,
		"water-level": getWaterLevel,
		"intensity": getLightIntensity,
	}
}
