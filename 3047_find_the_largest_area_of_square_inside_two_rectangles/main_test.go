package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestSquareArea(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		bottomLeft := [][]int{{1, 1}, {2, 2}, {3, 1}}
		topRight := [][]int{{3, 3}, {4, 4}, {6, 6}}
		expected := int64(1)

		assert.Equal(t, expected, largestSquareArea(bottomLeft, topRight))
	})

	t.Run("test case 2", func(t *testing.T) {
		bottomLeft := [][]int{{1, 1}, {1, 3}, {1, 5}}
		topRight := [][]int{{5, 5}, {5, 7}, {5, 9}}
		expected := int64(4)

		assert.Equal(t, expected, largestSquareArea(bottomLeft, topRight))
	})

	t.Run("test case 3", func(t *testing.T) {
		bottomLeft := [][]int{{1, 1}, {2, 2}, {1, 2}}
		topRight := [][]int{{3, 3}, {4, 4}, {3, 4}}
		expected := int64(1)

		assert.Equal(t, expected, largestSquareArea(bottomLeft, topRight))
	})

	t.Run("test case 4", func(t *testing.T) {
		bottomLeft := [][]int{{1, 1}, {3, 3}, {3, 1}}
		topRight := [][]int{{2, 2}, {4, 4}, {4, 2}}
		expected := int64(0)

		assert.Equal(t, expected, largestSquareArea(bottomLeft, topRight))
	})
}
