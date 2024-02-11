package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCherryPickup(t *testing.T) {
	t.Run("should return 24", func(t *testing.T) {
		grid := [][]int{
			{3, 1, 1},
			{2, 5, 1},
			{1, 5, 5},
			{2, 1, 1},
		}
		expected := 24

		assert.Equal(t, expected, cherryPickup(grid))
	})

	t.Run("should return 28", func(t *testing.T) {
		grid := [][]int{
			{1, 0, 0, 0, 0, 0, 1},
			{2, 0, 0, 0, 0, 3, 0},
			{2, 0, 9, 0, 0, 0, 0},
			{0, 3, 0, 5, 4, 0, 0},
			{1, 0, 2, 3, 0, 0, 6},
		}
		expected := 28

		assert.Equal(t, expected, cherryPickup(grid))
	})

}
