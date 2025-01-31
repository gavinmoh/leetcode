package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestIsland(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		grid := [][]int{{1, 0}, {0, 1}}
		expected := 3

		assert.Equal(t, expected, largestIsland(grid))
	})

	t.Run("test case 2", func(t *testing.T) {
		grid := [][]int{{1, 1}, {1, 0}}
		expected := 4

		assert.Equal(t, expected, largestIsland(grid))
	})

	t.Run("test case 3", func(t *testing.T) {
		grid := [][]int{{1, 1}, {1, 1}}
		expected := 4

		assert.Equal(t, expected, largestIsland(grid))
	})
}
