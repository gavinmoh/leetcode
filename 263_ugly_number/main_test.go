package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsUgly(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		n := 6

		assert.True(t, isUgly(n))
	})

	t.Run("test case 2", func(t *testing.T) {
		n := 1

		assert.True(t, isUgly(n))
	})

	t.Run("test case 3", func(t *testing.T) {
		n := 14

		assert.False(t, isUgly(n))
	})

	t.Run("test case 4", func(t *testing.T) {
		n := 0

		assert.False(t, isUgly(n))
	})

	t.Run("test case 5", func(t *testing.T) {
		n := -1

		assert.False(t, isUgly(n))
	})
}
