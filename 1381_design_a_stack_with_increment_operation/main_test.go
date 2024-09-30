package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomStack(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		maxSize := 3
		stack := Constructor(maxSize)
		stack.Push(1)
		stack.Push(2)
		assert.Equal(t, 2, stack.Pop())
		stack.Push(2)
		stack.Push(3)
		stack.Push(4)
		assert.Equal(t, 3, len(stack.Data))
		stack.Increment(5, 100)
		assert.Equal(t, 101, stack.Data[0])
		assert.Equal(t, 102, stack.Data[1])
		assert.Equal(t, 103, stack.Data[2])
		stack.Increment(2, 100)
		assert.Equal(t, 201, stack.Data[0])
		assert.Equal(t, 202, stack.Data[1])
		assert.Equal(t, 103, stack.Data[2])
		assert.Equal(t, 103, stack.Pop())
		assert.Equal(t, 202, stack.Pop())
		assert.Equal(t, 201, stack.Pop())
		assert.Equal(t, -1, stack.Pop())
	})
}
