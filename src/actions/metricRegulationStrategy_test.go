package actions

import (
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"os"
	"pcs/models"
	"testing"
	"time"
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
		assert.Equal(t, 10.0, threshold)
	})

	t.Run("Strategy should return the maximum good threshold when the metric is in the upper critical range", func(t *testing.T) {
		setup()
		strat := newNeededActionRegulationStrategy("testMetric", 1)
		threshold := strat.determineLevelNeeded(CRITICAL_HIGH)
		assert.Equal(t, 15.0, threshold)
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

	t.Run("Strat should determine that timer is expired", func(t *testing.T) {
		setup()

		strat := newNeededActionRegulationStrategy("testMetric", 1)
		strat.startTimer()
		time.Sleep(62 * time.Second)
		assert.True(t, strat.isTimerExpired())
	})

	t.Run("Strat should determine that timer is still running", func(t *testing.T) {
		setup()

		strat := newNeededActionRegulationStrategy("testMetric", 1)
		strat.startTimer()
		time.Sleep(10 * time.Second)
		assert.False(t, strat.isTimerExpired())
	})

	t.Run("Strat should decide against action if timer isn't expired", func(t *testing.T) {
		setup()
		strat := newNeededActionRegulationStrategy("testMetric", 1)
		strat.startTimer()
		time.Sleep(10 * time.Second)
		decision, critRange, _ := strat.decide(models.Snapshot{
			HealthProperties: models.ConvertedReadingsCollection{
				"testMetric": &models.HealthProperty{
					Level:          20.5,
					Unit:           "arb units",
					Interpretation: models.CRITICAL,
				},
			},
		})
		assert.False(t, decision)
		assert.Equal(t, NOT_CRITICAL, critRange)
	})

	t.Run("Strat should decide for action if timer is expired", func(t *testing.T) {
		setup()
		strat := newNeededActionRegulationStrategy("testMetric", 1)
		strat.startTimer()
		time.Sleep(62 * time.Second)
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

	t.Run("Strat should resolve an action", func(t *testing.T) {
		setup()
		strat := newNeededActionRegulationStrategy("testMetric", 1)
		userId, _ := uuid.Parse(os.Getenv("USER_ID"))

		criticalSnapshot := models.Snapshot{
			UserId:    userId,
			PlantId:   userId,
			Timestamp: time.Now(),
			HealthProperties: models.ConvertedReadingsCollection{
				"testMetric": &models.HealthProperty{
					Level:          3,
					Unit:           "arb units",
					Interpretation: models.CRITICAL,
				},
			},
		}
		strat.Regulate(strat, criticalSnapshot)
		assert.True(t, len(storeDict) > 0)
		assert.NotEqual(t, strat.activeActionId, uuid.Nil)

		var originalId uuid.UUID
		for id, _ := range storeDict {
			originalId = id
		}

		assert.Equal(t, originalId, strat.activeActionId)

		GetActionsStoreInstance().Execute()

		goodSnapshot := models.Snapshot{
			UserId:    userId,
			PlantId:   userId,
			Timestamp: time.Now(),
			HealthProperties: models.ConvertedReadingsCollection{
				"testMetric": &models.HealthProperty{
					Level:          14,
					Unit:           "arb units",
					Interpretation: models.GOOD,
				},
			},
		}
		strat.Regulate(strat, goodSnapshot)
		assert.True(t, len(storeDict) == 0)
		assert.Equal(t, strat.activeActionId, uuid.Nil)
	})

	t.Run("Strat shouldn't regulate if timer isn't expired yet", func(t *testing.T) {
		setup()
		strat := newNeededActionRegulationStrategy("testMetric", 1)
		userId := uuid.New()

		criticalSnapshot := models.Snapshot{
			UserId:    userId,
			PlantId:   userId,
			Timestamp: time.Now(),
			HealthProperties: models.ConvertedReadingsCollection{
				"testMetric": &models.HealthProperty{
					Level:          3,
					Unit:           "arb units",
					Interpretation: models.CRITICAL,
				},
			},
		}
		strat.Regulate(strat, criticalSnapshot)
		assert.True(t, len(storeDict) == 1)
		var actionId uuid.UUID
		for id, _ := range storeDict {
			actionId = id
		}

		err := GetActionsStoreInstance().Execute()
		assert.Nil(t, err)
		assert.True(t, len(storeDict) == 0)

		nextSnapshot := models.Snapshot{
			UserId:    userId,
			PlantId:   userId,
			Timestamp: time.Now(),
			HealthProperties: models.ConvertedReadingsCollection{
				"testMetric": &models.HealthProperty{
					Level:          4,
					Unit:           "arb units",
					Interpretation: models.CRITICAL,
				},
			},
		}
		strat.Regulate(strat, nextSnapshot)
		assert.True(t, len(storeDict) == 0)

		_, isInStore := storeDict[actionId]
		assert.False(t, isInStore)
	})

	t.Run("Strat should regulate if timer is expired yet", func(t *testing.T) {
		setup()
		strat := newNeededActionRegulationStrategy("testMetric", 1)
		userId := uuid.New()

		criticalSnapshot := models.Snapshot{
			UserId:    userId,
			PlantId:   userId,
			Timestamp: time.Now(),
			HealthProperties: models.ConvertedReadingsCollection{
				"testMetric": &models.HealthProperty{
					Level:          3,
					Unit:           "arb units",
					Interpretation: models.CRITICAL,
				},
			},
		}
		strat.Regulate(strat, criticalSnapshot)
		assert.True(t, len(storeDict) == 1)
		var actionId uuid.UUID
		for id, _ := range storeDict {
			actionId = id
		}

		err := GetActionsStoreInstance().Execute()
		assert.Nil(t, err)
		_, isInStore := storeDict[actionId]
		assert.False(t, isInStore)

		time.Sleep(62 * time.Second)

		nextSnapshot := models.Snapshot{
			UserId:    userId,
			PlantId:   userId,
			Timestamp: time.Now(),
			HealthProperties: models.ConvertedReadingsCollection{
				"testMetric": &models.HealthProperty{
					Level:          4,
					Unit:           "arb units",
					Interpretation: models.CRITICAL,
				},
			},
		}
		strat.Regulate(strat, nextSnapshot)
		assert.True(t, len(storeDict) == 1)

		assert.NotNil(t, strat.activeActionId)
		//assert.True(t, len(storeDict) == 0)
		//_, isInStore = storeDict[actionId]
		//assert.False(t, isInStore)
	})
}
