package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSubarrays(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 3, 5, 2, 7, 5}
		minK := 1
		maxK := 5
		expected := int64(2)

		assert.Equal(t, expected, countSubarrays(nums, minK, maxK))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 1, 1, 1}
		minK := 1
		maxK := 1
		expected := int64(10)

		assert.Equal(t, expected, countSubarrays(nums, minK, maxK))
	})
}
