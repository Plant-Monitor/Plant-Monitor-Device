package snapshots

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnapshotUpdater(t *testing.T) {
	t.Run("No initial SnapshotUpdater instance", func(t *testing.T) {
		rslt := GetSnapshotUpdaterInstance()
		assert.NotNil(t, rslt)
	})

	t.Run("There is an initial instance, returned instance should be the same", func(t *testing.T) {
		instance := GetSnapshotUpdaterInstance()
		instance2 := GetSnapshotUpdaterInstance()

		assert.Equal(t, instance, instance2)
	})
}
