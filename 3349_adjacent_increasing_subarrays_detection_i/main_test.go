package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasIncreasingSubarrays(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1}
		k := 3

		assert.True(t, hasIncreasingSubarrays(nums, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7}
		k := 5

		assert.False(t, hasIncreasingSubarrays(nums, k))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{-3, -19, -8, -16}
		k := 2

		assert.False(t, hasIncreasingSubarrays(nums, k))
	})

	t.Run("test case 4", func(t *testing.T) {
		nums := []int{-15, 19}
		k := 1

		assert.True(t, hasIncreasingSubarrays(nums, k))
	})

	t.Run("test case 5", func(t *testing.T) {
		nums := []int{5, 8, -2, -1}
		k := 2

		assert.True(t, hasIncreasingSubarrays(nums, k))
	})
}
