package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestLocal(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		grid := [][]int{{9, 9, 8, 1}, {5, 6, 2, 6}, {8, 2, 6, 4}, {6, 2, 2, 2}}
		expected := [][]int{{9, 9}, {8, 6}}

		assert.Equal(t, expected, largestLocal(grid))
	})

	t.Run("test case 2", func(t *testing.T) {
		grid := [][]int{{1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 2, 1, 1}, {1, 1, 1, 1, 1}, {1, 1, 1, 1, 1}}
		expected := [][]int{{2, 2, 2}, {2, 2, 2}, {2, 2, 2}}

		assert.Equal(t, expected, largestLocal(grid))
	})
}
