package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestSubarray(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 1, 0, 1}
		expected := 3

		assert.Equal(t, expected, longestSubarray(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{0, 1, 1, 1, 0, 1, 1, 0, 1}
		expected := 5

		assert.Equal(t, expected, longestSubarray(nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{1, 1, 1}
		expected := 2

		assert.Equal(t, expected, longestSubarray(nums))
	})
}
