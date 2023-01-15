package snapshots

import (
	"fmt"
	"os"
	"sync"
	"time"
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
			snapshotUpdateInterval := os.Getenv("SNAPSHOT_UPDATE_INTERVAL")
			snapshotUpdateIntervalDuration, _ := time.ParseDuration(snapshotUpdateInterval)
			snapshotUpdaterInstance = &SnapshotUpdater{NewPeriodicUpdateStrategy(snapshotUpdateIntervalDuration)}
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
