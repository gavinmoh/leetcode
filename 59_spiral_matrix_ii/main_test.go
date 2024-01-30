package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateMatrix(t *testing.T) {
	t.Run("should generate a 3x3 matrix", func(t *testing.T) {
		n := 3
		expected := [][]int{
			{1, 2, 3},
			{8, 9, 4},
			{7, 6, 5},
		}
		assert.Equal(t, expected, generateMatrix(n))
	})

	t.Run("should generate a 1x1 matrix", func(t *testing.T) {
		n := 1
		expected := [][]int{
			{1},
		}
		assert.Equal(t, expected, generateMatrix(n))
	})
}
