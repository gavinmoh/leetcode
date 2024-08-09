package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumMagicSquaresInside(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		grid := [][]int{{4, 3, 8, 4}, {9, 5, 1, 9}, {2, 7, 6, 2}}
		expected := 1

		assert.Equal(t, expected, numMagicSquaresInside(grid))
	})

	t.Run("test case 2", func(t *testing.T) {
		grid := [][]int{{8}}
		expected := 0

		assert.Equal(t, expected, numMagicSquaresInside(grid))
	})

	t.Run("test case 3", func(t *testing.T) {
		grid := [][]int{{10, 3, 5}, {1, 6, 11}, {7, 9, 2}}
		expected := 0

		assert.Equal(t, expected, numMagicSquaresInside(grid))
	})

	t.Run("test case 4", func(t *testing.T) {
		grid := [][]int{{7, 0, 5}, {2, 4, 6}, {3, 8, 1}}
		expected := 0

		assert.Equal(t, expected, numMagicSquaresInside(grid))
	})
}
