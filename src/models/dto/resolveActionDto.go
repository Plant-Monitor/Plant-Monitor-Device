package dto

import (
	"github.com/google/uuid"
	"pcs/models"
)

type ResolveActionDto struct {
	ActionId uuid.UUID       `json:"action_id"`
	Snapshot models.Snapshot `json:"current_snapshot"`
}

func CreateResolveActionDto(actionId uuid.UUID, snapshot models.Snapshot) *ResolveActionDto {
	return &ResolveActionDto{ActionId: actionId, Snapshot: snapshot}
}
