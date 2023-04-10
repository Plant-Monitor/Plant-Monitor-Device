package actions

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"pcs/models"
	"testing"
)

func TestNeededActionRegulation(t *testing.T) {
	setup := func() {
		err := godotenv.Load("../../.env")
		if err != nil {
			return
		}
	}

	t.Run("Strategy should return the minimum good threshold when the metric is in the lower critical range", func(t *testing.T) {
		setup()
		strat := newNeededActionRegulationStrategy("testMetric", 1)
		threshold := strat.determineLevelNeeded(CRITICAL_LOW)
		assert.Equal(t, float32(10), threshold)
	})

	t.Run("Strategy should return the maximum good threshold when the metric is in the upper critical range", func(t *testing.T) {
		setup()
		strat := newNeededActionRegulationStrategy("testMetric", 1)
		threshold := strat.determineLevelNeeded(CRITICAL_HIGH)
		assert.Equal(t, float32(15), threshold)
	})

	t.Run("Decide should return true when the metric is in lower critical range", func(t *testing.T) {
		setup()
		strat := newNeededActionRegulationStrategy("testMetric", 1)
		decision, _, _ := strat.decide(models.Snapshot{
			HealthProperties: models.ConvertedReadingsCollection{
				"testMetric": &models.HealthProperty{
					Level:          3,
					Unit:           "arb units",
					Interpretation: models.CRITICAL,
				},
			},
		})
		assert.True(t, decision)
	})

	t.Run("Decide should return true when the metric is in upper critical range", func(t *testing.T) {
		setup()
		strat := newNeededActionRegulationStrategy("testMetric", 1)
		decision, _, _ := strat.decide(models.Snapshot{
			HealthProperties: models.ConvertedReadingsCollection{
				"testMetric": &models.HealthProperty{
					Level:          20.4,
					Unit:           "arb units",
					Interpretation: models.CRITICAL,
				},
			},
		})
		assert.True(t, decision)
	})

	t.Run("Decide should return false when interpretation is OKAY", func(t *testing.T) {
		setup()
		strat := newNeededActionRegulationStrategy("testMetric", 1)
		decision, _, _ := strat.decide(models.Snapshot{
			HealthProperties: models.ConvertedReadingsCollection{
				"testMetric": &models.HealthProperty{
					Level:          19,
					Unit:           "arb units",
					Interpretation: models.OKAY,
				},
			},
		})
		assert.False(t, decision)
	})

	t.Run("Decide should return false when interpretation is GOOD", func(t *testing.T) {
		setup()
		strat := newNeededActionRegulationStrategy("testMetric", 1)
		decision, _, _ := strat.decide(models.Snapshot{
			HealthProperties: models.ConvertedReadingsCollection{
				"testMetric": &models.HealthProperty{
					Level:          13,
					Unit:           "arb units",
					Interpretation: models.GOOD,
				},
			},
		})
		assert.False(t, decision)
	})

	t.Run("Decide should return CRITICAL_LOW critical range when the metric is in lower critical range", func(t *testing.T) {
		setup()
		strat := newNeededActionRegulationStrategy("testMetric", 1)
		decision, critRange, _ := strat.decide(models.Snapshot{
			HealthProperties: models.ConvertedReadingsCollection{
				"testMetric": &models.HealthProperty{
					Level:          3,
					Unit:           "arb units",
					Interpretation: models.CRITICAL,
				},
			},
		})
		assert.True(t, decision)
		assert.Equal(t, CRITICAL_LOW, critRange)
	})

	t.Run("Decide should return CRITICAL_HIGH critical range when the metric is in upper critical range", func(t *testing.T) {
		setup()
		strat := newNeededActionRegulationStrategy("testMetric", 1)
		decision, critRange, _ := strat.decide(models.Snapshot{
			HealthProperties: models.ConvertedReadingsCollection{
				"testMetric": &models.HealthProperty{
					Level:          20.5,
					Unit:           "arb units",
					Interpretation: models.CRITICAL,
				},
			},
		})
		assert.True(t, decision)
		assert.Equal(t, CRITICAL_HIGH, critRange)
	})

	t.Run("Decide should return NOT_CRITICAL critical range when the metric is in lower critical range", func(t *testing.T) {
		setup()
		strat := newNeededActionRegulationStrategy("testMetric", 1)
		decision, critRange, _ := strat.decide(models.Snapshot{
			HealthProperties: models.ConvertedReadingsCollection{
				"testMetric": &models.HealthProperty{
					Level:          14,
					Unit:           "arb units",
					Interpretation: models.GOOD,
				},
			},
		})
		assert.False(t, decision)
		assert.Equal(t, NOT_CRITICAL, critRange)
	})

	t.Run("Strat should decide against taking action if check timer isn't expired yet", func(t *testing.T) {
		setup()
		strat := newNeededActionRegulationStrategy("testMetric", 1)
		
		decision, critRange, _ := strat.decide(models.Snapshot{
			HealthProperties: models.ConvertedReadingsCollection{
				"testMetric": &models.HealthProperty{
					Level:          14,
					Unit:           "arb units",
					Interpretation: models.GOOD,
				},
			},
		})
		assert.False(t, decision)
		assert.Equal(t, NOT_CRITICAL, critRange)
	})
}
