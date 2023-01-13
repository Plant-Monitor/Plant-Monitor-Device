package snapshots

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type SnapshotPublisher struct {
	subscribers  []SnapshotSubscriber
	currentState Snapshot
}

var instance *SnapshotPublisher

func getInstance() *SnapshotPublisher {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			fmt.Println("Creating snapshotPublisher instance now.")
			instance = &SnapshotPublisher{}
		} else {
			fmt.Println("snapshotPublisher instance already created.")
		}
	} else {
		fmt.Println("snapshotPublisher instance already created.")
	}

	return instance
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

}
