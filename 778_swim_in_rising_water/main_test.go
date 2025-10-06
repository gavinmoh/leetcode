package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwimInWater(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		grid := [][]int{{0, 2}, {1, 3}}
		expected := 3

		assert.Equal(t, expected, swimInWater(grid))
	})

	t.Run("test case 2", func(t *testing.T) {
		grid := [][]int{
			{0, 1, 2, 3, 4},
			{24, 23, 22, 21, 5},
			{12, 13, 14, 15, 16},
			{11, 17, 18, 19, 20},
			{10, 9, 8, 7, 6},
		}
		expected := 16

		assert.Equal(t, expected, swimInWater(grid))
	})
}
