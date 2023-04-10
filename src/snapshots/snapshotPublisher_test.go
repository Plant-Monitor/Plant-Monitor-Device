package snapshots

import (
	"github.com/joho/godotenv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnapshotPublisher(t *testing.T) {
	setup := func() {
		err := godotenv.Load("../../.env")
		if err != nil {
			return
		}
	}

	t.Run("No initial SnapshotPublisher instance", func(t *testing.T) {
		setup()
		rslt := GetSnapshotPublisherInstance()
		assert.NotNil(t, rslt)
	})

	t.Run("There is an initial instance, returned instance should be the same", func(t *testing.T) {
		setup()
		instance := GetSnapshotPublisherInstance()
		instance2 := GetSnapshotPublisherInstance()

		assert.Equal(t, instance, instance2)
	})

}
