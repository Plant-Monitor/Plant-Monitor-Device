package models

type DigitalReadingsCollection map[Metric]DigitalReading
type DigitalReading uint8 // ?: @timiagiri Are we expecting 8 bit resolution or higher
// type Pin string
type PeripheralNumber uint8
