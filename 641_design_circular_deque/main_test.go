package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyCircularDeque(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		k := 3
		deque := Constructor(k)

		assert.True(t, deque.InsertLast(1))
		assert.True(t, deque.InsertLast(2))
		assert.True(t, deque.InsertFront(3))
		assert.False(t, deque.InsertFront(4))
		assert.Equal(t, 2, deque.GetRear())
		assert.True(t, deque.IsFull())
		assert.True(t, deque.DeleteLast())
		assert.True(t, deque.InsertFront(4))
		assert.Equal(t, 4, deque.GetFront())
	})
}
