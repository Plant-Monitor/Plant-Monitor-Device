package pin_reading_conversion

import "pcs/models"

type digitalReadingConversionStrategy interface {
	convert(models.DigitalReading) models.HealthProperty
}
