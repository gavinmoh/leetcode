package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateMatrix(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		mat := [][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		}
		expected := [][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		}

		assert.Equal(t, expected, updateMatrix(mat))
	})

	t.Run("test case 2", func(t *testing.T) {
		mat := [][]int{
			{0, 0, 0},
			{0, 1, 0},
			{1, 1, 1},
		}
		expected := [][]int{
			{0, 0, 0},
			{0, 1, 0},
			{1, 2, 1},
		}

		assert.Equal(t, expected, updateMatrix(mat))
	})
}
