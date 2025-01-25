package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexicographicallySmallestArray(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 5, 3, 9, 8}
		limit := 2
		expected := []int{1, 3, 5, 8, 9}

		assert.Equal(t, expected, lexicographicallySmallestArray(nums, limit))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 7, 6, 18, 2, 1}
		limit := 3
		expected := []int{1, 6, 7, 18, 1, 2}

		assert.Equal(t, expected, lexicographicallySmallestArray(nums, limit))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{1, 7, 28, 19, 10}
		limit := 3
		expected := []int{1, 7, 28, 19, 10}

		assert.Equal(t, expected, lexicographicallySmallestArray(nums, limit))
	})
}
