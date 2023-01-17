package utils

import (
	"sync"
)

type Singleton struct{}

func GetSingletonInstance[T any](
	instance *T,
	lock *sync.Mutex,
	constructor func(...any) *T,
	initParams ...any,
) *T {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = constructor(initParams...)
		}
	}

	return instance
}
