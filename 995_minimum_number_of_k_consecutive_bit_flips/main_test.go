package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinKBitFlips(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{0, 1, 0}
		k := 1
		expected := 2

		assert.Equal(t, expected, minKBitFlips(nums, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 1, 0}
		k := 2
		expected := -1

		assert.Equal(t, expected, minKBitFlips(nums, k))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{0, 0, 0, 1, 0, 1, 1, 0}
		k := 3
		expected := 3

		assert.Equal(t, expected, minKBitFlips(nums, k))
	})
}
