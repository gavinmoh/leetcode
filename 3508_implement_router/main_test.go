package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		memoryLimit := 3
		router := Constructor(memoryLimit)

		assert.True(t, router.AddPacket(1, 4, 90))
		assert.True(t, router.AddPacket(2, 5, 90))
		assert.False(t, router.AddPacket(1, 4, 90))
		assert.True(t, router.AddPacket(3, 5, 95))
		assert.True(t, router.AddPacket(4, 5, 105))
		assert.Equal(t, []int{2, 5, 90}, router.ForwardPacket())
		assert.True(t, router.AddPacket(5, 2, 110))
		assert.Equal(t, 1, router.GetCount(5, 100, 110))
	})

	t.Run("test case 2", func(t *testing.T) {
		memoryLimit := 2
		router := Constructor(memoryLimit)

		assert.True(t, router.AddPacket(7, 4, 90))
		assert.Equal(t, []int{7, 4, 90}, router.ForwardPacket())
		assert.Equal(t, []int{}, router.ForwardPacket())
	})
}
