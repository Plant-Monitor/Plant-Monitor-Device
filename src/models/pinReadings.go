package models

type PinReadingsCollection map[Pin]PinReading
type PinReading uint8 // ?: @timiagiri Are we expecting 8 bit resolution or higher
type Pin string
