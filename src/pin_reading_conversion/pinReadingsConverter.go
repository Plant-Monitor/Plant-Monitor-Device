package pin_reading_conversion

import (
	"pcs/models"
	"pcs/utils"
	"sync"
)

type PinReadingsConverter struct {
	metricToConversionStrategy map[models.Metric]digitalReadingConversionStrategy
}

var pinReadingsConverterInstance *PinReadingsConverter
var pinReadingsConverterLock *sync.Mutex = &sync.Mutex{}

func newPinReadingsConverter(initParams ...any) *PinReadingsConverter {
	return &PinReadingsConverter{
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

func (converter *PinReadingsConverter) Convert(reads models.DigitalReadingsCollection) models.ConvertedReadingsCollection {
	convertedReads := models.ConvertedReadingsCollection{}

	for metric, read := range reads {
		convertedReads[metric] = converter.metricToConversionStrategy[metric].convert(read)
	}

	return convertedReads
}
