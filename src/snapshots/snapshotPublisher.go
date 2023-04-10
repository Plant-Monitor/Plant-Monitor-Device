package snapshots

import (
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
	for {
		publisher.updateState()
		publisher.notifySubscribers()
	}
}

func (publisher *SnapshotPublisher) setup() {
	publisher.loadSubscribers()
}

func (publisher *SnapshotPublisher) loadSubscribers() {
	tempSub := MetricSubscriber{
		updateStrategy: MetricSubscriberUpdateStrategy{
			analysisStrategy: analysis.NewThresholdAnalysisStrategy("temperature"),
		},
	}

	moistureSub := MetricSubscriber{
		updateStrategy: MetricSubscriberUpdateStrategy{
			analysisStrategy: analysis.NewThresholdAnalysisStrategy("moisture"),
		},
	}

	lightSub := MetricSubscriber{
		updateStrategy: MetricSubscriberUpdateStrategy{
			analysisStrategy: analysis.NewThresholdAnalysisStrategy("light"),
		},
	}

	tankLevelSub := MetricSubscriber{
		updateStrategy: MetricSubscriberUpdateStrategy{
			analysisStrategy: analysis.NewThresholdAnalysisStrategy("tank level"),
		},
	}

	humiditySub := MetricSubscriber{
		updateStrategy: MetricSubscriberUpdateStrategy{
			analysisStrategy: analysis.NewThresholdAnalysisStrategy("humidity"),
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
	for _, sub := range publisher.subscribers {
		sub.update(publisher.currentState)
	}
	publisher.updater.update(publisher.currentState)
}

// Update the state of the SnapshotPublisher by reading the GPIO pins
func (publisher *SnapshotPublisher) updateState() {
	convertedReads := publisher.pchClient.GetReadings()
	publisher.currentState = publisher.buildSnapshot(convertedReads)
}

// Build snapshot from the provided readings collection
func (publisher *SnapshotPublisher) buildSnapshot(readings models.ConvertedReadingsCollection) *models.Snapshot {
	return models.BuildSnapshot(readings)
}
