package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumSubmatrixSumTarget(t *testing.T) {
	t.Run("should return 4", func(t *testing.T) {
		matrix := [][]int{
			{0, 1, 0},
			{1, 1, 1},
			{0, 1, 0},
		}
		target := 0
		expected := 4

		assert.Equal(t, expected, numSubmatrixSumTarget(matrix, target))
	})

	t.Run("should return 5", func(t *testing.T) {
		matrix := [][]int{
			{1, -1},
			{-1, 1},
		}
		target := 0
		expected := 5

		assert.Equal(t, expected, numSubmatrixSumTarget(matrix, target))
	})

	t.Run("should return 0", func(t *testing.T) {
		matrix := [][]int{
			{904},
		}
		target := 0
		expected := 0

		assert.Equal(t, expected, numSubmatrixSumTarget(matrix, target))
	})
}
