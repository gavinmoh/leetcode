package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpiralMatrixIII(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		rows := 1
		cols := 4
		rStart := 0
		cStart := 0
		expected := [][]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}}

		assert.Equal(t, expected, spiralMatrixIII(rows, cols, rStart, cStart))
	})

	t.Run("test case 2", func(t *testing.T) {
		rows := 5
		cols := 6
		rStart := 1
		cStart := 4
		expected := [][]int{
			{1, 4}, {1, 5}, {2, 5}, {2, 4}, {2, 3}, {1, 3},
			{0, 3}, {0, 4}, {0, 5}, {3, 5}, {3, 4}, {3, 3},
			{3, 2}, {2, 2}, {1, 2}, {0, 2}, {4, 5}, {4, 4},
			{4, 3}, {4, 2}, {4, 1}, {3, 1}, {2, 1}, {1, 1},
			{0, 1}, {4, 0}, {3, 0}, {2, 0}, {1, 0}, {0, 0},
		}

		assert.Equal(t, expected, spiralMatrixIII(rows, cols, rStart, cStart))
	})

}
