package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinDifference(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{5, 3, 2, 4}
		expected := 0

		assert.Equal(t, expected, minDifference(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 5, 0, 10, 14}
		expected := 1

		assert.Equal(t, expected, minDifference(nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{3, 100, 20}
		expected := 0

		assert.Equal(t, expected, minDifference(nums))
	})
}
