package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranspose(t *testing.T) {
	t.Run("should return [[1,4,7],[2,5,8],[3,6,9]]", func(t *testing.T) {
		matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		expected := [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}
		assert.Equal(t, expected, transpose(matrix))
	})

	t.Run("should return [[1,4],[2,5],[3,6]]", func(t *testing.T) {
		matrix := [][]int{{1, 2, 3}, {4, 5, 6}}
		expected := [][]int{{1, 4}, {2, 5}, {3, 6}}
		assert.Equal(t, expected, transpose(matrix))
	})
}
