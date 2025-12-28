package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountNegatives(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		grid := [][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}}
		expected := 8

		assert.Equal(t, expected, countNegatives(grid))
	})

	t.Run("test case 2", func(t *testing.T) {
		grid := [][]int{{3, 2}, {1, 0}}
		expected := 0

		assert.Equal(t, expected, countNegatives(grid))
	})
}
