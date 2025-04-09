package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinOperations(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{5, 2, 5, 4, 5}
		k := 2
		expected := 2

		assert.Equal(t, expected, minOperations(nums, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{2, 1, 2}
		k := 2
		expected := -1

		assert.Equal(t, expected, minOperations(nums, k))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{9, 7, 5, 3}
		k := 1
		expected := 4

		assert.Equal(t, expected, minOperations(nums, k))
	})
}
