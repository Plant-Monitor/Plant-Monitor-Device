package models

import (
	"os"
	"time"

	"github.com/google/uuid"
)

type Snapshot struct {
	UserId           uuid.UUID                   `json:"user_id"`
	PlantId          uuid.UUID                   `json:"plant_id"`
	Timestamp        time.Time                   `json:"timestamp"`
	HealthProperties ConvertedReadingsCollection `json:"health_properties"`
}

type ConvertedReadingsCollection map[Metric]*HealthProperty

type Metric string

type HealthProperty struct {
	Level          float64        `json:"level"`
	Unit           string         `json:"unit"`
	Interpretation Interpretation `json:"interpretation"`
}

type Interpretation string

const (
	GOOD Interpretation = "GOOD"
	OKAY = "OKAY"
	CRITICAL = "CRITICAL"
)

func BuildSnapshot(readings ConvertedReadingsCollection) *Snapshot {
	userId, _ := uuid.Parse(os.Getenv("USER_ID"))
	plantId, _ := uuid.Parse(os.Getenv("PLANT_ID"))

	return &Snapshot{
		userId,
		plantId,
		time.Now(),
		readings,
	}
}
