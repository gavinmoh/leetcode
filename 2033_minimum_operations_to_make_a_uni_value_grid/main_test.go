package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinOperations(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		grid := [][]int{{2, 4}, {6, 8}}
		x := 2
		expected := 4

		assert.Equal(t, expected, minOperations(grid, x))
	})

	t.Run("test case 2", func(t *testing.T) {
		grid := [][]int{{1, 5}, {2, 3}}
		x := 1
		expected := 5

		assert.Equal(t, expected, minOperations(grid, x))
	})

	t.Run("test case 3", func(t *testing.T) {
		grid := [][]int{{1, 2}, {3, 4}}
		x := 2
		expected := -1

		assert.Equal(t, expected, minOperations(grid, x))
	})
}
