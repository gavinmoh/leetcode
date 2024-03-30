package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubarraysWithKDistinct(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 2, 1, 2, 3}
		k := 2
		expected := 7

		assert.Equal(t, expected, subarraysWithKDistinct(nums, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 2, 1, 3, 4}
		k := 3
		expected := 3

		assert.Equal(t, expected, subarraysWithKDistinct(nums, k))
	})
}
