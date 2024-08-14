package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestDistancePair(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 3, 1}
		k := 1
		expected := 0

		assert.Equal(t, expected, smallestDistancePair(nums, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 1, 1}
		k := 2
		expected := 0

		assert.Equal(t, expected, smallestDistancePair(nums, k))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{1, 6, 1}
		k := 3
		expected := 5

		assert.Equal(t, expected, smallestDistancePair(nums, k))
	})

	t.Run("test case 4", func(t *testing.T) {
		nums := []int{9, 10, 7, 10, 6, 1, 5, 4, 9, 8}
		k := 18
		expected := 2

		assert.Equal(t, expected, smallestDistancePair(nums, k))
	})
}
