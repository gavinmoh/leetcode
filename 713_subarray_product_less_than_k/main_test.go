package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumSubarrayProductLessThanK(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{10, 5, 2, 6}
		k := 100
		expected := 8

		assert.Equal(t, expected, numSubarrayProductLessThanK(nums, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 2, 3}
		k := 0
		expected := 0

		assert.Equal(t, expected, numSubarrayProductLessThanK(nums, k))
	})
}
