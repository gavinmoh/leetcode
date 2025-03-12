package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumCount(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{-2, -1, -1, 1, 2, 3}
		expected := 3

		assert.Equal(t, expected, maximumCount(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{-3, -2, -1, 0, 0, 1, 2}
		expected := 3

		assert.Equal(t, expected, maximumCount(nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{5, 20, 66, 1314}
		expected := 4

		assert.Equal(t, expected, maximumCount(nums))
	})
}
