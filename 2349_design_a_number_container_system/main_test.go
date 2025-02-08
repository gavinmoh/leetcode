package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberContainers(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		obj := Constructor()
		assert.Equal(t, -1, obj.Find(10))
		obj.Change(2, 10)
		obj.Change(1, 10)
		obj.Change(3, 10)
		obj.Change(5, 10)
		assert.Equal(t, 1, obj.Find(10))
		obj.Change(1, 20)
		assert.Equal(t, 2, obj.Find(10))
	})
}
