package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPowerOfThree(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 27

		assert.True(t, isPowerOfThree(n))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 0

		assert.False(t, isPowerOfThree(n))
	})

	t.Run("test case 3", func(t *testing.T) {
		n := -1

		assert.False(t, isPowerOfThree(n))
	})
}
