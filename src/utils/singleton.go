package utils

import (
	"sync"
)

type Singleton struct{}

var lock = &sync.Mutex{}

func GetSingletonInstance[T any](
	instance *T,
	constructor func(interface{}) *T,
	initParams interface{},
) *T {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = constructor(initParams)
		}
	}

	return instance
}
