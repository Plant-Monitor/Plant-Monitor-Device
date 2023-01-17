package models

import (
	"os"
	"time"

	"github.com/google/uuid"
)

type Snapshot struct {
	User_id           uuid.UUID          `json:"user_id"`
	Plant_id          uuid.UUID          `json:"plant_id"`
	Timestamp         time.Time          `json:"timestamp"`
	Health_properties ReadingsCollection `json:"health_properties"`
}

type ReadingsCollection map[Metric]HealthProperty

type Metric string

type HealthProperty struct {
	Level          float32        `json:"level"`
	Unit           string         `json:"unit"`
	Interpretation Interpretation `json:"interpretation"`
}

type Interpretation int64

func BuildSnapshot(readings ReadingsCollection) *Snapshot {
	user_id, _ := uuid.Parse(os.Getenv("USER_ID"))
	plant_id, _ := uuid.Parse(os.Getenv("PLANT_ID"))

	return &Snapshot{
		user_id,
		plant_id,
		time.Now(),
		readings,
	}
}
