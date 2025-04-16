package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoodTriplets(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums1 := []int{2, 0, 1, 3}
		nums2 := []int{0, 1, 2, 3}
		expected := int64(1)

		assert.Equal(t, expected, goodTriplets(nums1, nums2))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums1 := []int{4, 0, 1, 3, 2}
		nums2 := []int{4, 1, 0, 2, 3}
		expected := int64(4)

		assert.Equal(t, expected, goodTriplets(nums1, nums2))
	})
}
