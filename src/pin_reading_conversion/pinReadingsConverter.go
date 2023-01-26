package pin_reading_conversion

import (
	"pcs/config"
	"pcs/models"
	"pcs/utils"
	"sync"
)

type PinReadingsConverter struct {
	gpioConfig                 config.GpioConfig
	metricToConversionStrategy map[models.Metric]digitalReadingConversionStrategy
}

var pinReadingsConverterInstance *PinReadingsConverter
var pinReadingsConverterLock *sync.Mutex = &sync.Mutex{}

func newPinReadingsConverter(initParams ...any) *PinReadingsConverter {
	return &PinReadingsConverter{
		*config.GetGpioConfigInstance(),
		make(map[models.Metric]digitalReadingConversionStrategy),
	}
}

func GetPinReadingsConverterInstance() *PinReadingsConverter {
	return utils.GetSingletonInstance(
		pinReadingsConverterInstance,
		pinReadingsConverterLock,
		newPinReadingsConverter,
		nil,
	)
}

func (converter *PinReadingsConverter) Convert(pinReads models.PinReadingsCollection) models.ConvertedReadingsCollection {
	convertedReads := models.ConvertedReadingsCollection{}

	for pin, reading := range pinReads {
		metric := converter.gpioConfig[pin]
		convertedReads[metric] = converter.metricToConversionStrategy[metric].convert(reading)
	}

	return convertedReads
}
