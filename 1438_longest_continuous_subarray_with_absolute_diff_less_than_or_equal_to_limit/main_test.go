package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSubarray(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{8, 2, 4, 7}
		limit := 4
		expected := 2

		assert.Equal(t, expected, longestSubarray(nums, limit))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{10, 1, 2, 4, 7, 2}
		limit := 5
		expected := 4

		assert.Equal(t, expected, longestSubarray(nums, limit))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{4, 2, 2, 2, 4, 4, 2, 2}
		limit := 0
		expected := 3

		assert.Equal(t, expected, longestSubarray(nums, limit))
	})
}
