package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumValueSum(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 2, 1}
		k := 3
		edges := [][]int{{0, 1}, {0, 2}}
		expected := int64(6)

		assert.Equal(t, expected, maximumValueSum(nums, k, edges))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{2, 3}
		k := 7
		edges := [][]int{{0, 1}}
		expected := int64(9)

		assert.Equal(t, expected, maximumValueSum(nums, k, edges))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{7, 7, 7, 7, 7, 7}
		k := 3
		edges := [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}}
		expected := int64(42)

		assert.Equal(t, expected, maximumValueSum(nums, k, edges))
	})
}
