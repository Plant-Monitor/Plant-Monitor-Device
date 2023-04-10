package actions

import (
	"pcs/models"
	"pcs/utils"
)

type action interface {
	execute()
}

type automatedAction interface {
	action
}

type userAction struct {
	healthProperty models.HealthProperty
	safeLevel      float32
}

func (action *userAction) execute() {
	utils.GetServerClientInstance().WriteAction()
}

func newUserAction(healthProperty models.HealthProperty, safeLevel float32) {
	getuserActionsStoreInstance().add(userAction{
		healthProperty,
		safeLevel,
	})
}
