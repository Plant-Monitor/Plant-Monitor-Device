package snapshots

import (
	"fmt"
	"time"
	"pcs/actions"
	"pcs/analysis"
	"pcs/models"
	"pcs/pch"
	"pcs/utils"
	"sync"
)

var snapshotPublisherInstance *SnapshotPublisher
var snapshotPublisherLock *sync.Mutex = &sync.Mutex{}

type SnapshotPublisher struct {
	subscribers  []SnapshotSubscriber
	updater      SnapshotUpdater
	currentState *models.Snapshot
	pchClient    *pch.PCHClient
}

func (publisher *SnapshotPublisher) Run() {
	publisher.setup()
	for {
		publisher.updateState()
		publisher.notifySubscribers()
		//fmt.Println("[SnapshotPublisher] Snapshot has been published!")

		publisher.runActionsStore()
		//fmt.Println("[SnapshotPublisher] Actions store has been ran")
		time.Sleep(time.Second)
	}
}

func (publisher *SnapshotPublisher) setup() {
	publisher.loadSubscribers()
}

func (publisher *SnapshotPublisher) loadSubscribers() {
	tempSub := MetricSubscriber{
		updateStrategy: MetricSubscriberUpdateStrategy{
			analysisStrategy: analysis.NewThresholdAnalysisStrategy("temperature"),
			regulationStrategy: actions.NewNeededActionRegulationStrategy("temperature", 1),
		},
	}

	moistureSub := MetricSubscriber{
		updateStrategy: MetricSubscriberUpdateStrategy{
			analysisStrategy: analysis.NewThresholdAnalysisStrategy("moisture"),
			regulationStrategy: actions.NewNeededActionRegulationStrategy("moisture", 1),
		},
	}

	lightSub := MetricSubscriber{
		updateStrategy: MetricSubscriberUpdateStrategy{
			analysisStrategy: analysis.NewThresholdAnalysisStrategy("light intensity"),
			regulationStrategy: actions.NewNeededActionRegulationStrategy("light intensity", 1),
		},
	}

	tankLevelSub := MetricSubscriber{
		updateStrategy: MetricSubscriberUpdateStrategy{
			analysisStrategy: analysis.NewThresholdAnalysisStrategy("water level"),
			regulationStrategy: actions.NewNeededActionRegulationStrategy("water level", 1),
		},
	}

	humiditySub := MetricSubscriber{
		updateStrategy: MetricSubscriberUpdateStrategy{
			analysisStrategy: analysis.NewThresholdAnalysisStrategy("humidity"),
			regulationStrategy: actions.NewNeededActionRegulationStrategy("humidity", 1),
		},
	}

	publisher.Subscribe(&tempSub)
	publisher.Subscribe(&moistureSub)
	publisher.Subscribe(&lightSub)
	publisher.Subscribe(&tankLevelSub)
	publisher.Subscribe(&humiditySub)
}

func GetSnapshotPublisherInstance() *SnapshotPublisher {
	return utils.GetSingletonInstance(
		snapshotPublisherInstance,
		snapshotPublisherLock,
		newSnapshotPublisher,
		nil,
	)
}

func newSnapshotPublisher(initParams ...any) *SnapshotPublisher {
	return &SnapshotPublisher{
		subscribers: make([]SnapshotSubscriber, 0),
		updater:     *GetSnapshotUpdaterInstance(),
		pchClient:   pch.GetPCHClientInstance(),
	}
}

// Subscribe Add a subscriber to the publisher
func (publisher *SnapshotPublisher) Subscribe(sub SnapshotSubscriber) {
	publisher.subscribers = append(publisher.subscribers, sub)
}

// Notify the subscribers of the most current Snapshot stored in the publisher
func (publisher *SnapshotPublisher) notifySubscribers() {
	//fmt.Println("[SnapshotPublisher] Notifying subscribers")
	for _, sub := range publisher.subscribers {
		sub.update(publisher.currentState)
	}
	publisher.updater.update(publisher.currentState)
}

// Update the state of the SnapshotPublisher by reading the GPIO pins
func (publisher *SnapshotPublisher) updateState() {
	//fmt.Println("[SnapshotPublisher] Updating current snapshot")
	convertedReads := publisher.pchClient.GetReadings()
	publisher.currentState = publisher.buildSnapshot(convertedReads)

	//fmt.Printf("[SnapshotPublisher] Built snapshot %+v\n", publisher.currentState)
}

// Build snapshot from the provided readings collection
func (publisher *SnapshotPublisher) buildSnapshot(readings models.ConvertedReadingsCollection) *models.Snapshot {
	//fmt.Println("[SnapshotPublisher] Building snapshot")
	return models.BuildSnapshot(readings)
}

// Wrapper for ActionsStore.execute
func (publisher *SnapshotPublisher) runActionsStore() {
	err := actions.GetActionsStoreInstance().Execute()
	if err != nil {
		fmt.Printf("Error occured during actions execution %s\n", err)
		return
	}
}
