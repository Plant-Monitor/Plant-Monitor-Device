package snapshots

import (
	"pcs/actions"
	"pcs/analysis"
	"pcs/models"
	"pcs/utils"
	"time"
	"fmt"
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
		fmt.Println("[snapshots] Writing to server\n\n\n\n\n\n\n\n\n\n\n\n\n")
		_, err := perUpdateStrategy.serverClient.WriteSnapshot(snapshot)
		if err != nil {
			return false
		}
		timeNow := time.Now()
		perUpdateStrategy.lastUpdate = &timeNow
		return true
	}
	fmt.Println("[snapshots] Not writing to the server yet")
	return false
}

type MetricSubscriberUpdateStrategy struct {
	analysisStrategy   analysis.IMetricAnalysisStrategy
	regulationStrategy actions.IMetricRegulationStrategy
}

func newMetricSubscriberUpdateStrategy(
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
