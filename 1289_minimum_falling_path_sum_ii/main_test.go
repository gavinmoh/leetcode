package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinFallingPathSum(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		grid := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}
		expected := 13

		assert.Equal(t, expected, minFallingPathSum(grid))
	})

	t.Run("test case 2", func(t *testing.T) {
		grid := [][]int{{7}}
		expected := 7

		assert.Equal(t, expected, minFallingPathSum(grid))
	})

	t.Run("test case 3", func(t *testing.T) {
		grid := [][]int{
			{-73, 61, 43, -48, -36},
			{3, 30, 27, 57, 10},
			{96, -76, 84, 59, -15},
			{5, -49, 76, 31, -7},
			{97, 91, 61, -46, 67},
		}
		expected := -192

		assert.Equal(t, expected, minFallingPathSum(grid))
	})
}
