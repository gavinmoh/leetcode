package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountHillValley(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{2, 4, 1, 1, 6, 5}
		expected := 3

		assert.Equal(t, expected, countHillValley(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{6, 6, 5, 5, 4, 1}
		expected := 0

		assert.Equal(t, expected, countHillValley(nums))
	})
}
