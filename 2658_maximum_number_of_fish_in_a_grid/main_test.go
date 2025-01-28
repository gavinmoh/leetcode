package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxFish(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		grid := [][]int{{0, 2, 1, 0}, {4, 0, 0, 3}, {1, 0, 0, 4}, {0, 3, 2, 0}}
		expected := 7

		assert.Equal(t, expected, findMaxFish(grid))
	})

	t.Run("test case 2", func(t *testing.T) {
		grid := [][]int{{1, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 1}}
		expected := 1

		assert.Equal(t, expected, findMaxFish(grid))
	})
}
