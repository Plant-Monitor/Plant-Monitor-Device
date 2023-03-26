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
		_, err := perUpdateStrategy.serverClient.WriteSnapshot(snapshot)
		if err != nil {
			return false
		}
		timeNow := time.Now()
		perUpdateStrategy.lastUpdate = &timeNow
		return true
	}
	return false
}

type MetricSubscriberUpdateStrategy struct {
	analysisStrategy   analysis.IMetricAnalysisStrategy
	regulationStrategy actions.IMetricRegulationStrategy
}

func (strat *MetricSubscriberUpdateStrategy) create(
	analysisStrat analysis.IMetricAnalysisStrategy,
	regulationStrat actions.IMetricRegulationStrategy,
) *MetricSubscriberUpdateStrategy {
	return &MetricSubscriberUpdateStrategy{analysisStrat, regulationStrat}
}

func (strat *MetricSubscriberUpdateStrategy) update(snapshot *models.Snapshot) bool {
	strat.analysisStrategy.Interpret(strat.analysisStrategy, *snapshot)
	//strat.regulationStrategy.Regulate(*snapshot)

	return false
}
