package snapshots_test

import (
	"pcs/snapshots"
	"pcs/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnapshotUpdater(t *testing.T) {
	utils.LoadEnv()
	t.Run("No initial SnapshotUpdater instance", func(t *testing.T) {
		rslt := snapshots.GetSnapshotUpdaterInstance()
		assert.NotNil(t, rslt)
	})

	t.Run("There is an initial instance, returned instance should be the same", func(t *testing.T) {
		instance := snapshots.GetSnapshotUpdaterInstance()
		instance2 := snapshots.GetSnapshotUpdaterInstance()

		assert.Equal(t, instance, instance2)
	})
}
