package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrapRainWater(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		heightMap := [][]int{
			{1, 4, 3, 1, 3, 2},
			{3, 2, 1, 3, 2, 4},
			{2, 3, 3, 2, 3, 1},
		}
		expected := 4

		assert.Equal(t, expected, trapRainWater(heightMap))
	})

	t.Run("test case 2", func(t *testing.T) {
		heightMap := [][]int{
			{3, 3, 3, 3, 3},
			{3, 2, 2, 2, 3},
			{3, 2, 1, 2, 3},
			{3, 2, 2, 2, 3},
			{3, 3, 3, 3, 3},
		}
		expected := 10

		assert.Equal(t, expected, trapRainWater(heightMap))
	})
}
