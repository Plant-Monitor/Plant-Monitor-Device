package actions

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"pcs/models"
	"testing"
	"time"
)

func TestActionsStore(t *testing.T) {
	setup := func() {
		err := godotenv.Load("../../.env")
		if err != nil {
			return
		}
	}

	t.Run("Test that actions store is a singleton", func(t *testing.T) {
		setup()
		store1 := GetActionsStoreInstance()
		store2 := GetActionsStoreInstance()
		assert.Equal(t, store1, store2)

		actionId := uuid.New()
		act := &Action{ActionID: actionId}
		store2.add(act)

		assert.Equal(t, store1, store2)
		assert.Equal(t, store1.get(actionId), act)
	})

	t.Run("Test adding to ActionsStore", func(t *testing.T) {
		setup()
		store := GetActionsStoreInstance()
		actionId := uuid.New()
		action := &Action{
			ActionID:        actionId,
			Timestamp:       time.Now(),
			Type:            NEEDED,
			Status:          UNRESOLVED,
			Metric:          "temperature",
			LevelNeeded:     69.420,
			CurrentSnapshot: &models.Snapshot{},
			executeCallback: func() error { return nil },
		}
		store.add(action)

		assert.Equal(t, store.get(actionId), action)
	})

	t.Run("Test resolving an action in the store", func(t *testing.T) {
		setup()
		store := GetActionsStoreInstance()
		actionId := uuid.New()
		action := &Action{
			ActionID:        actionId,
			Timestamp:       time.Now(),
			Type:            NEEDED,
			Status:          UNRESOLVED,
			Metric:          "temperature",
			LevelNeeded:     69.420,
			CurrentSnapshot: &models.Snapshot{},
			executeCallback: func() error { return nil },
		}
		store.add(action)

		assert.Equal(t, store.get(actionId), action)

		err := store.resolve(action)
		assert.Nil(t, err)
		assert.Nil(t, store.get(actionId))
	})

	t.Run("Test that actions execute from the store", func(t *testing.T) {
		setup()
		actionId := newAction(
			NEEDED,
			420.69,
			"test3",
			&models.Snapshot{},
			func() error { return nil },
		)
		fmt.Printf("Generated actionId: %s\n", actionId)
		assert.NotNil(t, GetActionsStoreInstance().get(actionId))
		err := GetActionsStoreInstance().Execute()
		assert.Nil(t, err)
	})
}
