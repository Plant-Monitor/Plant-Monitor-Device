package snapshots

import (
	"pcs/models"
	"pcs/utils"
	"time"
)

type SnapshotWatcherUpdateStrategy interface {
	update(models.Snapshot) bool
}

type PeriodicUpdateStrategy struct {
	lastUpdate     *time.Time
	updateInterval time.Duration
	serverClient   *utils.ServerClient
}

func NewPeriodicUpdateStrategy(updateInterval string) *PeriodicUpdateStrategy {
	interval, _ := time.ParseDuration(updateInterval)

	return &PeriodicUpdateStrategy{
		nil,
		interval,
		utils.GetServerClientInstance(),
	}
}

func (perUpdateStrategy *PeriodicUpdateStrategy) update(snapshot models.Snapshot) (didUpdate bool) {
	if perUpdateStrategy.lastUpdate == nil || snapshot.Timestamp.Sub(*perUpdateStrategy.lastUpdate) >= perUpdateStrategy.updateInterval {
		perUpdateStrategy.serverClient.WriteSnapshot(&snapshot)
		timeNow := time.Now()
		perUpdateStrategy.lastUpdate = &timeNow
		return true
	}
	return false
}
