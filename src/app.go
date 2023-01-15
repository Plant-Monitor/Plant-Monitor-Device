package main

import (
	"pcs/snapshots"
)

func main() {
	snapshotPublisher := snapshots.GetSnapshotPublisherInstance()
	snapshotPublisher.Run()
}
