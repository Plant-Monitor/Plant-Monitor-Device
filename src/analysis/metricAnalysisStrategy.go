package analysis

import (
	config "pcs/config/metric"
	"pcs/models"
	"fmt"
)

type IMetricAnalysisStrategy interface {
	analyze(level float64) models.Interpretation
	Interpret(i IMetricAnalysisStrategy, snapshot models.Snapshot) models.Interpretation
}

type metricAnalysisStrategy struct {
	metric models.Metric
}

func (strat *metricAnalysisStrategy) Interpret(i IMetricAnalysisStrategy, snapshot models.Snapshot) models.Interpretation {
	fmt.Printf("[Analysis] Interpreting %s\n", strat.metric)
	healthProp := snapshot.HealthProperties[strat.metric]
	interpretation := i.analyze(healthProp.Level)
	healthProp.Interpretation = interpretation
	return interpretation
}

type ThresholdAnalysisStrategy struct {
	metricAnalysisStrategy
}

func NewThresholdAnalysisStrategy(metric models.Metric) *ThresholdAnalysisStrategy {
	return &ThresholdAnalysisStrategy{
		metricAnalysisStrategy{metric: metric},
	}
}

func (strat *ThresholdAnalysisStrategy) analyze(level float64) models.Interpretation {
	threshCollection := config.GetThresholdCollection(strat.metric)

	switch {
	case level <= threshCollection.LowerCriticalThreshold || level >= threshCollection.UpperCriticalThreshold:
		return models.CRITICAL
	case level >= threshCollection.GoodMinThreshold && level <= threshCollection.GoodMaxThreshold:
		return models.GOOD
	default:
		return models.OKAY
	}

}

// todo(pcs-51) Subclass of thresholdAnalysisStrategy that only uses the threshold analysis during certain hours of the day
type TimeSpannedThresholdAnalysisStrategy struct {
	ThresholdAnalysisStrategy
}

// todo(pcs-51) make constructor for TimeSpannedThresholdAnalysisStrategy
