package actions

import (
	"fmt"
	"github.com/google/uuid"
	"pcs/utils"
	"sync"
)

type ActionsStore struct {
}

var actionsStoreInstance *ActionsStore
var actionsStoreLock = &sync.Mutex{}
var storeDict = make(map[uuid.UUID]*Action)

func GetActionsStoreInstance() *ActionsStore {
	return utils.GetSingletonInstance(
		actionsStoreInstance,
		actionsStoreLock,
		newActionsStore,
		nil,
	)
}

func newActionsStore(initParams ...any) *ActionsStore {
	//instance := make(ActionsStore)
	//return &instance
	return &ActionsStore{}
}

func (store *ActionsStore) add(action *Action) {
	actionsStoreLock.Lock()
	defer actionsStoreLock.Unlock()
	storeDict[action.ActionID] = action
}

func (store *ActionsStore) resolve(action *Action) error {
	delete(storeDict, action.ActionID)
	return nil
}

func (store *ActionsStore) get(id uuid.UUID) *Action {
	return storeDict[id]
}

func (store *ActionsStore) Execute() error {
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
