package snapshots

import (
	"pcs/actions"
	"pcs/analysis"
	"pcs/models"
	"pcs/utils"
	"time"
)

type SnapshotWatcherUpdateStrategy interface {
	update(*models.Snapshot) bool
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

func (perUpdateStrategy *PeriodicUpdateStrategy) update(snapshot *models.Snapshot) (didUpdate bool) {
	if perUpdateStrategy.lastUpdate == nil || snapshot.Timestamp.Sub(*perUpdateStrategy.lastUpdate) >= perUpdateStrategy.updateInterval {
		perUpdateStrategy.serverClient.WriteSnapshot(snapshot)
		timeNow := time.Now()
		perUpdateStrategy.lastUpdate = &timeNow
		return true
	}
	return false
}

type MetricSubscriberUpdateStrategy struct {
	analysisStrategy   analysis.MetricAnalysisStrategy
	regulationStrategy actions.MetricRegulationStrategy
}

func (strat *MetricSubscriberUpdateStrategy) create(
	analysisStrat analysis.MetricAnalysisStrategy,
	regulationStrat actions.MetricRegulationStrategy,
) *MetricSubscriberUpdateStrategy {
	return &MetricSubscriberUpdateStrategy{analysisStrat, regulationStrat}
}

func (strat *MetricSubscriberUpdateStrategy) update(snapshot *models.Snapshot) bool {
	if strat.analysisStrategy.Interpret(snapshot) == models.CRITICAL {
		strat.regulationStrategy.DispatchAction()
		return true
	}
	return false
}
