package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNextGreaterElement(t *testing.T) {
	t.Run("should return [-1,3,-1]", func(t *testing.T) {
		nums1 := []int{4, 1, 2}
		nums2 := []int{1, 3, 4, 2}
		expected := []int{-1, 3, -1}
		assert.Equal(t, expected, nextGreaterElement(nums1, nums2))
	})

	t.Run("should return [3,-1]", func(t *testing.T) {
		nums1 := []int{2, 4}
		nums2 := []int{1, 2, 3, 4}
		expected := []int{3, -1}
		assert.Equal(t, expected, nextGreaterElement(nums1, nums2))
	})

}
