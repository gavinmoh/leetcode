package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPowerOfTwo(t *testing.T) {
	t.Run("should return true", func(t *testing.T) {
		assert.True(t, isPowerOfTwo(1))
		assert.True(t, isPowerOfTwo(16))
	})

	t.Run("should return false", func(t *testing.T) {
		assert.False(t, isPowerOfTwo(3))
	})
}
