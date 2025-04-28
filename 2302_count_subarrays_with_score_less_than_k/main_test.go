package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSubarrays(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{2, 1, 4, 3, 5}
		k := int64(10)
		expected := int64(6)

		assert.Equal(t, expected, countSubarrays(nums, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 1, 1}
		k := int64(5)
		expected := int64(5)

		assert.Equal(t, expected, countSubarrays(nums, k))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{5, 2, 6, 8, 9, 7}
		k := int64(50)
		expected := int64(13)

		assert.Equal(t, expected, countSubarrays(nums, k))
	})
}
