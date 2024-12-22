package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeftmostBuildingQueries(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		heights := []int{6, 4, 8, 5, 2, 7}
		queries := [][]int{{0, 1}, {0, 3}, {2, 4}, {3, 4}, {2, 2}}
		expected := []int{2, 5, -1, 5, 2}

		assert.Equal(t, expected, leftmostBuildingQueries(heights, queries))
	})

	t.Run("test case 2", func(t *testing.T) {
		heights := []int{5, 3, 8, 2, 6, 1, 4, 6}
		queries := [][]int{{0, 7}, {3, 5}, {5, 2}, {3, 0}, {1, 6}}
		expected := []int{7, 6, -1, 4, 6}

		assert.Equal(t, expected, leftmostBuildingQueries(heights, queries))
	})
}
