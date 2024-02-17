package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFurthestBuilding(t *testing.T) {
	t.Run("should return 4", func(t *testing.T) {
		heights := []int{4, 2, 7, 6, 9, 14, 12}
		bricks := 5
		ladders := 1
		expected := 4

		assert.Equal(t, expected, furthestBuilding(heights, bricks, ladders))
	})

	t.Run("should return 7", func(t *testing.T) {
		heights := []int{4, 12, 2, 7, 3, 18, 20, 3, 19}
		bricks := 10
		ladders := 2
		expected := 7

		assert.Equal(t, expected, furthestBuilding(heights, bricks, ladders))
	})

	t.Run("should return 3", func(t *testing.T) {
		heights := []int{14, 3, 19, 3}
		bricks := 17
		ladders := 0
		expected := 3

		assert.Equal(t, expected, furthestBuilding(heights, bricks, ladders))
	})

	t.Run("should return 5", func(t *testing.T) {
		heights := []int{1, 5, 1, 2, 3, 4, 10000}
		bricks := 4
		ladders := 1
		expected := 5

		assert.Equal(t, expected, furthestBuilding(heights, bricks, ladders))
	})
}
