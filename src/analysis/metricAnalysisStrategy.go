package analysis

import "pcs/models"

type iMetricAnalysisStrategy interface {
	analyze(snapshot *models.Snapshot) models.Interpretation
	Interpret(snapshot *models.Snapshot) models.Interpretation
}

type MetricAnalysisStrategy struct {
	metric models.Metric

	analyze func(snapshot *models.Snapshot) models.Interpretation
}

func (strat *MetricAnalysisStrategy) Interpret(snapshot *models.Snapshot) models.Interpretation {
	interpretation := strat.analyze(snapshot)
	snapshot.Health_properties[strat.metric].Interpretation = interpretation
	return interpretation
}
