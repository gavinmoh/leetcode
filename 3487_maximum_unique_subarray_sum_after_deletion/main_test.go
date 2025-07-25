package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSum(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		expected := 15

		assert.Equal(t, expected, maxSum(nums))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 1, 0, 1, 1}
		expected := 1

		assert.Equal(t, expected, maxSum(nums))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{1, 2, -1, -2, 1, 0, -1}
		expected := 3

		assert.Equal(t, expected, maxSum(nums))
	})

	t.Run("test case 4", func(t *testing.T) {
		nums := []int{-100}
		expected := -100

		assert.Equal(t, expected, maxSum(nums))
	})

	t.Run("test case 5", func(t *testing.T) {
		nums := []int{-20, 20}
		expected := 20

		assert.Equal(t, expected, maxSum(nums))
	})
}
