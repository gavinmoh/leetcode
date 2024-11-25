package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlidingPuzzle(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		board := [][]int{{1, 2, 3}, {4, 0, 5}}
		expected := 1

		assert.Equal(t, expected, slidingPuzzle(board))
	})

	t.Run("test case 2", func(t *testing.T) {
		board := [][]int{{1, 2, 3}, {5, 4, 0}}
		expected := -1

		assert.Equal(t, expected, slidingPuzzle(board))
	})

	t.Run("test case 3", func(t *testing.T) {
		board := [][]int{{4, 1, 2}, {5, 0, 3}}
		expected := 5

		assert.Equal(t, expected, slidingPuzzle(board))
	})
}
