package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleton(t *testing.T) {
	type TestingSingleton struct{}

	constructor := func(interface{}) *TestingSingleton {
		return &TestingSingleton{}
	}

	t.Run("inital singleton instance is nil, should initialize a value", func(t *testing.T) {
		var instance *TestingSingleton = nil

		rslt := GetSingletonInstance(instance, constructor, nil)

		assert.NotNil(t, rslt)
	})

	t.Run("initial singleton instance isn't nil, should not change value", func(t *testing.T) {
		instance := &TestingSingleton{}
		rslt := GetSingletonInstance(instance, constructor, nil)

		assert.Equal(t, instance, rslt)
	})
}
