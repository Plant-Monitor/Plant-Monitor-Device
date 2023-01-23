package snapshots

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnapshotPublisher(t *testing.T) {
	t.Run("No initial SnapshotPublisher instance", func(t *testing.T) {
		rslt := GetSnapshotPublisherInstance()
		assert.NotNil(t, rslt)
	})

	t.Run("There is an initial instance, returned instance should be the same", func(t *testing.T) {
		instance := GetSnapshotPublisherInstance()
		instance2 := GetSnapshotPublisherInstance()

		assert.Equal(t, instance, instance2)
	})

}
