package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestRange(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := [][]int{{4, 10, 15, 24, 26}, {0, 9, 12, 20}, {5, 18, 22, 30}}
		expected := []int{20, 24}

		assert.Equal(t, expected, smallestRange(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := [][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}
		expected := []int{1, 1}

		assert.Equal(t, expected, smallestRange(nums))
	})
}
