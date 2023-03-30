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
	Status          actionStatus     `json:"action_status"`
	Metric          models.Metric    `json:"metric"`
	LevelNeeded     float32          `json:"level_needed"`
	CurrentSnapshot *models.Snapshot `json:"current_snapshot"`

	executeCallback ActionExecutionCallback
}

type actionType int8
type actionStatus int8
type ActionExecutionCallback func() error

const (
	TAKEN actionType = iota
	NEEDED
)

const (
	RESOLVED actionStatus = iota
	UNRESOLVED
)

func newAction(
	actType actionType,
	levelNeeded float32,
	metric models.Metric,
	snapshot *models.Snapshot,
	executionCallback ActionExecutionCallback,
) (actionId uuid.UUID) {
	action := &Action{
		uuid.New(),
		time.Now(),
		actType,
		UNRESOLVED,
		metric,
		levelNeeded,
		snapshot,
		executionCallback,
	}

	getActionsStoreInstance().add(action)

	return action.ActionID
}

func (action *Action) execute() (serverErr error, execErr error) {
	serverErr = utils.GetServerClientInstance().CreateAction(action)

	return serverErr, action.executeCallback()
}
