package actions

import (
	"fmt"
	"github.com/google/uuid"
	"pcs/utils"
	"sync"
)

type actionsStore struct {
}

var actionsStoreInstance *actionsStore
var actionsStoreLock = &sync.Mutex{}
var storeDict = make(map[uuid.UUID]*Action)

func getActionsStoreInstance() *actionsStore {
	return utils.GetSingletonInstance(
		actionsStoreInstance,
		actionsStoreLock,
		newActionsStore,
		nil,
	)
}

func newActionsStore(initParams ...any) *actionsStore {
	//instance := make(actionsStore)
	//return &instance
	return &actionsStore{}
}

func (store *actionsStore) add(action *Action) {
	actionsStoreLock.Lock()
	defer actionsStoreLock.Unlock()
	storeDict[action.ActionID] = action
}

func (store *actionsStore) resolve(action *Action) error {
	delete(storeDict, action.ActionID)
	return nil
}

func (store *actionsStore) get(id uuid.UUID) *Action {
	return storeDict[id]
}

func (store *actionsStore) execute() error {
	for _, action := range storeDict {
		serverErr, execErr := action.execute()
		if serverErr != nil {
			fmt.Printf("Failed to create action on server: %s\n", serverErr)
		}
		if execErr != nil {
			fmt.Printf("Failed to execute action: %s\n", execErr)
		}
	}
	return nil
}
