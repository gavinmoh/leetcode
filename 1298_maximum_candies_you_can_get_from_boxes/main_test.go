package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxCandies(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		status := []int{1, 0, 1, 0}
		candies := []int{7, 5, 4, 100}
		keys := [][]int{{}, {}, {1}, {}}
		containedBoxes := [][]int{{1, 2}, {3}, {}, {}}
		initialBoxes := []int{0}
		expected := 16

		assert.Equal(t, expected, maxCandies(status, candies, keys, containedBoxes, initialBoxes))
	})

	t.Run("test case 2", func(t *testing.T) {
		status := []int{1, 0, 0, 0, 0, 0}
		candies := []int{1, 1, 1, 1, 1, 1}
		keys := [][]int{{1, 2, 3, 4, 5}, {}, {}, {}, {}, {}}
		containedBoxes := [][]int{{1, 2, 3, 4, 5}, {}, {}, {}, {}, {}}
		initialBoxes := []int{0}
		expected := 6

		assert.Equal(t, expected, maxCandies(status, candies, keys, containedBoxes, initialBoxes))
	})
}
