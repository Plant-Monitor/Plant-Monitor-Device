package dto

import (
	"github.com/google/uuid"
	"pcs/models"
	"time"
)

type ActionDto struct {
	ActionID        uuid.UUID       `json:"action_id"`
	Timestamp       time.Time       `json:"timestamp"`
	Type            actionType      `json:"action_type"`
	Status          actionStatus    `json:"action_status"`
	Metric          models.Metric   `json:"metric"`
	LevelNeeded     float32         `json:"level_needed"`
	CurrentSnapshot models.Snapshot `json:"current_snapshot"`
}

type actionType int8

const (
	TAKEN actionType = iota
	NEEDED
)

type actionStatus int8

const (
	RESOLVED actionStatus = iota
	UNRESOLVED
)
