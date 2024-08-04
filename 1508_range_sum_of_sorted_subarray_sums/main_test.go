package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeSum(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		n := 4
		left := 1
		right := 5
		expected := 13

		assert.Equal(t, expected, rangeSum(nums, n, left, right))
	})

	t.Run("test case 2", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		n := 4
		left := 3
		right := 4
		expected := 6

		assert.Equal(t, expected, rangeSum(nums, n, left, right))
	})

	t.Run("test case 3", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		n := 4
		left := 1
		right := 10
		expected := 50

		assert.Equal(t, expected, rangeSum(nums, n, left, right))
	})

}
