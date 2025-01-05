package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXorAllNums(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums1 := []int{2, 1, 3}
		nums2 := []int{10, 2, 5, 0}
		expected := 13

		assert.Equal(t, expected, xorAllNums(nums1, nums2))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums1 := []int{1, 2}
		nums2 := []int{3, 4}
		expected := 0

		assert.Equal(t, expected, xorAllNums(nums1, nums2))
	})
}
