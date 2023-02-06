package actions

import (
	"pcs/utils"
	"sync"
)

type iActionsStore interface {
	add(action action)
}

type actionsStore struct {
	actionsQueue []action
}

func (store *actionsStore) add(action action) {
	store.actionsQueue = append(store.actionsQueue, action)
}

type automatedActionsStore struct{ actionsStore }

var automatedActionsStoreInstance *automatedActionsStore
var automatedActionsStoreLock *sync.Mutex = &sync.Mutex{}

func getAutomatedActionsStoreInstance() *automatedActionsStore {
	return utils.GetSingletonInstance(
		automatedActionsStoreInstance,
		automatedActionsStoreLock,
		newAutomatedActionsStore,
		nil,
	)
}

func newAutomatedActionsStore(initParams ...any) *automatedActionsStore {
	return &automatedActionsStore{
		make([]automatedAction, 0),
	}
}

type userActionsStore struct{ actionsStore }

var userActionsStoreInstance *userActionsStore
var userActionsStoreLock *sync.Mutex = &sync.Mutex{}

func getuserActionsStoreInstance() *userActionsStore {
	return utils.GetSingletonInstance(
		userActionsStoreInstance,
		userActionsStoreLock,
		newUserActionsStore,
		nil,
	)
}

func newUserActionsStore(initParams ...any) *userActionsStore {
	return &userActionsStore{
		make([]userAction, 0),
	}
}
