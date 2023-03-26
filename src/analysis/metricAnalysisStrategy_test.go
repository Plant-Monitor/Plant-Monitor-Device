package analysis

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"pcs/models"
	"testing"
)

func TestThresholdAnalysisStrategy(t *testing.T) {
	setup := func() {
		err := godotenv.Load("../../.env")
		if err != nil {
			return
		}
	}

	t.Run(
		"Testing analyze",
		func(t *testing.T) {

			strategy := NewThresholdAnalysisStrategy("moisture")
			setup := func() {
				fmt.Println("saying hi before tests run")
				setup()
			}

			t.Run("Should return critical when level is too high", func(t *testing.T) {
				setup()
				assert.Equal(t, models.CRITICAL, strategy.analyze(21))
			})
			t.Run("Should return critical when level is too low", func(t *testing.T) {
				setup()
				assert.Equal(t, models.CRITICAL, strategy.analyze(4))
			})
			t.Run("Should return good when level is within the right range", func(t *testing.T) {
				setup()
				assert.Equal(t, models.GOOD, strategy.analyze(12))
			})
			t.Run("Should return OKAY when level is in lower range", func(t *testing.T) {
				setup()
				assert.Equal(t, models.OKAY, strategy.analyze(7))
			})
			t.Run("Should return OKAY when level is in upper range", func(t *testing.T) {
				setup()
				assert.Equal(t, models.OKAY, strategy.analyze(17.3))
			})
		},
	)

	t.Run("Testing ThresholdAnalysisStrategy.Interpret", func(t *testing.T) {
		var testedMetric models.Metric
		var strategy *ThresholdAnalysisStrategy

		setup := func() {
			setup()
			testedMetric = "moisture"
			strategy = NewThresholdAnalysisStrategy(testedMetric)

		}

		t.Run("Should return critical when level is too high", func(t *testing.T) {
			setup()
			snapshot := models.Snapshot{
				HealthProperties: models.ConvertedReadingsCollection{
					"moisture":    &models.HealthProperty{Level: 21, Unit: "arbitrary"},
					"temperature": &models.HealthProperty{Unit: "deg C"},
				},
			}

			strategy.Interpret(strategy, snapshot)
			assert.Equal(t, models.CRITICAL, snapshot.HealthProperties[testedMetric].Interpretation)
		})
		t.Run("Should return critical when level is too low", func(t *testing.T) {
			setup()
			snapshot := models.Snapshot{
				HealthProperties: models.ConvertedReadingsCollection{
					"moisture":    &models.HealthProperty{Level: 4, Unit: "arbitrary"},
					"temperature": &models.HealthProperty{Unit: "deg C"},
				},
			}

			strategy.Interpret(strategy, snapshot)
			assert.Equal(t, models.CRITICAL, snapshot.HealthProperties[testedMetric].Interpretation)
		})
		t.Run("Should return good when level is within the right range", func(t *testing.T) {
			setup()
			snapshot := models.Snapshot{
				HealthProperties: models.ConvertedReadingsCollection{
					"moisture":    &models.HealthProperty{Level: 12, Unit: "arbitrary"},
					"temperature": &models.HealthProperty{Unit: "deg C"},
				},
			}

			strategy.Interpret(strategy, snapshot)
			assert.Equal(t, models.GOOD, snapshot.HealthProperties[testedMetric].Interpretation)
		})
		t.Run("Should return OKAY when level is in lower range", func(t *testing.T) {
			setup()
			snapshot := models.Snapshot{
				HealthProperties: models.ConvertedReadingsCollection{
					"moisture":    &models.HealthProperty{Level: 7, Unit: "arbitrary"},
					"temperature": &models.HealthProperty{Unit: "deg C"},
				},
			}

			strategy.Interpret(strategy, snapshot)
			assert.Equal(t, models.OKAY, snapshot.HealthProperties[testedMetric].Interpretation)
		})
		t.Run("Should return OKAY when level is in upper range", func(t *testing.T) {
			setup()
			snapshot := models.Snapshot{
				HealthProperties: models.ConvertedReadingsCollection{
					"moisture":    &models.HealthProperty{Level: 17.3, Unit: "arbitrary"},
					"temperature": &models.HealthProperty{Unit: "deg C"},
				},
			}

			strategy.Interpret(strategy, snapshot)
			assert.Equal(t, models.OKAY, snapshot.HealthProperties[testedMetric].Interpretation)
		})
	})
}
