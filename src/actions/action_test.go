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

func TestActions(t *testing.T) {
	setup := func() {
		err := godotenv.Load("../../.env")
		if err != nil {
			return
		}
	}

	t.Run("Create a new action", func(t *testing.T) {
		setup()

		actionId := newAction(
			NEEDED,
			420.69,
			"temperature",
			&models.Snapshot{},
			func() error { return nil },
		)
		assert.NotNil(t, actionId)
		fmt.Printf("Generated actionId: %s\n", actionId)
	})

	// Setup: Server needs to be running
	t.Run("Execute action's callback (NEEDED action type)", func(t *testing.T) {
		setup()

		action := &Action{
			ActionID:        uuid.New(),
			Timestamp:       time.Now(),
			Type:            NEEDED,
			Status:          UNRESOLVED,
			Metric:          "test4",
			LevelNeeded:     69.420,
			CurrentSnapshot: &models.Snapshot{},
			executeCallback: func() error { return nil },
		}
		serverErr, execErr := action.execute()

		assert.Nil(t, serverErr)
		assert.Nil(t, execErr)
	})

	t.Run("Ensure that an action is store to the action store upon creation", func(t *testing.T) {
		setup()

		actionId := newAction(
			NEEDED,
			420.69,
			"test4",
			&models.Snapshot{},
			func() error { return nil },
		)

		fmt.Printf("Generated action id: %s\n", actionId)
		storedAction := getActionsStoreInstance().get(actionId)
		assert.NotNil(t, storedAction)
		assert.Equal(t, storedAction.ActionID, actionId)
	})
}
