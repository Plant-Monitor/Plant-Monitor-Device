package snapshots

import (
	"fmt"
	"pcs/gpio"
	"sync"
)

var snapshotPublisherLock = &sync.Mutex{}
var snapshotPublisherInstance *SnapshotPublisher

type SnapshotPublisher struct {
	subscribers  []SnapshotSubscriber
	currentState Snapshot
	gpioClient   *gpio.GpioClient
}

func (publisher *SnapshotPublisher) Run() {
	for {
		publisher.updateState()
		publisher.notifySubscribers()
	}
}

func getInstance() *SnapshotPublisher {
	if snapshotPublisherInstance == nil {
		snapshotPublisherLock.Lock()
		defer snapshotPublisherLock.Unlock()
		if snapshotPublisherInstance == nil {
			fmt.Println("Creating snapshotPublisher instance now.")
			snapshotPublisherInstance = &SnapshotPublisher{gpioClient: gpio.GetGpioClientInstance()}
		} else {
			fmt.Println("snapshotPublisher instance already created.")
		}
	} else {
		fmt.Println("snapshotPublisher instance already created.")
	}

	return snapshotPublisherInstance
}

// Add a subscriber to the publisher
func (publisher *SnapshotPublisher) Subscribe(sub SnapshotSubscriber) {
	publisher.subscribers = append(publisher.subscribers, sub)
}

// Notify the subscribers of the most current Snapshot stored in the publisher
func (publisher *SnapshotPublisher) notifySubscribers() {
	for _, sub := range publisher.subscribers {
		sub.update(publisher.currentState)
	}
}

// Update the state of the SnapshotPublisher by reading the GPIO pins
func (publisher *SnapshotPublisher) updateState() {
	currentReadings := publisher.gpioClient.Read()
	publisher.currentState = publisher.buildSnapshot(currentReadings)
}

func (publisher *SnapshotPublisher) buildSnapshot(readings map[string]float32) Snapshot {
	return Snapshot{}
}