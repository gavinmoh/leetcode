package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxAbsoluteSum(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, -3, 2, 3, -4}
		expected := 5

		assert.Equal(t, expected, maxAbsoluteSum(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{2, -5, 1, -4, 3, -2}
		expected := 8

		assert.Equal(t, expected, maxAbsoluteSum(nums))
	})
}
