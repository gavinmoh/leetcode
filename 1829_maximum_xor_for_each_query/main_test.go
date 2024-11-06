package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMaximumXor(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{0, 1, 1, 3}
		maximumBit := 2
		expected := []int{0, 3, 2, 3}

		assert.Equal(t, expected, getMaximumXor(nums, maximumBit))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{2, 3, 4, 7}
		maximumBit := 3
		expected := []int{5, 2, 6, 5}

		assert.Equal(t, expected, getMaximumXor(nums, maximumBit))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{0, 1, 2, 2, 5, 7}
		maximumBit := 3
		expected := []int{4, 3, 6, 4, 6, 7}

		assert.Equal(t, expected, getMaximumXor(nums, maximumBit))
	})
}
