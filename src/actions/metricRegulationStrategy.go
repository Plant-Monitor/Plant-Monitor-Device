package actions

import "pcs/models"

type IMetricRegulationStrategy interface {
	dispatchAction()
	decide(snapshot models.Snapshot) bool
	Regulate(i IMetricRegulationStrategy, snapshot models.Snapshot)
}

type metricRegulationStrategy struct {
	//iMetricRegulationStrategy
	actionsStore *actionsStore
	metric       models.Metric
}

// Probably subject to overriding or template method
func (strat *metricRegulationStrategy) dispatchAction() {
	strat.actionsStore.add(strat.actionFactory.create())
}

func (strat *metricRegulationStrategy) Regulate(i IMetricRegulationStrategy, snapshot models.Snapshot) {
	if i.decide(snapshot) {
		strat.dispatchAction()
	}
}
