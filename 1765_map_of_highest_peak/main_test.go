package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHighestPeak(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		isWater := [][]int{{0, 1}, {0, 0}}
		expected := [][]int{{1, 0}, {2, 1}}

		assert.Equal(t, expected, highestPeak(isWater))
	})

	t.Run("test case 2", func(t *testing.T) {
		isWater := [][]int{{0, 0, 1}, {1, 0, 0}, {0, 0, 0}}
		expected := [][]int{{1, 1, 0}, {0, 1, 1}, {1, 2, 2}}

		assert.Equal(t, expected, highestPeak(isWater))
	})

	t.Run("test case 3", func(t *testing.T) {
		isWater := [][]int{
			{0, 0, 0, 0, 0, 0, 1, 0},
			{0, 1, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 1, 0},
			{0, 0, 1, 0, 0, 0, 0, 0},
		}
		expected := [][]int{
			{2, 1, 2, 3, 2, 1, 0, 1},
			{1, 0, 1, 2, 3, 2, 1, 2},
			{2, 1, 2, 3, 4, 3, 2, 3},
			{3, 2, 3, 4, 5, 4, 3, 4},
			{4, 3, 3, 4, 4, 3, 2, 3},
			{4, 3, 2, 3, 3, 2, 1, 2},
			{3, 2, 1, 2, 2, 1, 0, 1},
			{2, 1, 0, 1, 2, 2, 1, 2},
		}

		assert.Equal(t, expected, highestPeak(isWater))
	})
}
