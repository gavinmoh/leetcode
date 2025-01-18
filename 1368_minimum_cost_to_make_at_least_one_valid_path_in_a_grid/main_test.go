package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinCost(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		grid := [][]int{
			{1, 1, 1, 1},
			{2, 2, 2, 2},
			{1, 1, 1, 1},
			{2, 2, 2, 2},
		}
		expected := 3

		assert.Equal(t, expected, minCost(grid))
	})

	t.Run("test case 2", func(t *testing.T) {
		grid := [][]int{
			{1, 1, 3},
			{3, 2, 2},
			{1, 1, 4},
		}
		expected := 0

		assert.Equal(t, expected, minCost(grid))
	})

	t.Run("test case 3", func(t *testing.T) {
		grid := [][]int{
			{1, 2},
			{4, 3},
		}
		expected := 1

		assert.Equal(t, expected, minCost(grid))
	})
}
