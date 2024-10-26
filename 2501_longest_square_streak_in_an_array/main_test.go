package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSquareStreak(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{4, 3, 6, 16, 8, 2}
		expected := 3

		assert.Equal(t, expected, longestSquareStreak(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{2, 3, 5, 6, 7}
		expected := -1

		assert.Equal(t, expected, longestSquareStreak(nums))
	})
}
