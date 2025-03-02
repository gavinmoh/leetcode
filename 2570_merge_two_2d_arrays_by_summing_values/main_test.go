package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeArrays(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums1 := [][]int{{1, 2}, {2, 3}, {4, 5}}
		nums2 := [][]int{{1, 4}, {3, 2}, {4, 1}}
		expected := [][]int{{1, 6}, {2, 3}, {3, 2}, {4, 6}}

		assert.Equal(t, expected, mergeArrays(nums1, nums2))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums1 := [][]int{{2, 4}, {3, 6}, {5, 5}}
		nums2 := [][]int{{1, 3}, {4, 3}}
		expected := [][]int{{1, 3}, {2, 4}, {3, 6}, {4, 3}, {5, 5}}

		assert.Equal(t, expected, mergeArrays(nums1, nums2))
	})
}
