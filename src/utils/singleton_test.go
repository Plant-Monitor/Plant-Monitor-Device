package utils_test

import (
	"pcs/utils"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleton(t *testing.T) {
	type TestingSingleton struct{}

	constructor := func(...any) *TestingSingleton {
		return &TestingSingleton{}
	}
	lock := &sync.Mutex{}

	t.Run("inital singleton instance is nil, should initialize a value", func(t *testing.T) {
		var instance *TestingSingleton = nil

		rslt := utils.GetSingletonInstance(instance, lock, constructor, nil)

		assert.NotNil(t, rslt)
	})

	t.Run("initial singleton instance isn't nil, should not change value", func(t *testing.T) {
		instance := &TestingSingleton{}
		rslt := utils.GetSingletonInstance(instance, lock, constructor, nil)

		assert.Equal(t, instance, rslt)
	})
}
