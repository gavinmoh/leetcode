package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanTraverseAllPairs(t *testing.T) {
	t.Run("should return true", func(t *testing.T) {
		nums := []int{2, 3, 6}

		assert.True(t, canTraverseAllPairs(nums))
	})

	t.Run("should return false", func(t *testing.T) {
		nums := []int{3, 9, 5}

		assert.False(t, canTraverseAllPairs(nums))
	})

	t.Run("should return true", func(t *testing.T) {
		nums := []int{4, 3, 12, 8}

		assert.True(t, canTraverseAllPairs(nums))
	})
}
