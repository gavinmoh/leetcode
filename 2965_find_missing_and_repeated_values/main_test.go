package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMissingAndRepeatedValues(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		grid := [][]int{{1, 3}, {2, 2}}
		expected := []int{2, 4}

		assert.Equal(t, expected, findMissingAndRepeatedValues(grid))
	})

	t.Run("test case 2", func(t *testing.T) {
		grid := [][]int{{9, 1, 7}, {8, 9, 2}, {3, 4, 6}}
		expected := []int{9, 5}

		assert.Equal(t, expected, findMissingAndRepeatedValues(grid))
	})
}
