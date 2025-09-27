package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestTriangleArea(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		points := [][]int{{0, 0}, {0, 1}, {1, 0}, {0, 2}, {2, 0}}
		expected := 2.0

		assert.Equal(t, expected, largestTriangleArea(points))
	})

	t.Run("test case 2", func(t *testing.T) {
		points := [][]int{{1, 0}, {0, 0}, {0, 1}}
		expected := 0.5

		assert.Equal(t, expected, largestTriangleArea(points))
	})
}
