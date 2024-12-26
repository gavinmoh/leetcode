package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSumOfThreeSubarrays(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 2, 1, 2, 6, 7, 5, 1}
		k := 2
		expected := []int{0, 3, 5}

		assert.Equal(t, expected, maxSumOfThreeSubarrays(nums, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 2, 1, 2, 1, 2, 1, 2, 1}
		k := 2
		expected := []int{0, 2, 4}

		assert.Equal(t, expected, maxSumOfThreeSubarrays(nums, k))
	})
}
