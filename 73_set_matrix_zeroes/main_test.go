package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetZeroes(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
		expected := [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}
		setZeroes(matrix)

		assert.Equal(t, expected, matrix)
	})

	t.Run("test case 2", func(t *testing.T) {
		matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
		expected := [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}
		setZeroes(matrix)

		assert.Equal(t, expected, matrix)
	})
}
