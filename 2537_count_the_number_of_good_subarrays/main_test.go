package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountGood(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 1, 1, 1, 1}
		k := 10
		expected := int64(1)

		assert.Equal(t, expected, countGood(nums, k))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{3, 1, 4, 3, 2, 2, 4}
		k := 2
		expected := int64(4)

		assert.Equal(t, expected, countGood(nums, k))
	})
}
