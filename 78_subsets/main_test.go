package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubsets(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 2, 3}
		expected := [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}
		result := subsets(nums)

		for _, subset := range result {
			assert.Contains(t, expected, subset)
		}
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{0}
		expected := [][]int{{}, {0}}
		result := subsets(nums)

		for _, subset := range result {
			assert.Contains(t, expected, subset)
		}
	})
}
