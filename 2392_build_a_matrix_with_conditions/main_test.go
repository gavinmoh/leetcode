package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildMatrix(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		k := 3
		rowConditions := [][]int{{1, 2}, {3, 2}}
		colConditions := [][]int{{2, 1}, {3, 2}}
		expected := [][]int{
			{3, 0, 0},
			{0, 0, 1},
			{0, 2, 0},
		}

		assert.Equal(t, expected, buildMatrix(k, rowConditions, colConditions))
	})

	t.Run("test case 2", func(t *testing.T) {
		k := 3
		rowConditions := [][]int{{1, 2}, {2, 3}, {3, 1}, {2, 3}}
		colConditions := [][]int{{2, 1}}
		expected := [][]int{}

		assert.Equal(t, expected, buildMatrix(k, rowConditions, colConditions))
	})

	t.Run("test case 3", func(t *testing.T) {
		k := 6
		rowConditions := [][]int{{1, 4}, {4, 3}, {1, 2}, {1, 4}}
		colConditions := [][]int{{6, 2}, {5, 3}, {1, 3}, {5, 3}, {6, 4}, {2, 3}}
		expected := [][]int{
			{6, 0, 0, 0, 0, 0},
			{0, 5, 0, 0, 0, 0},
			{0, 0, 0, 0, 1, 0},
			{0, 0, 0, 2, 0, 0},
			{0, 0, 4, 0, 0, 0},
			{0, 0, 0, 0, 0, 3},
		}

		assert.Equal(t, expected, buildMatrix(k, rowConditions, colConditions))
	})
}
