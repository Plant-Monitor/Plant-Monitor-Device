package actions

import (
	"github.com/google/uuid"
	"pcs/models"
	"pcs/utils"
	"time"
)

type Action struct {
	ActionID        uuid.UUID        `json:"action_id"`
	Timestamp       time.Time        `json:"timestamp"`
	Type            actionType       `json:"action_type"`
	Status          actionStatus     `json:"status"`
	Metric          models.Metric    `json:"metric"`
	LevelNeeded     float32          `json:"level_needed"`
	CurrentSnapshot *models.Snapshot `json:"current_snapshot"`
	CriticalRange   criticalRange    `json:"critical_range"`

	executeCallback ActionExecutionCallback
}

type actionType string
type actionStatus int8
type criticalRange int8
type ActionExecutionCallback func() error

const (
	TAKEN  actionType = "TAKEN"
	NEEDED            = "NEEDED"
)

const (
	RESOLVED actionStatus = iota
	UNRESOLVED
)

const (
	NOT_CRITICAL criticalRange = iota
	CRITICAL_HIGH
	CRITICAL_LOW
)

func newAction(
	actType actionType,
	levelNeeded float32,
	metric models.Metric,
	snapshot *models.Snapshot,
	executionCallback ActionExecutionCallback,
) (actionId uuid.UUID) {
	action := &Action{
		ActionID:        uuid.New(),
		Timestamp:       time.Now(),
		Type:            actType,
		Status:          UNRESOLVED,
		Metric:          metric,
		LevelNeeded:     levelNeeded,
		CurrentSnapshot: snapshot,
		executeCallback: executionCallback,
	}

	store := getActionsStoreInstance()
	store.add(action)

	return action.ActionID
}

func (action *Action) execute() (serverErr error, execErr error) {
	serverErr = utils.GetServerClientInstance().CreateAction(action)

	return serverErr, action.executeCallback()
}
