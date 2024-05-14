package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMaximumGold(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		grid := [][]int{{0, 6, 0}, {5, 8, 7}, {0, 9, 0}}
		expected := 24

		assert.Equal(t, expected, getMaximumGold(grid))
	})

	t.Run("test case 1", func(t *testing.T) {
		grid := [][]int{{1, 0, 7}, {2, 0, 6}, {3, 4, 5}, {0, 3, 0}, {9, 0, 20}}
		expected := 28

		assert.Equal(t, expected, getMaximumGold(grid))
	})
}
