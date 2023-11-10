package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestoreArray(t *testing.T) {

	t.Run("should return [1,2,3,4]", func(t *testing.T) {
		adjacentPairs := [][]int{{2, 1}, {3, 4}, {3, 2}}
		expected := [][]int{{1, 2, 3, 4}, {4, 3, 2, 1}}
		assert.Contains(t, expected, restoreArray(adjacentPairs))
	})

	t.Run("should return [-2,4,1,-3]", func(t *testing.T) {
		adjacentPairs := [][]int{{4, -2}, {1, 4}, {-3, 1}}
		expected := [][]int{{-2, 4, 1, -3}, {-3, 1, 4, -2}}
		assert.Contains(t, expected, restoreArray(adjacentPairs))
	})

	t.Run("should return [100000,-100000]", func(t *testing.T) {
		adjacentPairs := [][]int{{100000, -100000}}
		expected := [][]int{{100000, -100000}, {-100000, 100000}}
		assert.Contains(t, expected, restoreArray(adjacentPairs))
	})
}
