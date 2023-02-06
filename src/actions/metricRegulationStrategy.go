package actions

type iMetricRegulationStrategy interface {
	dispatchAction()
}

type MetricRegulationStrategy struct {
	actionFactory ActionFactory
	actionsStore  ActionsStore
}

// Probably subject to overriding or template method
func (strat *MetricRegulationStrategy) dispatchAction() {
	actionsStore.add(actionFactory.create())
}
