package config

import (
	"pcs/models"
	"pcs/utils"
	"sync"
)

type MetricAnalysisThresholdMap map[models.Metric]ThresholdCollection

var thresholdMapInstance *MetricAnalysisThresholdMap
var thresholdMapLock = &sync.Mutex{}

type ThresholdCollection struct {
	LowerCriticalThreshold float32 `json:"lower_critical_threshold"`
	UpperCriticalThreshold float32 `json:"upper_critical_threshold"`
	GoodMinThreshold       float32 `json:"good_min_threshold"`
	GoodMaxThreshold       float32 `json:"good_max_threshold"`
}

func GetMetricAnalysisThresholdMapInstance() *MetricAnalysisThresholdMap {
	return utils.GetSingletonInstance(
		thresholdMapInstance,
		thresholdMapLock,
		newThresholdMap,
		nil,
	)
}

func newThresholdMap(initParams ...any) *MetricAnalysisThresholdMap {
	var instance MetricAnalysisThresholdMap
	// todo(pcs-51) complete json loading and unmarshalling

	return &instance
}