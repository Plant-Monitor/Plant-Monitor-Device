package snapshots

import "time"

type SnapshotWatcherUpdateStrategy interface {
	update(Snapshot)
}

type PeriodicUpdateStrategy struct {
	lastUpdate     time.Time
	updateInterval time.Duration
}

func (perUpdateStrategy *PeriodicUpdateStrategy) update(snapshot Snapshot) {

}
