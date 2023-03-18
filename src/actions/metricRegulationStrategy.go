package actions

import "pcs/models"

type iMetricRegulationStrategy interface {
	dispatchAction()
	decide(snapshot models.Snapshot) bool
	Regulate(snapshot models.Snapshot)
}

type MetricRegulationStrategy struct {
	iMetricRegulationStrategy
	actionFactory actionFactory
	actionsStore  *actionsStore
	metric        models.Metric
}

// Probably subject to overriding or template method
func (strat *MetricRegulationStrategy) dispatchAction() {
	strat.actionsStore.add(strat.actionFactory.create())
}

func (strat *MetricRegulationStrategy) Regulate(snapshot models.Snapshot) {
	if strat.decide(snapshot) {
		strat.dispatchAction()
	}
}
