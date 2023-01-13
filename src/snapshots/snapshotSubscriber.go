package snapshots

import (
	"fmt"
	"sync"
)

type SnapshotSubscriber interface {
	update(Snapshot)
}

// Implements SnapshotSubscriber. Role is to update DB after a configured timelapse
type SnapshotUpdater struct {
	updateStrategy SnapshotWatcherUpdateStrategy
}

var snapshotUpdaterlock = &sync.Mutex{}
var snapshotUpdaterInstance *SnapshotUpdater

func getSnapshotUpdaterInstance() *SnapshotUpdater {
	if snapshotUpdaterInstance == nil {
		snapshotUpdaterlock.Lock()
		defer snapshotUpdaterlock.Unlock()
		if snapshotUpdaterInstance == nil {
			fmt.Println("Creating SnapshotUpdater instance now.")
			snapshotUpdaterInstance = &SnapshotUpdater{new(PeriodicUpdateStrategy)}
		} else {
			fmt.Println("SnapshotUpdater instance already created.")
		}
	} else {
		fmt.Println("SnapshotUpdater instance already created.")
	}

	return snapshotUpdaterInstance
}

func (snapshotUpdater *SnapshotUpdater) update(snapshot Snapshot) {
	snapshotUpdater.updateStrategy.update(snapshot)
}
