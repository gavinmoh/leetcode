package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersection(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums1 := []int{1, 2, 2, 1}
		nums2 := []int{2, 2}
		expected := []int{2}

		assert.Equal(t, expected, intersection(nums1, nums2))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums1 := []int{4, 9, 5}
		nums2 := []int{9, 4, 9, 8, 4}
		expected := []int{9, 4}
		result := intersection(nums1, nums2)

		assert.Equal(t, 2, len(result))
		for _, num := range expected {
			assert.Contains(t, result, num)
		}
	})
}
