package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubarrayLength(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 2, 3, 1, 2, 3, 1, 2}
		k := 2
		expected := 6

		assert.Equal(t, expected, maxSubarrayLength(nums, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 2, 1, 2, 1, 2, 1, 2}
		k := 1
		expected := 2

		assert.Equal(t, expected, maxSubarrayLength(nums, k))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{5, 5, 5, 5, 5, 5, 5}
		k := 4
		expected := 4

		assert.Equal(t, expected, maxSubarrayLength(nums, k))
	})
}
