package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxAdjacentDistance(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 2, 4}
		expected := 3

		assert.Equal(t, expected, maxAdjacentDistance(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{-5, -10, -5}
		expected := 5

		assert.Equal(t, expected, maxAdjacentDistance(nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{-2, 1, -5}
		expected := 6

		assert.Equal(t, expected, maxAdjacentDistance(nums))
	})
}
