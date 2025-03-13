package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinZeroArray(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{2, 0, 2}
		queries := [][]int{{0, 2, 1}, {0, 2, 1}, {1, 1, 3}}
		expected := 2

		assert.Equal(t, expected, minZeroArray(nums, queries))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{4, 3, 2, 1}
		queries := [][]int{{1, 3, 2}, {0, 2, 1}}
		expected := -1

		assert.Equal(t, expected, minZeroArray(nums, queries))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{0}
		queries := [][]int{{0, 0, 2}, {0, 0, 4}, {0, 0, 4}, {0, 0, 3}, {0, 0, 5}}
		expected := 0

		assert.Equal(t, expected, minZeroArray(nums, queries))
	})

	t.Run("test case 4", func(t *testing.T) {
		nums := []int{10}
		queries := [][]int{
			{0, 0, 5}, {0, 0, 3}, {0, 0, 2}, {0, 0, 1},
			{0, 0, 4}, {0, 0, 1}, {0, 0, 4}, {0, 0, 1},
			{0, 0, 4}, {0, 0, 5}, {0, 0, 3}, {0, 0, 4},
			{0, 0, 1},
		}
		expected := 3

		assert.Equal(t, expected, minZeroArray(nums, queries))
	})
}
