package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxEqualRowsAfterFlips(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		matrix := [][]int{{0, 1}, {1, 1}}
		expected := 1

		assert.Equal(t, expected, maxEqualRowsAfterFlips(matrix))
	})

	t.Run("test case 2", func(t *testing.T) {
		matrix := [][]int{{0, 1}, {1, 0}}
		expected := 2

		assert.Equal(t, expected, maxEqualRowsAfterFlips(matrix))
	})

	t.Run("test case 3", func(t *testing.T) {
		matrix := [][]int{{0, 0, 0}, {0, 0, 1}, {1, 1, 0}}
		expected := 2

		assert.Equal(t, expected, maxEqualRowsAfterFlips(matrix))
	})
}
