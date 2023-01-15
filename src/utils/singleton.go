package utils

import (
	"fmt"
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
			fmt.Println("Creating Singleton instance now.")
			instance = constructor(initParams)
		} else {
			fmt.Println("Singleton instance already created.")
		}
	} else {
		fmt.Println("Singleton instance already created.")
	}

	return instance
}
