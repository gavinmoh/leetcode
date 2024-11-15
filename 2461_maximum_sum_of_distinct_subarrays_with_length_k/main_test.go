package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumSubarraySum(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 5, 4, 2, 9, 9, 9}
		k := 3
		expected := int64(15)

		assert.Equal(t, expected, maximumSubarraySum(nums, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{4, 4, 4}
		k := 3
		expected := int64(0)

		assert.Equal(t, expected, maximumSubarraySum(nums, k))
	})
}
