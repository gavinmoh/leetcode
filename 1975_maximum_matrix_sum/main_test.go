package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxMatrixSum(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		matrix := [][]int{{1, -1}, {-1, 1}}
		expected := int64(4)

		assert.Equal(t, expected, maxMatrixSum(matrix))
	})

	t.Run("test case 2", func(t *testing.T) {
		matrix := [][]int{{1, 2, 3}, {-1, -2, -3}, {1, 2, 3}}
		expected := int64(16)

		assert.Equal(t, expected, maxMatrixSum(matrix))
	})
}
