package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindKDistantIndices(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{3, 4, 9, 1, 3, 9, 5}
		key := 9
		k := 1
		expected := []int{1, 2, 3, 4, 5, 6}

		assert.Equal(t, expected, findKDistantIndices(nums, key, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{2, 2, 2, 2, 2}
		key := 2
		k := 2
		expected := []int{0, 1, 2, 3, 4}

		assert.Equal(t, expected, findKDistantIndices(nums, key, k))
	})
}
