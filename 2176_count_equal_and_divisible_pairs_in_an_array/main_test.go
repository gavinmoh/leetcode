package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountPairs(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{3, 1, 2, 2, 2, 1, 3}
		k := 2
		expected := 4

		assert.Equal(t, expected, countPairs(nums, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		k := 1
		expected := 0

		assert.Equal(t, expected, countPairs(nums, k))
	})
}
