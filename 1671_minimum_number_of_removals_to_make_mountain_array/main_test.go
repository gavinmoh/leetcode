package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumMountainRemovals(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 3, 1}
		expected := 0

		assert.Equal(t, expected, minimumMountainRemovals(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{2, 1, 1, 5, 6, 2, 3, 1}
		expected := 3

		assert.Equal(t, expected, minimumMountainRemovals(nums))
	})
}
