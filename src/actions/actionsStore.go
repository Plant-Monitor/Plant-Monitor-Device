package actions

import (
	"pcs/utils"
	"sync"
)

/* actionsStore ABSTRACT CLASS */

type iActionsStore interface {
	add(action action)
}

type actionsStore struct {
	actionsQueue []action
}

func (store *actionsStore) add(action action) {
	store.actionsQueue = append(store.actionsQueue, action)
}

/* automatedActionsStore CLASS */

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
		actionsStore: actionsStore{
			actionsQueue: make([]action, 0),
		},
	}
}

func (store *automatedActionsStore) add(action automatedAction) {
	store.actionsStore.add(action)
}

/* userActionsStore CLASS */

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
		actionsStore: actionsStore{
			actionsQueue: make([]action, 0),
		},
	}
}

func (store *userActionsStore) add(action userAction) {
	store.actionsStore.add(action)
}
