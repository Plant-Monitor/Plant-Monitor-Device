package analysis

import (
	config "pcs/config/metric"
	"pcs/models"
)

type iMetricAnalysisStrategy interface {
	analyze(level float32) models.Interpretation
	Interpret(snapshot models.Snapshot) models.Interpretation
}

type metricAnalysisStrategy struct {
	metric models.Metric

	analyze func(level float32) models.Interpretation
}

func (strat *metricAnalysisStrategy) Interpret(snapshot models.Snapshot) models.Interpretation {
	healthProp := snapshot.HealthProperties[strat.metric]
	interpretation := strat.analyze(healthProp.Level)
	healthProp.Interpretation = interpretation
	return interpretation
}

type ThresholdAnalysisStrategy struct {
	metricAnalysisStrategy
}

func NewThresholdAnalysisStrategy(metric models.Metric) *ThresholdAnalysisStrategy {
	instance := &ThresholdAnalysisStrategy{
		metricAnalysisStrategy{metric: metric},
	}
	instance.metricAnalysisStrategy.analyze = instance.analyze

	return instance
}

func (strat *ThresholdAnalysisStrategy) analyze(level float32) models.Interpretation {
	threshCollection := (*config.GetMetricAnalysisThresholdMapInstance())[strat.metric]

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
