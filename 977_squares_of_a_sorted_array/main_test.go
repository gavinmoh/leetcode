package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedSquares(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{-4, -1, 0, 3, 10}
		expected := []int{0, 1, 9, 16, 100}

		assert.Equal(t, expected, sortedSquares(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{-7, -3, 2, 3, 11}
		expected := []int{4, 9, 9, 49, 121}

		assert.Equal(t, expected, sortedSquares(nums))
	})
}
