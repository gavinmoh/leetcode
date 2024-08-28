package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSubIslands(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		grid1 := [][]int{{1, 1, 1, 0, 0}, {0, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 1, 1}}
		grid2 := [][]int{{1, 1, 1, 0, 0}, {0, 0, 1, 1, 1}, {0, 1, 0, 0, 0}, {1, 0, 1, 1, 0}, {0, 1, 0, 1, 0}}
		expected := 3

		assert.Equal(t, expected, countSubIslands(grid1, grid2))
	})

	t.Run("test case 2", func(t *testing.T) {
		grid1 := [][]int{{1, 0, 1, 0, 1}, {1, 1, 1, 1, 1}, {0, 0, 0, 0, 0}, {1, 1, 1, 1, 1}, {1, 0, 1, 0, 1}}
		grid2 := [][]int{{0, 0, 0, 0, 0}, {1, 1, 1, 1, 1}, {0, 1, 0, 1, 0}, {0, 1, 0, 1, 0}, {1, 0, 0, 0, 1}}
		expected := 2

		assert.Equal(t, expected, countSubIslands(grid1, grid2))
	})
}
