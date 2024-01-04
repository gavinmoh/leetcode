package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinOperations(t *testing.T) {
	t.Run("should return 4", func(t *testing.T) {
		nums := []int{2, 3, 3, 2, 2, 4, 2, 3, 4}
		expected := 4
		assert.Equal(t, expected, minOperations(nums))
	})

	t.Run("should return -1", func(t *testing.T) {
		nums := []int{2, 1, 2, 2, 3, 3}
		expected := -1
		assert.Equal(t, expected, minOperations(nums))
	})

	t.Run("should return 7", func(t *testing.T) {
		nums := []int{14, 12, 14, 14, 12, 14, 14, 12, 12, 12, 12, 14, 14, 12, 14, 14, 14, 12, 12}
		expected := 7
		assert.Equal(t, expected, minOperations(nums))
	})

}
