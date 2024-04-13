package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestRectangleArea(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		heights := []int{2, 1, 5, 6, 2, 3}
		expected := 10

		assert.Equal(t, expected, largestRectangleArea(heights))
	})

	t.Run("test case 2", func(t *testing.T) {
		heights := []int{2, 4}
		expected := 4

		assert.Equal(t, expected, largestRectangleArea(heights))
	})

	t.Run("test case 3", func(t *testing.T) {
		heights := []int{2, 1, 2}
		expected := 3

		assert.Equal(t, expected, largestRectangleArea(heights))
	})

}
