package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinOperations(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{0, 1, 1, 1, 0, 0}
		expected := 3

		assert.Equal(t, expected, minOperations(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{0, 1, 1, 1}
		expected := -1

		assert.Equal(t, expected, minOperations(nums))
	})
}
