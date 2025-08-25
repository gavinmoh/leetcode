package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDiagonalOrder(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		expected := []int{1, 2, 4, 7, 5, 3, 6, 8, 9}

		assert.Equal(t, expected, findDiagonalOrder(mat))
	})

	t.Run("test case 2", func(t *testing.T) {
		mat := [][]int{{1, 2}, {3, 4}}
		expected := []int{1, 2, 3, 4}

		assert.Equal(t, expected, findDiagonalOrder(mat))
	})
}
