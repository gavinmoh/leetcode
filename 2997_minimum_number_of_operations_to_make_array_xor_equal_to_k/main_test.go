package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinOperations(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{2, 1, 3, 4}
		k := 1
		expected := 2

		assert.Equal(t, expected, minOperations(nums, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{2, 0, 2, 0}
		k := 0
		expected := 0

		assert.Equal(t, expected, minOperations(nums, k))
	})
}
