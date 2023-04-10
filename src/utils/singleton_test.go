package utils_test

import (
	"pcs/utils"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleton(t *testing.T) {
	type TestingSingleton struct {
		match int
		dict  map[int]int
	}

	constructor := func(...any) *TestingSingleton {
		return &TestingSingleton{match: 0, dict: make(map[int]int)}
	}
	lock := &sync.Mutex{}

	t.Run("inital singleton instance is nil, should initialize a value", func(t *testing.T) {
		var instance TestingSingleton

		rslt := utils.GetSingletonInstance(&instance, lock, constructor, nil)

		assert.NotNil(t, rslt)
		assert.NotNil(t, instance)
		assert.Equal(t, instance, *rslt)

		instance.match = 3

		assert.Equal(t, instance, *rslt)
		assert.Equal(t, instance.match, rslt.match)

		rslt.match = 6

		assert.Equal(t, instance, *rslt)
		assert.Equal(t, instance.match, rslt.match)
	})

	t.Run("initial singleton instance isn't nil, should not change value", func(t *testing.T) {
		instance := &TestingSingleton{}
		rslt := utils.GetSingletonInstance(instance, lock, constructor, nil)

		assert.Equal(t, instance, rslt)

		instance.match = 3

		assert.Equal(t, instance, rslt)
		assert.Equal(t, instance.match, rslt.match)
	})

	//t.Run("inital singleton instance is nil, should initialize a value (with maps)", func(t *testing.T) {
	//	var instance TestingSingleton
	//
	//	rslt := utils.GetSingletonInstance(&instance, lock, constructor, nil)
	//
	//	assert.NotNil(t, rslt)
	//	assert.NotNil(t, instance)
	//	assert.Equal(t, instance, *rslt)
	//
	//	instance.dict[1] = 3
	//
	//	assert.Equal(t, instance, *rslt)
	//	assert.Equal(t, instance.dict[1], rslt.dict[1])
	//
	//	rslt.match = 6
	//
	//	assert.Equal(t, instance, *rslt)
	//	assert.Equal(t, instance.match, rslt.match)
	//})
}
