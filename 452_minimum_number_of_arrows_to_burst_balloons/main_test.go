package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMinArrowShots(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		points := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
		expected := 2

		assert.Equal(t, expected, findMinArrowShots(points))
	})

	t.Run("test case 2", func(t *testing.T) {
		points := [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
		expected := 4

		assert.Equal(t, expected, findMinArrowShots(points))
	})

	t.Run("test case 3", func(t *testing.T) {
		points := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}
		expected := 2

		assert.Equal(t, expected, findMinArrowShots(points))
	})

	t.Run("test case 4", func(t *testing.T) {
		points := [][]int{{-2147483648, 2147483647}}
		expected := 1

		assert.Equal(t, expected, findMinArrowShots(points))
	})
}
