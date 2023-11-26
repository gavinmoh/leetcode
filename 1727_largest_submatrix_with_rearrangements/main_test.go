package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestSubmatrix(t *testing.T) {
	t.Run("should return 4", func(t *testing.T) {
		matrix := [][]int{{0, 0, 1}, {1, 1, 1}, {1, 0, 1}}
		expected := 4
		assert.Equal(t, expected, largestSubmatrix(matrix))
	})

	t.Run("should return 3", func(t *testing.T) {
		matrix := [][]int{{1, 0, 1, 0, 1}}
		expected := 3
		assert.Equal(t, expected, largestSubmatrix(matrix))
	})

	t.Run("should return 2", func(t *testing.T) {
		matrix := [][]int{{1, 1, 0}, {1, 0, 1}}
		expected := 2
		assert.Equal(t, expected, largestSubmatrix(matrix))
	})
}
