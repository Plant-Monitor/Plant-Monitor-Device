package models

import "time"

type Snapshot struct {
	Timestamp time.Time `json:"timestamp"`
}

func BuildSnapshot() *Snapshot {
	return &Snapshot{
		time.Now(),
	}
}
