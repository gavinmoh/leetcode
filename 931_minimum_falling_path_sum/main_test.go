package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinFallingPathSum(t *testing.T) {
	t.Run("should return 13", func(t *testing.T) {
		matrix := [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}
		expected := 13
		assert.Equal(t, expected, minFallingPathSum(matrix))
	})

	t.Run("should return -59", func(t *testing.T) {
		matrix := [][]int{{-19, 57}, {-40, -5}}
		expected := -59
		assert.Equal(t, expected, minFallingPathSum(matrix))
	})
}
