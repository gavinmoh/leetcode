package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortestSubarray(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1}
		k := 1
		expected := 1

		assert.Equal(t, expected, shortestSubarray(nums, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 2}
		k := 4
		expected := -1

		assert.Equal(t, expected, shortestSubarray(nums, k))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{2, -1, 2}
		k := 3
		expected := 3

		assert.Equal(t, expected, shortestSubarray(nums, k))
	})

	t.Run("test case 4", func(t *testing.T) {
		nums := []int{48, 99, 37, 4, -31}
		k := 140
		expected := 2

		assert.Equal(t, expected, shortestSubarray(nums, k))
	})

	t.Run("test case 5", func(t *testing.T) {
		nums := []int{17, 85, 93, -45, -21}
		k := 150
		expected := 2

		assert.Equal(t, expected, shortestSubarray(nums, k))
	})
}
