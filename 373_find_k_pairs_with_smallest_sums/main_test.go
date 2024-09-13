package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKSmallestPairs(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums1 := []int{1, 7, 11}
		nums2 := []int{2, 4, 6}
		k := 3
		expected := [][]int{{1, 2}, {1, 4}, {1, 6}}

		assert.Equal(t, expected, kSmallestPairs(nums1, nums2, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums1 := []int{1, 1, 2}
		nums2 := []int{1, 2, 3}
		k := 2
		expected := [][]int{{1, 1}, {1, 1}}

		assert.Equal(t, expected, kSmallestPairs(nums1, nums2, k))
	})
}
