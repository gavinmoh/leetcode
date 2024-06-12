package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortColors(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{2, 0, 2, 1, 1, 0}
		expected := []int{0, 0, 1, 1, 2, 2}

		sortColors(nums)
		assert.Equal(t, expected, nums)
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{2, 0, 1}
		expected := []int{0, 1, 2}

		sortColors(nums)
		assert.Equal(t, expected, nums)
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 2, 0}
		expected := []int{0, 1, 2}

		sortColors(nums)
		assert.Equal(t, expected, nums)
	})
}
