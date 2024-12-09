package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumBeauty(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{4, 6, 1, 2}
		k := 2
		expected := 3

		assert.Equal(t, expected, maximumBeauty(nums, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 1, 1, 1}
		k := 10
		expected := 4

		assert.Equal(t, expected, maximumBeauty(nums, k))
	})
}
