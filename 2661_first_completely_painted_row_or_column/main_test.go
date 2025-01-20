package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstCompleteIndex(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		arr := []int{1, 3, 4, 2}
		mat := [][]int{{1, 4}, {2, 3}}
		expected := 2

		assert.Equal(t, expected, firstCompleteIndex(arr, mat))
	})

	t.Run("test case 2", func(t *testing.T) {
		arr := []int{2, 8, 7, 4, 1, 3, 5, 6, 9}
		mat := [][]int{{3, 2, 5}, {1, 4, 6}, {8, 7, 9}}
		expected := 3

		assert.Equal(t, expected, firstCompleteIndex(arr, mat))
	})

	t.Run("test case 3", func(t *testing.T) {
		arr := []int{1, 4, 5, 2, 6, 3}
		mat := [][]int{{4, 3, 5}, {1, 2, 6}}
		expected := 1

		assert.Equal(t, expected, firstCompleteIndex(arr, mat))
	})
}
