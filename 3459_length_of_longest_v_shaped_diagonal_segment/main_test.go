package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLenOfVDiagonal(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		grid := [][]int{{2, 2, 1, 2, 2}, {2, 0, 2, 2, 0}, {2, 0, 1, 1, 0}, {1, 0, 2, 2, 2}, {2, 0, 0, 2, 2}}
		expected := 5

		assert.Equal(t, expected, lenOfVDiagonal(grid))
	})

	t.Run("test case 2", func(t *testing.T) {
		grid := [][]int{{2, 2, 2, 2, 2}, {2, 0, 2, 2, 0}, {2, 0, 1, 1, 0}, {1, 0, 2, 2, 2}, {2, 0, 0, 2, 2}}
		expected := 4

		assert.Equal(t, expected, lenOfVDiagonal(grid))
	})

	t.Run("test case 3", func(t *testing.T) {
		grid := [][]int{{1, 2, 2, 2, 2}, {2, 2, 2, 2, 0}, {2, 0, 0, 0, 0}, {0, 0, 2, 2, 2}, {2, 0, 0, 2, 0}}
		expected := 5

		assert.Equal(t, expected, lenOfVDiagonal(grid))
	})

	t.Run("test case 4", func(t *testing.T) {
		grid := [][]int{{1}}
		expected := 1

		assert.Equal(t, expected, lenOfVDiagonal(grid))
	})
}
