package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCommon(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums1 := []int{1, 2, 3}
		nums2 := []int{2, 4}
		expected := 2

		assert.Equal(t, expected, getCommon(nums1, nums2))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums1 := []int{1, 2, 3, 6}
		nums2 := []int{2, 3, 4, 5}
		expected := 2

		assert.Equal(t, expected, getCommon(nums1, nums2))
	})
}
