package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumObstacles(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		grid := [][]int{{0, 1, 1}, {1, 1, 0}, {1, 1, 0}}
		expected := 2

		assert.Equal(t, expected, minimumObstacles(grid))
	})

	t.Run("test case 2", func(t *testing.T) {
		grid := [][]int{{0, 1, 0, 0, 0}, {0, 1, 0, 1, 0}, {0, 0, 0, 1, 0}}
		expected := 0

		assert.Equal(t, expected, minimumObstacles(grid))
	})

	t.Run("test case 3", func(t *testing.T) {
		grid := [][]int{{0, 1, 1}, {0, 1, 1}, {1, 1, 0}}
		expected := 2

		assert.Equal(t, expected, minimumObstacles(grid))
	})
}
