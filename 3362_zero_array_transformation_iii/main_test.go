package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxRemoval(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{2, 0, 2}
		queries := [][]int{{0, 2}, {0, 2}, {1, 1}}
		expected := 1

		assert.Equal(t, expected, maxRemoval(nums, queries))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 1, 1, 1}
		queries := [][]int{{1, 3}, {0, 2}, {1, 3}, {1, 2}}
		expected := 2

		assert.Equal(t, expected, maxRemoval(nums, queries))
	})
}
