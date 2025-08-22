package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumArea(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		grid := [][]int{{0, 1, 0}, {1, 0, 1}}
		expected := 6

		assert.Equal(t, expected, minimumArea(grid))
	})

	t.Run("test case 2", func(t *testing.T) {
		grid := [][]int{{1, 0}, {0, 0}}
		expected := 1

		assert.Equal(t, expected, minimumArea(grid))
	})
}
