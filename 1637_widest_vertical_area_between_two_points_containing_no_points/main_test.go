package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxWidthOfVerticalArea(t *testing.T) {
	t.Run("should return 1", func(t *testing.T) {
		points := [][]int{{8, 7}, {9, 9}, {7, 4}, {9, 7}}
		expected := 1
		assert.Equal(t, expected, maxWidthOfVerticalArea(points))
	})

	t.Run("should return 3", func(t *testing.T) {
		points := [][]int{{3, 1}, {9, 0}, {1, 0}, {1, 4}, {5, 3}, {8, 8}}
		expected := 3
		assert.Equal(t, expected, maxWidthOfVerticalArea(points))
	})
}
