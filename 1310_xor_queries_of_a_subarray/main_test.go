package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXorQueries(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		arr := []int{1, 3, 4, 8}
		queries := [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}
		expected := []int{2, 7, 14, 8}

		assert.Equal(t, expected, xorQueries(arr, queries))
	})

	t.Run("test case 2", func(t *testing.T) {
		arr := []int{4, 8, 2, 10}
		queries := [][]int{{2, 3}, {1, 3}, {0, 0}, {0, 3}}
		expected := []int{8, 0, 4, 4}

		assert.Equal(t, expected, xorQueries(arr, queries))
	})

	t.Run("test case 3", func(t *testing.T) {
		arr := []int{16}
		queries := [][]int{{0, 0}, {0, 0}, {0, 0}}
		expected := []int{16, 16, 16}

		assert.Equal(t, expected, xorQueries(arr, queries))
	})

	t.Run("test case 4", func(t *testing.T) {
		arr := []int{1, 15, 8, 4}
		queries := [][]int{{2, 2}, {0, 0}, {0, 3}, {1, 2}}
		expected := []int{8, 1, 2, 7}

		assert.Equal(t, expected, xorQueries(arr, queries))
	})
}
