package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"pcs/models"
	"pcs/utils"
	"sync"
)

type MetricAnalysisThresholdMap map[models.Metric]ThresholdCollection

var thresholdMapInstance *MetricAnalysisThresholdMap
var thresholdMapLock = &sync.Mutex{}

type ThresholdCollection struct {
	LowerCriticalThreshold float64 `json:"lower_critical_threshold"`
	UpperCriticalThreshold float64 `json:"upper_critical_threshold"`
	GoodMinThreshold       float64 `json:"good_min_threshold"`
	GoodMaxThreshold       float64 `json:"good_max_threshold"`
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
	return loadThresholdMap()
}

func GetThresholdCollection(metric models.Metric) ThresholdCollection {
	return (*GetMetricAnalysisThresholdMapInstance())[metric]
}

func loadThresholdMap() *MetricAnalysisThresholdMap {

	content, err := os.ReadFile(
		fmt.Sprintf(
			"%s/src/config/metric/analysisThresholds.json",
			os.Getenv("PATH_TO_PROJECT"),
		),
	)

	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// Unmarshalling data
	var payload MetricAnalysisThresholdMap
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return &payload
}
