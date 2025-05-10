package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinSum(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums1 := []int{3, 2, 0, 1, 0}
		nums2 := []int{6, 5, 0}
		expected := int64(12)

		assert.Equal(t, expected, minSum(nums1, nums2))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums1 := []int{2, 0, 2, 0}
		nums2 := []int{1, 4}
		expected := int64(-1)

		assert.Equal(t, expected, minSum(nums1, nums2))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums1 := []int{0, 7, 28, 17, 18}
		nums2 := []int{1, 2, 6, 26, 1, 0, 27, 3, 0, 30}
		expected := int64(98)

		assert.Equal(t, expected, minSum(nums1, nums2))
	})

	t.Run("test case 4", func(t *testing.T) {
		nums1 := []int{8, 13, 15, 18, 0, 18, 0, 0, 5, 20, 12, 27, 3, 14, 22, 0}
		nums2 := []int{29, 1, 6, 0, 10, 24, 27, 17, 14, 13, 2, 19, 2, 11}
		expected := int64(179)

		assert.Equal(t, expected, minSum(nums1, nums2))
	})
}
