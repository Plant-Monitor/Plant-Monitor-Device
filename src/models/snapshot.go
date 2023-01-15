package models

import (
	"time"

	"github.com/google/uuid"
)

type Snapshot struct {
	user_id           uuid.UUID
	plant_id          uuid.UUID
	timestamp         time.Time
	health_properties ReadingsCollection
}

type ReadingsCollection map[Metric]HealthProperty

type Metric string

type HealthProperty struct {
	level          float32
	unit           string
	interpretation Interpretation
}

type Interpretation int64
