package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWaysToSplitArray(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{10, 4, -8, 7}
		expected := 2

		assert.Equal(t, expected, waysToSplitArray(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{2, 3, 1, 0}
		expected := 2

		assert.Equal(t, expected, waysToSplitArray(nums))
	})
}
