package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSubarrays(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 3, 2, 3, 3}
		k := 2
		expected := int64(6)

		assert.Equal(t, expected, countSubarrays(nums, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 4, 2, 1}
		k := 3
		expected := int64(0)

		assert.Equal(t, expected, countSubarrays(nums, k))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{2, 2, 2, 2, 1, 3, 3, 2, 2, 1, 1, 3, 1, 1, 2, 3, 2, 1, 1, 2, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 3}
		k := 5
		expected := int64(31)

		assert.Equal(t, expected, countSubarrays(nums, k))
	})
}
