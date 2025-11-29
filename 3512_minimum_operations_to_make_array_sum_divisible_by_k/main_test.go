package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinOperations(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{3, 9, 7}
		k := 5
		expected := 4

		assert.Equal(t, expected, minOperations(nums, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{4, 1, 3}
		k := 4
		expected := 0

		assert.Equal(t, expected, minOperations(nums, k))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{3, 2}
		k := 6
		expected := 5

		assert.Equal(t, expected, minOperations(nums, k))
	})
}
