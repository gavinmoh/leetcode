package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinOperations(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 1, 1}
		expected := 3

		assert.Equal(t, expected, minOperations(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 5, 2, 4, 1}
		expected := 14

		assert.Equal(t, expected, minOperations(nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{8}
		expected := 0

		assert.Equal(t, expected, minOperations(nums))
	})
}
