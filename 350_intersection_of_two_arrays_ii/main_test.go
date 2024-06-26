package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntersect(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums1 := []int{1, 2, 2, 1}
		nums2 := []int{2, 2}
		expected := []int{2, 2}
		output := intersect(nums1, nums2)

		sort.Ints(expected)
		sort.Ints(output)
		assert.Equal(t, expected, output)
	})

	t.Run("test case 2", func(t *testing.T) {
		nums1 := []int{4, 9, 5}
		nums2 := []int{9, 4, 9, 8, 4}
		expected := []int{4, 9}
		output := intersect(nums1, nums2)

		sort.Ints(expected)
		sort.Ints(output)
		assert.Equal(t, expected, output)
	})
}
