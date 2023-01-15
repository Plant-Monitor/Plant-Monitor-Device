package models

import (
	"os"
	"time"

	"github.com/google/uuid"
)

type Snapshot struct {
	user_id           uuid.UUID
	plant_id          uuid.UUID
	timestamp         time.Time
	health_properties ReadingsCollection
}

func NewSnapshot(readings ReadingsCollection) *Snapshot {
	user_id, _ := uuid.Parse(os.Getenv("USER_ID"))
	plant_id, _ := uuid.Parse(os.Getenv("PLANT_ID"))

	return &Snapshot{
		user_id,
		plant_id,
		time.Now(),
		readings,
	}
}

type ReadingsCollection map[Metric]HealthProperty

type Metric string

type HealthProperty struct {
	level          float32
	unit           string
	interpretation Interpretation
}

type Interpretation int64
