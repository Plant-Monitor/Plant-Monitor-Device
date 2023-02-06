package actions

type iMetricRegulationStrategy interface {
	dispatchAction()
}

type MetricRegulationStrategy struct {
	actionFactory actionFactory
	actionsStore  *actionsStore
}

// Probably subject to overriding or template method
func (strat *MetricRegulationStrategy) DispatchAction() {
	strat.actionsStore.add(strat.actionFactory.create())
}
