package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxMoves(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		grid := [][]int{{2, 4, 3, 5}, {5, 4, 9, 3}, {3, 4, 2, 11}, {10, 9, 13, 15}}
		expected := 3

		assert.Equal(t, expected, maxMoves(grid))
	})

	t.Run("test case 2", func(t *testing.T) {
		grid := [][]int{{3, 2, 4}, {2, 1, 9}, {1, 1, 7}}
		expected := 0

		assert.Equal(t, expected, maxMoves(grid))
	})
}
