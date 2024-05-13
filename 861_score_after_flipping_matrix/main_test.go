package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatrixScore(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		grid := [][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}}
		expected := 39

		assert.Equal(t, expected, matrixScore(grid))
	})

	t.Run("test case 2", func(t *testing.T) {
		grid := [][]int{{0}}
		expected := 1

		assert.Equal(t, expected, matrixScore(grid))
	})
}
