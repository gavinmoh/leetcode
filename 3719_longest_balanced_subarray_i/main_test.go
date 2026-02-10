package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestBalanced(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{2, 5, 4, 3}
		expected := 4

		assert.Equal(t, expected, longestBalanced(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{3, 2, 2, 5, 4}
		expected := 5

		assert.Equal(t, expected, longestBalanced(nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{1, 2, 3, 2}
		expected := 3

		assert.Equal(t, expected, longestBalanced(nums))
	})
}
