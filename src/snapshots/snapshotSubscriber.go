package snapshots

import (
	"pcs/utils"
)

type SnapshotSubscriber interface {
	update(Snapshot)
}

// Implements SnapshotSubscriber. Role is to update DB after a configured timelapse
type SnapshotUpdater struct {
	updateStrategy SnapshotWatcherUpdateStrategy
}

var snapshotUpdaterInstance *SnapshotUpdater

func GetSnapshotUpdaterInstance() *SnapshotUpdater {
	return utils.GetSingletonInstance(
		snapshotUpdaterInstance,
		newSnapshotUpdater,
		nil,
	)
}

func newSnapshotUpdater(initParams interface{}) *SnapshotUpdater {
	return &SnapshotUpdater{new(PeriodicUpdateStrategy)}
}

func (snapshotUpdater *SnapshotUpdater) update(snapshot Snapshot) {
	snapshotUpdater.updateStrategy.update(snapshot)
}
