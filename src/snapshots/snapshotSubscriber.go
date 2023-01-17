package snapshots

import (
	"os"
	"pcs/models"
	"pcs/utils"
	"sync"
)

type SnapshotSubscriber interface {
	update(models.Snapshot)
}

// Implements SnapshotSubscriber. Role is to update DB after a configured timelapse
type SnapshotUpdater struct {
	updateStrategy SnapshotWatcherUpdateStrategy
}

var snapshotUpdaterInstance *SnapshotUpdater
var snapshotUpdaterLock *sync.Mutex = &sync.Mutex{}

func GetSnapshotUpdaterInstance() *SnapshotUpdater {
	updateInterval := os.Getenv("SNAPSHOT_UPDATE_INTERVAL")

	return utils.GetSingletonInstance(
		snapshotUpdaterInstance,
		snapshotUpdaterLock,
		newSnapshotUpdater,
		updateInterval,
	)
}

func newSnapshotUpdater(initParams ...any) *SnapshotUpdater {
	updateStrategy := NewPeriodicUpdateStrategy(initParams[0].(string))
	return &SnapshotUpdater{updateStrategy: updateStrategy}
}

func (snapshotUpdater *SnapshotUpdater) update(snapshot models.Snapshot) {
	snapshotUpdater.updateStrategy.update(snapshot)
}
