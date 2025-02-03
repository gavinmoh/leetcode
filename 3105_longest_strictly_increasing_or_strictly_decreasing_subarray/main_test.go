package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestMonotonicSubarray(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 4, 3, 3, 2}
		expected := 2

		assert.Equal(t, expected, longestMonotonicSubarray(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{3, 3, 3, 3}
		expected := 1

		assert.Equal(t, expected, longestMonotonicSubarray(nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{3, 2, 1}
		expected := 3

		assert.Equal(t, expected, longestMonotonicSubarray(nums))
	})

	t.Run("test case 4", func(t *testing.T) {
		nums := []int{1}
		expected := 1

		assert.Equal(t, expected, longestMonotonicSubarray(nums))
	})
}
