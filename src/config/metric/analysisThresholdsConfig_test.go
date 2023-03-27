package config

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestMetricAnalysisThresholdMap(t *testing.T) {
	err := godotenv.Load("../../.env")
	if err != nil {
		return
	}

	t.Run("Test loadThresholdMap", func(t *testing.T) {
		expected := MetricAnalysisThresholdMap{
			"moisture":    {5, 20, 10, 15},
			"temperature": {5, 20, 10, 15},
			"soil_health": {5, 20, 10, 15},
			"tank_level":  {5, 20, 10, 15},
			"light":       {5, 20, 10, 15},
		}

		rslt := loadThresholdMap()
		assert.True(t, reflect.DeepEqual(*rslt, expected))
	})

	t.Run("Test newThresholdMap", func(t *testing.T) {
		expected := MetricAnalysisThresholdMap{
			"moisture":    {5, 20, 10, 15},
			"temperature": {5, 20, 10, 15},
			"soil_health": {5, 20, 10, 15},
			"tank_level":  {5, 20, 10, 15},
			"light":       {5, 20, 10, 15},
		}

		rslt := newThresholdMap()
		assert.True(t, reflect.DeepEqual(*rslt, expected))
	})

	t.Run("Test GetMetricAnalysisThresholdMapInstance", func(t *testing.T) {
		expected := MetricAnalysisThresholdMap{
			"moisture":    {5, 20, 10, 15},
			"temperature": {5, 20, 10, 15},
			"soil_health": {5, 20, 10, 15},
			"tank_level":  {5, 20, 10, 15},
			"light":       {5, 20, 10, 15},
		}
		assert.True(t, reflect.DeepEqual(*GetMetricAnalysisThresholdMapInstance(), expected))
	})
}
