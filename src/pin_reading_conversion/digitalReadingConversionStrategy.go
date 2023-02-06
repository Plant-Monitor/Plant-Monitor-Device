package pin_reading_conversion

import "pcs/models"

type digitalReadingConversionStrategy interface {
	convert(models.PinReading) *models.HealthProperty
}
