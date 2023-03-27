package snapshots_test

import (
	"pcs/snapshots"
	"pcs/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSnapshotUpdater(t *testing.T) {
	utils.LoadEnv()
	t.Run("No initial SnapshotUpdater instance", func(t *testing.T) {
		rslt := snapshots.GetSnapshotUpdaterInstance()
		assert.NotNil(t, rslt)
	})

	t.Run("There is an initial instance, returned instance should be the same", func(t *testing.T) {
		instance := snapshots.GetSnapshotUpdaterInstance()
		instance2 := snapshots.GetSnapshotUpdaterInstance()

		assert.Equal(t, instance, instance2)
	})
}

func TestMetricSubscriberUpdateStrategy(t *testing.T) {
	//setup := func() {
	//	err := godotenv.Load("../../.env")
	//	if err != nil {
	//		return
	//	}
	//}
	//
	//t.Run("Should update the snapshot metric's interpretation", func(t *testing.T) {
	//	var testedMetric models.Metric
	//	var updateStrategy *snapshots.MetricSubscriberUpdateStrategy
	//
	//	setup := func() {
	//		setup()
	//		testedMetric = "moisture"
	//		updateStrategy = create(
	//			analysis.NewThresholdAnalysisStrategy(testedMetric),
	//			actions.IMetricRegulationStrategy
	//			)
	//		strategy = analysis.NewThresholdAnalysisStrategy(testedMetric)
	//	}
	//
	//	t.Run("HealthProperty should be given the interpretation CRITICAL", func(t *testing.T){
	//		setup()
	//		snapshot :=
	//	})
	//})
}
