package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubsequence(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{2, 1, 3, 3}
		k := 2
		expected := []int{3, 3}

		assert.Equal(t, expected, maxSubsequence(nums, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{-1, -2, 3, 4}
		k := 3
		expected := []int{-1, 3, 4}

		assert.Equal(t, expected, maxSubsequence(nums, k))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{3, 4, 3, 3}
		k := 2
		expected := []int{4, 3}

		assert.Equal(t, expected, maxSubsequence(nums, k))
	})
}
