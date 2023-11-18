package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxFrequency(t *testing.T) {
	t.Run("should return 3", func(t *testing.T) {
		nums := []int{1, 2, 4}
		k := 5
		expected := 3
		assert.Equal(t, expected, maxFrequency(nums, k))
	})

	t.Run("should return 2", func(t *testing.T) {
		nums := []int{1, 4, 8, 13}
		k := 5
		expected := 2
		assert.Equal(t, expected, maxFrequency(nums, k))
	})

	t.Run("should return 1", func(t *testing.T) {
		nums := []int{3, 9, 6}
		k := 2
		expected := 1
		assert.Equal(t, expected, maxFrequency(nums, k))
	})
}
