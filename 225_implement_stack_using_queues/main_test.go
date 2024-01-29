package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyStack(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		stack := Constructor()
		stack.Push(1)
		stack.Push(2)
		assert.Equal(t, 2, stack.Top())
		assert.Equal(t, 2, stack.Pop())
		assert.False(t, stack.Empty())
	})
}
