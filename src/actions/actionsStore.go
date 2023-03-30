package actions

import (
	"github.com/google/uuid"
	"pcs/utils"
	"sync"
)

type actionsStore struct {
	actionsQueue      []Action
	unresolvedActions map[uuid.UUID]Action
}

var actionsStoreInstance *actionsStore
var actionsStoreLock = &sync.Mutex{}

func getActionsStoreInstance() *actionsStore {
	return utils.GetSingletonInstance(
		actionsStoreInstance,
		actionsStoreLock,
		newActionsStore,
		nil,
	)
}

func newActionsStore(initParams ...any) *actionsStore {
	return &actionsStore{
		actionsQueue:      make(0, []Action),
		unresolvedActions: make(map[uuid.UUID]Action),
	}
}

func (store *actionsStore) add(action *Action) {
	store.actionsQueue = append(store.actionsQueue, *action)
}

func (store *actionsStore) resolve(actionId uuid.UUID) error {
	return nil
}

func (store *actionsStore) execute() []error {
	//errs := make(0, []error)

	return nil
}