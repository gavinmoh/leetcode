package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxPoints(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		grid := [][]int{{1, 2, 3}, {2, 5, 7}, {3, 5, 1}}
		queries := []int{5, 6, 2}
		expected := []int{5, 8, 1}

		assert.Equal(t, expected, maxPoints(grid, queries))
	})

	t.Run("test case 2", func(t *testing.T) {
		grid := [][]int{{5, 2, 1}, {1, 1, 2}}
		queries := []int{3}
		expected := []int{0}

		assert.Equal(t, expected, maxPoints(grid, queries))
	})
}
