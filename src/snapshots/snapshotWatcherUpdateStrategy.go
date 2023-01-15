package snapshots

import (
	"context"
	"os"
	"pcs/utils"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type SnapshotWatcherUpdateStrategy interface {
	update(Snapshot)
}

type PeriodicUpdateStrategy struct {
	lastUpdate      *time.Time
	updateInterval  time.Duration
	mongoCollection *mongo.Collection
}

func NewPeriodicUpdateStrategy(updateInterval time.Duration) *PeriodicUpdateStrategy {
	env := os.Getenv("ENV")

	return &PeriodicUpdateStrategy{
		nil,
		updateInterval,
		utils.SetupMongoConnection().Database(env).Collection("snapshotdocobjects"),
	}
}

func (perUpdateStrategy *PeriodicUpdateStrategy) update(snapshot Snapshot) {
	if perUpdateStrategy.lastUpdate == nil || snapshot.Timestamp.Sub(*perUpdateStrategy.lastUpdate) >= perUpdateStrategy.updateInterval {
		perUpdateStrategy.mongoCollection.InsertOne(
			context.TODO(),
			snapshot,
		)
		*perUpdateStrategy.lastUpdate = time.Now()
	}
}
