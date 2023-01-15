package snapshots

import (
	"pcs/models"
	"time"
)

type SnapshotWatcherUpdateStrategy interface {
	update(models.Snapshot)
}

type PeriodicUpdateStrategy struct {
	lastUpdate     time.Time
	updateInterval time.Duration
}

func (perUpdateStrategy *PeriodicUpdateStrategy) update(snapshot models.Snapshot) {

}
