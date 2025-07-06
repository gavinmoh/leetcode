package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindSumPairs(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums1 := []int{1, 1, 2, 2, 2, 3}
		nums2 := []int{1, 4, 5, 2, 5, 4}

		obj := Constructor(nums1, nums2)
		assert.Equal(t, 8, obj.Count(7))

		obj.Add(3, 2)
		assert.Equal(t, 2, obj.Count(8))
		assert.Equal(t, 1, obj.Count(4))

		obj.Add(0, 1)
		obj.Add(1, 1)
		assert.Equal(t, 11, obj.Count((7)))
	})
}
