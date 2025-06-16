package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumDifference(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{7, 1, 5, 4}
		expected := 4

		assert.Equal(t, expected, maximumDifference(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{9, 4, 3, 2}
		expected := -1

		assert.Equal(t, expected, maximumDifference(nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{1, 5, 2, 10}
		expected := 9

		assert.Equal(t, expected, maximumDifference(nums))
	})
}
