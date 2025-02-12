package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumSum(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{18, 43, 36, 13, 7}
		expected := 54

		assert.Equal(t, expected, maximumSum(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{10, 12, 19, 14}
		expected := -1

		assert.Equal(t, expected, maximumSum(nums))
	})
}
