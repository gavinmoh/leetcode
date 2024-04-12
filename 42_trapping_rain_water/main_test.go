package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrap(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
		expected := 6

		assert.Equal(t, expected, trap(height))
	})

	t.Run("test case 2", func(t *testing.T) {
		height := []int{4, 2, 0, 3, 2, 5}
		expected := 9

		assert.Equal(t, expected, trap(height))
	})
}
