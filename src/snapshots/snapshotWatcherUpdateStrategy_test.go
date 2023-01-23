package snapshots

import (
	"pcs/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPeriodicUpdateStrategy(t *testing.T) {

	t.Run(
		"Should write to server since there hasn't been an update yet",
		func(t *testing.T) {
			updateInterval := "4h"

			strat := NewPeriodicUpdateStrategy(updateInterval)
			snapshot1 := *models.BuildSnapshot(models.ConvertedReadingsCollection{})

			assert.True(t, strat.update(snapshot1))
		},
	)

	t.Run(
		"Should write to server since time since snapshot timestamp is older than update interval",
		func(t *testing.T) {
			duration, _ := time.ParseDuration("4h25m")
			updateInterval := "4h"

			strat := NewPeriodicUpdateStrategy(updateInterval)
			snapshot1 := *models.BuildSnapshot(models.ConvertedReadingsCollection{})
			snapshot2 := models.Snapshot{
				Timestamp: time.Now().Add(duration),
			}

			assert.True(t, strat.update(snapshot1))
			assert.True(t, strat.update(snapshot2))
		},
	)

	t.Run(
		"Should not write to server since time since snapshot timestamp is newer than update interval",
		func(t *testing.T) {
			duration, _ := time.ParseDuration("3h")
			updateInterval := "4h"

			strat := NewPeriodicUpdateStrategy(updateInterval)
			snapshot1 := *models.BuildSnapshot(models.ConvertedReadingsCollection{})
			snapshot2 := models.Snapshot{
				Timestamp: time.Now().Add(duration),
			}

			assert.True(t, strat.update(snapshot1))
			assert.False(t, strat.update(snapshot2))
		},
	)
}
