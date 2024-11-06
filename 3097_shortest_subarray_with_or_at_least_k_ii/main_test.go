package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumSubarrayLength(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 2, 3}
		k := 2
		expected := 1

		assert.Equal(t, expected, minimumSubarrayLength(nums, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{2, 1, 8}
		k := 10
		expected := 3

		assert.Equal(t, expected, minimumSubarrayLength(nums, k))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{1, 2}
		k := 0
		expected := 1

		assert.Equal(t, expected, minimumSubarrayLength(nums, k))
	})
}
