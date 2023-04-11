package actions

import (
	"github.com/google/uuid"
	config "pcs/config/metric"
	"pcs/models"
	"pcs/models/dto"
	"pcs/utils"
	"time"
	"fmt"
)

type IMetricRegulationStrategy interface {
	decide(snapshot models.Snapshot) (
		decision bool,
		critRange criticalRange,
		actType actionType,
	)
	Regulate(i IMetricRegulationStrategy, snapshot models.Snapshot)
	determineCallback(snapshot models.Snapshot, critRange criticalRange) ActionExecutionCallback
	determineResolution(snapshot models.Snapshot) bool
	isTimerExpired() bool
}

type metricRegulationStrategy struct {
	//iMetricRegulationStrategy
	metric         models.Metric
	checkInterval  time.Duration // expressed in hours
	checkTimer     *time.Timer
	//cooldown	   *time.Timer
	activeActionId uuid.UUID
	pendingAction  bool
}

func (strat *metricRegulationStrategy) Regulate(i IMetricRegulationStrategy, snapshot models.Snapshot) {
	if strat.determineResolution(snapshot) {
		return
	}

	decision, critRange, actType := i.decide(snapshot)
	if decision {
		levelNeeded := strat.determineLevelNeeded(critRange)
		strat.activeActionId = newAction(
			actType,
			levelNeeded,
			strat.metric,
			&snapshot,
			i.determineCallback(snapshot, critRange),
		)
		strat.pendingAction = true
		fmt.Printf("[Actions] New action has been initialized\n")
	}
}

func (strat *metricRegulationStrategy) determineResolution(snapshot models.Snapshot) bool {
	interpretation := snapshot.HealthProperties[strat.metric].Interpretation
	if (strat.activeActionId != uuid.Nil) && (interpretation == models.GOOD) {
		strat.resolveActiveAction(snapshot)
		return true
	}
	return false
}

func (strat *metricRegulationStrategy) resolveActiveAction(snapshot models.Snapshot) {
	fmt.Printf("[Actions] Resolving %s\n", strat.metric)
	err := utils.GetServerClientInstance().ResolveAction(
		dto.CreateResolveActionDto(
			strat.activeActionId,
			snapshot,
		),
	)
	if err != nil {
		return
	}
	strat.resetRegulation()
}

func (strat *metricRegulationStrategy) resetRegulation() {
	strat.activeActionId = uuid.Nil
	strat.checkTimer = nil
	
}

func (strat *metricRegulationStrategy) determineLevelNeeded(critRange criticalRange) float64 {
	switch critRange {
	case CRITICAL_LOW:
		return config.GetThresholdCollection(strat.metric).GoodMinThreshold
	case CRITICAL_HIGH:
		return config.GetThresholdCollection(strat.metric).GoodMaxThreshold
	}

	return 0
}

type neededActionRegulationStrategy struct {
	metricRegulationStrategy
}

func NewNeededActionRegulationStrategy(metric models.Metric, checkInterval time.Duration) *neededActionRegulationStrategy {
	return &neededActionRegulationStrategy{
		metricRegulationStrategy: metricRegulationStrategy{
			metric:         metric,
			checkInterval:  checkInterval,
			checkTimer:     nil,
			//cooldown:		nil,
			activeActionId: uuid.Nil,
			pendingAction:  false,
		},
	}
}

func (strat *neededActionRegulationStrategy) decide(snapshot models.Snapshot) (
	decision bool,
	critRange criticalRange,
	actType actionType,
) {
	healthProp := snapshot.HealthProperties[strat.metric]
	if strat.isTimerExpired() && !strat.pendingAction {
		return strat.determineDecision(healthProp)
	}
	return false, NOT_CRITICAL, NEEDED

}

func (strat *metricRegulationStrategy) isTimerExpired() bool {
	if strat.checkTimer == nil {
		return true
	}
	select {
	case <-strat.checkTimer.C:
		return true
	default:
		return false
	}
}

func (strat *metricRegulationStrategy) startTimer() {
	strat.checkTimer = time.NewTimer(time.Minute * strat.checkInterval)
	//strat.cooldown = time.NewTimer(time.Second * 30)
}

func (strat *neededActionRegulationStrategy) determineDecision(healthProp *models.HealthProperty) (
	decision bool,
	critRange criticalRange,
	actType actionType,
) {
	if healthProp.Interpretation == models.CRITICAL {
		if healthProp.Level <= config.GetThresholdCollection(strat.metric).LowerCriticalThreshold {
			return true, CRITICAL_LOW, NEEDED
		}
		return true, CRITICAL_HIGH, NEEDED
	}

	return false, NOT_CRITICAL, NEEDED
}

func (strat *neededActionRegulationStrategy) determineCallback(snapshot models.Snapshot, critRange criticalRange) ActionExecutionCallback {
	return func() error {
		fmt.Printf("[Actions] Action execution for %s has been called!\n",strat.metric)
		strat.startTimer()
		strat.pendingAction = false
		return nil
	}
}
