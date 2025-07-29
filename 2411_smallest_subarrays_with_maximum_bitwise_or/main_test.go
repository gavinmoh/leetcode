package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallestSubarrays(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 0, 2, 1, 3}
		expected := []int{3, 3, 2, 2, 1}

		assert.Equal(t, expected, smallestSubarrays(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 2}
		expected := []int{2, 1}

		assert.Equal(t, expected, smallestSubarrays(nums))
	})
}
