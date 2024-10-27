package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSquares(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		matrix := [][]int{
			{0, 1, 1, 1},
			{1, 1, 1, 1},
			{0, 1, 1, 1},
		}
		expected := 15

		assert.Equal(t, expected, countSquares(matrix))
	})

	t.Run("test case 2", func(t *testing.T) {
		matrix := [][]int{
			{1, 0, 1},
			{1, 1, 0},
			{1, 1, 0},
		}
		expected := 7

		assert.Equal(t, expected, countSquares(matrix))
	})
}
