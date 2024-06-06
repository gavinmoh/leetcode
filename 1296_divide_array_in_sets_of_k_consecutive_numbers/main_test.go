package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPossibleDivide(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 2, 3, 3, 4, 4, 5, 6}
		k := 4

		assert.True(t, isPossibleDivide(nums, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{3, 2, 1, 2, 3, 4, 3, 4, 5, 9, 10, 11}
		k := 3

		assert.True(t, isPossibleDivide(nums, k))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		k := 3

		assert.False(t, isPossibleDivide(nums, k))
	})
}
