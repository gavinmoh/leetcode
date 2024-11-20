package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumTime(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		grid := [][]int{{0, 1, 3, 2}, {5, 1, 2, 5}, {4, 3, 8, 6}}
		expected := 7

		assert.Equal(t, expected, minimumTime(grid))
	})

	t.Run("test case 2", func(t *testing.T) {
		grid := [][]int{{0, 2, 4}, {3, 2, 1}, {1, 0, 4}}
		expected := -1

		assert.Equal(t, expected, minimumTime(grid))
	})
}
