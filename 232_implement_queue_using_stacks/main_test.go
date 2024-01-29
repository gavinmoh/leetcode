package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyQueue(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		myQueue := Constructor()
		myQueue.Push(1)
		myQueue.Push(2)
		assert.Equal(t, 1, myQueue.Peek())
		assert.Equal(t, 1, myQueue.Pop())
		assert.False(t, myQueue.Empty())
	})
}
