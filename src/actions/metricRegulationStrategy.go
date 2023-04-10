package actions

import (
	config "pcs/config/metric"
	"pcs/models"
	"time"
)

type IMetricRegulationStrategy interface {
	dispatchAction(actType actionType)
	decide(snapshot models.Snapshot) (
		decision bool,
		critRange criticalRange,
		actType actionType,
	)
	Regulate(i IMetricRegulationStrategy, snapshot models.Snapshot)
	determineCallback(snapshot models.Snapshot, critRange criticalRange) ActionExecutionCallback
}

type metricRegulationStrategy struct {
	//iMetricRegulationStrategy
	metric models.Metric
}

// Probably subject to overriding or template method
//func (strat *metricRegulationStrategy) dispatchAction(
//	actType actionType,
//	critRange criticalRange,
//) {
//	newAction()
//}

func (strat *metricRegulationStrategy) Regulate(i IMetricRegulationStrategy, snapshot models.Snapshot) {
	var levelNeeded float32
	decision, critRange, actType := i.decide(snapshot)
	if decision {
		levelNeeded = strat.determineLevelNeeded(critRange)
		newAction(
			actType,
			levelNeeded,
			strat.metric,
			&snapshot,
			i.determineCallback(snapshot, critRange),
		)
	}
}

func (strat *metricRegulationStrategy) determineLevelNeeded(critRange criticalRange) float32 {
	switch critRange {
	case CRITICAL_LOW:
		return config.GetThresholdCollection(strat.metric).GoodMinThreshold
	case CRITICAL_HIGH:
		return config.GetThresholdCollection(strat.metric).GoodMaxThreshold
	}

	return 0
}

type neededActionRegulationStrategy struct {
	metricRegulationStrategy
	checkInterval time.Duration // expressed in hours
	checkTimer    *time.Timer
}

func newNeededActionRegulationStrategy(metric models.Metric, checkInterval time.Duration) *neededActionRegulationStrategy {
	return &neededActionRegulationStrategy{
		metricRegulationStrategy: metricRegulationStrategy{
			metric: metric,
		},
		checkInterval: checkInterval,
	}
}

func (strat *neededActionRegulationStrategy) decide(snapshot models.Snapshot) (
	decision bool,
	critRange criticalRange,
	actType actionType,
) {
	healthProp := snapshot.HealthProperties[strat.metric]
	if strat.checkTimer != nil {
		select {
		case <-strat.checkTimer.C:
			return strat.determineDecision(healthProp)
		default:
			return false, NOT_CRITICAL, NEEDED
		}
	}

	return strat.determineDecision(healthProp)

}

func (strat *neededActionRegulationStrategy) determineDecision(healthProp *models.HealthProperty) (
	decision bool,
	critRange criticalRange,
	actType actionType,
) {
	if healthProp.Interpretation == models.CRITICAL {
		if healthProp.Level <= config.GetThresholdCollection(strat.metric).LowerCriticalThreshold {
			return true, CRITICAL_LOW, NEEDED
		}
		return true, CRITICAL_HIGH, NEEDED
	}

	return false, NOT_CRITICAL, NEEDED
}

func (strat *neededActionRegulationStrategy) determineCallback(snapshot models.Snapshot, critRange criticalRange) ActionExecutionCallback {
	return func() error {
		strat.checkTimer = time.NewTimer(time.Minute * strat.checkInterval)
		return nil
	}
}
