package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnesMinusZeros(t *testing.T) {
	t.Run("should return [[0,0,4],[0,0,4],[-2,-2,2]]", func(t *testing.T) {
		grid := [][]int{{0, 1, 1}, {1, 0, 1}, {0, 0, 1}}
		expected := [][]int{{0, 0, 4}, {0, 0, 4}, {-2, -2, 2}}
		assert.Equal(t, expected, onesMinusZeros(grid))
	})

	t.Run("should return [[5,5,5],[5,5,5]]", func(t *testing.T) {
		grid := [][]int{{1, 1, 1}, {1, 1, 1}}
		expected := [][]int{{5, 5, 5}, {5, 5, 5}}
		assert.Equal(t, expected, onesMinusZeros(grid))
	})
}
