package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGridGame(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		grid := [][]int{{2, 5, 4}, {1, 5, 1}}
		expected := int64(4)

		assert.Equal(t, expected, gridGame(grid))
	})

	t.Run("test case 2", func(t *testing.T) {
		grid := [][]int{{3, 3, 1}, {8, 5, 2}}
		expected := int64(4)

		assert.Equal(t, expected, gridGame(grid))
	})

	t.Run("test case 3", func(t *testing.T) {
		grid := [][]int{{1, 3, 1, 15}, {1, 3, 3, 1}}
		expected := int64(7)

		assert.Equal(t, expected, gridGame(grid))
	})
}
