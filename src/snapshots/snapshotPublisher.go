package snapshots

import (
	"pcs/gpio"
	"pcs/models"
	"pcs/pin_reading_conversion"
	"pcs/utils"
	"sync"
)

var snapshotPublisherInstance *SnapshotPublisher
var snapshotPublisherLock *sync.Mutex = &sync.Mutex{}

type SnapshotPublisher struct {
	subscribers       []SnapshotSubscriber
	updater           SnapshotUpdater
	currentState      models.Snapshot
	gpioClient        *gpio.GpioClient
	readingsConverter *pin_reading_conversion.PinReadingsConverter
}

func (publisher *SnapshotPublisher) Run() {
	for {
		publisher.updateState()
		publisher.notifySubscribers()
	}
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
		subscribers:       make([]SnapshotSubscriber, 0),
		updater:           *GetSnapshotUpdaterInstance(),
		gpioClient:        gpio.GetGpioClientInstance(),
		readingsConverter: pin_reading_conversion.GetPinReadingsConverterInstance(),
	}
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
	publisher.updater.update(publisher.currentState)
}

// Update the state of the SnapshotPublisher by reading the GPIO pins
func (publisher *SnapshotPublisher) updateState() {
	pinReads := publisher.gpioClient.Read()
	convertedReads := publisher.readingsConverter.Convert(pinReads)
	publisher.currentState = publisher.buildSnapshot(convertedReads)
}

// Build snapshot from the provided readings collection
func (publisher *SnapshotPublisher) buildSnapshot(readings models.ConvertedReadingsCollection) models.Snapshot {
	return *models.BuildSnapshot(readings)
}
