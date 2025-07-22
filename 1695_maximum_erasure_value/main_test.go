package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumUniqueSubarray(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{4, 2, 4, 5, 6}
		expected := 17

		assert.Equal(t, expected, maximumUniqueSubarray(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{5, 2, 1, 2, 5, 2, 1, 2, 5}
		expected := 8

		assert.Equal(t, expected, maximumUniqueSubarray(nums))
	})
}
