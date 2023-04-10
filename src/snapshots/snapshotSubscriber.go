package snapshots

import (
	"os"
	"pcs/models"
	"pcs/utils"
	"sync"
	"fmt"
)

type SnapshotSubscriber interface {
	update(snapshot *models.Snapshot)
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
	fmt.Printf("[snapshots] Setting update interval to %s\n", initParams[0].(string))
	updateStrategy := NewPeriodicUpdateStrategy(initParams[0].(string))
	return &SnapshotUpdater{updateStrategy: updateStrategy}
}

func (snapshotUpdater *SnapshotUpdater) update(snapshot *models.Snapshot) {
	snapshotUpdater.updateStrategy.update(snapshot)
}

type MetricSubscriber struct {
	updateStrategy MetricSubscriberUpdateStrategy
}

func newMetricSubscriber(metric models.Metric) *MetricSubscriber {
	return &MetricSubscriber{}
}

func (sub *MetricSubscriber) update(snapshot *models.Snapshot) {
	sub.updateStrategy.update(snapshot)
}
